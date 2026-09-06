---
title: "ui.Interface.setType"
source: "fgl-topics/c_fgl_ClassInterface_setType.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setType"
type: "concept"
---

# ui.Interface.setType

> Defines the type of the program for the front-end.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

## Syntax

```
ui.Interface.setType(
   type STRING )
```

1. type is the identifier of the program. Possible values are:
   `normal`, `container`, `child`.

## Usage

Use the `ui.Interface.setType()` class method to define the type for the program
to be used by the front-ends, to define parent/child relation between programs displayed on the same
front-end.

Possible values are:

- `normal`: Defines a regular program, not attached to any container.
- `container`: Defines a container program that can have children.
- `child`: Defines a child program, to be combined with [`ui.Interface.setContainer()`](3118-ui-interface-setcontainer.md "Define the parent container for the current program.").

The type passed to this method will be passed to the front-end in order to define the
parent/child relationship of programs.

Call the method at the beginning of the program, before any interactive instruction.

## Related links

**Related concepts**  

[ui.Interface.setText](3122-ui-interface-settext.md "Defines the title for the program.")

[ui.Interface.setImage](3119-ui-interface-setimage.md "Defines the icon image of the program.")
