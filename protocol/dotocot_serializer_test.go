package protocol

import (
	"crypto/sha512"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	sender := []byte{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	consumer := []byte{5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8}
	payload := []byte{9, 10, 11, 12}
	resultArray := append(sender, consumer...)
	resultArray = append(resultArray, payload...)

	hasher := sha512.New()

	hasher.Write(resultArray)

	hash := hasher.Sum(nil)

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

	if len(*serializedBytes) != 142 {
		t.Errorf("invalid array length, expected %d got %d", 142, len(*serializedBytes))
	}

	if (*serializedBytes)[0] != Version {
		t.Errorf("incorrect protocol version, expected %d got %d", Version, (*serializedBytes)[0])
	}

	deserialized := Dotocot{}
	deserialized.Deserialize(*serializedBytes)

	if deserialized.Version != dotocotMessage.Version {
		t.Errorf("incorrect protocol version, expected %d got %d", dotocotMessage.Version, deserialized.Version)
	}

	if string(deserialized.Sender) != string(dotocotMessage.Sender) {
		t.Errorf("incorrect sender, expected %s got %s", string(dotocotMessage.Sender), string(deserialized.Sender))
	}

	if string(deserialized.TargetConsumer) != string(dotocotMessage.TargetConsumer) {
		t.Errorf("incorrect consumer, expected %s got %s", string(dotocotMessage.TargetConsumer), string(deserialized.TargetConsumer))
	}

	if deserialized.PayloadType != dotocotMessage.PayloadType {
		t.Errorf("incorrect payload type, expected %d got %d", dotocotMessage.PayloadType, deserialized.PayloadType)
	}

	if string(deserialized.Payload) != string(dotocotMessage.Payload) {
		t.Errorf("incorrect payload, expected %s got %s", string(dotocotMessage.Payload), string(deserialized.Payload))
	}
}
