---
title: "The MessageServer class"
source: "fgl-topics/c_fgl_ClassMessageServer.html"
breadcrumb: "Library reference > Built-in packages > The base package > The MessageServer class"
type: "concept"
---

# The MessageServer class

> The base.MessageServer class allows a program to send a key action over the network to other programs using this service.

This class can be used to join a group of programs to be notified by simple messages (key
events). The programs can run on different machines connected together in a network.

> **Warning:**
>
> **Security Note:**
>
> Improper use of the `base.MessageServer` can lead to a denial-of-service (DoS)
> attack. Consider using a solution based on [`ON
> IDLE`](../11_user-interface/1897-on-idle-block.md) to achieve inter-process communication.

> **Important:**
>
> This feature is experimental and subject to change.

The `base.MessageServer` uses network API capabilities with Sockets and the
UDP protocol. The computers must be configured with a network. The UDP protocol does not
guarantee the transmission of datagrams, therefore messages sent with the MessageServer can
arrive out of order, duplicated, or go missing without notice.

The UDP port is `6600` and the IP address group is `224.0.1.1`.
These cannot be changed.

> **Important:**
>
> This feature is only supported in direct connection with the GDC
> front-end. It is not supported when using other front-ends or when using the GAS.

To implement inter-process communication, consider using a "pulling" method solution based on the
[`ON IDLE`](../11_user-interface/1897-on-idle-block.md) clause in dialogs: When the
current dialog is waiting for a user action, after some seconds of inactivity, `ON
IDLE` block is executed, and the program code can then for example query a database table of
events, where message producers insert new event rows. This solution is more reliable since the
events are persistent in the database, and can also contain detailed information.

## Child topics

- [base.MessageServer methods](3087-base-messageserver-methods.md)
- [Examples](3092-examples.md): base.MessageServer usage examples.
