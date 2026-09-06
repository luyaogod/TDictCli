---
title: "DISPLAY BY NAME"
source: "fgl-topics/c_fgl_record_display_DISPLAY_BY_NAME.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > DISPLAY BY NAME"
type: "concept"
---

# DISPLAY BY NAME

> The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.

## Syntax

```
DISPLAY BY NAME { variable | record.* } [,...]
  [ {ATTRIBUTE|ATTRIBUTES} ( display-attribute [,...] ) ]
```

where display-attribute is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | NORMAL
| REVERSE | BLINK | UNDERLINE
}
```

1. variable is a program variable that has the same name as a form field.
2. record.\* is a record variable that has members with the same names as form
   fields.

## Usage

A `DISPLAY BY NAME` statement copies the data from program variables to the form
fields associated to the variables by name. The program variables used in `DISPLAY BY NAME`
must have the same name as the form fields where they have to be displayed. The language ignores any
record structure name prefix when matching the names. The names must be unique and unambiguous; if not,
the instruction raises an error.

For example, the following statement displays the values for the specified variables in the form
fields with corresponding names `cust_company` and `cust_address`:

```
DISPLAY BY NAME p_customer.cust_company,
                p_customer.cust_address
```

The `DISPLAY BY NAME` instruction is usually not needed if the program is always in
the context of a dialog controlling the form fields.

## DISPLAY BY NAME uses the default screen record

Unlike the `DISPLAY TO` instruction where you can explicitly specify a [screen record or screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), `DISPLAY
BY NAME` displays data to the screen fields of the default screen records. The default
screen records are those having the names of the tables defined in the `TABLES`
section of the form specification file. When the form fields define a record list in the layout,
only the first row can be referenced with the default screen record. In the next example, the form
contains a static record list definition in the
layout.

```
SCHEMA custdemo
SCREEN
{
[f01  |f02                  ]
[f01  |f02                  ]
[f01  |f02                  ]
[f01  |f02                  ]
}
END
TABLES
customer
END
ATTRIBUTES
f01 = customer.cust_num;
f02 = customer.cust_name;
END
```

In the program, a `DISPLAY BY NAME` statement will display the data in the first
line of the record list in the form:

```
DISPLAY BY NAME record_cust.*
```

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

[DISPLAY TO](1880-display-to.md "The DISPLAY ... TO instruction displays data to specific form fields.")
