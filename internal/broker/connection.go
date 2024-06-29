package broker

import "net"

type Connection struct {
	NodeId string
	Socket net.Conn
}
