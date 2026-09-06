---
title: "Example 1: MENU with abstract action options"
source: "fgl-topics/c_fgl_menus_example_1.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Examples > Example 1: MENU with abstract action options"
type: "concept"
description: "MAIN MENU ON ACTION new CALL newFile() ON ACTION open CALL openFile() ON ACTION save CALL saveFile() ON ACTION import LOAD FROM \"infile.dat\" INSERT INTO table ON ACTION quit EXIT PROGRAM END MENU END ..."
---

# Example 1: MENU with abstract action options

```
MAIN
  MENU
    ON ACTION new 
        CALL newFile()
      ON ACTION open 
        CALL openFile()
    ON ACTION save 
        CALL saveFile()
    ON ACTION import 
        LOAD FROM "infile.dat" INSERT INTO table 
    ON ACTION quit 
        EXIT PROGRAM
  END MENU
END MAIN
```
