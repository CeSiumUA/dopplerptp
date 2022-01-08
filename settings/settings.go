package settings

type SettingsProvider interface {
	GetNodeAddresses() []string
}

func GetNodeAddresses() []string {
	return make([]string, 0)
}

func GetDefaultPort() int32 {
	return 4395
}

func GetSender() []byte {
	return make([]byte, 32)
}
