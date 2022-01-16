package networking

import (
	"fmt"
	"net"

	"github.com/CeSiumUA/dopplerptp/settings"
)

func ExploreAddresses(addresses []string) {
	for _, addr := range addresses {
		go exploreDotocotPackage(addr)
	}
}

func exploreDotocotPackage(address string) {
	address = fmt.Sprintf("%s:%d", address, settings.NetworkingSettings.GetDefaultPort())
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}

	dotocotConnection := Connection{
		Connection: &conn,
	}

	err = dotocotConnection.PerformHandshake()

	if err != nil {
		return
	}
	AddConnection(&dotocotConnection)
}
