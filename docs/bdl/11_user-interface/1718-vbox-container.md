---
title: "VBOX container"
source: "fgl-topics/c_fgl_FormSpecFiles_VBOX_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > VBOX container"
type: "concept"
---

# VBOX container

> Packs child layout elements vertically.

## Syntax

```
VBOX [identifier]  [ ( attribute [,...] ) ]
  layout-container
  [...] 
END
```

1. identifier defines the name of the element, it is optional but
   recommended.
2. attribute is an attribute for the element.
3. layout-container is another child container.

## Form attributes

`COMMENT`, `FONTPITCH`, `HIDDEN`, `NOSWIPE`, `ORIENTATION`, `STYLE`, `SPLIT` , `SPLITTER`, `TAG`.

## Style attributes

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

Class-specific:
[navigationArrows](1653-vbox-style-attributes.md),
[navigationDots](1653-vbox-style-attributes.md),
[packed](1653-vbox-style-attributes.md).

## Can hold

`VBOX`, `HBOX`, `GROUP`, `FOLDER`, `GRID`, `SCROLLGRID`, `TABLE`, `TREE`.

## Usage

A `VBOX` container packs other form items together, in the vertical direction.

For more details about this item type, see [VBOX item type](1707-vbox-item-type.md "Defines a layout area to render child elements in vertical direction.").

## Example

```
VBOX
  GROUP ( TEXT = "Customer" )
   ...
  END
  TABLE
   ...
  END
END
```

## Related links

**Related concepts**  

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[Layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.")

[Automatic HBox/VBox with splitter](1554-automatic-hbox-vbox-with-splitter.md "Horizontal and vertical boxes with splitter are created automatically when stretchable elements are set side by side.")
