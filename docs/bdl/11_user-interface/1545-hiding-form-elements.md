---
title: "Hiding form elements"
source: "fgl-topics/c_fgl_form_responsive_hidden.html"
breadcrumb: "User interface > Form definitions > Form rendering > Responsive Layout > Hiding form elements"
type: "concept"
---

# Hiding form elements

> Forms elements can be automatically hidden depending on the screen size.

## Form element visibility with responsive layout

By hiding some form elements, applications can adapt to the desktop or web browser window size
and to mobile device screen size and orientation, to provide an optimal user experience.

Form elements can be hidden automatically for a given screen size, depending on an abstract size
specification, with the [`HIDDEN@screen-size`](1543-layout-structure-for-responsive.md) attribute syntax.

> **Important:**
>
> Make sure that elements using the
> `HIDDEN@screen-size` attribute do not hide fields used for input
> by dialog instructions: The current field that has the focus must always be visible.

In the next example, the `customer.cust_address` field will be hidden for small
screens, while `customer.cust_company` will be hidden for small and medium sized
screens, while `customer.cust_id` and `customer.cust_name` fields are
the most important data, and must always be visible:

```
...
ATTRIBUTES
EDIT f01 = customer.cust_id;
EDIT f02 = customer.cust_name;
EDIT f03 = customer.cust_address, HIDDEN@SMALL;
EDIT f04 = customer.cust_company, HIDDEN@SMALL, HIDDEN@MEDIUM;
...
END
```

The above example assumes that there are no other form elements related to the address and
company fields. If labels are used for these fields, they should use the same
`HIDDEN@screen-size` attributes.

Containers such as `GRID`, `TABLE`, `TREE`,
`SCROLLGRID`, as well as `HBOX`, `VBOX`,
`FOLDER` and `GROUP` can be configured with the
`HIDDEN@screen-size` attribute, to hide all the content for a
given screen size:

```
LAYOUT
VBOX
  GRID ( HIDDEN@SMALL )
      ...
  END
...
END
...
```

## Hiding form elements by program

Form elements can be hidden or shown by program, depending on application rules. For example, a
given application user may not have permissions to see some data. This can be implemented with the
[`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.") and [`ui.Form.setFieldHidden()`](../15_library-reference/3161-ui-form-setfieldhidden.md "Show or hide a form field.")
methods.

Hiding form elements by program can be used in conjunction with [`HIDDEN@screen-size`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.")
attributes, to get responsive layout:

- When hiding an element by program (value=1 or 2), it will always be hidden to the user, no
  matter the `HIDDEN@screen-size` definitions.
- When showing an element by program (value=0), it will be visible or hidden, depending on the
  `HIDDEN@screen-size` definitions.

## Example using `HIDDEN@screen-size`

The following form definition file uses the `HIDDEN` attribute with screen-size
specifiers, to indicated what elements must disappear, depending on the screen or container window
size. Note that `field11` gets a [`STRETCH=X`](1544-horizontal-stretching.md "Define stretchable form elements to achieve responsive layout.") attribute, to occupy all the space when the second grid is
hidden.

```
LAYOUT (STYLE="main2")
HBOX
GRID grid1
{
[l11       |f11         ]
[l12       |f12         ]
}
END
GRID grid2 (HIDDEN@SMALL,HIDDEN@MEDIUM)
{
[l21       |f21         ]
[l22       |f22         ]
}
END
END
END
ATTRIBUTES
LABEL l11: label11, TEXT="Field11";
EDIT f11 = FORMONLY.field11, STRETCH=X;
LABEL l12: label12, TEXT="Field12", HIDDEN@SMALL;
DATEEDIT f12 = FORMONLY.field12, HIDDEN@SMALL;
LABEL l21: label21, TEXT="Field21";
EDIT f21 = FORMONLY.field21;
LABEL l22: label22, TEXT="Field22";
DATEEDIT f22 = FORMONLY.field22;
END
```

Rendering on a small screen: The grid2 and label12+field12 elements are hidden:

![Rendering on a small screen](../_images/rl_hidden_s1_small.jpg)

*Rendering on a small screen*

Rendering on a medium screen: The grid2 element is hidden:

![Rendering on a medium screen](../_images/rl_hidden_s1_medium.jpg)

*Rendering on a medium screen*

Rendering on a large screen: All elements are visible:

![Rendering on a large screen](../_images/rl_hidden_s1_large.jpg)

*Rendering on a large screen*
