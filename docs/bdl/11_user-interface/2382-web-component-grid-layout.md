---
title: "Web component grid layout"
source: "fgl-topics/c_fgl_webcomponent_layout_grid.html"
breadcrumb: "User interface > User interface programming > Web components > Controlling the web component layout > Web component grid layout"
type: "concept"
description: "The item tag of the WEBCOMPONENT defines the default dimensions of the web component area: LAYOUT GRID { <GROUP g1 > [f1 ][f2 ] [f3 ] ... < > [f5 ] [ ] [ ] [ ] } END In the ATTRIBUTES section, use the ..."
---

# Web component grid layout

The item tag of the `WEBCOMPONENT` defines the default dimensions of the web
component area:

```
LAYOUT
GRID
{
<GROUP g1                 >
[f1           ][f2        ]
[f3                       ]
 ...
<                         >
[f5                       ]
[                         ]
[                         ]
[                         ]
}
END
```

In the `ATTRIBUTES` section, use the [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") and [`HEIGHT`](1782-height-attribute.md "The HEIGHT attribute forces an explicit height for a form element.")
attributes, to define the sizing policy of a web component field:

By default, the `WEBCOMPONENT` widget gets the size of the form item. This is the
behavior of `SIZEPOLICY=FIXED`, and is the only option. This defines the minimum size
of the web component when it is stretchable.

`SIZEPOLICY=DYNAMIC` or `SIZEPOLICY=INITIAL` are not supported with
`WEBCOMPONENT` fields.

By default, `WEBCOMPONENT` widget are stretchable in both directions
(`STRETCH=BOTH`). This default behavior can be changed by using another value for the
`STRETCH` attribute. In the next example, the web component will get its initial size
from the form item size, and will keep this size even if the window is resized because
`STRETCH=NONE`:

```
WEBCOMPONENT f5 = FORMONLY.mymap,
   SIZEPOLICY = FIXED,
   STRETCH = NONE;
```

## Related links

**Related concepts**  

[WEBCOMPONENT item type](1708-webcomponent-item-type.md "Defines a specialized form item that holds an external component.")
