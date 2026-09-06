---
title: "Wait for front-end ping timeout"
source: "fgl-topics/c_fgl_feconn_ping_wait.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Wait for front-end ping timeout"
type: "concept"
description: "An end user may leave its computer for a while without using it. The network policy (firewall) might force a close of the TCP connection after a given period of inactivity. To avoid such connection ..."
---

# Wait for front-end ping timeout

An end user may leave its computer for a while without using it. The network policy (firewall)
might force a close of the TCP connection after a given period of inactivity. To avoid such
connection shutdown when there is no GUI exchange, the front-end sends successive ping events after
a timeout period of inactivity, to keep the TCP connection alive. The front-end ping is a normal
behavior and part of the GUI client/server protocol, this is a built-in "keep alive" solution
provided by Genero.

> **Note:**
>
> This feature is not supported when running on mobile devices, or when displaying applications on
> mobile devices.

The inactivity timeout to send a ping event can be configured in front-ends. On the runtime
system side, it is also possible to configure the wait-for-ping timeout with the
`gui.protocol.pingTimeout` FGLPROFILE setting.

With this keep alive technique, a front-end connection always remains open, even if the user
leaves the workstation for several hours. If your network connection comes at a cost to support this
type of connection, it is possible to configure the front-end and turn off the ping event, when the
front-end allows it. Check the front-end configuration documentation for more details.

The wait-for-ping timeout period of the runtime system can be configured with the
`gui.protocol.pingTimeout` fglprofile entry (default is 600
seconds = 10 minutes):

```
gui.protocol.pingTimeout = 800
```

If you set the wait-for-ping timeout to a value lower than the ping delay of the front-end, the
program will stop with the -8063 error after that timeout, even if the TCP connection is still
alive. For example, when a front-end has a ping delay of 5 minutes, the recommended minimum value
for this parameter is about 330 seconds (5 minutes + 30 seconds to make sure the client ping
arrives).

If the front-end program is not stopped properly as happens when, for example, it is killed by a
system reboot, or when there is network failure, the TCP connection is lost and the runtime system
no longer receives ping events. In this case, the runtime system waits for a specified time, before
it raises the error [**-8063**](../15_library-reference/4483-genero-bdl-errors.md). By default (when no exception handler is defined), the runtime system just
stops. However, the error -8063 can be trapped with a `WHENEVER ERROR` or
`TRY/CATCH` block, to perform some finalization tasks. The program should then
immediately stop with `EXIT PROGRAM n` (where
`n > 0`).

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
