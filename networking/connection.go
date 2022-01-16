package networking

import (
	"bytes"
	"fmt"
	"net"

	"github.com/CeSiumUA/dopplerptp/logging"
	"github.com/CeSiumUA/dopplerptp/protocol"
	"github.com/CeSiumUA/dopplerptp/settings"
)

type Connection struct {
	Connection        *net.Conn
	AssociatedAddress [protocol.PublicKeyLength]byte
}

func (cn *Connection) PerformHandshake() error {
	address := (*cn.Connection).RemoteAddr().String()
	dotocotHandshake := protocol.CreateDotocotHandshakeMessage(settings.NetworkingSettings.GetSender(), address)
	serializedBytes := dotocotHandshake.Serialize()
	wroteBytes, err := (*cn.Connection).Write(*serializedBytes)
	logging.GlobalLogger.LogInfo("wrote %d bytes to network channel %s", wroteBytes, (*cn.Connection).RemoteAddr().String())
	if err != nil {
		return err
	}
	readBuffer := make([]byte, 8192)
	read, err := (*cn.Connection).Read(readBuffer)
	logging.GlobalLogger.LogInfo("read %d bytes from network channel %s", read, (*cn.Connection).RemoteAddr().String())
	if err != nil {
		return err
	}
	readBuffer = readBuffer[0:read]
	dotocotMessage := protocol.Dotocot{}
	err = dotocotMessage.Deserialize(readBuffer)
	if err != nil {
		return err
	}
	if dotocotMessage.PayloadType != protocol.HANDSHAKE || !bytes.Equal(dotocotMessage.TargetConsumer, settings.NetworkingSettings.GetSender()) {
		return fmt.Errorf("incorrect handshake")
	}
	return err
}

func (cn *Connection) AcceptHandshake() error {
	readBuffer := make([]byte, 8192)
	read, err := (*cn.Connection).Read(readBuffer)
	if err != nil {
		return err
	}
	readBuffer = readBuffer[0:read]
	dotocotMessage := protocol.Dotocot{}
	err = dotocotMessage.Deserialize(readBuffer)
	if err != nil {
		return err
	}
	if dotocotMessage.PayloadType != protocol.HANDSHAKE || !bytes.Equal(dotocotMessage.TargetConsumer, settings.NetworkingSettings.GetSender()) {
		return fmt.Errorf("incorrect handshake")
	}
	address := (*cn.Connection).RemoteAddr().String()
	dotocotHandshake := protocol.CreateDotocotHandshakeMessage(settings.NetworkingSettings.GetSender(), address)
	serializedBytes := dotocotHandshake.Serialize()
	_, err = (*cn.Connection).Write(*serializedBytes)
	return err
}
