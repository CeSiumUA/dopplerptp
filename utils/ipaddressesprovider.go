package utils

type IpAddressesProvider interface {
	GetAddresses() ([]string, error)
}
