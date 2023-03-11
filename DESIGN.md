World

World: SessionHandler
- listens for session events (connect, disconnect, input)
- performs actions on world based on events, for example:
  - Add/Remove Characters from rooms
  - Move Characters from one room to another
  - Message Characters in the room

Server: ConnectionHandler
- listens for client connections
- upon accepting connection, creates thread to:
  - create a session
  - tail connection for user input
  - read input and send event to event channel
