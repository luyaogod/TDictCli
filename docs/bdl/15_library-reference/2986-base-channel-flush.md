---
title: "base.Channel.flush"
source: "fgl-topics/c_fgl_ClassChannel_flush.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.flush"
type: "concept"
---

# base.Channel.flush

> Flushes the channel.

## Syntax

```
flush()
```

## Usage

Call the `flush()` method when you want the written data to be flushed to the
output immediately, especially when using a bidirectional channel (`"u"` mode).

When the channel is opened with the unidirectional `"w"` mode, the flush is done
automatically when writing data to the output.

## Example

A typical use case of the `flush()` method is when using sockets and the server
program waits for client data, in a bidirectional socket.

The server.4gl file:

```
MAIN
    DEFINE ch base.Channel, s STRING
    LET ch = base.Channel.create()
    DISPLAY "server: open socket..."
    CALL ch.openServerSocket("127.0.0.1", 12345, "u")
    WHILE NOT ch.isEof()
        DISPLAY "server: waiting for input..."
        LET s = ch.readLine()
        IF NOT ch.isEof() THEN
            DISPLAY SFMT("Read: %1", s)
        END IF
    END WHILE
    DISPLAY "server: channel EOF..."
    CALL ch.close()
END MAIN
```

The client.4gl file:

```
MAIN
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    DISPLAY "client: open socket..."
    CALL ch.openClientSocket("127.0.0.1", 12345, "u", 5)
    DISPLAY "client: write line 1..."
    CALL ch.writeLine("Data line 1")
    DISPLAY "client: sleep 3 secs..."
    SLEEP 3
    DISPLAY "client: flush..."
    CALL ch.flush()
    DISPLAY "client: sleep 3 secs..."
    SLEEP 3
    CALL ch.close()
END MAIN
```

Start the server program, then the client program in separate terminals.

Output:

| Terminal 1: server progam | Terminal 1: client progam |
| --- | --- |
| $ fglrun server.42m server: open socket... server: waiting for input... |  |
|  | $ fglrun client.42m client: open socket... client: write line 1... client: sleep 3 secs... |
|  | client: flush... |
| Read: Data line 1 server: waiting for input... |  |
|  | client: sleep 3 secs... |
| server: channel EOF... |  |

## Related links

**Related concepts**  

[base.Channel.close](2983-base-channel-close.md "Closes the channel.")
