---
title: "The INPUT sub-dialog"
source: "fgl-topics/c_fgl_DIALOG_subdialog_INPUT_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Structure of a declarative DIALOG block > The INPUT sub-dialog"
type: "concept"
---

# The INPUT sub-dialog

> The INPUT sub-dialog implements single record input in fields of the current form.

## Program variable to form field binding

Each record member variable is bound to the corresponding field of a [screen record](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), in order to manipulate the
values that the user enters in the form fields.

The `INPUT` clause
can be used in two forms:

1. `INPUT BY NAME variable-list`
2. `INPUT variable-list FROM field-list`

The `BY NAME` clause implicitly binds the fields to the variables that have the
same identifiers as the field names. The variables must be declared with the same names as the
fields from which they accept input. The runtime system ignores any record name prefix when making
the match. The unqualified names of the variables and of the fields must be unique and unambiguous
within their respective domains. If they are not, the runtime system generates an exceptions, and
sets the `status` variable to a negative value.

```
PUBLIC TYPE t_cust RECORD
       cust_num INTEGER,
       cust_name VARCHAR(50),
       cust_address VARCHAR(100)
   END RECORD
   ...
DIALOG input_customer(p_cust t_cust INOUT)
   INPUT BY NAME p_cust.*
      BEFORE FIELD cust_name 
       ...
   END INPUT
   ...
END DIALOG
```

The `FROM` clause explicitly binds the fields in the screen record to a list of
program variables by position. The number of variables or record members must equal the number of
fields listed in the `FROM` clause. Each variable must be of the same (or a
compatible) data type as the corresponding screen field. When the user enters data, the runtime
system checks the entered value against the data type of the variable, not the data type of the
screen field.

```
DEFINE c_name VARCHAR(50)
       c_addr VARCHAR(100)
   ...
DIALOG input_user_info()
  INPUT c_name,
        c_addr
   FROM FORMONLY.field01,
        FORMONLY.field02 
    BEFORE FIELD cust_name 
       ...
  END INPUT
   ...
END DIALOG
```

## Identifying an INPUT sub-dialog

The name of an `INPUT` sub-dialog can be used to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.") with a prefix.

In order to identify the `INPUT` sub-dialog with a specific name, you can
use the `ATTRIBUTES` clause to set the `NAME` attribute:

```
INPUT BY NAME p_cust.*
      ATTRIBUTES (NAME = "cust")
  ...
```

## Control blocks in INPUT

Simple record input declared with the `INPUT` sub-dialog can raise the
following triggers:

- [BEFORE INPUT](1940-before-input-block.md)
- [BEFORE FIELD](1942-before-field-block.md)
- [ON CHANGE](1943-on-change-block.md)
- [AFTER FIELD](1944-after-field-block.md)
- [AFTER INPUT](1941-after-input-block.md)

In the singular `INPUT` instruction, `BEFORE INPUT` and
`AFTER INPUT` blocks are typically used as initialization and finalization
blocks. In an `INPUT` sub-dialog of a `DIALOG` block,
`BEFORE INPUT` and `AFTER INPUT` blocks will be executed
each time the focus goes to (`BEFORE`) or leaves (`AFTER`)
the group of fields defined by this sub-dialog.

## Related links

**Related concepts**  

[INPUT ATTRIBUTES clause](2090-input-attributes-clause.md "INPUT specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")
