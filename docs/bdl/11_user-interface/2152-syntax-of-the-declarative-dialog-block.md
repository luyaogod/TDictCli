---
title: "Syntax of the declarative DIALOG block"
source: "fgl-topics/c_fgl_ui_syntax_decl_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Syntax of the declarative DIALOG block"
type: "concept"
---

# Syntax of the declarative DIALOG block

> The declarative DIALOG block defines an interactive instruction that can be used by a parent DIALOG using the SUBDIALOG clause.

## Syntax

```
[ PRIVATE | PUBLIC ] DIALOG dialog-name (
     { parameter-name type-specification
     | record-name record-type INOUT
     }
     [,...]
  )
   [ define-block ]
   { record-input-block
   | construct-block
   | display-array-block
   | input-array-block
   }
END DIALOG
```

1. dialog-name defines the identifier for the declarative
   `DIALOG` block.
2. parameter-name is the name of a formal argument of the sub-dialog.
3. type-specification can be one of:
   - A [primitive type](../08_language-basics/0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](../08_language-basics/0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](../08_language-basics/0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
4. record-name is the name of a parameter defined as
   record-type.
5. record-type is a the name of a user-defined [`TYPE`](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") declared as a [`RECORD`](../08_language-basics/0717-record.md "The RECORD keyword defines a structured type or variable.") structure.

where define-block is a [local
variable declaration block](../08_language-basics/0688-define.md "The DEFINE instruction declares a program variable with a given type.").

A declarative `DIALOG` block holds a single dialog sub-block, that can be one of:

- record-input-block
- construct-block
- display-array-block
- input-array-block

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
