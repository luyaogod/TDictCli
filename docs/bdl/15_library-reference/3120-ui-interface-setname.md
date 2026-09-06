---
title: "ui.Interface.setName"
source: "fgl-topics/c_fgl_ClassInterface_setName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setName"
type: "concept"
---

# ui.Interface.setName

> Define the name of the current program for the front-end.

## Syntax

```
ui.Interface.setName(
   name STRING )
```

1. name is the identifier of the program.

## Usage

Use the `ui.Interface.setName()` class method to define the identifier for the
program to be used by the front-ends.

The name passed to this method will be passed to the front-end in order to identify the
program.

Call the method at the beginning of the program, before any interactive instruction.

By default, it is the program name (without .42m or
.42r extension).

## Related links

**Related concepts**  

[ui.Interface.setText](3122-ui-interface-settext.md "Defines the title for the program.")

[ui.Interface.setImage](3119-ui-interface-setimage.md "Defines the icon image of the program.")
