package networking

import (
	"dopplerptp/settings"
	"fmt"
	"net"
	"sync"
)

func ExploreAddresses(addresses []string) {
	var mutex sync.Mutex
	for _, addr := range addresses {
		go exploreDotocotPackage(addr, &mutex)
	}
}

func exploreDotocotPackage(address string, mutex *sync.Mutex) {
	address = fmt.Sprintf("%s:%d", address, settings.GetDefaultPort())
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

	mutex.Lock()
	connections = append(connections, &dotocotConnection)
	mutex.Unlock()
}
