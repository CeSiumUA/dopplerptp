package utils

type DotocotSerializer interface {
	Serialize() *[]byte
	Deserialize([]byte) error
	Verify([]byte) bool
}
