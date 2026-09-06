---
title: "Front-end identification"
source: "fgl-topics/c_fgl_feconn_fe_ident.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Front-end identification"
type: "concept"
description: "To start a program from the front-end platform, the front-end can open a terminal session on the application server. This is done for example by using a ssh , rlogin , or telnet terminal session. When ..."
---

# Front-end identification

To start a program from the front-end platform, the front-end can open a terminal session on the
application server. This is done for example by using a ssh,
rlogin, or telnet terminal session. When the terminal session
is open, the front-end sends some shell commands to set environment variables, like FGLSERVER,
before starting the program to display the application forms on the front-end where the terminal
session was initiated.

In this configuration, front-end identification takes place. The front-end identification
prevents the display of application forms on a front-end that did not start the program on the
server. If the front-end is not identified, a serious security issue can arise, as anyone could run
a fake program to display on any front-end and ask for a password.

> **Important:**
>
> **Front-end identification is achieved automatically by an
> initial protocol handshake. However, there can be a security hole if regular operating system users
> on the application server can overwrite the program or the shell script started by the front-end
> terminal session. Malicious programs can try to display the application on another workstation to
> read confidential data. As long as basic application users do not have read and write privileges on
> the program files, there is no risk. To make sure that program files on the server side are
> protected from basic users, create a special user on the server to manage the application program
> files, and give other users only read access to those files. As long as basic users cannot modify
> programs on the server side, there is no security issue.**

## Related links

**Related concepts**  

[FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.")
