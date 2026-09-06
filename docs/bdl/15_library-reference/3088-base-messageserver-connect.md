---
title: "base.MessageServer.connect"
source: "fgl-topics/c_fgl_ClassMessageServer_connect.html"
breadcrumb: "Library reference > Built-in packages > The base package > The MessageServer class > base.MessageServer methods > base.MessageServer.connect"
type: "concept"
---

# base.MessageServer.connect

> Connects to the group of programs to be notified by a message.

## Syntax

> **Warning:**
>
> **Security Note:**
>
> Improper use of the `base.MessageServer` can lead to a denial-of-service (DoS)
> attack. Consider using a solution based on [`ON
> IDLE`](../11_user-interface/1897-on-idle-block.md) to achieve inter-process communication.

```
base.MessageServer.connect()
```

## Usage

Use the `connect()` method to join the group of programs that can be notified by a key event
message.
