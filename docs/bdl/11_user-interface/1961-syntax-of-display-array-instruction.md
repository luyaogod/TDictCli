---
title: "Syntax of DISPLAY ARRAY instruction"
source: "fgl-topics/c_fgl_DisplayArray_012.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Syntax of DISPLAY ARRAY instruction"
type: "concept"
---

# Syntax of DISPLAY ARRAY instruction

> The DISPLAY ARRAY instruction controls the display of a program array on the screen.

## Syntax

```
DISPLAY ARRAY array TO screen-array.*
  [ HELP help-number ]
  [ {ATTRIBUTE|ATTRIBUTES} ( { display-attribute
                   | control-attribute }
                       [,...]) ]
  [  dialog-control-block
  [...] 
END DISPLAY ]
```

where dialog-control-block is one of:

```
{ BEFORE DISPLAY
| AFTER DISPLAY
| BEFORE ROW
| AFTER ROW
| ON IDLE seconds
| ON TIMER seconds
| ON ACTION action-name
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-display-array ) ]
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
| ON DRAG_ENTER ( dnd-object )
| ON DRAG_OVER ( dnd-object )
| ON DROP ( dnd-object )
| ON KEY ( key-name [,...] )
}
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

where dialog-statement is one
of:

```
{ statement
| ACCEPT DISPLAY
| EXIT DISPLAY
| CONTINUE DISPLAY 
| NEXT FIELD  { CURRENT | NEXT | PREVIOUS | field-spec }
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
| KEEP CURRENT ROW [ = boolean ]
| HELP = help-number
| COUNT = row-count
| UNBUFFERED [ = boolean ]
| DOUBLECLICK = action-name
| FOCUSONFIELD
| CURRENT ROW DISPLAY = current-row-attributes
}
```

1. array is a static or dynamic array containing the records you want to
   display.
2. screen-array is the name of the screen array used to display data.
3. help-number is an integer that associates a help message number with the
   instruction.
4. action-name identifies an action that can be executed by the user.
5. seconds is an integer literal or variable that defines a number of
   seconds.
6. row-index identifies the program variable which holds the row index
   corresponding to the tree view node that has been expanded or collapsed.
7. dnd-object references a `ui.DragDrop` variable defined in the
   scope of the dialog.
8. key-name is an hot-key identifier (such as `F11` or
   `Control-z`).
9. statement is any instruction supported by the language.
10. row-count defines the total number of rows for a static array.
11. boolean is a boolean expression that evaluates to `TRUE` or
    `FALSE`.
12. action-attributes are dialog-specific action attributes.
13. current-row-attributes is a string expression of a comma-separated list of
    display-attribute keywords.
