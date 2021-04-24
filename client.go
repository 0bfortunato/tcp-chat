package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/gookit/color"
)

type client struct {
	conn net.Conn
	nick string
	room *room
	// commands only receive from command type
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
				color:  color.BgBlack,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
				color:  color.FgDarkGray,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
				color:  color.Magenta,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
				color:  color.BgDefault,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
				color:  color.BgBlue,
			}

		default:
			c.err(fmt.Errorf("Unknow command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("Error: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("-> " + msg + "\n"))
}
