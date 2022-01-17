package networking

import (
	"fmt"
	"sync"
)

var connections []*Connection

var mutex sync.Mutex = sync.Mutex{}

func AddConnection(cn *Connection) {

	mutex.Lock()
	connections = append(connections, cn)
	fmt.Printf("Added new connection: %s\n", (*cn.Connection).RemoteAddr())
	mutex.Unlock()
}
