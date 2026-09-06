---
title: "HBOX container"
source: "fgl-topics/c_fgl_FormSpecFiles_HBOX_container.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > HBOX container"
type: "concept"
---

# HBOX container

> Packs child layout elements horizontally.

## Syntax

```
HBOX [identifier]  [ ( attribute [,...] ) ]
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
[navigationArrows](1640-hbox-style-attributes.md),
[navigationDots](1640-hbox-style-attributes.md),
[packed](1640-hbox-style-attributes.md).

## Can hold

`VBOX`, `HBOX`, `GROUP`, `FOLDER`, `GRID`, `SCROLLGRID`, `TABLE`, `TREE`.

## Usage

An `HBOX` container packs other form items together, in the horizontal
direction.

For more details about this item type, see [HBOX item type](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.").

## Example

```
HBOX
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

[Using hbox tags to align form items](1556-using-hbox-tags-to-align-form-items.md "The hbox tag concept has been introduced to bypass the limitations of the character-based grid in forms.")

[Automatic HBox/VBox with splitter](1554-automatic-hbox-vbox-with-splitter.md "Horizontal and vertical boxes with splitter are created automatically when stretchable elements are set side by side.")
