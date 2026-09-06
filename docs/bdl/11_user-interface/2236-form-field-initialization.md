---
title: "Form field initialization"
source: "fgl-topics/c_fgl_prog_dialogs_field_init.html"
breadcrumb: "User interface > User interface programming > Input fields > Form field initialization"
type: "concept"
---

# Form field initialization

> Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.

## Basics of the WITHOUT DEFAULTS option

The `INPUT` and `INPUT ARRAY` dialogs provide the `WITHOUT
DEFAULTS` option, to use program variable values when the dialog starts, or to apply the
`DEFAULT` attribute defined in
forms. The semantics of this option is slightly different in `INPUT` and
`INPUT ARRAY` dialogs. Use of the `WITHOUT DEFAULTS` clause is always
recommended in `INPUT ARRAY`.

The `WITHOUT DEFAULTS` option can be used in the binding clause or as an
`ATTRIBUTES` option. When used in the binding clause, the option is defined
statically at compile time as `TRUE`. When used as an `ATTRIBUTES`
option, it can be specified with an boolean expression, that is evaluated when the
`DIALOG` interactive instruction starts:

```
INPUT BY NAME p_cust.* ATTRIBUTES (WITHOUT DEFAULTS = NOT new)
   ...
END INPUT
```

## The WITHOUT DEFAULTS clause in INPUT

In the default mode, an `INPUT` clears the program variables and assigns the values defined by the
`DEFAULT` attribute in the form file (or indirectly, the default value defined in the
database schema files). This mode is typically used to input and insert a new record in the
database. The `REQUIRED` field
attributes are checked, to make sure that the user has entered all data that is mandatory. Note that
`REQUIRED` only forces the user to enter the field, the value can be
`NULL` unless the `NOT
NULL` attribute is used. Therefore, if you have an `AFTER FIELD` or
`ON CHANGE` control block with validation rules, you can use the
`REQUIRED` attribute to force the user to enter the field and trigger that block.

In contrast, the `WITHOUT DEFAULTS` option starts the `INPUT`
dialog with the existing values of program variables. This mode is typically used in order to update
an existing database row. Existing values are considered valid, thus the `REQUIRED`
attributes are ignored when this option is used.

The `NOT NULL` field attribute is always checked at dialog validation, even if the
`WITHOUT DEFAULTS` option is set.

## The WITHOUT DEFAULTS clause in INPUT ARRAY

