---
title: "ui.Interface.setText"
source: "fgl-topics/c_fgl_ClassInterface_setText.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setText"
type: "concept"
---

# ui.Interface.setText

> Defines the title for the program.

## Syntax

```
ui.Interface.setText(
   title STRING )
```

1. title is the text to be used as program title.

## Usage

Use the `ui.Interface.setText()` class method to define the title for the program
to be used by the front-ends, to easily identify the program in the window
container.

Call the method at the beginning of the program, before any
interactive instruction.

## Related links

**Related concepts**  

[Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.")

[ui.Interface.setName](3120-ui-interface-setname.md "Define the name of the current program for the front-end.")

[ui.Interface.setImage](3119-ui-interface-setimage.md "Defines the icon image of the program.")

[ui.Interface.getText](3108-ui-interface-gettext.md "Returns the title of the program.")
