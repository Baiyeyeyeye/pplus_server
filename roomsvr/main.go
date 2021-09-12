package main

import (
	"log"

	"github.com/panjf2000/gnet"
)

type roomServer struct {
	*gnet.EventServer
}

func (es *roomServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = []byte("hello world")
	return out, gnet.None
}

func main() {
	echo := new(roomServer)
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}
