---
title: "Binding scrollgrids to arrays in dialogs"
source: "fgl-topics/c_fgl_ui_scrollgrid_screen_array.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Binding scrollgrids to arrays in dialogs"
type: "concept"
---

# Binding scrollgrids to arrays in dialogs

> Program arrays act as data model that are bound to form scrollgrids, when implementing list dialogs.

## Identifying list views in program dialogs

In list dialogs such as the `INPUT ARRAY` or `DISPLAY ARRAY`, the
screen array identifies the record list element in the current form to be bound to the program array
used by the dialog.

In this example, the `INPUT ARRAY` uses the `custlist` screen array
of the form, and binds the `custarr` program `ARRAY` with the
code:

```
INPUT ARRAY custarr FROM custlist.*
```

The screen array members are associated with the program array record members by position. The
order and number of the screen array elements is important as they are bound by position to the
members of the program array.

The position of the `SCROLLGRID` fields, however, can differ from the members of
the screen array and program array.

To omit fields in the `SCROLLGRID` layout, yet include them in the definition of
the screen array, and define the fields as [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).") in the form definition file. The program array can then be
defined from the database table definition with the `DEFINE LIKE` instruction:

```
DEFINE custarr DYNAMIC ARRAY OF RECORD LIKE customer.*
```

The array is usually defined with a flat list of members with `ARRAY OF RECORD / END
RECORD`. However, the array can be structured with sub-records and still be used with a list
dialog. This is especially useful when you need to define arrays from database tables, and
additional information needs to be managed at runtime (for example to hold image resource for each
row, to be displayed with the `IMAGECOLUMN`
attribute):

```
SCHEMA shop
DEFINE a_items DYNAMIC ARRAY OF RECORD
                   item_data RECORD LIKE items.*,
                   it_image STRING,
                   it_count INTEGER
               END RECORD
...
DISPLAY ARRAY a_items TO sr.*
   ...
```

## Defining screen arrays

The `SCROLLGRID` container is bound to a screen array defined in the
`INSTRUCTIONS` section, by the name of the form fields used in the screen array
definition.

The form data type and additional properties are defined in the `ATTRIBUTES`
section as form fields:

```
LAYOUT
SCROLLGRID (WANTFIXEDPAGESIZE = NO)
{
 Id:   [f1    ]
 Name: [f2           ][f3           ]
}
END
END
...
ATTRIBUTES
EDIT f1 = customer.cust_num;
EDIT f2 = customer.cust_fname, 
EDIT f3 = customer.cust_lname;
...
```

Each form field of the scrollgrid must be grouped in the
`INSTRUCTIONS` section in a [`SCREEN RECORD`](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
definition.

```
INSTRUCTIONS
SCREEN RECORD custlist( cust_num, cust_fname, cust_lname );
END
```

## Related links

**Related concepts**  

[Variable binding in DISPLAY ARRAY](1964-variable-binding-in-display-array.md "Variable binding in DISPLAY ARRAY")

[Variable binding in INPUT ARRAY](2010-variable-binding-in-input-array.md "Variable binding in INPUT ARRAY")
