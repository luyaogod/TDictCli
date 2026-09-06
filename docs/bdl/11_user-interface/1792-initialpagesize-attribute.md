---
title: "INITIALPAGESIZE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_INITIALPAGESIZE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > INITIALPAGESIZE attribute"
type: "concept"
---

# INITIALPAGESIZE attribute

> The INITIALPAGESIZE attribute defines the initial page size of a list element.

## Syntax

```
INITIALPAGESIZE = lines
```

1. lines is an integer that defines the initial page size for the list
   element.

## Usage

Some list containers such as resizable `SCROLLGRID` containers display by default
a single row, if no other elements in the form layout forces the container size implicitly.

In such case, the `INITIALPAGESIZE` attribute can be used to specify a default
initial number of rows to be displayed.

## Related links

**Related concepts**  

[WANTFIXEDPAGESIZE attribute](1847-wantfixedpagesize-attribute.md "The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.")

[SCROLLGRID item type](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.")
