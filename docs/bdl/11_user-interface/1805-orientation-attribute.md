---
title: "ORIENTATION attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ORIENTATION.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ORIENTATION attribute"
type: "concept"
---

# ORIENTATION attribute

> The ORIENTATION attribute defines whether an element displays vertically or horizontally.

## Syntax

```
ORIENTATION[@screen-size] = {VERTICAL|HORIZONTAL}
```

where screen-size can
be:

```
{ SMALL | MEDIUM | LARGE }
```

1. screen-size is a screen size selector that indicates how the element is
   oriented, vertically or horizontally, depending on the size of the screen. Several
   `ORIENTATION@screen-size` attributes can be used for the same
   element.
2. With `HBOX` / `VBOX` elements, the
   `@screen-size` selector is mandatory.

## Usage

The `ORIENTATION` attribute defines how a form element must be oriented, when the
element can be rendered in both horizontal and vertical directions.

The `ORIENTATION` attribute can be used in the definition of a
`RADIOGROUP` form item, to specify that radio items must be displayed horizontally or
vertically. Default is vertical arrangement. With `RADIOGROUP`, you can also use the
`@screen-size` modifier, to adapt the orientation of the radio
group items to the screen size:

```
RADIOGROUP f1 = customer.status, ITEMS=(...),
    ORIENTATION=HORIZONTAL;
RADIOGROUP f2 = order.options, ITEMS=(...),
    ORIENTATION@LARGE=HORIZONTAL;
```

A `VBOX` or `HBOX` container can be defined with the
`ORIENTATION` attribute, in conjunction with a
`@screen-size` modifier, to get the appropriate rendering,
depending on the screen size:

```
VBOX vb1 ( ORIENTATION@LARGE = HORIZONTAL )
```

- A `VBOX` renders by default vertically, and the orientation can be changed to
  horizontal.
- An `HBOX` renders by default horizontally, and its orientation can be changed to
  vertical.

## Example

```
VBOX vb1 ( ORIENTATION@LARGE = HORIZONTAL )
GRID grid1
...
TABLE table1
...
END
```

## Related links

**Related concepts**  

[HBOX/VBOX orientation](1546-hbox-vbox-orientation.md "HBOX/VBOX containers can be defined to adapt to the screen size.")

[RADIOGROUP item type](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field.")

[HBOX item type](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.")

[VBOX item type](1707-vbox-item-type.md "Defines a layout area to render child elements in vertical direction.")
