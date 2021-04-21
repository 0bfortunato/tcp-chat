package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "Anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.msg("Nick is required. Usage: /nick <NAME>")
		return
	}

	c.nick = args[1]
	c.msg(fmt.Sprintf("All right, I will call you %s", c.nick))

}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.msg("Room name is required. Usage: /join <ROOM-NAME>")
		return
	}
}
