---
title: "base.Channel.openServerSocket"
source: "fgl-topics/c_fgl_ClassChannel_openServerSocket.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.openServerSocket"
type: "concept"
---

# base.Channel.openServerSocket

> Open a TCP server socket channel.

## Syntax

```
openServerSocket(
   interface STRING,
   port INTEGER,
   mode STRING )
```

1. interface is the name of the network interface to be used. This can be a
   hostname or an IPv4 address.
2. port is the port number of the service.
3. mode is the open mode. Only `"u"` is allowed (combined with
   `"b"` if needed).

## Usage

The `openServerSocket()` method initializes the channel object to
listen to a given TCP interface and port.

The server socket accepts multiple client connections. After calling the
`openServerSocket()` method, a call to `readLine()` waits until the
first client connects and returns after reading a complete line. Only one client connection can be
serviced at time: it's not possible to select a specific client connection. A client connection must
be closed by writing the EOF character to the channel. The EOF character is ASCII 26. Do not call
`base.Channel.close()` to close a client/server connection; this would close the
server socket and reject any pending client connection. The next call to `readLine()`
after writing EOF will wait until the next client connects or select the next pending client.

Pay attention to the character set used by the network protocol you want to use by
opening a channel with this method. The protocol must be based on ASCII, or must use the same
character set as the application.

The interface parameter indicates the hostname or the IPv4 address
of the network interface(s) to be used, in case the server uses different network adapters. Use
`NULL` to listen to all network interfaces, or when the server has only one network
interface. Specify `"0.0.0.0"` to listen to all interfaces.

The port parameter defines the TCP port to listen to.

The opening mode must be `"u"`, to read and write
from/to the socket. The method will raise error [-8085](4483-genero-bdl-errors.md) if the mode is different from
`"u"`.

The `"u"` mode can be combined with the `"b"` binary mode, to avoid
CR/LF translation on Windows®
platforms. The binary mode is only required in specific cases, and will only take effect when
writing data.

The method raises error [-8084](4483-genero-bdl-errors.md)
if the socket cannot be opened.

## Example

```
MAIN
    DEFINE io base.Channel
    DEFINE s STRING
    LET io = base.Channel.create()
    CALL io.openServerSocket("127.0.0.1", 4711, "u")
    WHILE TRUE
        LET s = io.readLine() 
        CALL io.writeLine(s)
        -- next line closes the current connection
        CALL io.writeLine(ASCII 26) -- EOF
    END WHILE
END MAIN
```

## Related links

**Related concepts**  

[Setup a TCP socket channel](3007-setup-a-tcp-socket-channel.md "Setup a TCP socket channel")

[base.Channel.openClientSocket](2989-base-channel-openclientsocket.md "Open a TCP client socket channel.")
