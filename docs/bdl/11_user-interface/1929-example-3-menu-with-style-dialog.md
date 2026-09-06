---
title: "Example 3: MENU with STYLE=\"dialog\""
source: "fgl-topics/c_fgl_menus_example_3.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Examples > Example 3: MENU with STYLE=\"dialog\""
type: "concept"
description: "This code example implements typical message box utility functions implemented with MENU dialogs: FUNCTION mbox_ync(title,msg) DEFINE title, msg STRING DEFINE res SMALLINT MENU title ..."
---

# Example 3: MENU with STYLE="dialog"

This code example implements typical message box utility functions implemented with
`MENU` dialogs:

```
FUNCTION mbox_ync(title,msg)
    DEFINE title, msg STRING
    DEFINE res SMALLINT
    MENU title ATTRIBUTES(STYLE="dialog",COMMENT=msg)
        ON ACTION yes     LET res = 1
        ON ACTION no      LET res = 0
        ON ACTION cancel  LET res = -1
    END MENU
    RETURN res
END FUNCTION

FUNCTION mbox_yn(title,msg)
    DEFINE title, msg STRING
    DEFINE res BOOLEAN
    MENU title ATTRIBUTES(STYLE="dialog",COMMENT=msg)
        ON ACTION yes LET res = TRUE
        ON ACTION no  LET res = FALSE
    END MENU
    RETURN res
END FUNCTION

FUNCTION mbox_ok(title,msg)
    DEFINE title, msg STRING
    MENU title ATTRIBUTES(STYLE="dialog",COMMENT=msg)
        ON ACTION accept
    END MENU
END FUNCTION
```
