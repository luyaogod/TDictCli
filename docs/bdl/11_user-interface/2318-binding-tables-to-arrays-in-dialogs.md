---
title: "Binding tables to arrays in dialogs"
source: "fgl-topics/c_fgl_ui_tables_screen_array.html"
breadcrumb: "User interface > User interface programming > Table views > Binding tables to arrays in dialogs"
type: "concept"
---

# Binding tables to arrays in dialogs

> Program arrays act as data model that are bound to form tables, when implementing list dialogs.

## Identifying list views in program dialogs

In list dialogs such as the `INPUT ARRAY` or `DISPLAY ARRAY`, the
screen array identifies the record list element in the current form to be bound to the program array
used by the dialog.

In this example, the `INPUT ARRAY` uses the `custlist` screen array
of the form, and binds the `custarr` program `ARRAY`
with:

```
INPUT ARRAY custarr FROM custlist.*
```

The screen array members are associated to the program array record members by position. The
order and number of the screen array elements matter as they are bound by position to the members
the program array. The position of the `TABLE` columns, however, can differ from the
members of the screen array and program array.

To omit columns in the `TABLE` layout, yet include them in the definition of the
screen array, and define the columns as [`PHANTOM` fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).") in the form definition file. The program array can then be
defined from the database table definition with the `DEFINE LIKE` instruction:

```
DEFINE custarr DYNAMIC ARRAY OF RECORD LIKE customer.*
```

Note that the array is usually defined with a flat list of members with `ARRAY OF RECORD /
END RECORD`. However, the array can be structured with sub-records and still be used with a
list dialog. This is especially useful when you need to define arrays from database tables, and
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

## Defining screen arrays in TABLES

The `TABLE` container is bound to a screen array defined in the
`INSTRUCTION` section, by the name of the form fields used in the screen array
definition.

The column data type and additional column properties are defined in the `ATTRIBUTES`
section as form fields:

```
LAYOUT
...
TABLE
{
[c1    |c2     |c3              ]
[c1    |c2     |c3              ]
[c1    |c2     |c3              ]
[c1    |c2     |c3              ]
}
END
...

ATTRIBUTES
EDIT c1 = customer.cust_num;
EDIT c2 = customer.cust_name, 
EDIT c3 = customer.cust_cdate;
...
```

Each form field of the table must be grouped in the `INSTRUCTIONS` section in a
[`SCREEN RECORD`](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
definition.

```
SCREEN RECORD custlist( cust_num, cust_name, cust_cdate );
```

## Related links

**Related concepts**  

[Array binding in list controllers](2301-array-binding-in-list-controllers.md "Program array elements are bound to screen arrays elements in the definition of the DISPLAY ARRAY or INPUT ARRAY list dialog.")
