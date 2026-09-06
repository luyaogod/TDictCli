---
title: "INPUT ARRAY behavior"
source: "fgl-topics/c_fgl_MigI4GL_044.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > INPUT ARRAY behavior"
type: "concept"
---

# INPUT ARRAY behavior

> This topic describes INPUT ARRAY differences between I4GL and Genero BDL.

## DISPLAY TO and field modification flag

During an `INPUT ARRAY` dialog, to change the value of a form field by program,
existing code typically sets the value in the program array field, and then displays this value to
the corresponding form field
with:

```
DISPLAY prog-array[arr_curr()].colname
   TO screen-array[scr_line()].fieldname
```

In such
case, IBM® Informix®
4GL will not mark the field as modified, while Genero BDL will set the [modification flag](../11_user-interface/2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") of the field, like when the
user modifies the field interactively.

As result, with Genero BDL, the row is considered as modified by the user, and row validation
will occur as if the user entered the value, while I4GL will consider the displayed value as valid,
even if it does not fullfill the form constraints, allowing the user navigate in the record list
without correcting invalid data.

This behavior has also an impact on new temporary rows created at the end of the list: If the
user does not enter values or when no `DISPLAY TO` is performed, the temporary row
disappears when moving back in the list. If a `DISPLAY TO` is used, the row is
touched and will remain in the list.

When setting a form field value by program, the value should always satisfy field
constraints.

Best practice to change the data visible to the user is to configure the `INPUT
ARRAY` dialog with the [The buffered and unbuffered modes](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")`UNBUFFERED` attribute. This way, changing the program array will automatically
sync the form fields, without the need to use a `DISPLAY TO` instruction, and without
setting the field modification flag.

## INCLUDE constraint behavior

When defining a form field with an [`INCLUDE`](../11_user-interface/1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") list and without a `NOT NULL` constraint, for
example:

```
f1 = FORMONLY.f1 TYPE CHAR, INCLUDE = ("a","b");
```

During an `INPUT ARRAY`, when modifying an existing row, IBM Informix 4GL allows the user to
empty the field with a space. With Genero BDL, the user must enter one of the values listed in the
`INCLUDE` attribute.

Consequently, I4GL allows the the end-user to clear such fields on existing rows, but forces the
field be not be null when creating new rows, which is not consistent.

The Genero BDL behavior is consistent, while the I4GL behavior can be considered as a bug.

Best practice is to add a `NOT NULL` constraint on the field, to make sure that
the field must contain one of the values defined by the `INCLUDE` list.

## Related links

**Related concepts**  

[Input fields](../11_user-interface/2230-input-fields.md "Describes various concepts related to form field management in dialogs")

[Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
