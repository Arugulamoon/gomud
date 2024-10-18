# gomud

Proof of concept MUD (multi-user dungeon) in golang using sockets and TLS certs. References listed to put it together.

Start the server in one terminal window; start up additional clients in separate terminals; clients enter commands to move from a room to another; interact by waving etc

## Sample Output

### Server Terminal
```bash
$ go run ./cmd/server/main.go
2024/10/18 14:35:07 Starting async tcp server to receive messages
2024/10/18 14:36:05 Server accepted connection and created session: 1
2024/10/18 14:36:23 Received message on session 1: /goto Hallway
2024/10/18 14:37:02 Server accepted connection and created session: 2
2024/10/18 14:37:14 Received message on session 2: /goto LivingRoom
2024/10/18 14:37:14 target room not found
2024/10/18 14:37:24 Received message on session 2: /goto Hallway
2024/10/18 14:37:35 Received message on session 2: /goto LivingRoom
2024/10/18 14:37:41 Error handling connection read tcp [::1]:7324->[::1]:52099: wsarecv: An existing con
nection was forcibly closed by the remote host.
```

### Client 1 Terminal
```bash
$ go run ./cmd/client/main.go
You have entered your bedroom. There is a door leading out! (type "/goto Hallway
" to leave the bedroom)
Items:
  Book
Welcome Character 27!
/goto Hallway
You have entered a hallway with doors at either end. (type "/goto LivingRoom" to
 enter the living room or "/goto Bedroom" to enter the bedroom)
Character 2 entered the world.
Character 2 entered the room.
Character 2 left the room.
Character 2 left the world.
```

### Client 2 Terminal
```bash
$ go run ./cmd/client/main.go
You have entered your bedroom. There is a door leading out! (type "/goto Hallway
" to leave the bedroom)
Items:
  Book
Welcome Character 2!
/goto LivingRoom
There is no one around with that name...
/goto Hallway
You have entered a hallway with doors at either end. (type "/goto LivingRoom" to
 enter the living room or "/goto Bedroom" to enter the bedroom)
/goto LivingRoom
You have entered the living room. (type "/goto Hallway" to enter the hallway)
exit status 0xc000013a
```

## Setup: Generate Self-Signed SSL Certificate

### You must have a CA
```bash
openssl genrsa -out ca.key 2048
openssl req \
  -new \
  -x509 \
  -days 3650 \
  -key ca.key \
  -subj "/C=CA/ST=ON/L=Ottawa/O=Eden-Walker/CN=Eden-Walker Root CA" \
  -out ca.crt
```

On Windows Git Bash (difference in slashes):
```bash
openssl req -new -x509 -days 3650 -key ca.key -subj "//C=CA\ST=ON\L=Ottawa\O=Eden-Walker\CN=Eden-Walker Root CA" -out ca.crt
```

### Create a server CSR with 'localhost' in CN
```bash
openssl req \
  -newkey rsa:2048 \
  -nodes \
  -keyout server.key \
  -subj "/C=CA/ST=ON/L=Ottawa/O=Eden-Walker/CN=localhost" \
  -out server.csr
```

On Windows Git Bash (difference in slashes):
```bash
openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "//C=CA\ST=ON\L=Ottawa\O=Eden-Walker\CN=localhost" -out server.csr
```

### Finally sign server cert by CA and pass the subjectAltName when you signing server cert
```bash
openssl x509 \
  -req \
  -extfile <(printf "subjectAltName=DNS:localhost") \
  -days 3650 \
  -in server.csr \
  -CA ca.crt \
  -CAkey ca.key \
  -CAcreateserial \
  -out server.crt
```

On Windows Git Bash (issue with printf command):
```bash
echo subjectAltName=DNS:localhost > openssl.cnf
openssl x509 -req -extfile openssl.cnf -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
```

## Run

### Server Terminal
```bash
go run cmd/server/main.go
```

### Client Terminal
```bash
go run cmd/client/main.go
```

### Alternative Client Terminal
```bash
openssl s_client -connect localhost:7324
```

## References:
* https://github.com/shuklalok/Mywork/tree/master/tls
* https://gist.github.com/denji/12b3a568f092ab951456
* https://superuser.com/questions/346958/can-the-telnet-or-netcat-clients-communicate-over-ssl
* https://eli.thegreenplace.net/2021/go-https-servers-with-tls/
* https://eli.thegreenplace.net/2021/go-socket-servers-with-tls/
* https://dev.to/hgsgtk/how-go-handles-network-and-system-calls-when-tcp-server-1nbd
* https://github.com/reiver/go-telnet
