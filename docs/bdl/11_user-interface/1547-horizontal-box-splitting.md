---
title: "Horizontal box splitting"
source: "fgl-topics/c_fgl_form_responsive_hbox_split.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Horizontal box splitting"
type: "concept"
---

# Horizontal box splitting

> HBOX/VBOX containers can be defined to display a single child container.

## HBOX/VBOX single content display

By default, `HBOX` and `VBOX` elements show all child containers.
When the `HBOX` or `VBOX` is oriented horizontally (see [HBOX/VBOX orientation](1546-hbox-vbox-orientation.md "HBOX/VBOX containers can be defined to adapt to the screen size.")), you can force the box to display a single child
container at a time, by using the `SPLIT[@screen-size]`
attribute.

> **Note:**
>
> The `SPLIT` attribute takes only effect when the `HBOX` or
> `VBOX` is oriented horizontally.

The `SPLIT` attribute can be specified with
`@screen-size` modifier, to get the required rendering, depending
on the size of the screen. For example, on mobile devices with small screens, there is usually only
room to display a single child container.

When `SPLIT` is in action with desktop / web browser front end, is it possible to
move between child containers by using a dedicated widget. The type of widget can be controlled with
the [`navigationArrows`](1640-hbox-style-attributes.md) / [`navigationDots`](1640-hbox-style-attributes.md) style attributes.

On mobile devices, the end user can use the swipe gesture to switch between the child containers.
To deny the swiping (and thus force the first child container to be the only visible element), use
the [`NOSWIPE`](1802-noswipe-attribute.md "The NOSWIPE attribute denies swipe gestures on the element.") attribute for the
`HBOX` or `VBOX` element.

For the complete syntax, see [SPLIT attribute](1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.").

## Form example defining a swipable splitview

Form definition file, note that:

- The `HBOX` element gets a `SPLIT` attribute that should only have
  an effect on small screens.
- The `TABLE` uses `STRETCHCOLUMNS` to occupy all the space
  available for that container.
- The field12 `EDIT` is defined with `STRETCH=X` to stretch all
  fields of the parent `GRID`.

```
LAYOUT ( STYLE="main2" )
HBOX ( SPLIT@SMALL )
GRID grid1
{
Field11: [f11         ]
Field12: [f12     ]
Field13: [f13         ]
}
END
TABLE table1 ( STRETCHCOLUMNS )
{
 Col21  Col22     Col23
[c21   |c22      |c23       ]
[c21   |c22      |c23       ]
[c21   |c22      |c23       ]
}
END
END -- HBOX
END

ATTRIBUTES
EDIT f11 = FORMONLY.field11, STRETCH=X;
COMBOBOX f12 = FORMONLY.field12, ITEMS=("aaa","bbb");
DATEEDIT f13 = FORMONLY.field13;
EDIT c21 = FORMONLY.col21;
EDIT c22 = FORMONLY.col22;
DATEEDIT c23 = FORMONLY.col23;
END

INSTRUCTIONS
SCREEN RECORD sr1(FORMONLY.col21 THRU FORMONLY.col23);
END
```

Visual result with a small screen size (`SPLIT` enters in action):

![Screenshot of form showing only grid container](../_images/rl_hbox_split_1.jpg)

*Form with horizontal box in small window: GRID container*

![Screenshot of form showing only table container](../_images/rl_hbox_split_2.jpg)

*Form with horizontal box in small window: TABLE container*

Visual result with a medium screen size (`SPLIT` is ignored):

![Screenshot of form with hbox showing grid and table side by side](../_images/rl_hbox_split_3.jpg)

*Form with horizontal box in medium window*
