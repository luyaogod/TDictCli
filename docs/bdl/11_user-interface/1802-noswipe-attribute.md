---
title: "NOSWIPE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_NOSWIPE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > NOSWIPE attribute"
type: "concept"
---

# NOSWIPE attribute

> The NOSWIPE attribute denies swipe gestures on the element.

## Syntax

```
NOSWIPE
```

## Usage

The `NOSWIPE` attribute disables swipe gestures on touchscreens, for the element
where the attribute is used.

This attribute can be used for:

- `HBOX` and `VBOX` using the [`SPLIT`](1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.") attribute, to disable swipe
  gestures and force a given child element to remain on the screen.
- `FOLDER` containers, to disable the swipe gesture when folder page content also
  allows swipe gestures (a web component with a map for example), or when using nested folders.

## Example

```
HBOX hb1 ( SPLIT, NOSWIPE )
GRID grid1
...
TABLE table1
...
END
```

## Related links

**Related concepts**  

[Horizontal box splitting](1547-horizontal-box-splitting.md "HBOX/VBOX containers can be defined to display a single child container.")

[HBOX container](1717-hbox-container.md "Packs child layout elements horizontally.")

[VBOX container](1718-vbox-container.md "Packs child layout elements vertically.")

[FOLDER container](1720-folder-container.md "Defines the parent container for folder pages.")
