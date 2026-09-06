---
title: "SPLIT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SPLIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SPLIT attribute"
type: "concept"
---

# SPLIT attribute

> The SPLIT attribute forces a horizontal box to show only one child container.

## Syntax

```
SPLIT[@screen-size]
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates when the box is to be
   split, depending on the size of the screen. Several
   `SPLIT@screen-size` attributes can be used for the same
   element.

## Usage

When the `SPLIT` attribute takes effect, the `HBOX` or
`VBOX` displays only one child container at a time. The child containers can for
example be a `GRID`, a `TABLE` or a `SCROLLGRID`.

The `SPLIT` attribute only takes effect when the [orientation](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally.") of the box is horizontal. This
applies to:

- `HBOX` elements without an `ORIENTATION` attribute.
- `HBOX` or `VBOX` elements with
  `ORIENTATION=HORIZONTAL`
- `HBOX` or `VBOX` elements with
  `ORIENTATION@screen-size=HORIZONTAL` and the
  screen-size applies.

On mobile devices, when the box is split, the end user can use the swipe gesture to move between
child containers. To deny the swipe gesture, use the [`NOSWIPE`](1802-noswipe-attribute.md "The NOSWIPE attribute denies swipe gestures on the element.") attribute.

When `SPLIT` is in action with desktop / web browser front ends, is it possible to
move between child containers by using a specific widget. The type of widget can be controlled with
the [`navigationArrows`](1640-hbox-style-attributes.md) / [`navigationDots`](1640-hbox-style-attributes.md) style attributes.

The `SPLIT` and [`SPLITTER`](1819-splitter-attribute.md "The SPLITTER attribute forces the container to use a splitter widget between each child element.") attributes are similar concepts, but are not to be used together:
`SPLITTER` makes only sense when several child containers are displayed at the same
time.

## Example

```
HBOX hb1 ( SPLIT@SMALL )
GRID grid1
...
TABLE table1
...
END
```

## Related links

**Related concepts**  

[Horizontal box splitting](1547-horizontal-box-splitting.md "HBOX/VBOX containers can be defined to display a single child container.")

[HBOX item type](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.")

[VBOX item type](1707-vbox-item-type.md "Defines a layout area to render child elements in vertical direction.")
