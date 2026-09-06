---
title: "Form field input differences"
source: "fgl-topics/c_fgl_MigI4GL_051.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Form field input differences"
type: "concept"
---

# Form field input differences

> This topic describes differences in field input between I4GL and Genero BDL.

## PICTURE chars disappear with null field

With IBM® Informix® 4GL, the mask characters defined in the `PICTURE` attribute of the
field are still displayed when you leave the field with no value.

With Genero BDL, the form field is left completly empty, without any character from the
`PICTURE` input mask. When re-entering the form field with the
`PICTURE` attribute, the mask characters appear again.

For example:

```
DATABASE FORMONLY
SCREEN
{          
[f01      ]
[f02      ]
}   
END 
ATTRIBUTES
f01 = FORMONLY.field1, PICTURE="AA-###-AA";
f02 = FORMONLY.field2;
END
```

```
MAIN
    DEFINE rec RECORD
           field1 CHAR(9),
           field2 CHAR(9)
       END RECORD
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    INPUT BY NAME rec.*
    DISPLAY "rec.field1 = ", rec.field1
END MAIN
```

Genero BDL implements a consistent behavior while I4GL does not: The mask characters of a field
with a `PICTURE` attribute are part of the value. If a field is
`NULL`, no mask characters must be shown to the user to have a visual result that
corresponds to the actual field value.

## Space clears whole DATETIME and INTERVAL field

With IBM Informix 4GL, if a `DATETIME` or `INTERVAL` field contains a
value, a single space input at the beginning of the field does not erase the whole value.

With Genero BDL, a single space input at the beginning of the field erases the current
`DATETIME` or `INTERVAL` value, as done for numeric and
`DATE` types.

For example:

```
DATABASE FORMONLY
SCREEN
{          
[f01                        ]
[f02                        ]
[f03                        ]
[f04                        ]
}   
END 
ATTRIBUTES
f01 = FORMONLY.field1;
f02 = FORMONLY.field2;
f03 = FORMONLY.field3;
f04 = FORMONLY.field4;
END
```

```
MAIN
    DEFINE rec RECORD
           field1 INTEGER,
           field2 DATE,
           field3 DATETIME YEAR TO SECOND,
           field4 INTERVAL HOUR(6) TO MINUTE
       END RECORD
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    LET rec.field1 = 9999
    LET rec.field2 = MDY(12,24,2002)
    LET rec.field3 = "2012-11-24 11:22:33"
    LET rec.field4 = "-9999:55"
    INPUT BY NAME rec.* WITHOUT DEFAULTS
END MAIN
```

The input behavior of Genero BDL for `DATETIME` and `INTERVAL`
fields is consistent with numeric and `DATE` types: Entering a space must erase all
the value.

## get\_fldbuf() returns NULL for empty fields

With IBM Informix 4GL, after clearing the form field with `CLEAR FORM` or `CLEAR
field-name`, the `get_fldbuf()` function returns a set of
space characters for a `CHAR` field. For other types `INTEGER` or
`DATE`, I4GL returns `NULL` when the field is empty.

In such case, with Genero BDL, `get_fldbuf()` returns `NULL` when
the field is empty.

The Genero BDL behavior is more consistent than I4GL: Returned `NULL` reflects the
empty field seen by the end user.

## get\_flgbuf() returns stars when overflow

With IBM Informix 4GL, when a numeric field is to show to hold the value, your can see
`***` stars to indicate an overflow, but the `get_fldbuf()` function
returns the actual value of the program variable.

The Genero BDL behavior is more consistent than I4GL: In such case, `get_fldbuf()`
returns `***` stars as seen by the user.

Review the form definition by increasing the width of the form field, or use values that fit into
the field size.

## Related links

**Related concepts**  

[PICTURE attribute](../11_user-interface/1807-picture-attribute.md "The PICTURE attribute specifies a character pattern for data entry in a text field, and prevents entry of values that conflict with the specified pattern.")
