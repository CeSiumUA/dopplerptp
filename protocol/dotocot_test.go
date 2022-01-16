package protocol

import (
	"math/rand"
	"testing"
)

var basePackageSize int32 = 170

func TestSerializeDeserialize(t *testing.T) {

	var payloadSize int32 = 4

	dotocotMessage := createDotocotPackage(payloadSize)

	serializedBytes := dotocotMessage.Serialize()

	if int32(len(*serializedBytes)) != (basePackageSize + payloadSize) {
		t.Errorf("invalid array length, expected %d got %d", 142, len(*serializedBytes))
	}

	if (*serializedBytes)[0] != Version {
		t.Errorf("incorrect protocol version, expected %d got %d", Version, (*serializedBytes)[0])
	}

	deserialized := Dotocot{}
	err := deserialized.Deserialize(*serializedBytes)

	if err != nil {
		t.Error(err)
	}

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

func TestVerify(t *testing.T) {
	dtct := Dotocot{}
	rawBytes := make([]byte, 100)

	err := dtct.Deserialize(rawBytes)
	if err == nil {
		t.Errorf("no error invoked on invalid package")
	}
}

func TestCreatePackage(t *testing.T) {
	sender := []byte{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	consumer := []byte{5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8}
	payload := make([]byte, 40)
	rand.Read(payload)
	CreateDotocotProtocolMessage(sender, consumer, payload, 1)
}

func BenchmarkSerialize(b *testing.B) {
	for x := 0; x < 10; x++ {
		protocol := createDotocotPackage(int32(b.N))
		protocol.Serialize()
	}
}

func createDotocotPackage(payloadLength int32) *Dotocot {
	sender := []byte{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	consumer := []byte{5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8, 5, 6, 7, 8}
	payload := make([]byte, payloadLength)

	return CreateDotocotProtocolMessage(sender, consumer, payload, 1)
}
