---
title: "Array binding in list controllers"
source: "fgl-topics/c_fgl_prog_dialogs_array_binding.html"
breadcrumb: "User interface > User interface programming > List dialogs > Array binding in list controllers"
type: "concept"
---

# Array binding in list controllers

> Program array elements are bound to screen arrays elements in the definition of the DISPLAY ARRAY or INPUT ARRAY list dialog.

## Array elements are bound to screen array fields by position

A [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") groups a
set of fields to define list container columns in a .per form file.

When using a [program array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") in `DISPLAY
ARRAY` or `INPUT ARRAY`, the elements of the array are bound by position to
the fields of the associated screen array.

In the form
file:

```
...
INSTRUCTIONS
SCREEN RECORD sa (
    FORMONLY.col_pkey,
    FORMONLY.col_name,
    FORMONLY.col_details
  );
END
```

In the program file (note the name of the record elements can be different from screen array
element names):

```
DEFINE arr DYNAMIC ARRAY OF RECORD
           pkey INTEGER,
           name VARCHAR(50),
           details VARCHAR(50)
       END RECORD
...
DISPLAY ARRAY arr TO sa.*
   ...
END DISPLAY
```

> **Tip:**
>
> The order of the screen array elements must match the order of program array
> elements used by the list dialog. However, in the list container of the `LAYOUT`
> section, the order of the columns ([field item
> tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")) does not need to match the order of the fields in the screen array. To get the tabbing
> order defined by the form, you will however have to use [`OPTIONS FIELD ORDER FORM`](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.") in the
> program.

## Using PHANTOM fields to get the same columns as the database table

In most cases, the record list data comes from a database table. A program array can be easily
defined with the same structure as its corresponding database table, by using a schema file and a
[DEFINE LIKE instruction](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types."):

```
SCHEMA stores
DEFINE arr_cust DYNAMIC ARRAY OF RECORD LIKE customer.*
```

However, you might not want to display all columns of the database table in the list container.
To hide some columns, use [`PHANTOM` field definition](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).") in the form. The screen array will hold all
columns of the table and program array. It is then possible to bind the program array to the screen
array: The number of elements in the program array and in the screen array will
match.

```
...
ATTRIBUTES
EDIT c1 = FORMONLY.col_pkey;  -- Column in TABLE container
EDIT c1 = FORMONLY.col_name;  -- Column in TABLE container
PHANTOM FORMONLY.col_details; -- Not used in LAYOUT (only in screen array)
END
...
```

## Array sub-records can be bound to flat screen arrays

If additional fields are required to hold data that is not stored in the database
table, it is possible to define the program array with a sub-record matching the database table
structure, and define volatile data fields beside this record.

```
SCHEMA stores
DEFINE arr DYNAMIC ARRAY OF RECORD
           checked CHAR(1),
           sql_data RECORD LIKE items.*,
           comment STRING
       END RECORD
MAIN
    ...
    INPUT ARRAY arr FROM sa.* ...
    ...
END MAIN
```

Here the "`checked`" and "`comment`" members are not part of the
database table, while the "`sql_data`" element is defined `LIKE` the
"`items`" table of the "`stores`" schema. All members defined in this
dynamic array can be bound to a flat screen array.

## Complete example using additional fields and phantom fields

In this example, the program array is defined with more elements than the corresponding database
table, and the form definition file uses phantom fields, to hide some database columns.

Form
file:

```
LAYOUT
GRID
{
<TABLE t1                              >
[c1    |c2                     |c3     ]
[c1    |c2                     |c3     ]
[c1    |c2                     |c3     ]
[c1    |c2                     |c3     ]
<                                      >
}
END
END
ATTRIBUTES
EDIT c1 = FORMONLY.pkey, TITLE="PKey", NOENTRY;
EDIT c2 = FORMONLY.name, TITLE="Name";
CHECKBOX c3 = FORMONLY.checked, TITLE="Checked";
PHANTOM FORMONLY.details; -- Not used in layout
PHANTOM FORMONLY.comment; -- Not used in layout
END
INSTRUCTIONS
SCREEN RECORD sa(
   FORMONLY.checked, -- Note order is different to the layout
   FORMONLY.pkey,
   FORMONLY.name,
   FORMONLY.details,
   FORMONLY.comment
  );
END
```

Program
file:

```
TYPE t_type RECORD
       checked CHAR(1),
       sql_data RECORD -- Could be RECORD LIKE items.*
         pkey INTEGER,
         name VARCHAR(50),
         details VARCHAR(200)
       END RECORD,
       comment STRING
     END RECORD

DEFINE arr DYNAMIC ARRAY OF t_type

MAIN
    OPTIONS INPUT WRAP, FIELD ORDER FORM
    CALL create_db()
    CALL fill_array()
    OPEN FORM f1 FROM "form1"
    DISPLAY FORM f1
    INPUT ARRAY arr FROM sa.* ATTRIBUTES(WITHOUT DEFAULTS)
          BEFORE ROW
              MESSAGE arr[arr_curr()].sql_data.details
    END INPUT
END MAIN

FUNCTION create_db()
    DEFINE rec t_type
    CONNECT TO ":memory:+driver='dbmsqt'"
    CREATE TABLE items (
        pkey INTEGER PRIMARY KEY,
        name VARCHAR(50),
        details VARCHAR(200)
    )
    FOR rec.sql_data.pkey=100 TO 150
        LET rec.sql_data.name = SFMT("Item %1",rec.sql_data.pkey)
        LET rec.sql_data.details = SFMT("Details for %1",rec.sql_data.pkey)
        INSERT INTO items VALUES ( rec.sql_data.* )
    END FOR
END FUNCTION

FUNCTION fill_array()
    DEFINE x INTEGER
    DECLARE c1 CURSOR FOR
        SELECT 'N', items.*, '' FROM items ORDER BY pkey
    CALL arr.clear()
    LET x = 1
    FOREACH c1 INTO arr[x].*
        LET x = x+1
    END FOREACH
    CALL arr.deleteElement(x)
    FREE c1
END FUNCTION
```

## Related links

**Related concepts**  

[Binding tables to arrays in dialogs](2318-binding-tables-to-arrays-in-dialogs.md "Program arrays act as data model that are bound to form tables, when implementing list dialogs.")

[Variable binding in DISPLAY ARRAY](1964-variable-binding-in-display-array.md "Variable binding in DISPLAY ARRAY")

[Variable binding in INPUT ARRAY](2010-variable-binding-in-input-array.md "Variable binding in INPUT ARRAY")

[Populating a DISPLAY ARRAY](2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.")
