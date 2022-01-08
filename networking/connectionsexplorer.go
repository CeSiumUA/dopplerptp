package networking

import (
	"dopplerptp/protocol"
	"dopplerptp/utils"
	"fmt"
	"net"
)

func ExploreAddresses(addresses []string) {
	for _, addr := range addresses {
		go exploreDotocotPackage(addr)
	}
}

func exploreDotocotPackage(address string) {
	address = fmt.Sprintf("%s:%d", address, utils.GetDefaultPort())
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}

	defer conn.Close()

	handshakeRequest := protocol.CreateDotocotHandshakeMessage(utils.GetSender(), address)
	serializedMessage := handshakeRequest.Serialize()

	_, err = conn.Write(*serializedMessage)
	if err != nil {
		return
	}
}
