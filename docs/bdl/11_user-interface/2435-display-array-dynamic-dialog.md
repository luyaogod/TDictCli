---
title: "DISPLAY ARRAY dynamic dialog"
source: "fgl-topics/c_fgl_dynamic_dialogs_da.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > DISPLAY ARRAY dynamic dialog"
type: "concept"
---

# DISPLAY ARRAY dynamic dialog

> DISPLAY ARRAY can be implemented with dynamic dialogs.

## Row data methods for DISPLAY ARRAY dynamic dialog

To append, insert, delete the data rows of a `DISPLAY ARRAY` list dialog with the
dynamic dialog API, use the [`appendRow()`](../15_library-reference/3185-ui-dialog-appendrow.md "Appends a new row in the specified list."), [`insertRow()`](../15_library-reference/3205-ui-dialog-insertrow.md "Inserts a new row in the specified list."), [`deleteRow()`](../15_library-reference/3192-ui-dialog-deleterow.md "Deletes a row from the specified list."), [`deleteAllRows()`](../15_library-reference/3190-ui-dialog-deleteallrows.md "Deletes all rows from the specified list.") methods of the `ui.Dialog` class.

To set/get the values of the cells of a given row, first call [`setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list."), then set the
cell values with [`setFieldValue()`](../15_library-reference/3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.") or [`getFieldValue()`](../15_library-reference/3199-ui-dialog-getfieldvalue.md "Returns the value of a field controlled by a dynamic dialog."). If the row does not yet exist, use
`appendRow()` or `insertRow()` before
`setCurrentRow()`.

## Code example

The following example implements a singular list dynamic dialog.

Form definition file form.per:

```
LAYOUT
TABLE table1
{
[c11      |c12            ]
[c11      |c12            ]
[c11      |c12            ]
[c11      |c12            ]
}
END
END

ATTRIBUTES
EDIT c11 = FORMONLY.cust_id, TITLE="Num";
EDIT c12 = FORMONLY.cust_name, TITLE="Name";
END

INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

Program file main.4gl:

```
TYPE t_fields DYNAMIC ARRAY OF RECORD
    name STRING,
    type STRING
END RECORD

MAIN
    DEFINE d ui.Dialog
    DEFINE t STRING
    DEFINE fields t_fields
    DEFINE x INT

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    LET fields[1].name = "cust_id"
    LET fields[1].type = "INTEGER"
    LET fields[2].name = "cust_name"
    LET fields[2].type = "VARCHAR(50)"
    LET d = ui.Dialog.createDisplayArrayTo(fields,"sr")
    CALL d.addTrigger("ON ACTION re_fill")
    CALL d.addTrigger("ON ACTION add_row")
    CALL d.addTrigger("ON ACTION del_row")
    CALL d.addTrigger("ON ACTION clear_all")
    CALL d.addTrigger("ON ACTION accept")

    CALL fill_rows(d,10)

    WHILE (t := d.nextEvent()) IS NOT NULL
        CASE t
            WHEN "ON ACTION re_fill"
                CALL fill_rows(d,100)
                CALL d.setCurrentRow("sr",1)
            WHEN "ON ACTION add_row"
                CALL add_row(d)
            WHEN "ON ACTION del_row"
                CALL d.deleteRow("sr",d.getCurrentRow("sr"))
            WHEN "ON ACTION clear_all"
                CALL d.deleteAllRows("sr")
            WHEN "ON ACTION accept"
                CALL d.accept()
            WHEN "AFTER DISPLAY"
                EXIT WHILE
        END CASE
    END WHILE
    CALL d.close()

END MAIN

FUNCTION fill_rows(d ui.Dialog, cnt INT)
    DEFINE x INTEGER
    CALL d.deleteAllRows("sr")
    FOR x=1 TO cnt
        CALL d.appendRow("sr")
        CALL d.setCurrentRow("sr",x)
        CALL d.setFieldValue("sr.cust_id",x)
        CALL d.setFieldValue("sr.cust_name",SFMT("Name %1",x))
    END FOR
END FUNCTION

FUNCTION add_row(d ui.Dialog)
    DEFINE x INTEGER
    CALL d.appendRow("sr")
    LET x = d.getArrayLength("sr")
    CALL d.setCurrentRow("sr",x)
    CALL d.setFieldValue("sr.cust_id",x)
    CALL d.setFieldValue("sr.cust_name",SFMT("New %1",x))
END FUNCTION
```
