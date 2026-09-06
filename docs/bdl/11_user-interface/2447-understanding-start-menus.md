---
title: "Understanding start menus"
source: "fgl-topics/c_fgl_startmenus_005.html"
breadcrumb: "User interface > User interface programming > Start menus > Understanding start menus"
type: "concept"
---

# Understanding start menus

> This is an introduction to start menus.

The start menu defines a tree of commands that start programs on the application server where the
runtime system executes.

> **Important:**
>
> Startmenus are not supported with embedded mobile apps: Since startmenu command execute child
> processes, this feature cannot be supported on mobile when the app executes on the device.

![Form with startmenu screenshot](../_images/StartMenu01_gbc.jpg)

*StartMenu rendering in GBC SideBarDrawer*

It is recommended that you create a specific program dedicated to running the start menu. This
program must create (or load) a start menu, and then perform an interactive instruction to enter the
interaction loop.

The start menu must be defined in the abstract user interface tree under the
"`UserInterface`" root node.

The start menu is unique for a program and cannot be redefined.

When a start menu command is selected by the user, the runtime system automatically starts a
child process with the command specified in the command attribute.

## Related links

**Related concepts**  

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")
