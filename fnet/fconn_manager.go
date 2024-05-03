package fnet

import (
	"github.com/frpc/fiface"
	"sync"
)

type ConnManager struct {
	connectionMap map[int64]fiface.IConnection
	connLock      sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connectionMap: make(map[int64]fiface.IConnection),
	}
}

func (c *ConnManager) Add(conn fiface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connectionMap[conn.GetConnID()] = conn
}

func (c *ConnManager) Remove(conn fiface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connectionMap, conn.GetConnID())
}

func (c *ConnManager) Get(connID int64) (fiface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	return c.connectionMap[connID], nil
}

func (c *ConnManager) Len() int {
	return len(c.connectionMap)
}

func (c *ConnManager) ClearAll() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connectionMap {
		conn.Stop()
		delete(c.connectionMap, connID)
	}
}
