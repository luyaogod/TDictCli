---
title: "ui.Window.setImage"
source: "fgl-topics/c_fgl_ClassWindow_setImage.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.setImage"
type: "concept"
---

# ui.Window.setImage

> Set the window icon.

## Syntax

```
setImage(
   image STRING )
```

1. image is the image name for the icon of the window.

## Usage

The `setImage()` method specifies the icon of the window.

By default, the icon of a window is defined by the
[`IMAGE`](../11_user-interface/1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.") attribute of
the [`LAYOUT`](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
definition in form files.
