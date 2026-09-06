---
title: "base.MessageServer.send"
source: "fgl-topics/c_fgl_ClassMessageServer_send.html"
breadcrumb: "Library reference > Built-in packages > The base package > The MessageServer class > base.MessageServer methods > base.MessageServer.send"
type: "concept"
---

# base.MessageServer.send

> Sends a key event to the group of programs connected together.

## Syntax

> **Warning:**
>
> **Security Note:**
>
> Improper use of the `base.MessageServer` can lead to a denial-of-service (DoS)
> attack. Consider using a solution based on [`ON
> IDLE`](../11_user-interface/1897-on-idle-block.md) to achieve inter-process communication.

```
base.MessageServer.send(
   message STRING )
```

1. message is a string expression defining the
   key event to be sent over the network.

## Usage

Once connected to the message server group with `base.MessageServer.connect()`,
a program calls the `base.MessageServer.send()` class method
to notify other programs registered to the group.

```
CALL base.MessageServer.send("f1")
```

All
programs registered to the message server group are notified, including
the program which has sent the message. The messages can be treated
by the current dialog with a simple `ON KEY()` interaction
block.
