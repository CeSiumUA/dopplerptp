package networking

import (
	"dopplerptp/protocol"
	"dopplerptp/settings"
	"net"
)

type Connection struct {
	Connection *net.Conn
}

func (cn *Connection) PerformHandshake() error {
	address := (*cn.Connection).RemoteAddr().String()
	dotocotHandshake := protocol.CreateDotocotHandshakeMessage(settings.GetSender(), address)
	bytes := dotocotHandshake.Serialize()
	_, err := (*cn.Connection).Write(*bytes)
	if err != nil {
		return err
	}
	readBuffer := make([]byte, 8192)
	read, err := conn.Read(readBuffer)
	if err != nil {
		return err
	}
	readBuffer = readBuffer[0:read]
	dotocotMessage := protocol.Dotocot{}
	err = dotocotMessage.Deserialize(readBuffer)
	return err
}
