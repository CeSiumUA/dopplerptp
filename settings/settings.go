package settings

type SettingsProvider interface {
	GetNodeAddresses() []string
}

/*func GetNodeAddresses() []string {
	return make([]string, 0)
}*/

func GetDefaultPort() int32 {
	return 4395
}

func GetSender() []byte {
	return make([]byte, 32)
}

func GetVerifyFunction() func([]byte) bool {
	return func(bytes []byte) bool {
		return true
	}
}

func GetSigningFunction() func([]byte) []byte {
	return func(bytes []byte) []byte {
		return make([]byte, 32)
	}
}
