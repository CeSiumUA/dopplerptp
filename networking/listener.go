package networking

import (
	"fmt"
	"net"

	"github.com/CeSiumUA/dopplerptp/settings"
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
	for {
		conn, err := (*listener).Accept()
		if err != nil {
			continue
		}
		go acceptClient(&conn)
	}
}

func acceptClient(conn *net.Conn) {
	netConn := Connection{
		Connection: conn,
	}
	err := netConn.PerformHandshake()
	if err != nil {
		return
	}
	AddConnection(&netConn)
}
