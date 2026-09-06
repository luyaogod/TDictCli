---
title: "ui.Dialog.validate"
source: "fgl-topics/c_fgl_ClassDialog_validate.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.validate"
type: "concept"
---

# ui.Dialog.validate

> Checks form level validation rules.

## Syntax

```
validate(
   formFieldList )
  RETURNS INTEGER
```

1. formFieldList is [a variable
   parameter list](../08_language-basics/0643-variable-parameter-list.md "Variable parameter list delimiters.") of strings (using `[ ]` notation), to define the fields to be
   checked. The `[ ]` square brackets are optional, if only one element is provided. See
   also [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

## Usage

Use the `validate()` method in order to execute
[`NOT NULL`](../11_user-interface/1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."),
[`REQUIRED`](../11_user-interface/1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") and
[`INCLUDE`](../11_user-interface/1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.")
validation rules defined in the [form
specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.").

The method takes a variable list parameter with the `[ ]` square brace notation,
containing a comma-separated list of fields or screen records. The `[ ]` square
brackets are optional, if only one element is to be validated:

```
DIALOG.validate("customer.cust_name")
DIALOG.validate("cust_rec.*")
DIALOG.validate( [ "cust_rec.*", "ord_rec.*" ] )
```

There are different notations to identify form fields in a dialog. For details, see
[Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).

The method returns zero if success, or the input error code of the first field which does not
satisfy the validation rules.

The current field is always checked, even if it is not part of the validation field list. This is
mandatory, otherwise the current field may be left with invalid data.

If an error occurs, the `validate()` method automatically displays the
corresponding error message, and registers the [next field](3210-ui-dialog-nextfield.md "Registers the next field to go to.") to jump to when the interactive
instruction gets the control back.

The `validate()` method does not stop code execution if an error is detected.
You must execute a `CONTINUE DIALOG` or `CONTINUE INPUT` instruction
to cancel the code execution.

A typical usage is for a "save" action:

```
ON ACTION save 
   IF DIALOG.validate("cust_rec.*") < 0 THEN
      CONTINUE DIALOG
   END IF
   CALL customer_save()
```

## Related links

**Related concepts**  

[ui.Dialog.accept](3178-ui-dialog-accept.md "Validates and terminates the dialog.")

[ui.Dialog.cancel](3188-ui-dialog-cancel.md "Cancels a parent dialog from a sub-dialog.")
