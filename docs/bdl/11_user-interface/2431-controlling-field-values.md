---
title: "Controlling field values"
source: "fgl-topics/c_fgl_dynamic_dialogs_fields.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Controlling field values"
type: "concept"
---

# Controlling field values

> Fields values in dynamic dialogs can be manipulated dynamically.

## Unbuffered mode is the default

Dynamic dialogs do not use program variables and thus behave by default as static dialogs using
the `UNBUFFERED` mode: When an action is fired, and the corresponding trigger handler
is executed, the field is validated and the value is available with the [`ui.Dialog.getFieldValue()`](../15_library-reference/3199-ui-dialog-getfieldvalue.md "Returns the value of a field controlled by a dynamic dialog.")
method. Changing the value by program with [`ui.Dialog.setFieldValue()`](../15_library-reference/3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.") is automatically displayed to the corresponding
form-field and visible when the control goes back to the end user.

For more details about the `UNBUFFERED` attribute, see [The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.").

## Default form-field values

A dynamic input dialog created with [`ui.Dialog.createInputByName()`](../15_library-reference/3174-ui-dialog-createinputbyname.md "Creates an ui.Dialog object to implement a dynamic INPUT BY NAME.") behaves like a static `INPUT`
dialog using the `WITHOUT DEFAULTS` option: The `DEFAULT` attribute of
the form-field is not used.

A dynamic input array dialog created with [`ui.Dialog.createInputArrayFrom()`](../15_library-reference/3173-ui-dialog-createinputarrayfrom.md "Creates an ui.Dialog object to implement a dynamic INPUT ARRAY FROM.") behaves like a static `INPUT
ARRAY` using the `WITHOUT DEFAULTS` option: The values set in the internal
rows before starting the dialog will be used. However, like with a static `INPUT
ARRAY`, when adding a new row, the `DEFAULT` attributes of the form-fields
are used.

For more details about the `WITHOUT DEFAULTS` clause, see [Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.").

## Setting and getting field values

A dynamic dialog stores field values in internal buffers based on the field definitions provided
in the creation method. Access to these values is required, to implement the dynamic dialog.

For example, to set default values before entering the dialog loop, modifying and/or querying
values during the dialog loop, and to get the entered values after dialog termination when accepted
by the user.

To set or get values of fields controlled by a dynamic dialog, use respectively the [`ui.Dialog.setFieldValue()`](../15_library-reference/3228-ui-dialog-setfieldvalue.md "Sets the value of a field controlled by the dialog object.") and
[`ui.Dialog.getFieldValue()`](../15_library-reference/3199-ui-dialog-getfieldvalue.md "Returns the value of a field controlled by a dynamic dialog.")
methods. These methods take a form field name as parameter, that can be provided in different
notations. See [Identifying fields in ui.Dialog methods](../15_library-reference/3239-identifying-fields-in-ui-dialog-methods.md) for more details.

When implementing a display array or input array dynamic dialog handling a record list, the
set/get field value methods apply to the current row. If you want to set or get field values of a
particular row, first move to the row with the [`ui.Dialog.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.") method.

This example copies the values from the fields in the current row of a display array dynamic
dialog (`d_list`), to the field buffers of a record input dynamic dialog
(`d_rec`):

```
CALL d_list.setCurrentRow("sr_custlist", index)
FOR i=1 TO fields.getLength()
    CALL d_rec.setFieldValue( fields[i].name,
               d_list.getFieldValue(fields[i].name)
         )
END FOR
```

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

## Getting query conditions for a field

A dynamic dialog created with [`ui.Dialog.createConstructByName()`](../15_library-reference/3171-ui-dialog-createconstructbyname.md "Creates an ui.Dialog object to implement a dynamic CONSTRUCT BY NAME.") handles query by example input (like
`CONSTRUCT`).

To generate the SQL condition from the search value entered in a construct field, use the [`ui.Dialog.getQueryFromField()`](../15_library-reference/3201-ui-dialog-getqueryfromfield.md "Returns the SQL condition of a field used in a query by example dialog.") method, by passing the field name as
parameter:

```
LET field_condition = DIALOG.getQueryFromField("customer.cust_name")
```

To build the complete `WHERE` part for the `SELECT` statement,
iterate through all form fields and concatenate the form field condition by separating with the
`AND` or with the `OR`
operator:

```
FOR i=1 TO fields.getLength()
    LET field_condition = d.getQueryFromField(fields[i].name)
    IF field_condition IS NOT NULL THEN
       IF where_clause IS NOT NULL THEN
          LET where_clause = where_clause, " AND "
       END IF
       LET where_clause = where_clause, field_condition
    END IF
END FOR
```
