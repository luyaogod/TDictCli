---
title: "Example 3: DIALOG with SUBDIALOG"
source: "fgl-topics/c_fgl_multiple_dialogs_example_3.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Examples > Example 3: DIALOG with SUBDIALOG"
type: "concept"
description: "Form file \" comment.per \": LAYOUT GRID { [cmt ] [ ] [ ] [cck ] } END END ATTRIBUTES TEXTEDIT cmt = FORMONLY.the_comment, STRETCH=BOTH; CHECKBOX cck = FORMONLY.is_checked, NOT NULL, TEXT = \"Verified\", ..."
---

# Example 3: DIALOG with SUBDIALOG

Form file "comment.per":

```
LAYOUT
GRID
{
[cmt                              ]
[                                 ]
[                                 ]
                [cck              ]
}
END
END
ATTRIBUTES
TEXTEDIT cmt = FORMONLY.the_comment, STRETCH=BOTH;
CHECKBOX cck = FORMONLY.is_checked, NOT NULL,
  TEXT = "Verified",
  VALUECHECKED=TRUE, VALUEUNCHECKED=FALSE;
END
```

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

Form file "form1.per":

```
LAYOUT
VBOX
GRID
{
Id:   [f1        ]
Name: [f2                                        ]
}
END
FORM "comment"
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.cust_id TYPE INTEGER;
EDIT f2 = FORMONLY.cust_name TYPE VARCHAR;
END
```

Program source code:

```
IMPORT FGL comment

DEFINE cust RECORD
            cust_id INTEGER,
            cust_name VARCHAR(50)
    END RECORD

DEFINE cust_comment comment.t_comment

MAIN

    OPTIONS INPUT WRAP

    OPEN FORM f1 FROM "form1"
    DISPLAY FORM f1

    LET cust.cust_id = 199
    LET cust.cust_name = "Mike Patterson"
    LET cust_comment.the_comment = "Some comment about this customer..."
    LET cust_comment.is_checked = FALSE

    DIALOG ATTRIBUTES(FIELD ORDER FORM, UNBUFFERED)
        INPUT BY NAME cust.* ATTRIBUTES(WITHOUT DEFAULTS)
        END INPUT
        SUBDIALOG comment_input(cust_comment, TRUE, (FUNCTION event_callback))
        ON ACTION quit ATTRIBUTES(TEXT="Quit")
           EXIT DIALOG
    END DIALOG

END MAIN

FUNCTION event_callback(event SMALLINT)
    DEFINE msg STRING
    CASE event
      WHEN comment.c_comment_changed
         LET msg = "Comment changed..."
      WHEN comment.c_comment_checked
         LET msg = "Comment checked..."
    END CASE
    MESSAGE SFMT("%1: %2",CURRENT HOUR TO FRACTION(3),msg)
END FUNCTION
```
