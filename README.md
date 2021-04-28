# TCP chat made with golang

## Commands

- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time
- `/rooms` - show list of available rooms to join.
- `/msg <msg>` - broadcast message to everyone in a room.
- `/quit` - disconnects from the chat server.

### Run with docker
To run with docker, you will need navigate to root directory and run the command `docker build -t <name:tag> .`

After this, you'll need run an container with this image, so:
`docker run -d -it --name <container-name> <name:tag>`

e.g: `docker build -t tcp-chat:v1` and after `docker run -d -it --name go-chat tcp-chat:v1`

If you want test outside the container, you can use telnet command, but for this you will need get the container IP, so:
    `docker inspect tcp-chat | grep IAddress`, 

The IP is like `172.17.0.x`

And finally, `telnet 172.17.0.x 8080`
Where 172.17.0.x is the container ip, and the 8080 is the exposed port

### Used Packages 
Just the color package by fatih: `https://github.com/fatih/color`

