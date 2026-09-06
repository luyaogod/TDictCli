---
title: "ui.Interface.setContainer"
source: "fgl-topics/c_fgl_ClassInterface_setContainer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setContainer"
type: "concept"
---

# ui.Interface.setContainer

> Define the parent container for the current program.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

## Syntax

```
ui.Interface.setContainer(
   name STRING )
```

1. name is the name of the parent container.

## Usage

The `ui.Interface.setContainer(name)` class method is used to specify the name of
the parent program for the current program. This creates a parent/child relation between two
independent programs running with distinct fglrun processes.

The type of the program must be set to "child" with the [`ui.Interface.setType()`](3123-ui-interface-settype.md "Defines the type of the program for the front-end.") method.

Each program must be identified by a name, to be set with the [`ui.Interface.setName()`](3120-ui-interface-setname.md "Define the name of the current program for the front-end.") class
method.
