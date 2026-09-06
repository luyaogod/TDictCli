---
title: "Widget width inside hbox tags"
source: "fgl-topics/c_fgl_form_rendering_hbox_tags_size.html"
breadcrumb: "User interface > Form definitions > Form rendering > Grid-based layout > Using hbox tags to align form items > Widget width inside hbox tags"
type: "concept"
description: "The form item width is used to define the widget size, and the maximum input length, according to the data type of the variable bound to the field. For details about maximum input length, see Input ..."
---

# Widget width inside hbox tags

The form item width is used to define the widget size, and the maximum input length, according to
the data type of the variable bound to the field. For details about maximum input length, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

By default, the form item width of [`BUTTONEDIT`](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action."), [`DATEEDIT`](1688-dateedit-item-type.md "Defines a line-edit with a calendar widget to pick a date.") and [`COMBOBOX`](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.") widgets are computed by fglform as follows:

```
if item-tag-width > 2
   width = item-tag-width - 2
else
   width = item-tag-width
```

Where item-tag-width represents the number of cells used in the form layout by
the item tag.

For example, with the following form definition, all form item tags have a width or 10
cells:

```
LAYOUT
GRID
{
 1234567890
[f1        ]
[f2        ]
[f3        ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.field1;
BUTTONEDIT f2 = FORMONLY.field2;
DATEEDIT f3 = FORMONLY.field3;
END
```

The resulting .42f file will define the following `width` and
`gridWidth` attributes:

```
...
      <Edit width="10" ... gridWidth="10" ... />
...
      <ButtonEdit width="8" ... gridWidth="10" ... />
...
      <DateEdit width="8" ... gridWidth="10" ... />
...
```

At runtime, the maximum input length will depend on the data type of the associated variable,
based on the `width` attribute.

If needed, it is possible to specify the width of a field element within an hbox tag, by using
the `-` (hyphen) marker. When using the hyphen marker, the
`item-tag-width-2` rule does not apply.

The hbox tag is created when using the `:` (colon) and the `-`
(hyphen) marker is to be placed after the last cell that will define the actual field width.

For example:

```
LAYOUT
GRID
{
      EDIT [f1     ]  width = 7
BUTTONEDIT [f2     ]  width = 5 (7-2)
BUTTONEDIT [f3:    ]  width = 2
BUTTONEDIT [f4   : ]  width = 3 (5-2)
BUTTONEDIT [f5  -: ]  width = 4
}
END
END
ATTRIBUTES
      EDIT f1 = FORMONLY.field1;
BUTTONEDIT f2 = FORMONLY.field2;
BUTTONEDIT f3 = FORMONLY.field3;
BUTTONEDIT f4 = FORMONLY.field4;
BUTTONEDIT f5 = FORMONLY.field5;
END
```

In the above form definition:

- The `f1` item tag occupies 7 grid columns and gets a width of 7.
- The `f2` item tag occupies 7 grid columns and gets a width of 5 (7-2), because
  it's a `BUTTONEDIT`, to save space for the button.
- The `f3` item tag occupies 7 grid columns, defines and hbox because a
  `:` colon is used, and gets a width of 2, based on the position of the colon inside
  the square breaces.
- The `f4` item tag occupies 7 grid columns, defines and hbox with
  `:`, and gets a width of 3, because the `:` colon limits the element
  width to 5 cells inside the hbox. Since it's a `BUTTONEDIT`, 5-2 = 3.
- The `f5` item tag occupies 7 grid columns, defines and hbox with
  `:`, and gets a width of 4, because the `-` hyphen defines the exact
  width for the field. When using the hyphen marker, the -2 rule does not apply.

Visual result:

![Form using field width specifier screenshot](../_images/layout72_gbc.jpg)

*Form using field width specifier*

## Related links

**Related concepts**  

[Hbox tags](1680-hbox-tags.md "Hbox tags group several item tags within the same horizontal layout box, inside a grid-based container (GRID).")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
