---
title: "VBOX item type"
source: "fgl-topics/c_fgl_FormSpecFiles_VBOX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > VBOX item type"
type: "concept"
---

# VBOX item type

> Defines a layout area to render child elements in vertical direction.

## VBOX item basics

The `VBOX` container automatically packs the contained elements vertically from
top to bottom.

Contained elements are packed in the order in which they appear in the `LAYOUT`
section of the form file.

No decoration (border) is added when you use a `VBOX` container.

By combining `VBOX` and `HBOX` containers, you can define any
alignment you choose.

## Defining an VBOX

A `VBOX` is defined in a `LAYOUT` tree, as [`VBOX` container](1718-vbox-container.md "Packs child layout elements vertically.") for other form
items. It can for example be combined with an [`HBOX` container](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.") to pack form elements to be displayed in vertical and
horizontal directions:

```
LAYOUT
  HBOX
    VBOX
      GROUP
         ...
      END
      TABLE
         ...
      END
    END
    VBOX
      ...
    END
  END
END
```

`HBOX` and `VBOX` containers can be combined to
implement responsive layout, by using the [`ORIENTATION@screen-size`](1546-hbox-vbox-orientation.md "HBOX/VBOX containers can be defined to adapt to the screen size.") attribute.

When the `HBOX` or `VBOX` has an horizontal orientation, it can be
rendered as a splitview by setting the [`SPLIT@screen-size`](1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.") attribute.

`HBOX` and `VBOX` with horizontal orientation and
`SPLIT` attribute can be decorated with nativation arrows and/or navigation dots.
These can respectively be controlled with the [`navigationArrows`](1640-hbox-style-attributes.md) and [`navigationDots`](1640-hbox-style-attributes.md) style attributes.

By default, when there is no stretchable child element such as a `TABLE`, space is
inserted between any two elements. If needed, you can set the [`packed`](1640-hbox-style-attributes.md) style attribute to `"yes"`, in order to pack the
elements to the left or top of the box container, depending on the orientation. For a vertical box,
elements are packed to the top. For an horizontal box, elements are packed to the left.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items."), [HBox style attributes](1640-hbox-style-attributes.md "HBox presentation style attributes apply to an HBox element.") and [VBox style attributes](1653-vbox-style-attributes.md "VBox presentation style attributes apply to an VBox element.").
