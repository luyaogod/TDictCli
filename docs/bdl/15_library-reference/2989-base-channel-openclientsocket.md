---
title: "base.Channel.openClientSocket"
source: "fgl-topics/c_fgl_ClassChannel_openClientSocket.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.openClientSocket"
type: "concept"
---

# base.Channel.openClientSocket

> Open a TCP client socket channel.

## Syntax

```
openClientSocket(
   host STRING,
   port INTEGER,
   mode STRING,
   timeout INTEGER )
```

1. host identifies the server machine you want to connect to. This can be a
   hostname or an IPv4 address.
2. port is the port number of the service.
3. mode is the open mode. Can be `"r"`, `"w"` or
   `"u"` (combined with `"b"` if needed).
4. timeout is the timeout in seconds. -1 indicates no timeout (wait
   forever).

## Usage

Use the `openClientSocket()` method to establish a TCP connection to a server.

Pay attention to the character set used by the network protocol you want to use by opening a
channel with this method. The protocol must be based on ASCII, or must use the same character set as
the application.

The host parameter defines the host name of the server.

The port parameter defines the TCP port to connect to.

The opening mode can be one of the following:

- `r`: For read mode: only to read from the socket
- `w`: For write mode: only to write to the socket
- `u`: For read and write mode: To read and write from/to the socket

Any of these modes can be followed by `b`, to use binary mode and
avoid CR/LF translation on Windows®
platforms. The binary mode is only required in specific cases, and will only take effect when
writing data.

If the opening mode is not one of the above letters, the method will
raise error [-8085](4483-genero-bdl-errors.md).

When the timeout parameter is -1, the connection waits forever. Otherwise, if
the timeout is expires, the error [-8086](4483-genero-bdl-errors.md) is raised.

The method raises error [-8084](4483-genero-bdl-errors.md) if the channel cannot be opened.

## Example

```
CALL ch.openClientSocket( "localhost", 80, "u", 5 )
```

For a complete example, see [Example 4: Communicating with an HTTP server](3012-example-4-communicating-with-an-http-server.md).

## Related links

**Related concepts**  

[Setup a TCP socket channel](3007-setup-a-tcp-socket-channel.md "Setup a TCP socket channel")

[base.Channel.openServerSocket](2992-base-channel-openserversocket.md "Open a TCP server socket channel.")
