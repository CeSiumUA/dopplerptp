package settings

var NetworkingSettings SettingsProvider = CreateDefaultSettingsProvider()

type SettingsProvider interface {
	GetNodeAddresses() []string
	GetDefaultPort() int32
	GetSender() []byte
	GetVerifyFunction() func([]byte) bool
	GetSigningFunction() func([]byte) []byte
}

type DefaultSettingsProvider struct {
	NodeAddresses   []string
	DefaultPort     int32
	Sender          []byte
	VerifyFunction  func([]byte) bool
	SigningFunction func([]byte) []byte
}

func (settings *DefaultSettingsProvider) GetNodeAddresses() []string {
	return settings.NodeAddresses
}

func (settings *DefaultSettingsProvider) GetDefaultPort() int32 {
	return settings.DefaultPort
}

func (settings *DefaultSettingsProvider) GetSender() []byte {
	return settings.Sender
}

func (settings *DefaultSettingsProvider) GetVerifyFunction() func([]byte) bool {
	return settings.VerifyFunction
}

func (settings *DefaultSettingsProvider) GetSigningFunction() func([]byte) []byte {
	return settings.SigningFunction
}

func CreateDefaultSettingsProvider() *DefaultSettingsProvider {
	return &DefaultSettingsProvider{
		NodeAddresses: make([]string, 0),
		DefaultPort:   4395,
		Sender:        make([]byte, 32),
		VerifyFunction: func(bytes []byte) bool {
			return true
		},
		SigningFunction: func(bytes []byte) []byte {
			return make([]byte, 32)
		},
	}
}
