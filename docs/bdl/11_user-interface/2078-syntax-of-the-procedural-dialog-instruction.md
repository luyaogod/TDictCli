---
title: "Syntax of the procedural DIALOG instruction"
source: "fgl-topics/c_fgl_ui_syntax_proc_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Syntax of the procedural DIALOG instruction"
type: "concept"
---

# Syntax of the procedural DIALOG instruction

> The DIALOG block is an interactive instruction that executes several sub-dialogs simultaneously.

## Syntax

```
DIALOG
   [ {ATTRIBUTE|ATTRIBUTES} ( dialog-control-attribute [,...] ) ]

   { record-input-block
   | construct-block
   | display-array-block
   | input-array-block
   | SUBDIALOG [module-name.]dialog-name [ ( param [,...] ) ]
   }
   [...]

   [
   dialog-control-block
   [...]
   ]

END DIALOG
```

where dialog-control-attribute
is:

```
{ FIELD ORDER FORM
| UNBUFFERED [ = boolean ]
}
```

where
dialog-name in the `SUBDIALOG` clause is the name of a [declarative dialog block](2152-syntax-of-the-declarative-dialog-block.md "The declarative DIALOG block defines an interactive instruction that can be used by a parent DIALOG using the SUBDIALOG clause.") defined outside the scope of
the current function, in another module identified by module-name.

where dialog-control-block is one
of:

```
{ BEFORE DIALOG
| ON ACTION action-name
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-dialog ) ]
| ON KEY ( key-name [,...] )
| ON IDLE seconds
| ON TIMER seconds
| COMMAND option-name
        [ option-comment ]
        [ HELP help-number ]
| COMMAND KEY ( key-name  [,...] ) option-name
        [ option-comment ]
        [ HELP help-number ]
| AFTER DIALOG
}
    dialog-statement
    [...]
```

where action-attributes-dialog
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

where record-input-block
is:

```
INPUT { BY NAME { variable | record.* } [,...]
      | { variable | record.* } [,...] FROM field-list
      }
   [ {ATTRIBUTE|ATTRIBUTES} ( input-control-attribute [,...] ) ]
   [ input-control-block
       [...]
   ]
END INPUT
```

where input-control-attribute
is:

```
{ HELP = help-number
| NAME = "sub-dialog-name"
| WITHOUT DEFAULTS [ = boolean ]
}
```

where
input-control-block is one of:

```
{ BEFORE INPUT
| BEFORE FIELD field-spec [,...]
| ON CHANGE field-spec [,...]
| AFTER FIELD field-spec [,...]
| AFTER INPUT
| ON ACTION action-name
             [ INFIELD field-spec ]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-input ) ]
| ON KEY ( key-name [,...] )}
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

where construct-block
is:

```
CONSTRUCT { BY NAME variable ON column-list
          | variable ON column-list FROM field-list
          }
   [ {ATTRIBUTE|ATTRIBUTES} ( construct-control-attribute [,...] ) ]
   [ construct-control-block
       [...]
   ]
END CONSTRUCT
```

where construct-control-attribute
is:

```
{ HELP = help-number
| NAME = "sub-dialog-name"
}
```

where
construct-control-block is one of:

```
{ BEFORE CONSTRUCT
| BEFORE FIELD field-spec [,...]
| ON CHANGE field-spec [,...]
| AFTER FIELD field-spec [,...]
| AFTER CONSTRUCT
| ON ACTION action-name
             [INFIELD field-spec]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-construct ) ]
| ON KEY ( key-name [,...] )}
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

where display-array-block
is:

```
DISPLAY ARRAY array TO screen-array.*
   [ {ATTRIBUTE|ATTRIBUTES} ( display-array-control-attribute [,...] ) ]
   [ display-array-control-block
       [...]
   ]
END DISPLAY
```

where display-array-control-attribute
is:

```
{ HELP = help-number
| COUNT = row-count
| KEEP CURRENT ROW = [ = boolean ]
| DOUBLECLICK = action-name
| FOCUSONFIELD
}
```

where display-array-control-block is one
of:

