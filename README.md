# gomud

## Server Terminal
```bash
go run cmd/server/main.go
```

## Client Terminal
```bash
go run cmd/client/main.go
```

## Alternative Client: Native Telnet 
```bash
telnet localhost 8080
```

AI/NPCs also have session?

## Models

### Server
- Start(chan session.SessionEvent) error

### SessionEvent
- Session *Session
- Event   interface{}

#### Event
- SessionCreatedEvent
- SessionDisconnectedEvent
- SessionInputEvent
  - Input

### Session
- Id            string
- Connection    net.Conn
- EventChannel  chan SessionEvent
- User          *User
- New(net.Conn, chan SessionEvent) *Session
- SessionId() string
- WriteLine(string) error
- Stream() error

### SessionHandler
- World         *World
- EventChannel  <-chan SessionEvent
- NewSessionHandler(*World, <-chan SessionEvent) *SessionHandler
- Start()

### World
- Characters        []*Character
- Rooms             []*Room
- NewWorld() *World
- Init()
- HandleCharacterJoined(*Session)
- HandleCharacterLeft(*Session)
- HandleCharacterInput(*Session, string)
- MoveCharacter(*Session, *Room)

### Entity
- Id  string
- EntityId() string

### User
- Character *Character
- ADD: Session *Session
- GenerateName() string
- Request|AcceptSession(char Character) *Session ???

### Character
- Name  string
- Room  *Room

### Room
- Id          string
- Description string // Appearance
- Links       []*RoomLink
- Sessions    []*Session
- SendMessage(*Session, string)
- ContainsCharacter(string)
- AddCharacter(*Session)
- RemoveCharacter(*Session)

### RoomLink
- Verb    string
- RoomId  string

Quest: Level, Tasks, Goals, Rewards
Currency: Gil, Tokens
Long Walkway -> Enter Context (Room)
