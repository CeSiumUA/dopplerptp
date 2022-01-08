package networking

import (
	"dopplerptp/settings"
	"fmt"
	"net"
	"sync"
)

func StartListener() error {
	localAddress := fmt.Sprintf("0.0.0.0:%d", settings.GetDefaultPort())
	listener, err := net.Listen("tcp", localAddress)
	if err != nil {
		return err
	}
	go listenConnections(&listener)
	return nil
}

func listenConnections(listener *net.Listener) {
	var mutex sync.Mutex
	for {
		conn, err := (*listener).Accept()
		if err != nil {
			continue
		}
		go acceptClient(&conn, &mutex)
	}
}

func acceptClient(conn *net.Conn, mutex *sync.Mutex) {
	netConn := Connection{
		Connection: conn,
	}
	err := netConn.PerformHandshake()
	if err != nil {
		return
	}
	mutex.Lock()
	connections = append(connections, &netConn)
	mutex.Unlock()
}
