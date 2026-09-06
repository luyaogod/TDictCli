---
title: "Example 2: MENU with text-mode options"
source: "fgl-topics/c_fgl_menus_example_2.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Examples > Example 2: MENU with text-mode options"
type: "concept"
description: "MAIN MENU \"File\" COMMAND KEY ( CONTROL-N ) \"New\" \"Creates New File\" HELP 101 CALL newFile() COMMAND KEY ( CONTROL-O ) \"Open\" \"Open existing File\" HELP 102 CALL openFile() COMMAND KEY ( CONTROL-S ) ..."
---

# Example 2: MENU with text-mode options

```
MAIN
  MENU "File"
    COMMAND KEY ( CONTROL-N ) "New" "Creates New File" HELP 101
      CALL newFile()
    COMMAND KEY ( CONTROL-O ) "Open" "Open existing File" HELP 102
      CALL openFile()
    COMMAND KEY ( CONTROL-S ) "Save" "Save Current File" HELP 103
      CALL saveFile()
    COMMAND "Import"
      LOAD FROM "infile.dat" INSERT INTO table
    COMMAND KEY ( CONTROL-Q ) "Quit" "Quit Program" HELP 201
      EXIT PROGRAM
  END MENU
END MAIN
```