With an [`INPUT ARRAY`](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records."),
the `WITHOUT DEFAULT` option defines whether the program array is populated when the
dialog begins. Once the dialog is started, existing rows are always handled as records to be updated
in the database (`WITHOUT DEFAULTS=TRUE`), while newly created rows are handled as
records to be inserted in the database (`WITHOUT DEFAULTS=FALSE`). In other words,
column default values defined in the form specification file or the database schema files are only
used for newly-created rows.

It is unusual to implement an `INPUT ARRAY` with no `WITHOUT
DEFAULTS` option, because the program array would be cleared and the list would appear
empty.

> **Important:**
>
> The default in `INPUT ARRAY` used inside `DIALOG` is
> `WITHOUT DEFAULTS=TRUE`, while in a singular `INPUT ARRAY` dialog, the
> default is `WITHOUT DEFAULTS=FALSE`.

## Variable initialization with WITHOUT DEFAULTS

When using the `WITHOUT DEFAULTS` clause, program variables are typically
initialized with default values before the `INPUT` dialog starts or in the
`BEFORE INSERT` block of an `INPUT ARRAY`, in order to set default
values for new created rows.

When initializing variables before an new record input with an
`INPUT` or `INPUT ARRAY` dialog, if a field must have no default
value, set the corresponding variable to `NULL`. If the value contains only space
characters (ASCII 32), or any other non-visible whitespace characters such as TAB (ASCII 9), this
will be considered as a non-null value, when the field defines an [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") attribute with
`NULL` in the possible values:

```
EDIT f02 = order.ord_valid TYPE CHAR, INCLUDE=("Y","N",NULL)
```

With a space or
TAB character, the field will appear empty to the user, but the current value will not match the
`INCLUDE` constraint. Note that fields with a boolean value such as
`"Y"/"N"` should be initialized with one of these values and deny nulls with [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values."), except in some rare cases
where the status is undefined.

A typical mistake is to expect `NULL`, when assigning a variable from a [substring expression](../08_language-basics/0633-substring-s-e.md "The [] (square brackets) operator extracts a substring from a variable."):

```
DEFINE c5 CHAR(5)
DEFINE c1 CHAR(1)
LET c5 = "ABC"     -- value is ABC__ (with 2 trailing spaces)
LET c1 = c5[5,5]   -- c1 will contain a space, not NULL!
```

## Example

The next code example shows how to use the `WITHOUT DEFAULTS` option in the
`ATTRIBUTES` block, initialized by a boolean variable depending on the context, to
create a new record or update an existing record.

Main form file "main\_form.per":

```
LAYOUT
GRID
{
[b1                        ]
[b2                        ]
[b3                        ]
[f1                        ]
[                          ]
[                          ]
}
END
END
ATTRIBUTES
BUTTON b1: new, TEXT = "New record";
BUTTON b2: mod, TEXT = "Update record";
BUTTON b3: cancel, TEXT = "Exit program";
TEXTEDIT f1 = FORMONLY.info, STRETCH=BOTH;
END
```

The edit form file "edit\_form.per":

```
LAYOUT
GRID
{
Pkey:[f1  ] Name:[f2                            ]
        Creation:[f3               ]
<TABLE t1                                       >
[c1                                             ]
[c1                                             ]
[c1                                             ]
<                                               >
}
END
ATTRIBUTES
EDIT f1 = FORMONLY.PKey, NOENTRY;
EDIT f2 = FORMONLY.Name, REQUIRED, NOT NULL,
  PLACEHOLDER="<Enter a name>";
DATEEDIT f3 = FORMONLY.crea, REQUIRED;
EDIT c1 = FORMONLY.comment, REQUIRED,
  PLACEHOLDER="<Enter a comment>";
END
INSTRUCTIONS
SCREEN RECORD sr_com (FORMONLY.comment);
END
```

The program code "dialog.4gl":

```
IMPORT util

TYPE type1 RECORD
           pkey INTEGER,
           name VARCHAR(40),
           crea DATE
       END RECORD
TYPE type2 DYNAMIC ARRAY OF RECORD
           comment STRING
       END RECORD

DEFINE rec type1
DEFINE com type2

MAIN
    DEFINE info STRING

    OPEN FORM f1 FROM "main_form"
    DISPLAY FORM f1

    INPUT BY NAME info ATTRIBUTES(UNBUFFERED,ACCEPT=FALSE)

        ON ACTION new
            CALL edit_data(FALSE)
            LET info = "rec = ", util.JSON.stringify(rec),
                       "\ncom = ", util.JSON.stringify(com)
        ON ACTION mod
            LET rec.pkey = 101
            LET rec.name = "Paul McCalloug"
            LET rec.crea = MDY(12,24,2012)
            CALL com.clear()
            LET com[1].comment = "first comment...."
            LET com[2].comment = "second comment...."
            LET com[3].comment = "third comment...."
            CALL edit_data(TRUE)
            LET info = "rec = ", util.JSON.stringify(rec),
                       "\ncom = ", util.JSON.stringify(com)
    END INPUT

END MAIN

FUNCTION edit_data(mod BOOLEAN)
    DEFINE l_rec type1
    DEFINE l_com type2

    OPEN WINDOW w1 WITH FORM "edit_form"

    DIALOG ATTRIBUTES(UNBUFFERED)

        INPUT BY NAME l_rec.* ATTRIBUTES(WITHOUT DEFAULTS=mod)
        END INPUT

        INPUT ARRAY l_com FROM sr_com.* ATTRIBUTES(WITHOUT DEFAULTS=mod)
        END INPUT

        BEFORE DIALOG
            IF mod THEN
                -- Copy current record values
                LET l_rec = rec
                CALL com.copyTo(l_com)
            ELSE
                -- Set default values by program
                LET l_rec.pkey = 999
                CALL l_com.clear()
                LET l_com[1].comment = "xxxxxxxx"
            END IF

        AFTER DIALOG
            IF NOT int_flag THEN
                LET rec = l_rec
                CALL l_com.copyTo(com)
            END IF

        ON ACTION accept
            LET int_flag = FALSE
            ACCEPT DIALOG

        ON ACTION cancel
            LET int_flag = TRUE
            EXIT DIALOG

    END DIALOG

    CLOSE WINDOW w1

END FUNCTION
```

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
