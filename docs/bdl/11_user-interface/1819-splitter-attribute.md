---
title: "SPLITTER attribute"
source: "fgl-topics/c_fgl_FSFAttributes_SPLITTER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > SPLITTER attribute"
type: "concept"
---

# SPLITTER attribute

> The SPLITTER attribute forces the container to use a splitter widget between each child element.

## Syntax

```
SPLITTER
```

## Usage

This attribute indicates that the container (typically, a `VBOX` or
`HBOX`) must have a splitter between each child element held by the container.

If a container is defined with a splitter and if the children are stretchable (like
`TABLE` or `TEXTEDIT`), users can resize the child elements inside the
container.

## Example

```
VBOX ( SPLITTER )
```

## Related links

**Related concepts**  

[HBOX item type](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.")

[VBOX item type](1707-vbox-item-type.md "Defines a layout area to render child elements in vertical direction.")
