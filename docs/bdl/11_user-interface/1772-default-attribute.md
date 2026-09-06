---
title: "DEFAULT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_DEFAULT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > DEFAULT attribute"
type: "concept"
---

# DEFAULT attribute

> The DEFAULT attribute defines a default value to a field during data entry.

## Syntax

```
DEFAULT = value
```

1. value can be any literal expression supported by the form compiler,
   as long as it matches the form field type.
2. value can be `TODAY` to specify the current system
   date as default.
3. value can be `CURRENT` to specify the current system
   datetime as default.

## Usage

The `DEFAULT` attribute defines a default value for the form field.

The literal constant specified as default value must match the form field type. For example,
when defining a numeric field, use a numeric decimal constant, for character string fields,
use a double-quoted character literal.

The effect of the `DEFAULT` attribute depends on the `WITHOUT
DEFAULTS` configuration option of the dialog using the form:

- With the [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") statement, form
  default values are ignored when using the `WITHOUT DEFAULTS` option. With this
  option, the runtime system displays the values in the program variables to the screen. Otherwise,
  the form default values will be displayed when the dialog starts.
- With the [`INPUT ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") statement,
  the `WITHOUT DEFAULTS` option indicates if the existing program array elements have
  to be used. The form default values are always used for new rows inserted by the user. Default
  values can be set by program in the `BEFORE INSERT` block.

Default values can also be specified in the [database
schema file](../09_advanced-features/0796-column-validation-file-val.md "The .val database schema file holds functional and display attributes of database table columns."), for form fields defined with database column reference.

When not using `WITHOUT NULL INPUT` option in the `DATABASE`
section of a form, all fields default to null values unless a `DEFAULT` attribute
is specified in the `ATTRIBUTES` section.

If the field is `FORMONLY`, you must also specify a data type when you assign
the `DEFAULT` attribute to a field.

If both the `DEFAULT` attribute and the `REQUIRED` attribute are
defined for a field, the `REQUIRED` attribute is ignored.

Note that `DATETIME` and `INTERVAL` literals are not supported
in the `DEFAULT` attribute.

## Example

```
EDIT f001 = order.orderdate, DEFAULT = TODAY;
EDIT f012 = FORMONLY.discount TYPE DECIMAL(5,2), DEFAULT=0.10;
```

## Related links

**Related concepts**  

[Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.")

[REQUIRED attribute](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.")

[EDIT item type](1690-edit-item-type.md "Defines a simple line-edit field.")
