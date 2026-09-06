---
title: "Example 1: Comment sub-dialog"
source: "fgl-topics/c_fgl_declarative_dialogs_example_1.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Examples > Example 1: Comment sub-dialog"
type: "concept"
description: "The module \" comment.4gl \": PUBLIC TYPE t_comment RECORD the_comment VARCHAR(200), is_checked BOOLEAN END RECORD PUBLIC TYPE t_event_callback FUNCTION (event SMALLINT) PRIVATE DEFINE _event_callback ..."
---

# Example 1: Comment sub-dialog

The module "comment.4gl":

```
PUBLIC TYPE t_comment RECORD
        the_comment VARCHAR(200),
        is_checked BOOLEAN
    END RECORD

PUBLIC TYPE t_event_callback FUNCTION (event SMALLINT)
PRIVATE DEFINE _event_callback t_event_callback
PUBLIC CONSTANT c_comment_changed = 101
PUBLIC CONSTANT c_comment_checked = 102

DIALOG comment_input(
    cr t_comment INOUT,
    can_clear STRING,
    event_callback t_event_callback
)
    INPUT BY NAME cr.* ATTRIBUTES(WITHOUT DEFAULTS)
        BEFORE INPUT
           CALL DIALOG.setActionActive("clear",can_clear)
           LET _event_callback = event_callback
        ON ACTION clear ATTRIBUTES(TEXT="Clear")
           LET cr.the_comment = NULL
           LET cr.is_checked = FALSE
           CALL _signal_parent(c_comment_changed)
        ON CHANGE the_comment
           CALL _signal_parent(c_comment_changed)
        ON CHANGE is_checked
           CALL _signal_parent(c_comment_checked)
    END INPUT
END DIALOG

PRIVATE FUNCTION _signal_parent(ne SMALLINT) RETURNS ()
    IF _event_callback IS NOT NULL THEN
        CALL _event_callback(ne)
    END IF
END FUNCTION
```

For the complete example, see [Example 3: DIALOG with SUBDIALOG](2149-example-3-dialog-with-subdialog.md).
