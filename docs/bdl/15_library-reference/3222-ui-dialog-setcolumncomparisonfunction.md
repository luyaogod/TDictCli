---
title: "ui.Dialog.setColumnComparisonFunction"
source: "fgl-topics/c_fgl_ClassDialog_setColumnComparisonFunction.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setColumnComparisonFunction"
type: "concept"
---

# ui.Dialog.setColumnComparisonFunction

> Associates a comparison function to a form field of a list dialog.

## Syntax

```
setColumnComparisonFunction(
    name STRING,
    compare FUNCTION (
                 s1 STRING,
                 s2 STRING
             ) RETURNS INTEGER
 )
```

1. name is the name of form field that corresponds to a member of a structured
   array used by a [list dialog](../11_user-interface/2300-understanding-list-dialogs.md "List dialogs are dialogs controlling a list of records rendered in a list container such as TABLE, TREE or SCROLLGRID."). See [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md).
2. compare is a [function
   reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") of a user-defined sort function.

## Usage

The `setColumnComparisonFunction()` method binds a user-defined comparison function to
a form field used by a `DISPLAY ARRAY` or `INPUT ARRAY` list dialog.
This comparison function will be used to sort the rows when column sorting is used during this
dialog.

This method can be used several times for the same dialog, to associate different comparison
functions for each column.

The first parameter defines the column form field that corresponds to the record member of the
array bound to the dialog.

The last parameter is a [function
reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") to a user-defined function that compares two values s1 and
s2. The comparison function must return a negative `INTEGER`, if
value s1 is less then s2, a positive `INTEGER`,
if s1 is greater than s2, and return zero when
s1 and s2 are equal. Take care with `INTEGER`
overflows, when substracting large numbers to produce the negative or positive return value.
> **Tip:**
>
> The comparison function can be a reference to the [`util.Strings.collate()`](3557-util-strings-collate.md "Compares two strings using locale collation rules.") or [`util.Strings.collateNumeric()`](3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits.") functions: `FUNCTION
> util.Strings.collateNumeric`

Pay attention to `NULL` values: When both s1
and s2 are `NULL`, they can be considered as equal and the
comparison function should return zero. If only one of the values is `NULL`, it's up
to the comparison function to consider `NULL` as lower or as greater than any
non-null value. The [`NVL()`](../08_language-basics/0615-nvl-function.md "The NVL() operator returns the second parameter if the first argument evaluates to NULL.") operator is very useful in such case.

To distinguish several `DISPLAY ARRAY` or `INPUT ARRAY` sub-dialogs
in a [multiple-dialog](../11_user-interface/2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."), specify the
screen-array name as prefix for the field name bound to the comparison function. For example, with
`DIALOG.setColumnComparisonFunction("sr.doctype")` the sub-dialog is identified by
the "`sr`" screen array.

## Example

In this example, the "`doctype`" column defines a document type, that must be
ordered in a specific way, instead of the alphabetical order:

Form file "form.per":

```
LAYOUT
TABLE
{              
[c1     |c2                      |c3          ]
[c1     |c2                      |c3          ]
[c1     |c2                      |c3          ]
}   
END
END
ATTRIBUTES
c1 = FORMONLY.id, TITLE="Id";
c2 = FORMONLY.name, TITLE="Name";
c3 = FORMONLY.doctype, TITLE="Type";
END 
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

Program file "main.4gl":

```
MAIN
    DEFINE files DYNAMIC ARRAY OF RECORD
               id INTEGER,
               name STRING,
               doctype STRING
           END RECORD =
    [
        (id:101, name:"picture1.png", doctype:"Image"),
        (id:102, name:"config.json",  doctype:"JSON"),
        (id:103, name:"picture2.png", doctype:"Image"),
        (id:104, name:"myvideo1.mp4", doctype:"Video"),
        (id:105, name:"profile.txt",  doctype:"Profile"),
        (id:106, name:"picture3.png", doctype:"Image")
    ]

    OPEN FORM f FROM "form"
    DISPLAY FORM f

    DISPLAY ARRAY files TO sr.*
        BEFORE DISPLAY
            CALL DIALOG.setColumnComparisonFunction("sr.doctype", FUNCTION compareDocType)
    END DISPLAY

END MAIN

DEFINE dt2int DICTIONARY OF INTEGER =
       ( "Video":10, "Profile":20, "Image":30, "JSON":40 )

FUNCTION compareDocType(s1 STRING, s2 STRING) RETURNS INTEGER
    RETURN NVL(dt2int[s1],-100) - NVL(dt2Int[s2],-100)
END FUNCTION
```

## Related links

**Related concepts**  

[util.Strings.collateNumeric](3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits.")
