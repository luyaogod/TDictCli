---
title: "Setup a TCP socket channel"
source: "fgl-topics/c_fgl_ClassChannel_tcp_sockets.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Setup a TCP socket channel"
type: "concept"
description: "The base.Channel class provides methods to implement basic TCP client and server programs. Important: Consider character set encodings when designing such programs. No implicit character set ..."
---

# Setup a TCP socket channel

The `base.Channel` class provides methods to implement basic TCP client and server
programs.

> **Important:**
>
> Consider character set encodings when designing such programs. No implicit
> character set conversion is done by the runtime system. Both client and server must use the same
> character set and length semantics.

The following code example implements a client program connecting to a TCP port, using the [`openClientSocket()`](2989-base-channel-openclientsocket.md "Open a TCP client socket channel.") method:

```
MAIN
  DEFINE ch base.Channel,
         time DATETIME HOUR TO SECOND,
         data STRING
  LET ch = base.Channel.create()
  CALL ch.openClientSocket("localhost",99999,"u",3)
  CALL ch.writeLine("get_time")
  LET time = ch.readLine()
  DISPLAY "client 1: ", time
  CALL ch.writeLine("get_string")
  LET data = ch.readLine()
  DISPLAY "client 2: ", data
  CALL ch.writeLine("disconnect")
  CALL ch.close()
END MAIN
```

The following code example implements the server program that can be used with the above client
program. The server program uses the [`openServerSocket()`](2992-base-channel-openserversocket.md "Open a TCP server socket channel.") and `readLine()` methods to listen to a
given TCP interface/port. Note that the connection with a client must be ended by sending an EOF
character (ASCII 26) to the client, the next `readLine()` call will wait for a new
client connection, or select a pending client connection:

```
MAIN
  DEFINE ch base.Channel,
         cmd, data STRING
  LET ch = base.Channel.create()
  DISPLAY "starting server..."
  CALL ch.openServerSocket(null, 99999, "u")
  WHILE TRUE
      LET cmd = ch.readLine()
      IF ch.isEof() THEN
        DISPLAY "Connection ended by client..."
        EXIT WHILE
      END IF
      DISPLAY "cmd: ", cmd
      IF cmd == "get_time" THEN
        CALL ch.writeLine(CURRENT HOUR TO SECOND)
      END IF
      IF cmd == "get_string" THEN
        LET data = "This is a string..."
        CALL ch.writeLine(data)
      END IF
      IF cmd == "disconnect" THEN
         CALL ch.writeLine(ASCII 26) -- EOF
      END IF
  END WHILE
  DISPLAY "end of server..."
END MAIN
```

## Related links

**Related concepts**  

[Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.")
