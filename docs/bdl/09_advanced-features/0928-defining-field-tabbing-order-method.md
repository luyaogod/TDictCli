---
title: "Defining field tabbing order method"
source: "fgl-topics/c_fgl_programs_012.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining field tabbing order method"
type: "concept"
description: "Syntax OPTIONS FIELD ORDER { CONSTRAINED | UNCONSTRAINED | FORM } Usage Tabbing order is used in interactive instructions such as INPUT , INPUT ARRAY , or CONSTRUCT , where individual fields can get ..."
---

# Defining field tabbing order method

## Syntax

```
OPTIONS FIELD ORDER { CONSTRAINED | UNCONSTRAINED | FORM }
```

## Usage

Tabbing order is used in interactive instructions such as [INPUT](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form."), [INPUT ARRAY](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form."), or [CONSTRUCT](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form."), where individual fields can get the focus.

The `FIELD ORDER` runtime option defines the default behavior when moving from
field to field with the TAB and SHIFT-TAB keys in GUI mode, and with the Up / Down arrow keys in TUI
mode.

The `OPTIONS FIELD ORDER` defines the global field order mode. The field
order mode can also be defined at the dialog level, with the [`FIELD ORDER`](../11_user-interface/2089-dialog-attributes-clause.md) dialog
attribute.

By default, the tabbing order is defined by the list of fields used by the program
instruction. This corresponds to `FIELD ORDER CONSTRAINED` option, which is the
default.

When using `FIELD ORDER UNCONSTRAINED` in TUI mode, the Up and Down
arrow keys will move the cursor to the field above or below the current field, respectively. When
using the default `FIELD ORDER CONSTRAINED` option, the Up and Down arrow keys move
the cursor to the previous or next field, respectively. If `FIELD ORDER
UNCONSTRAINED` is used, the `Dialog.fieldOrder` [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") entry is ignored.

The `UNCONSTRAINED` option can only be supported in TUI mode, with a simple form layout.
It is not recommended to use this option in GUI mode.

The `FIELD ORDER FORM`
option instructs interactive instructions to use the tabbing order defined by the `TABINDEX` attributes of the current form
fields. With this option, tabbing order can be defined in the layout of the form, independently from
the program instruction. This is the preferred way in GUI mode. When `FIELD ORDER
FORM` is used, the `Dialog.fieldOrder` FGLPROFILE entry is ignored.

## Example

Form "form1.per":

```
LAYOUT
GRID
{
  First name:  [f001            ] Last name:   [f002            ]
  Address:     [f003                                            ]
}
END
END

ATTRIBUTES
EDIT f001 = FORMONLY.fname, TABINDEX = 2;
EDIT f002 = FORMONLY.lname, TABINDEX = 1;
EDIT f003 = FORMONLY.address, TABINDEX = 0;
END
```

Module
"main.4gl":

```
MAIN
  DEFINE rec RECORD
             fname VARCHAR(20),
             lname VARCHAR(20),
             address VARCHAR(50)
         END RECORD

  OPTIONS INPUT WRAP

  OPEN FORM f1 FROM "form1"
  DISPLAY FORM f1

  OPTIONS FIELD ORDER CONSTRAINED
  INPUT BY NAME rec.*

  OPTIONS FIELD ORDER UNCONSTRAINED
  INPUT BY NAME rec.*

  OPTIONS FIELD ORDER FORM
  INPUT BY NAME rec.*

END MAIN
```

## Related links

**Related concepts**  

[Defining the field input loop](0927-defining-the-field-input-loop.md "The OPTIONS INPUT [NO] WRAP instructions defines field wrapping in dialogs.")

[Defining the tabbing order](../11_user-interface/2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.")
