package protocol

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
)

const Version = 1
const PublicKeyLength = 32
const HashLength = 64

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

	isPackageValid := dtct.Verify(rawData)

	if !isPackageValid {
		return fmt.Errorf("package is not valid")
	}

	index := 0
	dtct.Version = rawData[0]
	index++

	dtct.Sender = rawData[index : index+PublicKeyLength]
	index += PublicKeyLength

	dtct.TargetConsumer = rawData[index : index+PublicKeyLength]
	index += PublicKeyLength

	dtct.Hash = rawData[index : index+HashLength]
	index += HashLength

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

	hashedBytes := dtct.Sender
	hashedBytes = append(hashedBytes, dtct.TargetConsumer...)
	hashedBytes = append(hashedBytes, dtct.Payload...)
	hashedBytes = append(hashedBytes, dtct.PayloadType)

	verifyHasher := sha512.New()
	verifyHasher.Write(hashedBytes)
	hash := verifyHasher.Sum(nil)

	hashCompareResult := bytes.Compare(dtct.Hash, hash)

	if hashCompareResult != 0 {
		return fmt.Errorf("hash is not valid")
	}

	return nil
}

func (dtct *Dotocot) Verify(rawData []byte) bool {
	minimalPackageLength := 1 + (2 * PublicKeyLength) + (HashLength) + 1 + 1 + 1
	atualPackageLength := len(rawData)
	return atualPackageLength >= minimalPackageLength
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

func CreateDotocotHandshakeMessage(sender []byte, address string) *Dotocot {
	payload := []byte(address)
	payloadLength := len(payload)

	targetConsumer := make([]byte, PublicKeyLength)
	payloadType := HANDSHAKE

	bytesToHash := make([]byte, 2*PublicKeyLength+payloadLength+1)

	index := 0
	copy(bytesToHash[index:PublicKeyLength], sender)
	index += PublicKeyLength
	copy(bytesToHash[index:PublicKeyLength], targetConsumer)
	index += PublicKeyLength
	copy(bytesToHash[index:payloadLength], payload)
	index += payloadLength
	copy(bytesToHash[index:1], []byte{payloadType})

	hasher512 := sha512.New()
	hasher512.Write(bytesToHash)
	hash := hasher512.Sum(nil)

	dotocot := Dotocot{
		Version:        Version,
		Sender:         sender,
		TargetConsumer: targetConsumer,
		Payload:        payload,
		PayloadLength:  uint64(len(payload)),
		PayloadType:    payloadType,
		Hash:           hash,
	}

	return &dotocot
}
