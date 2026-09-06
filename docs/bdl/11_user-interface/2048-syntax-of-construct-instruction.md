---
title: "Syntax of CONSTRUCT instruction"
source: "fgl-topics/c_fgl_Construct_005.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Syntax of CONSTRUCT instruction"
type: "concept"
---

# Syntax of CONSTRUCT instruction

> The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT.

## Syntax

```
CONSTRUCT { BY NAME variable ON column-list
          | variable ON column-list FROM field-list
          }
    [ {ATTRIBUTE|ATTRIBUTES} ( { display-attribute
                     | control-attribute }
                     [,...] ) ]
    [ HELP help-number ]
[ dialog-control-block
  [...] 
END CONSTRUCT ]
```

where column-list defines a list of database columns
as:

```
{ column-name
| table-name.*
|  table-name. column-name
} [,...]
```

where field-list defines a list of fields with one or more
of:

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

where dialog-control-block is one
of:

```
{ BEFORE CONSTRUCT
| AFTER CONSTRUCT
| BEFORE FIELD field-spec [,...]
| ON CHANGE field-spec [,...] 
| AFTER FIELD field-spec [,...]
| ON IDLE seconds
| ON TIMER seconds
| ON ACTION action-name
             [ INFIELD field-spec ]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-construct ) ]
| ON KEY ( key-name [,...] )
}
    dialog-statement
    [...]
```

where action-attributes-construct
is:

```
{ TEXT = string
| COMMENT = string
| IMAGE = string
| ACCELERATOR = string
| DEFAULTVIEW = { YES | NO | AUTO }
| CONTEXTMENU = { YES | NO | AUTO }
    [,...] }
```

where dialog-statement is one
of:

```
{ statement
| ACCEPT CONSTRUCT
| EXIT CONSTRUCT
| CONTINUE CONSTRUCT
| NEXT FIELD  { CURRENT | NEXT | PREVIOUS | field-spec }
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
| CANCEL [ = boolean ]
| FIELD ORDER FORM
| HELP = help-number
| NAME = "dialog-name"
}
```

1. variable is the variable that will contain the SQL condition built by the
   `CONSTRUCT` instruction.
2. column-name is the identifier of a database column of the current form.
3. table-name is the identifier of a database table of the current form.
4. field-name is the identifier of a field of the current form.
5. screen-array is the screen array that will be used in the current form.
6. line is a screen array line in the form.
7. screen-record is the identifier of a screen record of the current form.
8. help-number is an integer that allows you to associate a help message number
   with the instruction.
9. key-name is a hot-key identifier (like `F11` or
   `Control-z`).
10. dialog-name is the identifier of the dialog.
11. action-name identifies an action that can be executed by the user.
12. seconds is an integer literal or variable that defines a number of
    seconds.
13. statement is any instruction supported by the language.
14. action-attributes are dialog-specific action attributes.
