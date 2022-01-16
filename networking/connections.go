package networking

import "sync"

var connections []*Connection

var mutex sync.Mutex = sync.Mutex{}

func AddConnection(cn *Connection) {

	mutex.Lock()
	connections = append(connections, cn)
	mutex.Unlock()
}
