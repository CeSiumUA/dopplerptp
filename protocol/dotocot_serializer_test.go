package protocol

import (
	"crypto"
	_ "crypto/sha512"
	"testing"
)

func TestSerialize(t *testing.T) {
	sender := []byte{1, 2, 3, 4}
	consumer := []byte{5, 6, 7, 8}
	payload := []byte{9, 10, 11, 12}
	resultArray := append(sender, consumer...)
	resultArray = append(resultArray, payload...)
	hash := crypto.SHA512.New().Sum(resultArray)
	dotocotMessage := Dotocot{
		Version:        Version,
		Sender:         sender,
		TargetConsumer: consumer,
		PayloadType:    1,
		Payload:        payload,
		Hash:           hash,
		PayloadLength:  4,
	}

	serializedBytes := dotocotMessage.Serialize()

	if len(*serializedBytes) != 98 {
		t.Errorf("invalid array length, expected %d got %d", 38, len(*serializedBytes))
	}

	if (*serializedBytes)[0] != Version {
		t.Errorf("incorrect protocol version, expected %d got %d", Version, (*serializedBytes)[0])
	}
}
