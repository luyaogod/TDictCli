---
title: "The user interface"
source: "fgl-topics/c_fgl_intro_BDL_017.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > The user interface"
type: "concept"
---

# The user interface

> The Genero user interface technology is based on the sharing of an abstract representation between the runtime system and the front-end.

When a program starts, the runtime system creates the abstract user interface
(AUI) tree and passes this tree to the front-end. The front-end renders the abstract element
as real graphical objects on the workstation.

When an interaction statement takes control of the application, the tree on the front-end
is automatically synchronized with the runtime system tree. Runtime system and front-ends
communicate with the front end protocol, through the computer network. The AUI
tree and the protocol are using XML standards.

![AUI tree synchronization diagram](../_images/DUIFig01.jpg)

*AUI tree synchronization*

Resource files describe the appearance (decoration) of some of the graphic objects. Default
resource files (`default.4ad`, `default.4st`) are provided and
can be customized, or replaced with your own versions.

The elements of the AUI tree can be manipulated at runtime with built-in utilities.

## Related links

**Related concepts**  

[User interface](../11_user-interface/1507-user-interface.md "These topics cover programming the user interface (UI) with the Genero Business Development Language.")
