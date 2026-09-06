---
title: "Form-level validation rules"
source: "fgl-topics/c_fgl_prog_dialogs_field_validation.html"
breadcrumb: "User interface > User interface programming > Input fields > Form-level validation rules"
type: "concept"
---

# Form-level validation rules

> Form-level validation rules can be defined for each field controlled by a dialog.

Form-level validation can be specified at the form field level with attributes such as
`NOT NULL`, `REQUIRED` and `INCLUDE`. These
attributes are part of the business rules of the application and must be checked before saving
data into the database.

## Implicit validation rule checking

An `INPUT` or `INPUT ARRAY` block automatically executes
form-level validation rules in the following cases:

- The `NOT NULL` attribute is satisfied when a value is in the field.

  `NOT NULL` is checked:
  - when the user moves to a different row in a list controlled by an `INPUT ARRAY`.
    However, if the row is temporary and none of the fields is touched, the attribute is ignored.
  - in a `DIALOG` block, when focus leaves the sub-dialog controlling the field.
  - in a `DIALOG` block, when `NEXT FIELD` gives the focus to a field
    in a different sub-dialog than the current sub-dialog.
  - when the dialog instruction is ended, for example when a procedural `DIALOG` is
    ended with `ACCEPT DIALOG`, or when an singular `INPUT` is ended with
    `ACCEPT INPUT` or with the automatic accept action.
- When the `WITHOUT DEFAULTS` option is not used or set to `FALSE`,
  the `REQUIRED` attribute is satisfied, if the field modification flag is true, if the
  field value is assigned by a `DEFAULT` form attribute, or set to a non-NULL value at
  dialog initialization.

  `REQUIRED` is checked:
  - when the user moves to a different row in a list controlled by an `INPUT ARRAY`.
    However, if the row is temporary and none of the fields is touched, the attribute is ignored.
  - in a `DIALOG` block, when focus leaves the sub-dialog controlling the field.
  - in a `DIALOG` block, when `NEXT FIELD` gives the focus to a field
    in a different sub-dialog than the current sub-dialog.
  - when the dialog instruction is ended, for example when a procedural `DIALOG` is
    ended with `ACCEPT DIALOG`, or when a singular `INPUT` is ended with
    `ACCEPT INPUT` or with the automatic accept action.
- The `INCLUDE` attribute is satisfied, if the current field value matches one of
  the values in the list defined by the attribute.

  `INCLUDE` is checked when the target program variable must be assigned, this happens:
  - when `UNBUFFERED` mode is used, focus is in the field, and an action is
    invoked.
  - when the focus leaves the field.
  - when the user moves to a different row in a list controlled by an `INPUT ARRAY`.
    However, if the row is temporary and none of the fields is touched, the attribute is ignored.
  - in a `DIALOG` block, when focus leaves the sub-dialog controlling the field.
  - in a `DIALOG` block, when `NEXT FIELD` gives the focus to a field
    in a different sub-dialog than the current sub-dialog.
  - when the dialog instruction is ended, for example when a procedural `DIALOG` is
    ended with `ACCEPT DIALOG`, or when a singular `INPUT` is ended with
    `ACCEPT INPUT` or with the automatic accept action.

## Performing validation rules explicitly

Singular input dialogs (`INPUT` / `INPUT ARRAY`) create
default accept / cancel actions. The form-level validation rules are typically performed when the
accept action is triggered.

The `DIALOG` procedural instruction can be used as in singular interactive
instructions, with the typical OK / Cancel buttons (accept / cancel actions) to finish the
instruction. The accept/cancel action handlers would respectively execute the `ACCEPT
DIALOG` and `EXIT DIALOG` instructions. This solution allows the user to
input or modify one record at a time, and the program flow must reenter the`DIALOG`
instruction to edit or create another record. Alternatively, the `DIALOG` instruction
can let the user input / modify multiple records without leaving the dialog. In this case, you need
a way to execute the form-level validation rules defined for each field, before saving the data to
the database.

To validate a subset of fields controlled by the `DIALOG` instruction, use the
`ui.Dialog.validate("field-list")` method, as shown in this
example:

```
   ON ACTION save 
      IF DIALOG.validate("cust.*") < 0 THEN
         CONTINUE DIALOG
      END IF
      CALL customer_save()
```

This method automatically
displays an error message and registers the next field in case of
error. It is mandatory to execute a `CONTINUE DIALOG` instruction
if the function returns an error.

Within singular input dialogs, form-level validation rules can also be explicitly performed
with the `ACCEPT INPUT` instruction, or with the
`DIALOG.validate("*")` API call, followed by a `CONTINUE
INPUT` in case of error.

## Related links

**Related concepts**  

[Appending rows in INPUT ARRAY](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")

[Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.")

[ACCEPT DIALOG instruction](2142-accept-dialog-instruction.md "ACCEPT DIALOG instruction")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[Input field modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.")
