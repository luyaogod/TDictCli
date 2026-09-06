---
title: "Paged mode DISPLAY ARRAY"
source: "fgl-topics/c_fgl_dynamic_dialogs_da_ofb.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Paged mode DISPLAY ARRAY"
type: "concept"
---

# Paged mode DISPLAY ARRAY

> Dynamic dialogs using DISPLAY ARRAY can use ON FILL BUFFER events to fill pages of rows.

## The `ON FILL BUFFER` event

The concept of paged mode and `ON FILL BUFFER` trigger is explained in [Paged mode of DISPLAY ARRAY](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."). Read this first.

With dynamic dialogs, the data rows must be provided when the `ON FILL BUFFER`
event occurs. The program code must provide all field values for each row, from a row offset
position and for a given number of rows that represent a page.

Field values are set with the [`ui.Dialog.setFieldValue()`](../15_library-reference/3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.") dialog method. To select the row to be assigned,
you must call the [`ui.Dialog.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") dialog method before
`setFieldValue()` calls. When used in the context of the `ON FILL
BUFFER` event, `setCurrentRow()` will only define the current row to set
field values. After executing the `ON FILL BUFFER` code, the runtime system will
automatically reset the current row in the list.

> **Important:**
>
> When the type of the `setFieldValue()` target field is [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") or [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds."), the runtime system does not make a
> full deep copy of the LOB data: Only the locator/handle is copied. As result, both original
> `TEXT`/`BYTE` and the dialog field will point to the same LOB data.
> This is especially important in list dialogs. In order to get distinct LOB data in each row of a
> list dialog, before passing the `TEXT`/`BYTE` reference to the
> `setFieldValue()` method, you must create a LOB object, initialize its location with
> `LOCATE`, and set the LOB data from the source LOB. For more details, read the [TEXT/BYTE assignments](../08_language-basics/0555-byte.md)
> section.

When implementing a singular `DISPLAY ARRAY` dynamic dialog, the event string
returned by [`nextEvent()`](../15_library-reference/3209-ui-dialog-nextevent.md "Waits for a dialog event.")
contains "`ON FILL BUFFER`". When implementing a multiple-dialog with one or more
`DISPLAY ARRAY` sub-dialogs, the `nextEvent()` returns "`ON
FILL BUFFER screen-record-name`".

## Code example

The following example implements a singular list dynamic dialog using the paged mode.

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

Program file main.4gl: This version of the sample provides the total number
of rows before the dialog loop starts:

```
TYPE t_fields DYNAMIC ARRAY OF RECORD
    name STRING,
    type STRING
END RECORD

MAIN
    DEFINE d ui.Dialog
    DEFINE t STRING
    DEFINE fields t_fields
    DEFINE x, ofb_offset, ofb_length INT

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    LET fields[1].name = "cust_id"
    LET fields[1].type = "INTEGER"
    LET fields[2].name = "cust_name"
    LET fields[2].type = "VARCHAR(50)"
    LET d = ui.Dialog.createDisplayArrayTo(fields,"sr")
    CALL d.addTrigger("ON FILL BUFFER")
    CALL d.addTrigger("ON ACTION accept")
    CALL d.setArrayLength("sr",2000)

    WHILE (t := d.nextEvent()) IS NOT NULL
        CASE t
            WHEN "ON FILL BUFFER"
                LET ofb_offset = fgl_dialog_getbufferstart()
                LET ofb_length = fgl_dialog_getbufferlength()
                DISPLAY SFMT("ON FILL BUFFER offset:%1 length:%2",ofb_offset,ofb_length)
                FOR x = ofb_offset TO ofb_offset + ofb_length -1
                    CALL d.setCurrentRow("sr",x)
                    CALL d.setFieldValue("sr.cust_id",x)
                    CALL d.setFieldValue("sr.cust_name",SFMT("Name %1",x))
                END FOR
            WHEN "ON ACTION accept"
                CALL d.accept()
            WHEN "AFTER DISPLAY"
                EXIT WHILE
        END CASE
    END WHILE
    CALL d.close()

END MAIN
```

If the total number of rows is not known before the dialog executes, set the array length to -1
and stop providing rows in `ON FILL BUFFER` when the end of the result set is
reached:

```
TYPE t_fields DYNAMIC ARRAY OF RECORD
    name STRING,
    type STRING
END RECORD

MAIN
    DEFINE d ui.Dialog
    DEFINE t STRING
    DEFINE fields t_fields
    DEFINE x, ofb_offset, ofb_length INT

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    LET fields[1].name = "cust_id"
    LET fields[1].type = "INTEGER"
    LET fields[2].name = "cust_name"
    LET fields[2].type = "VARCHAR(50)"
    LET d = ui.Dialog.createDisplayArrayTo(fields,"sr")
    CALL d.addTrigger("ON FILL BUFFER")
    CALL d.addTrigger("ON ACTION accept")
    CALL d.setArrayLength("sr",-1)

    WHILE (t := d.nextEvent()) IS NOT NULL
        CASE t
            WHEN "ON FILL BUFFER"
                LET ofb_offset = fgl_dialog_getbufferstart()
                LET ofb_length = fgl_dialog_getbufferlength()
                DISPLAY SFMT("ON FILL BUFFER offset:%1 length:%2",ofb_offset,ofb_length)
                FOR x = ofb_offset TO ofb_offset + ofb_length -1
                    IF x>100 THEN -- Simulate end of result set
                        EXIT FOR
                    END IF
                    CALL d.setCurrentRow("sr",x)
                    CALL d.setFieldValue("sr.cust_id",x)
                    CALL d.setFieldValue("sr.cust_name",SFMT("Name %1",x))
                END FOR
            WHEN "ON ACTION accept"
                CALL d.accept()
            WHEN "AFTER DISPLAY"
                EXIT WHILE
        END CASE
    END WHILE
    CALL d.close()

END MAIN
```
