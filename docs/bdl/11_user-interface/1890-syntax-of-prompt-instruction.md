---
title: "Syntax of PROMPT instruction"
source: "fgl-topics/c_fgl_prompt_003.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Syntax of PROMPT instruction"
type: "concept"
---

# Syntax of PROMPT instruction

> The PROMPT statement assigns a user-supplied value to a variable.

## Syntax

```
PROMPT question
    [ {ATTRIBUTE|ATTRIBUTES} ( display-attribute [,...] ) ]
    FOR [CHAR[ACTER]] variable
    [ HELP number ]
    [ {ATTRIBUTE|ATTRIBUTES} ( control-attribute [,...] ) ]
[ dialog-control-block
   [...]
   END PROMPT ]
```

where dialog-control-block is one
of:

```
{ ON IDLE seconds
| ON TIMER seconds
| ON ACTION action-name
| ON KEY ( key-name [,...] )
}  
    statement
    [...]
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
| CENTURY = "century-spec"
| FORMAT = "format-spec"
| PICTURE = "picture-spec"
| SHIFT = { "up" | "down" }
| HELP = help-number
| UNBUFFERED [ = boolean ]
| WITHOUT DEFAULTS [ = boolean ]
}
```

1. question is a string expression displayed as
   a message for the input of the value.
2. variable is the name of the variable that receives
   the data typed by the user.
3. The `FOR CHAR` clause exits the prompt statement
   when the first character has been typed.
4. number is the help message number to be displayed
   when the user presses the help key.
5. key-name is an hot-key identifier (such as `F11` or `Control-z`).
6. action-name identifies an action that can be
   executed by the user.
7. seconds is an integer literal or variable that defines a number of
   seconds.
8. statement is an instruction that is executed
   when the user presses the key defined by key-name.
9. century-spec is a string specifying the century
   input rule, like the `CENTURY` attribute.
10. format-spec is a string defining the display
    format for the prompt field, like the `FORMAT` attribute.
11. picture-spec is a string defining the input format for the prompt field, like
    the `PICTURE` attribute.
