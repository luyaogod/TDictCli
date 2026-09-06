---
title: "FGLSERVER"
source: "fgl-topics/c_fgl_EnvVariables_FGLSERVER.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLSERVER"
type: "concept"
---

# FGLSERVER

> Defines the graphical front-end for the application.

In GUI mode, FGLSERVER defines the host name and port of the graphical front-end the runtime
system will connect to in order to display application forms.

If FGLSERVER is not defined, the runtime system (fglrun) assumes that the
front-end executes on the same computer.

The values for the FGLSERVER environment variable must be specified with the following
syntax:

```
{hostname|ip-address}[:server-num]
```

1. hostname is the name of a machine on the network.
2. ip-address is the IPv4 address ( Ex: 10:0:0:105 ).
3. server-num identifies the front-end.

The server-num parameter defines the front-end server number (first is 0,
second is 1, and so on). This defines implicitly the TCP port number the front-end is listening to,
as an offset for the base port 6400. For example, FGLSERVER=cobra:1 will use the TCP port 6401 (6400
+ 1). This parameter is optional, when not specified, it defaults to zero (port 6400).

## Related links

**Related concepts**  

[Genero user interface modes](../11_user-interface/1516-genero-user-interface-modes.md "User interface modes allow you to adapt the application form rendering to different types of displays.")

[Automatic front-end startup](../11_user-interface/1530-automatic-front-end-startup.md "Automatic front-end startup")

[Front-end connection](0541-front-end-connection.md "To execute a Genero program with a graphical user interface, you need to specify the front-end (i.e. the graphical server) to the runtime system.")
