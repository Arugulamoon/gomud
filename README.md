# gomud

## Server Terminal
```bash
go run main.go
```

## Client Terminal
```bash
telnet localhost 8080
```

AI/NPCs also have session

## Models

Server
- Start() error

SessionEvent
- Session *Session
- Event   interface{}

SessionCreatedEvent

SessionDisconnectedEvent

SessionInputEvent
- Input

Session
- Id          string
- Connection  net.Conn
- SessionId() string
- WriteLine() error
- GenerateSessionId() string

SessionHandler
- World         *World
- EventChannel  <-chan SessionEvent
- Users         map[string]*User
- NewSessionHandler() *SessionHandler
- Start()

World
- map[string]Entity ???
- Characters        []*Character
- Rooms             []*Room
- NewWorld() *World
- Init()
- HandleCharacterJoined()
- HandleCharacterLeft()
- HandleCharacterInput()
- MoveCharacter()

Entity
- Id  string
- EntityId() string

User
- Session   *Session
- Character *Character
- GenerateName() string
- Request|AcceptSession(char Character) *Session ???

Character
- Entity ???
- Name  string
- User  *User
- Room  *Room
- SendMessage()

Room
- Entity      ???
- []Entity    ???
- Id          string
- Description string
- RoomLinks   []*RoomLink
- Characters  []*Character
- SendMessage()
- AddCharacter()
- RemoveCharacter()

RoomLink
- Verb    string
- RoomId  string


Indexes
- Session -> User
