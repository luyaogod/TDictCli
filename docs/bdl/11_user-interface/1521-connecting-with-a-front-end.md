---
title: "Connecting with a front-end"
source: "fgl-topics/c_fgl_feconn_connection.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Connecting with a front-end"
type: "concept"
description: "In graphical mode, depending on the front-end technology that is used (desktop client, mobile client, web server client), there are different solutions to establish the connection between the runtime ..."
---

# Connecting with a front-end

In graphical mode, depending on the front-end technology that is used (desktop client, mobile
client, web server client), there are different solutions to establish the connection between the
runtime system and the front-end.

This topic describes the development context case, where programs are executed directly with
fglrun. In a production environment, programs will typically be started with
another technology, since the execution of programs will be triggered by the end user interacting
with the front-end. Read front-end specific documentation for more details.

From the point of view of the runtime system, the front-end acts as a graphical server and thus
the programs must connect to that GUI server in order to display forms and get user input.

The runtime system will try to connect to the front-end only when the first interactive
instruction like `MENU` or `INPUT` is reached.

For the runtime system, the front-end is identified by the [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.") environment variable. This variable
defines the host name of the machine where the front-end resides, and the number of the front-end
instances to be used.

The syntax for FGLSERVER is:

```
{hostname|ip-address}[:servernum]
```

For example:

```
$ FGLSERVER=fox:1
$ fglrun myprog
```

The servernum parameter is a whole number that defines the instance of the
front-end. It is actually defining a TCP port number the front-end is listening to,
starting from 6400. For example, if servernum equals 2, the TCP port
number used is 6402 (6400+2).

This is the standard/basic connection technique, but you can set up different types of
configurations. For example, you can have the front-end connect to an application server via ssh, to
pass through firewalls over the internet. Refer to the front-end documentation for more details.

There is an exception to the standard FGLSERVER specification, if the front-end is denied
permission to listen to a TCP port (this is the case with the Genero Mobile Development Client). If
you need to revert the connection principle in this particular case, use the
`--gui-listen` option of fglrun. With this option, the runtime
system will listen to the specified port, so the front-end can bind to the program and start to use
the GUI protocol. The procedure to work in such configuration is the following:

1. Start the program
   with:

   ```
    fglrun --gui-listen=tcp-port prog-name
   ```
2. Connect from the front-end, for example, with a URL using the following
   format:

   ```
   fgl://dev-server-hostname:tcp-port
   ```

## Related links

**Related concepts**  

[Mobile development mode](../17_mobile-applications/5083-mobile-development-mode.md "Set up a development environment to display app forms on a mobile front-end.")

[Graphical mode rendering (GUI mode)](1518-graphical-mode-rendering-gui-mode.md "Graphical mode rendering (GUI mode)")
