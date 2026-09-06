---
title: "Layout items"
source: "fgl-topics/c_fgl_FormSpecFiles_Layout_Items.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Layout items"
type: "concept"
---

# Layout items

> Layout items are containers with a body that can hold other form items.

Layout items can be specified as a tree of nested containers, or as layout tags within a single
[`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.")
container.

The example shows a tree of nested containers, where a `GRID` and
`TABLE` are included in a
`VBOX`:

```
LAYOUT
VBOX
 GRID ...
 {
 }
 END
 TABLE ...
 {
 }
 END
END
```

The example shows a `GRID` container including [layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container."). The layout tags group form fields in
dedicated areas. This syntax is usually more convenient to describe application forms:

```
LAYOUT
GRID
{
<g g1                     >
 Name: [f01              ]
<                         >
<t t1                     >
[c1  |c2                  ]
<                         >
}
END
END
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
