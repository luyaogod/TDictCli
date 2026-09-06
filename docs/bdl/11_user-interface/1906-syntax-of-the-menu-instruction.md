---
title: "Syntax of the MENU instruction"
source: "fgl-topics/c_fgl_menus_005.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Syntax of the MENU instruction"
type: "concept"
---

# Syntax of the MENU instruction

> The MENU instruction defines a set of options the end user can select to trigger actions in a program.

## Syntax

```
MENU [title]
   [ {ATTRIBUTE|ATTRIBUTES} ( menu-attribute [,...] ) ]
   [ BEFORE MENU
       menu-statement
     [...]
   ]
   menu-option
   [...]
END MENU
```

where menu-option is one of:

```
{ COMMAND option-name
      [option-comment] [ HELP help-number ]
    menu-statement
    [...]
| COMMAND KEY ( key-name ) option-name
      [option-comment] [ HELP help-number ]
    menu-statement
    [...]
| COMMAND KEY ( key-name )
    menu-statement
    [...]
| ON ACTION action-name
             [ {ATTRIBUTE|ATTRIBUTES} ( action-attributes-menu ) ]
    menu-statement
    [...]
| ON IDLE seconds 
    menu-statement
    [...]
| ON TIMER seconds
    menu-statement
    [...]
}
```

where action-attributes-menu
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

where menu-statement is:

```
{ statement
|  CONTINUE MENU
|  EXIT MENU
|  NEXT OPTION option 
|  SHOW OPTION { ALL | option [,...] } 
|  HIDE OPTION { ALL | option [,...] } 
}
```

where menu-attribute is:

```
{ STYLE = { "default" | "popup" | "dialog" | "user-defined-style" }
| COMMENT = "string"
| IMAGE = "string"
}
```

1. title is a string expression defining the title of the menu.
2. menu-attribute is an attribute that defines the behavior and presentation
   of the menu.
3. key-name is a hot-key identifier (like `F11` or
   `Control-z`).
4. option-name is a string expression defining the label of the menu option
   and identifying the action that can be executed by the user.
5. option-comment is a string expression containing a description for the menu
   option, displayed when option-name is the current.
6. help-number is an integer that allows you to associate a help message
   number with the menu option.
7. action-name identifies an action that can be executed by the user.
8. seconds is an integer literal or variable that defines a number of seconds.
9. action-name identifies an action that can be executed by the user.
10. action-attributes are dialog-specific action attributes.
11. user-defined-style is a custom style name.
