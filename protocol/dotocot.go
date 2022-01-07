package protocol

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
)

const Version = 1
const publicKeyLength = 32
const hashLength = 64

type Dotocot struct {
	Version        byte
	Sender         []byte
	TargetConsumer []byte
	Hash           []byte
	PayloadType    byte
	PayloadLength  uint64
	Payload        []byte
}

func (dtct *Dotocot) Serialize() *[]byte {
	longBuff := make([]byte, 8)
	binary.PutUvarint(longBuff, dtct.PayloadLength)
	totalLength := 1 + len(dtct.Sender) + len(dtct.TargetConsumer) + len(dtct.Hash) + 1 + len(longBuff) + len(dtct.Payload)
	resultBytes := make([]byte, totalLength)
	index := 0
	resultBytes[0] = dtct.Version
	index++

	senderLength := len(dtct.Sender)
	copy(resultBytes[index:index+senderLength], dtct.Sender)
	index += senderLength

	consumerLength := len(dtct.TargetConsumer)
	copy(resultBytes[index:index+consumerLength], dtct.TargetConsumer)
	index += consumerLength

	hashLength := len(dtct.Hash)
	copy(resultBytes[index:index+hashLength], dtct.Hash)
	index += hashLength

	resultBytes[index] = dtct.PayloadType
	index++

	copy(resultBytes[index:index+8], longBuff)
	index += 8

	payLoadLength := len(dtct.Payload)
	copy(resultBytes[index:index+payLoadLength], dtct.Payload)

	return &resultBytes
}

func (dtct *Dotocot) Deserialize(rawData []byte) error {
	index := 0
	dtct.Version = rawData[0]
	index++

	dtct.Sender = rawData[index : index+publicKeyLength]
	index += publicKeyLength

	dtct.TargetConsumer = rawData[index : index+publicKeyLength]
	index += publicKeyLength

	dtct.Hash = rawData[index : index+hashLength]
	index += hashLength

	dtct.PayloadType = rawData[index]
	index++

	payLoadSize := rawData[index : index+8]
	index += 8
	buffer := bytes.NewBuffer(payLoadSize)
	payLoadLength, err := binary.ReadUvarint(buffer)
	if err != nil {
		return err
	}
	dtct.PayloadLength = payLoadLength

	dtct.Payload = rawData[index:]

	return nil
}

func CreateDotocotProtocolMessage(sender, targetConsumer, payload []byte, payloadType byte) *Dotocot {
	hashedBytes := sender
	hashedBytes = append(hashedBytes, targetConsumer...)
	hashedBytes = append(hashedBytes, payload...)
	hashedBytes = append(hashedBytes, payloadType)

	hasher512 := sha512.New()
	hasher512.Write(hashedBytes)
	hash := hasher512.Sum(nil)

	dotocot := Dotocot{
		Version:        Version,
		Sender:         sender,
		TargetConsumer: targetConsumer,
		PayloadLength:  uint64(len(payload)),
		Payload:        payload,
		PayloadType:    payloadType,
		Hash:           hash,
	}
	return &dotocot
}