```
{ BEFORE DISPLAY
| BEFORE ROW
| AFTER ROW
| AFTER DISPLAY
| ON ACTION action-name
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-display-array ) ]
| ON KEY ( key-name [,...] )
| ON FILL BUFFER
| ON SELECTION CHANGE
| ON SORT
| ON APPEND [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-listmod-triggers ) ]
| ON INSERT [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-listmod-triggers ) ]
| ON UPDATE [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-listmod-triggers ) ]
| ON DELETE [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-listmod-triggers ) ]
| ON EXPAND ( row-index )
| ON COLLAPSE ( row-index )
| ON DRAG_START ( dnd-object )
| ON DRAG_FINISHED ( dnd-object )
| ON DRAG_ENTER( dnd-object )
| ON DRAG_OVER ( dnd-object )
| ON DROP ( dnd-object ) }
    dialog-statement
    [...]
```

where action-attributes-display-array
is:

```
{ TEXT = string
| COMMENT = string
| IMAGE = string
| ACCELERATOR = string
| DEFAULTVIEW = { YES | NO | AUTO }
| CONTEXTMENU = { YES | NO | AUTO }
| ROWBOUND
    [,...] }
```

where action-attributes-listmod-triggers
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

where input-array-block
is:

```
INPUT ARRAY array FROM screen-array.*
   [ {ATTRIBUTE|ATTRIBUTES} ( input-array-control-attribute [,...] ) ]
   [ input-array-control-block
       [...]
   ]
END INPUT
```

where input-array-control-attribute
is:

```
{ APPEND ROW [ = boolean ]
| AUTO APPEND [ = boolean ]
| COUNT = row-count
| DELETE ROW [ = boolean ]
| HELP = help-number
| INSERT ROW [ = boolean ]
| KEEP CURRENT ROW [ = boolean ]
| MAXCOUNT = max-row-count
| WITHOUT DEFAULTS [ = boolean ]
}
```

where
input-array-control-block is one
of:

```
{ BEFORE INPUT
| BEFORE ROW
| BEFORE FIELD [,...]
| ON CHANGE field-spec [,...]
| AFTER FIELD field-spec [,...]
| ON ROW CHANGE
| ON SORT
| AFTER ROW
| BEFORE DELETE
| AFTER DELETE
| BEFORE INSERT
| AFTER INSERT
| AFTER INPUT
| ON ACTION action-name
             [INFIELD field-spec]
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-input-array ) ]
| ON KEY ( key-name [,...] ) }
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
| ACCEPT DIALOG
| CANCEL DIALOG
| CONTINUE DIALOG
| EXIT DIALOG
| NEXT FIELD  { CURRENT | NEXT | PREVIOUS | field-spec }
}
```

where field-list defines a list of fields with one
or more
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

where column-list defines a list of database columns
as:

```
{ column-name
| table-name.*
| table-name.column-name
} [,...]
```

1. variable-definition is a variable declaration with data type as in a regular
   `DEFINE` statement.
2. array is the array of records used by the `DIALOG`
   statement.
3. help-number is an integer that allows you to associate a help message number
   with the command.
4. field-name is the identifier of a field of the current form.
5. option-name is a string expression defining the label of the action and
   identifying the action that can be executed by the user.
6. option-comment is a string expression containing a description for the menu
   option, displayed when option-name is the current.
7. column-name is the identifier of a database column of the current form.
8. table-name is the identifier of a database table of the current form.
9. variable is a simple program variable (not a record).
10. record is a program record (structured variable).
11. screen-array is the screen array that will be used in the current form.
12. line is a screen array line in the form.
13. screen-record is the identifier of a screen record of the current form.
14. action-name identifies an action that can be executed by the user.
15. seconds is an integer literal or variable that defines a number of
    seconds.
16. key-name is a hot-key identifier (like `F11` or
    `Control-z`).
17. row-index identifies the program variable which holds the row index
    corresponding to the tree node that has been expanded or collapsed.
18. dnd-object references a `ui.DragDrop` variable defined in the
    scope of the dialog.
19. statement is any instruction supported by the language.
20. action-attributes are dialog-specific action attributes for the action.
