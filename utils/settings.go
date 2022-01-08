package utils

type SettingsProvider interface {
	GetNodeAddresses() []string
}

type Settings struct {
}

func (s *Settings) GetNodeAddresses() []string {
	return make([]string, 0)
}
