package fiface

type IConnManager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID int64) (IConnection, error)
	Len() int
	ClearAll()
}
