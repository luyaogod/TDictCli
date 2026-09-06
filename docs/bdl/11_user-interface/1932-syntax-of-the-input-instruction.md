---
title: "Syntax of the INPUT instruction"
source: "fgl-topics/c_fgl_record_input_003.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Syntax of the INPUT instruction"
type: "concept"
---

# Syntax of the INPUT instruction

> The INPUT statement supports data entry in fields of the current form.

## Syntax

```
INPUT { BY NAME {variable|record.*} [,...] [WITHOUT DEFAULTS]
  | {variable|record.*} [,...] [WITHOUT DEFAULTS] FROM field-list
  }
  [ {ATTRIBUTE|ATTRIBUTES} (
          { display-attribute
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
| BEFORE FIELD field-spec [,...] 
| AFTER FIELD field-spec [,...] 
| ON CHANGE field-spec [,...] 
| ON IDLE seconds
| ON TIMER seconds
| ON ACTION action-name
             [ INFIELD field-spec ]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-input ) ]
| ON KEY ( key-name [,...] )
}
    dialog-statement
    [...]
```

where action-attributes-input
is:

```
{ TEXT = string
| COMMENT = string
| IMAGE = string
| ACCELERATOR = string
| DEFAULTVIEW = { YES | NO | AUTO }
| VALIDATE = NO
| CONTEXTMENU = { YES | NO | AUTO }
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
}
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

where field-spec identifies a unique field with one
of:

```
{ field-name
| table-name.field-name
| screen-array.field-name
| screen-record.field-name
}
```

where display-attribute is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | INVISIBLE | NORMAL
| REVERSE | BLINK | UNDERLINE
}
```

where control-attribute is:

```
{ ACCEPT [ = boolean ]
| CANCEL [ = boolean ]
| FIELD ORDER FORM
| HELP = help-number
| NAME = "dialog-name"
| UNBUFFERED [ = boolean ]
| WITHOUT DEFAULTS [ = boolean ]
}
```

1. variable is a program variable that will be filled by the
   `INPUT` statement.
2. record.\* is a record variable that will be filled by the
   `INPUT` statement.
3. help-number is an integer that allows you to associate a help message number
   with the instruction.
4. field-name is the identifier of a field of the current form.
5. table-name is the identifier of a database table of the current form.
6. screen-record is the identifier of a screen record of the current form.
7. screen-array is the screen array that will be used in the form.
8. line is a screen array line in the form.
9. key-name is a hot-key identifier (like `F11` or
   `Control-z`).
10. dialog-name is the identifier of the dialog.
11. seconds is an integer literal or variable that defines a number of
    seconds.
12. action-name identifies an action that can be executed by the user.
13. statement is any instruction supported by the language.
14. boolean is a boolean expression evaluated when the dialog starts.
15. action-attributes are dialog-specific action attributes.
