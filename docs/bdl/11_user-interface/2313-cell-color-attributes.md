---
title: "Cell color attributes"
source: "fgl-topics/c_fgl_prog_dialogs_cell_attr.html"
breadcrumb: "User interface > User interface programming > List dialogs > Cell color attributes"
type: "concept"
---

# Cell color attributes

> List controllers can display every cell in a specific color.

When using the `DISPLAY ARRAY` or `INPUT ARRAY`, you can assign
specific colors to cells of a `TABLE` or `TREE` rows with the [`DIALOG.setArrayAttributes()`](../15_library-reference/3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs).")
or [`DIALOG.setCellAttributes()`](../15_library-reference/3221-ui-dialog-setcellattributes.md "Define cell decoration attributes array for the specified list (singular dialog only).") method.

Call the method in the dialog initialization clause, for example, in `BEFORE
DISPLAY` for a singular `DISPLAY ARRAY` dialog.

The method takes an array as parameter. This array must have one of the following
structures:

- A `DYNAMIC ARRAY OF RECORD` (with the same structure as the data array, but using
  the `STRING` data type)
- A `DYNAMIC ARRAY WITH DIMENSION 2 OF STRING` (to define attributes in dynamic
  dialog when the row structure is defined at runtime)
- A `DYNAMIC ARRAY OF STRING` (to define attributes for complete lines instead of
  individual cells)

Cell attributes are defined by using the TTY attributes such as "red reverse" for example. See
method reference for all possible values.

If cell attributes are changed during the dialog execution, use the `UNBUFFERED`
mode to get automatic form synchronization. The unbuffered mode is not required if the cell
attributes are defined before executing the dialog, and are left unchanged until the dialog
ends.

## Example

This is the list.per form file defining the table view:

```
LAYOUT
TABLE
{
[c1         |c2             ]
}
END
END
ATTRIBUTES
c1 = FORMONLY.key;
c2 = FORMONLY.name;
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

This is the program code (main.4gl):

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name VARCHAR(100)
           END RECORD
    DEFINE att DYNAMIC ARRAY OF RECORD
               key STRING,
               name STRING
           END RECORD
    DEFINE I INT

    FOR i=1 TO 10
        LET arr[i].key = i
        LET arr[i].name = "Item "||i
        LET att[i].key = "red reverse"
        LET att[i].name = IIF(i MOD 2,"blue","green")
    END FOR

    OPEN FORM f1 FROM "list"
    DISPLAY FORM f1

    DISPLAY ARRAY arr TO sr.* ATTRIBUTES(UNBUFFERED)
        BEFORE DISPLAY
           CALL DIALOG.setCellAttributes(att)
       ON ACTION att_modify_cell 
         LET att[2].key = "red reverse"
       ON ACTION att_clear_cell 
         LET att[2].key = NULL
    END DISPLAY

END MAIN
```

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
