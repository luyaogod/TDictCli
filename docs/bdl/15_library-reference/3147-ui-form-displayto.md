---
title: "ui.Form.displayTo"
source: "fgl-topics/c_fgl_ClassForm_displayTo.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.displayTo"
type: "concept"
---

# ui.Form.displayTo

> Displays values to form fields or screen arrays.

## Syntax

```
ui.Form.displayTo(
    value { base-type | RECORD },
    formFieldName STRING,
    screenLine INTEGER,
    attributes STRING
  )
```

1. value is the scalar value or `RECORD` to be displayed, and
   base-type is a base data type of Genero (`INTEGER`,
   `DATE`, `VARCHAR`)
2. formFieldName is a simple form field name, like
   `"customer.cust_name"` or the name of a `SCREEN RECORD`
   defined in the current form.
3. screenLine is a line number in the screen array.
4. attributes is a space-separated list of TTY attribute names
   (like `"RED REVERSE"`).

## Usage

The `displayTo()` class method is equivalent to the `DISPLAY TO`
instruction.

Use of this method is only recommended, if the name of the form field, the name of the screen
record, and/or the display attributes are not known at compile-time. With regular (static) dialogs,
the `DISPLAY TO` or `DISPLAY BY NAME` instructions should be used
instead.

This method is typically used in the context of [dynamic
dialog programming](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.").

The `displayTo()` method can be used to:

- display a simple value to a form field,
- display a complete `RECORD` to all fields of a `SCREEN RECORD`,
- display a complete `RECORD` to all cells of a screen array line of a
  `SCREEN RECORD`.

> **Important:**
>
> When using a `RECORD` as value parameter, do not use the
> `.*` notation. However, in the second parameter of the
> `ui.Form.displayTo()` method, the screen record must be specified with a
> `.*` notation.

When displaying a complete `RECORD` to a screen array, you must specify the screen
line as third parameter. Otherwise, for simple form fields or flat screen records, the screen line
is ignored.

To display a simple value to a form field:

```
CALL ui.Form.displayTo("foo", "f1", NULL, NULL)
-- is the same as:
DISPLAY "foo" TO f1
```

Assuming that in the following code examples, the
`rec` variable is defined as a `RECORD`, and the `"sr"`
name is used in the definition of the `SCREEN RECORD` in the current form:

To display all elements of a complete `RECORD` to all fields grouped in a screen
record:

```
CALL ui.Form.displayTo(rec, "sr.*", NULL, NULL)
-- is the same as:
DISPLAY rec.* TO sr.*
```

To display a complete `RECORD` to a specific screen array
line:

```
CALL ui.Form.displayTo(rec, "sr.*", 2, "REVERSE, RED")
-- is the same as:
DISPLAY rec.* TO sr[2].* ATTRIBUTES(REVERSE, RED)
```

When passing `NULL` for the attributes parameter, the
`displayTo()` method will behave as `DISPLAY TO` without an
`ATTRIBUTE` clause, and the default TTY attributes at dialog, form or window level
will apply. If you want to display the value without any TTY attribute, specify the string
`"normal"` in the attributes parameter:

```
CALL ui.Form.displayTo("foo", "f1", NULL, "NORMAL")
-- is the same as:
DISPLAY "foo" TO f1 ATTRIBUTES(NORMAL)
```

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")

[Form fields](../11_user-interface/1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[DISPLAY TO / BY NAME in dialogs](../11_user-interface/2139-display-to-by-name-in-dialogs.md "DISPLAY TO / BY NAME in dialogs")
