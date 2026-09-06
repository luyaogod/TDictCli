---
title: "The front-end protocol"
source: "fgl-topics/c_fgl_feconn_protocol.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > The front-end protocol"
type: "concept"
description: "The front-end protocol (FEP) is an internal protocol used by the runtime system to synchronize the abstract user interface (AUI) representation on the front-end side. This protocol defines a simple ..."
---

# The front-end protocol

The front-end protocol (FEP) is an internal protocol used by the runtime system to
synchronize the abstract user interface (AUI) representation on the front-end side. This protocol
defines a simple set of operations to modify the AUI tree. This protocol is based on a command
processing principle (send command, receive answer) that can be serialized for transport over any
network protocol, like HTTP for example.

![Communication between the runtime system and the front-end diagram](../_images/DUIFig02.jpg)

*Typical communication between the Runtime System and the front-end*

1. Initialization phase: The runtime system sends the initial AUI tree.
2. The front-end builds the graphical user interface based on the AUI tree.
3. The front-end waits for a user interaction (mouse click, keyboard typing).
4. When the user performs some interaction, the front-end sends front-end events corresponding to
   the modifications made by the user.
5. Front-end events are analyzed and validated by the runtime system.
6. The runtime system sends back the result of the front-end requests, by way of AUI tree
   modification commands.
7. When receiving these commands, the front-end modifies its version of the AUI tree and updates
   the graphical user interface. It then waits for new user interactions (step 3).
