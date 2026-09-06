---
title: "Example 1: Help message file used in a MENU"
source: "fgl-topics/c_fgl_message_files_007.html"
breadcrumb: "User interface > Form definitions > Message files > Examples > Example 1: Help message file used in a MENU"
type: "concept"
description: "The message source file help.msg : .101 This is help about option 1 .102 This is help about help .103 This is help about My Menu Compiling the message file: $ fglmkmsg help.msg Program using the .iem ..."
---

# Example 1: Help message file used in a MENU

The message source file help.msg:

```
.101
This is help about option 1
.102
This is help about help 
.103
This is help about My Menu
```

Compiling the message file:

```
$ fglmkmsg help.msg
```

Program using the .iem compiled message file.

```
MAIN
    OPTIONS
        HELP FILE "help.iem"
    MENU "Sample"
        COMMAND "Option 1" HELP 101
            DISPLAY "Option 1 chosen"
        COMMAND "Help"
            CALL showhelp(103)
    END MENU
END MAIN
```
