Star Denizen
====================

This is meant to be a stupid little game for people to abuse.
I want to make a game will allow people to float about and kill each other.

Goals:
- Simple web server
    + Single real thread
    + Uses go to the max
    + Simple auth
    + websockets
    + memory based data for speed
    + backs up world state?
    + is completely scriptable
    + simple auth -> gate writes to having a token
- Frontend
    + simple auth -> landing page
    + lets you spectate other players or yourself ->
    + scoreboard -> [user, money, time_alive]

Tracer
====================

- [x] Meta
    + [x] Project setup
        * [x] fe
        * [x] cmd
            + [x] server
            + [x] test-client
        * [x] internal
    + [x] justfile
- [ ] Backend
    + [x] Serves ws
    + [x] accepts connections
    + [x] break out comms
        * [x] make server obj
        * [x] make client handler
        * [x] make a message queue for broadcast
    + [x] loop
        * [x] make game loop sends messages
        * [x] client can change message
        * [x] clients see the new message
    + [ ] game
        * [ ] entities
    + [ ] game loop
        * [ ] sends coords to all connections
        * [ ] updates coords for all entities
        * [ ] accepts input from connections
    + [ ] There's a way to profile latency
        * [ ] send -> proc -> ack -> measure
- [x] CLI Frontend
    + [x] Connects to bews
    + [x] Echos
    + [ ] Prints all the messages to stdout
- [ ] Frontend
    + [ ] Connects to bews
    + [ ] Recvs coords from be ws
    + [ ] renders ship from bews coords
    + [ ] can send movement information
        * [ ] rollback?
        * [ ] optimism?
    + [ ] There's a way to profile latency
        * [ ] send -> proc -> ack -> measure
    + [ ] stars
- [ ] Data Layer
