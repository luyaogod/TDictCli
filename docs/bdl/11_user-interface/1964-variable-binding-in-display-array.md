---
title: "Variable binding in DISPLAY ARRAY"
source: "fgl-topics/c_fgl_DisplayArray_015.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > Variable binding in DISPLAY ARRAY"
type: "concept"
description: "The DISPLAY ARRAY statement binds the members of the array of record to the screen array fields specified with the TO keyword. Array members and screen array fields are bound by position (not by ..."
---

# Variable binding in DISPLAY ARRAY

The `DISPLAY ARRAY` statement binds the members of the [array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") of record to the [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") fields specified with the
`TO` keyword. Array members and screen array fields are bound by position (not
by name). The number of members in the program array must match the number of fields in the screen
record (that is, in a single row of the screen
array).

```
SCHEMA stock
DEFINE cust_arr DYNAMIC ARRAY OF customer.*
...
DISPLAY ARRAY cust_arr TO sr.*
        ATTRIBUTES(UNBUFFERED)
   ...
END DISPLAY
```

The array is usually defined with a flat list of members with `ARRAY OF RECORD /
END RECORD`. However, the array can be structured with sub-records and still be used with a
`DISPLAY ARRAY` dialog. This is especially useful when you need to define arrays from
[database tables](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."), and additional information needs
to be managed at runtime (for example to hold image resource for each row, to be displayed with the
[`IMAGECOLUMN`](2320-displaying-column-images.md "You can use PHANTOM fields and the IMAGECOLUMN attribute to display images in a column, to the left of the column value.")
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

When using the `UNBUFFERED` attribute, the instruction is sensitive to program
variable changes. This means that you do not have to `DISPLAY` the
values; setting the program variable used by the dialog automatically displays the data
in the corresponding form
field.

```
ON ACTION change 
   LET arr[arr_curr()].field1 = newValue()
```

If the program array has the same structure as a database table (this is the case when the array
is defined with a `DEFINE LIKE` clause), you may not want to display/use
some of the columns. You can achieve this by using `PHANTOM` fields in
the screen array definition. Phantom fields will only be used to bind program variables,
and will not be transmitted to the front-end for display.

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
