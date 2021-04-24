package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
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
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	//log.Printf("New client has connected: %s", conn.RemoteAddr().String())
	color.New(color.FgBlue).Printf("new client has connected: %s\n", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "Anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	nick := strings.TrimSpace(strings.Join(args[1:], " "))
	if nick == "" {
		//c.err(fmt.Errorf("you must provide a nick"))
		error := color.New(color.FgRed, color.BgBlack).FprintfFunc()
		error(c.conn, "error: you must provide a nick\n")
		return
	}

	c.nick = nick
	//c.msg(fmt.Sprintf("All right, I will call you %s", c.nick))
	sucess := color.New(color.FgBlue).FprintfFunc()
	sucess(c.conn, "all right, i will call you %s\n", c.nick)

}

func (s *server) join(c *client, args []string) {

	if len(args) < 2 {
		c.msg("room name is required. Usage: /join <ROOM-NAME>")
		return

	}

	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomName] = r
	}

	r.members[c.conn.LocalAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

	c.msg(fmt.Sprintf("Welcome to %s", roomName))
}

func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("Avaible rooms: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	// Checks if c.room is nil, otherwise it will cause runtime error
	// `nil pointer reference`
	if c.room == nil {
		c.msg("You need join a room before send messages")
		return
	}

	if len(args) < 2 {
		c.msg("Message is required, usage: /msg <MSG>")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+"-> "+msg)
}

func (s *server) quit(c *client) {
	color.MagentaString("log: client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("Sad to see u go :/")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, color.HiMagentaString("log: %s has left the room", c.nick))
	}
}
