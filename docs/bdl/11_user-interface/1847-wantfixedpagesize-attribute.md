---
title: "WANTFIXEDPAGESIZE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WANTFIXEDPAGESIZE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WANTFIXEDPAGESIZE attribute"
type: "concept"
---

# WANTFIXEDPAGESIZE attribute

> The WANTFIXEDPAGESIZE attribute controls the vertical resizing of a list element.

## Syntax

```
WANTFIXEDPAGESIZE [ = NO ]
```

## Usage

The `WANTFIXEDPAGESIZE` attribute can be used for `SCROLLGRID`
containers to control the vertical resizing of the list element.

For backward compatibility, `WANTFIXEDPAGESIZE` can be used for
`TABLE`/`TREE` containers. In
`TABLE`/`TREE` definitions, use [`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") instead of
`WANTFIXEDPAGESIZE`.

By default, a `SCROLLGRID` container is not resizable in height. The number of
visible scrollgrid rows is defined by the form file. To allow the scrollgrid to stretch vertically,
use the attribute `WANTFIXEDPAGESIZE=NO`.

When using a stretchable `SCROLLGRID` with `WANTFIXEDPAGESIZE=NO`,
one can define the rendering with the [`customWidget`](1646-scrollgrid-style-attributes.md) style attribute.

## Related links

**Related concepts**  

[INITIALPAGESIZE attribute](1792-initialpagesize-attribute.md "The INITIALPAGESIZE attribute defines the initial page size of a list element.")

[SCROLLGRID container](1723-scrollgrid-container.md "Defines a scrollable grid view widget.")
