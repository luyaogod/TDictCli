---
title: "ui.Interface.setImage"
source: "fgl-topics/c_fgl_ClassInterface_setImage.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.setImage"
type: "concept"
---

# ui.Interface.setImage

> Defines the icon image of the program.

## Syntax

```
ui.Interface.setImage(
   image STRING )
```

1. image is the image filename to be used as program icon.

## Usage

Use the `ui.Interface.setImage()` class method to define the icon image for the
program to be used by the front-ends, to easily identify the program in the window
container.

Call the method at the beginning of the program, before any interactive
instruction.

## Related links

**Related concepts**  

[Containers for program windows](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.")

[ui.Interface.setName](3120-ui-interface-setname.md "Define the name of the current program for the front-end.")

[ui.Interface.setText](3122-ui-interface-settext.md "Defines the title for the program.")

[ui.Interface.getImage](3105-ui-interface-getimage.md "Returns the icon image of the program.")
