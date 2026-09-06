---
title: "The dynamic user interface"
source: "fgl-topics/c_fgl_DynamicUI_006.html"
breadcrumb: "User interface > User interface basics > The dynamic user interface"
type: "concept"
---

# The dynamic user interface

> The dynamic user interface is the base concept of the Genero user interaction components.

The dynamic user interface (DUI) concept implements a flexible graphical user
interface programming toolkit, based on the usage of [XML](http://www.w3.org/XML/) standards, to define an abstract representation of the
application forms. This abstract user interface representation can be displayed by different types
of display devices called front-ends, which execute on the user workstation or on the same platform
as the runtime system.

The abstract definition of the user interface can be manipulated at runtime as a tree of
interface objects. This tree is called the [abstract user
interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.").

The runtime system is in charge of the abstract user interface tree, and the front-end is in
charge of rendering this abstract tree on the screen. The front-end gets a copy of that tree which
is automatically synchronized by the runtime system by using the front-end protocol.

In development, application screens are defined by [form
specification files](1537-form-definitions.md "This section describes how to define application forms and program resources related to the presentation layer.").

The following schema describes the dynamic user interface concept, showing how the
abstract user Interface tree is shared by the runtime system and the front-end.

![AUI tree shared between the Runtime System and front-end diagram](../_images/DUIFig01.jpg)

*AUI tree shared between the runtime system and front-end*

When the AUI trees are synchronized, only the changes are sent to or received from the
front-end.

Note that when running on a mobile device, both front-end and runtime system execute on the same
platform. Still the AUI tree protocol takes place, and both components perform the tasks they are
dedicated to.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
