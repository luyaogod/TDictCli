---
title: "ui.Dialog.setFieldValue"
source: "fgl-topics/c_fgl_ClassDialog_setFieldValue.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setFieldValue"
type: "concept"
---

# ui.Dialog.setFieldValue

> Sets the value of a field controlled by the dialog object.

## Syntax

```
setFieldValue(
   name STRING,
   value fgl-type )
```

1. name is the name of the form field, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. value is the value to be set, where fgl-type
   is one of the [primitive data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").

## Usage

The `setFieldValue()` method can be used when implementing a dynamic
dialog, to set the value of a field:

```
DEFINE default_address STRING,
       default_creadate DATE
...
CALL d.setFieldValue( "customer.cust_addr", default_address )
CALL d.setFieldValue( "customer.cust_creadate", default_creadate )
```

The first parameter defines the field to be set.

In a dynamic dialog controlling a list of records (`INPUT ARRAY` / `DISPLAY
ARRAY`), this method sets the value for a field in the current row. To fill the list of
records before dynamic dialog execution, use [`setCurrentRow()`](3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") to set the current row, then set field (cell values) with
`setFieldValue()`.

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

During dialog execution, the `setFieldValue()` method must only be used to set the
value of a field for the current row. In a regular dynamic list dialog, calling the [`setCurrentRow()`](3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") method to change
the current row before calling `setFieldValue()` will have no effect. However, when
the `ON FILL BUFFER` event occurs in a `DISPLAY ARRAY` using the
[paged mode](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."), the
`setCurrentRow()` method must be used to set the current row before calling
`setFieldValue()` for each field to fill the data page. In this context, a
`setFieldValue()` will produce error [-8129](4483-genero-bdl-errors.md), if the current row is not
part of the visible page. In a `DISPLAY ARRAY` using the [full list mode](../11_user-interface/2308-full-list-mode-of-display-array.md "In order to handle short/medium result sets, use the full list mode of DISPLAY ARRAY."), the error -8129 is raised, if
there is no current row (when the array is empty).

## Example

The following code example implements a `FOR` loop to copy values of
all fields of the `d_disparr` dialog to the field of the
`d_recinp` dialog:

```
DEFINE row, i INTEGER,
       h base.SqlHandle,
       fields DYNAMIC ARRAY OF RECORD
                  name STRING,
                  type STRING
              END RECORD,
       d_rec ui.Dialog,
       d_list ui.Dialog
...
    -- Fill the array with rows from an SqlHandle object
    CALL h.open()
    LET row = 0
    CALL h.fetch()
    WHILE status == 0
        -- must set the current row before setting values
        CALL d_list.setCurrentRow("sr_custlist", row:=row+1 )
        FOR i = 1 TO h.getResultCount()
            CALL d_list.setFieldValue( h.getResultName(i), h.getResultValue(i) )
        END FOR
        CALL h.fetch()
    END WHILE
    CALL h.close()
    CALL d_list.setCurrentRow("sr_custlist", 1)
...
    -- Copy field values from d_list to d_rec dialog
    FOR i=1 TO fields.getLength()
        CALL d_rec.setFieldValue( fields[i].name,
               d_list.getFieldValue( fields[i].name )
             )
    END FOR
```

## Related links

**Related concepts**  

[ui.Dialog.getFieldValue](3199-ui-dialog-getfieldvalue.md "Returns the value of a field controlled by a dynamic dialog.")

[Dynamic Dialogs](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
