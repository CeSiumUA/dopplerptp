package utils

type DotocotSerializer interface {
	Serialize() *[]byte
	Deserialize([]byte)
}
