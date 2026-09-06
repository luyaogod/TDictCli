---
title: "Syntax of INPUT ARRAY instruction"
source: "fgl-topics/c_fgl_InputArray_005.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Syntax of INPUT ARRAY instruction"
type: "concept"
---

# Syntax of INPUT ARRAY instruction

> The INPUT ARRAY supports data entry by users into a screen array and stores the entered data in an array of records.

## Syntax

```
INPUT ARRAY array 
  [ WITHOUT DEFAULTS ]
  FROM screen-array.*   
  [ {ATTRIBUTE|ATTRIBUTES} ( { display-attribute
                   | control-attribute 
                   } [,...] ) ]
  [ HELP help-number ]
[ dialog-control-block 
   [...] 
END INPUT ]
```

where dialog-control-block is one
of:

```
{ BEFORE INPUT
| AFTER INPUT
| AFTER DELETE
| BEFORE ROW
| AFTER ROW
| BEFORE FIELD field-spec  [,...]
| AFTER FIELD field-spec  [,...]
| ON ROW CHANGE
| ON CHANGE field-spec [,...]
| ON IDLE seconds
| ON TIMER seconds
| ON ACTION action-name
             [ INFIELD field-spec ]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-input-array ) ]
| ON KEY ( key-name [,...] )
| BEFORE INSERT
| AFTER INSERT
| BEFORE DELETE
}     
  dialog-statement
  [...]
```

where action-attributes-input-array
is:

```
{ TEXT = string
| COMMENT = string
| IMAGE = string
| ACCELERATOR = string
| DEFAULTVIEW = { YES | NO | AUTO }
| VALIDATE = NO
| CONTEXTMENU = { YES | NO | AUTO }
| ROWBOUND
    [,...] }
```

where dialog-statement is one
of:

```
{ statement
| ACCEPT INPUT
| EXIT INPUT
| CONTINUE INPUT
| NEXT FIELD { CURRENT | NEXT | PREVIOUS | field-spec }
| CANCEL DELETE
| CANCEL INSERT
}
```

where field-spec identifies a unique field with one
of:

```
{ field-name
| table-name.field-name
| screen-array.field-name
| screen-record.field-name
}
```

where display-attribute
is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | INVISIBLE | NORMAL
| REVERSE | BLINK | UNDERLINE
}
```

where control-attribute
is:

```
{ ACCEPT [ = boolean ]
| APPEND ROW [ = boolean ]
| AUTO APPEND [ = boolean ]
| CANCEL [ = boolean ]
| COUNT = row-count
| DELETE ROW [ = boolean ]
| FIELD ORDER FORM
| HELP = help-number
| INSERT ROW [ = boolean ]
| KEEP CURRENT ROW [ = boolean ]
| MAXCOUNT = max-row-count
| UNBUFFERED [ = boolean ]
| WITHOUT DEFAULTS [ = boolean ]
| CURRENT ROW DISPLAY = current-row-attributes
}
```

1. array is the array of records that will be filled by the `INPUT
   ARRAY` statement.
2. help-number is an integer that allows you to associate a help message number
   with the instruction.
3. field-name is the identifier of a field of the current form.
4. table-name is the identifier of a database table of the current form.
5. screen-record is the identifier of a screen record of the current form.
6. screen-array is the screen array that will be used in the form.
7. action-name identifies an action that can be executed by the user.
8. seconds is an integer literal or variable that defines a number of
   seconds.
9. key-name is a hot-key identifier (like `F11` or
   `Control-z`).
10. statement is any instruction supported by the language.
11. row-count defines the initial number of rows for a static array.
12. max-row-count is the maximum number of rows that can be created.
13. boolean is a boolean expression that evaluates to `TRUE` or
    `FALSE`.
14. action-attributes are dialog-specific action attributes.
15. current-row-attributes is a string expression of a comma-separated list of
    display-attribute keywords.
