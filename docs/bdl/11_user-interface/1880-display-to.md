---
title: "DISPLAY TO"
source: "fgl-topics/c_fgl_record_display_DISPLAY_TO.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > DISPLAY TO"
type: "concept"
---

# DISPLAY TO

> The DISPLAY ... TO instruction displays data to specific form fields.

## Syntax

```
DISPLAY expression [,...] TO field-spec [,...]
  [ {ATTRIBUTE|ATTRIBUTES} ( display-attribute [,...] ) ]
```

where field-spec
is:

```
{ field-name
| table-name.*
| table-name.field-name
| screen-array[line].*
| screen-array[line].field-name
| screen-record.*
| screen-record.field-name
} [,...]
```

where display-attribute
is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | NORMAL
| REVERSE | BLINK | UNDERLINE
}
```

1. expression is any expression supported by the language.
2. field-name is the identifier of a field of the current form.
3. table-name is the identifier of a database table of the current form.
4. screen-record is the identifier of a screen record of the current form.
5. screen-array is the screen array that will be used in the form.

## Usage

A `DISPLAY TO` statement copies the data from program variables to the form fields
specified after the `TO` keyword.

When the program variables do not have the same names as the form fields, you must use the
`TO` clause to explicitly map the variables to form fields using [a screen record or screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."). You can list
the fields individually, or you can use the `screen-record.*` or
`screen-record[n].*` notation, where
`screen-record[n].*` specifies all the fields in line `n` of a
screen array.

In this example, the values in the `p_items` program record are displayed in the
first row of the `s_items` screen
array:

```
DISPLAY p_items.* TO s_items[1].*
```

The expanded list of screen fields must correspond in order and in number to the expanded list of
identifiers after the `DISPLAY` keyword. Identifiers and their corresponding fields
must have the same or compatible data types. For example, the next `DISPLAY`
statement displays the values in the `p_customer` program record in fields of the
`s_customer` screen
record:

```
DISPLAY p_customer.* TO s_customer.*
```

For this example, the `p_customer` program record and the
`s_customer` screen record require compatible declarations. The following
`DEFINE` statement declares the `p_customer` program
record:

```
DEFINE p_customer RECORD
  customer_num LIKE customer.customer_num,
  fname LIKE customer.fname,
  lname LIKE customer.lname,
  phone LIKE customer.phone
END RECORD
```

This fragment of a form specification declares the `s_customer` screen
record:

```
ATTRIBUTES
 f000 = customer.customer_num;
 f001 = customer.fname;
 f002 = customer.lname;
 f003 = customer.phone;
END
```

The `DISPLAY TO` instruction is usually not needed if the program is always in the
context of a dialog controlling the form fields.

## DISPLAY TO / BY NAME changes the touched flag

The `DISPLAY TO / BY NAME` statement changes the [modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") of the target fields. After
displaying a field value with a `DISPLAY TO / BY NAME` instruction, the `FIELD_TOUCHED()` operator returns true and
the [`ON CHANGE`](1943-on-change-block.md) and [`ON ROW CHANGE`](2018-on-row-change-block.md) triggers may be
invoked, if the current field value was changed.

In dialogs controlling field input such as `INPUT` or `INPUT
ARRAY`, use the `UNBUFFERED` attribute to display data to fields automatically without changing the
'touched' status of fields. The `UNBUFFERED` clause will perform automatic form field
and program variable synchronization. When using the `UNBUFFERED` mode, the touched
flag can be set with `DIALOG.setFieldTouched()`, if you want to get the same effect
as a `DISPLAY TO / BY NAME`.

## Specifying TTY attributes in the DISPLAY statements

The `ATTRIBUTES` clause temporarily overrides any default display attributes or
any attributes specified in the `OPTIONS` or `OPEN WINDOW` statements
for the fields. When the `DISPLAY TO / BY NAME` statement completes execution, the
default display attributes are restored. In a `DISPLAY TO / BY NAME` statement, any
screen attributes specified in the `ATTRIBUTES` clause apply to all the fields that
you specify after the `TO` keyword.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

The `REVERSE`, `BLINK`, `INVISIBLE`, and
`UNDERLINE` attributes are not sensitive to the color or monochrome status of the
terminal, if the terminal is capable of displaying these intensity modes. The
`ATTRIBUTES` clause can include zero or more of the `BLINK`,
`REVERSE`, and `UNDERLINE` attributes, and zero or one of the other
attributes. That is, all of the attributes except `BLINK`, `REVERSE`,
and `UNDERLINE` are mutually exclusive.

The `DISPLAY TO / BY NAME` statement ignores the `INVISIBLE`
attribute, regardless of whether you specify it in the `ATTRIBUTES` clause.

## Related links

**Related concepts**  

[DISPLAY BY NAME](1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.")
