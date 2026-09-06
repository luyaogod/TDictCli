---
title: "Client information"
source: "genero-install-topics/c_flm_clientinfo.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Client information"
type: "concept"
---

# Client information

> The client sends information when connecting to the DVM, which helps to identify the client for a session.

The `clientInfo` is sent to the DVM as part of the protocol between DVM and the
front-end.

You can output client information when using the command to show the registered active license
user list, `fglWrt -a info users`. The
information displayed may differ depending on the network architecture and the front-end. Typically,
the following information is displayed:

```
FrontEndId2   : frontendID
cx-mode       : cnType
app-id        : appId 
user-name     : uname 
host-addr     : ip 
hw-addr       : mac 
host-name     : hname
```

where:

**frontendID**

This identifies the front-end. For the web and desktop clients, the GUI server's IP address and
port, and the process identifier is displayed. For desktop clients, the Global Unique Identifier
(GUID) of the front-end is displayed.

**cnType**

This is the connection type. The value can be "gas" when the connection is through a web server
to the DVM, or "direct" when the connection is direct to the DVM via the Genero Desktop Client
(GDC).

**appId**

This is the Global Unique Identifier (GUID) of the application session number. This number is in
the URL /sua/appId/

**uname**

This is the local user name. The user that runs the front-end.

**ip**

This is the user computer IP address; this is the local IP address even if the connection goes
through a router or a proxy.

**mac**

This is the user computer mac address.

**hname**

This is the name of the computer that runs the DVM.

> **Note:**
>
> Depending on the proxy configuration or the web server configuration, hname
> and ip can be the proxy or the web server information instead of the user
> information.
