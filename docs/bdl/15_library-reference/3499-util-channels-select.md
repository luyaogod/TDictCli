---
title: "util.Channels.select"
source: "fgl-topics/c_fgl_ext_util_Channels_select.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.select"
type: "concept"
---

# util.Channels.select

> Waits for activity on a set of TCP socket listening base.Channel objects.

## Syntax

```
util.Channels.select(
     channels DYNAMIC ARRAY OF base.Channel )
RETURNS INTEGER
```

1. channels is an array of `base.Channel` objects opened with
   [`base.Channel.openServerSocket()`](2992-base-channel-openserversocket.md "Open a TCP server socket channel.") or provided by [`util.Channels.accept()`](3497-util-channels-accept.md "Returns the accepted base.Channel object for a given a server-socket.").
2. Returns the index of the channel that can be used to read data from.

## Usage

The `util.Channels.select()` takes an array of [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") objects, all opened as TCP
socket readers with `base.Channel.openServerSocket()`.

If one of the listening channels has some data available to read, the `select()`
method returns the index of that channel in the array passed as parameter.

If the returned index matches the index of the initial server channel passed in the first call to
`select()`, use the [`util.Channels.accept()`](3497-util-channels-accept.md "Returns the accepted base.Channel object for a given a server-socket.") method to get a new `base.Channel`
object to read from that port. The new `base.Channel` object must be added to the set
of channels passed to `select()`, and can then be used to read data from it with
`base.Channel.readLine()`.

## Example

The server.4gl file:

```
IMPORT util

FUNCTION main()
    DEFINE x, n INTEGER
    DEFINE channels DYNAMIC ARRAY OF base.Channel
    DEFINE s base.Channel
    DEFINE c base.Channel
    DEFINE buf STRING

    LET s = base.Channel.create()
    CALL s.openServerSocket("127.0.0.1", 4711, "u")
    LET channels[1] = s

    DISPLAY "server: listening on port 4711"
    WHILE (x := util.Channels.select(channels)) > 0
        DISPLAY SFMT("server: select() returned index: %1",x)
        IF x == 1 THEN
            LET c = util.Channels.accept(s)
            LET n = channels.getLength() + 1
            LET channels[n] = c
            DISPLAY SFMT("server: accept() returned new channel: %1", n)
        ELSE
            LET c = channels[x]
            LET buf = c.readLine()
            IF c.isEof() THEN
                DISPLAY SFMT("server: close(%1)", x)
                CALL c.close()
                CALL channels.deleteElement(x)
            ELSE
                DISPLAY SFMT("server: read(%1): %2", x, buf)
                CALL c.writeLine(buf)
                CALL c.flush()
            END IF
        END IF
    END WHILE
END FUNCTION
```

The client.4gl source:

```
MAIN
    DEFINE x, n INTEGER
    DEFINE c base.Channel
    DEFINE pid INTEGER
    DEFINE buf STRING
    LET n = NVL(arg_val(1), 20)
    LET pid = fgl_getpid()
    LET c = base.Channel.create()
    DISPLAY SFMT("client pid:%1 start", pid)
    CALL c.openClientSocket("127.0.0.1", 4711, "u", 10)
    FOR x = 1 TO n
        LET buf = SFMT("pid=%1 msgnum:%2", pid, x)
        DISPLAY SFMT("client pid:%1 write:'%2'", pid, buf)
        CALL c.writeLine(buf)
        LET buf = c.readLine()
        DISPLAY SFMT("client pid:%1 read :'%2'", pid, buf)
        SLEEP 1 
    END FOR
    DISPLAY SFMT("client pid:%1 stop", pid)
END MAIN
```

Compile both programs, start server in one dedicated terminal, then start several client
instances and watch the outputs.

## Related links

**Related concepts**  

[util.Channels.accept](3497-util-channels-accept.md "Returns the accepted base.Channel object for a given a server-socket.")

[util.Channels.selectWithTimeout](3500-util-channels-selectwithtimeout.md "Waits for activity on a set of TCP socket listening base.Channel objects and returns after a given period of inactivity.")

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
