package networking

import (
	"bytes"
	"dopplerptp/protocol"
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

	handshakeRequest := protocol.CreateDotocotHandshakeMessage(settings.GetSender(), address)
	serializedMessage := handshakeRequest.Serialize()

	_, err = conn.Write(*serializedMessage)
	if err != nil {
		conn.Close()
		return
	}

	readBuffer := make([]byte, 8192)
	read, err := conn.Read(readBuffer)
	if err != nil {
		conn.Close()
		return
	}
	readBuffer = readBuffer[0:read]
	dotocotMessage := protocol.Dotocot{}
	err = dotocotMessage.Deserialize(readBuffer)
	if err != nil {
		conn.Close()
		return
	}
	if dotocotMessage.PayloadType != protocol.HANDSHAKE || !bytes.Equal(dotocotMessage.TargetConsumer, settings.GetSender()) {
		conn.Close()
		return
	}
	mutex.Lock()
	dotocotConnection := Connection{
		Connection: &conn,
	}
	connections = append(connections, &dotocotConnection)
	mutex.Unlock()
}
