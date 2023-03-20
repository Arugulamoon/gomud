# gomud

Home is a center. Taturo does not go on missions, but she goes with the group to Far East. Moving base. Home. Center.

Being in the world is moving along main storyline, sidequests. All the while progressing through alternating sequences of witnessing something occuring, story, and performing actions. actions that effect change in the context.

Performing actions prompts stories to progress and potentially being a trigger for another story. Which has a beginning and an end. 

But effectively, you can freeze time on stories, wander around, do other things. Progress and Nurture?

Commands and CommandSets applied to a World by a Character. Characters have Commands, Worlds have AvailableCommands.

Quest: Level, Tasks, Goals, Rewards
Currency: Gil, Tokens
Long Walkway -> Enter Context (Room)

## Server Terminal
```bash
go run cmd/server/main.go
```

## Client Terminal
```bash
go run cmd/client/main.go
```

## Alternative Client Terminal
```bash
openssl s_client -connect localhost:7324
```

## Generate Self-Signed SSL Certificate

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

### Create a server CSR with 'localhost' in CN
```bash
openssl req \
  -newkey rsa:2048 \
  -nodes \
  -keyout server.key \
  -subj "/C=CA/ST=ON/L=Ottawa/O=Eden-Walker/CN=localhost" \
  -out server.csr
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

### References:
* https://github.com/shuklalok/Mywork/tree/master/tls
* https://gist.github.com/denji/12b3a568f092ab951456
* https://superuser.com/questions/346958/can-the-telnet-or-netcat-clients-communicate-over-ssl
* https://eli.thegreenplace.net/2021/go-https-servers-with-tls/
* https://eli.thegreenplace.net/2021/go-socket-servers-with-tls/
* https://dev.to/hgsgtk/how-go-handles-network-and-system-calls-when-tcp-server-1nbd
* https://github.com/reiver/go-telnet
