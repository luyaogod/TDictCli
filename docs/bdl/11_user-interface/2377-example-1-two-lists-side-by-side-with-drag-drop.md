---
title: "Example 1: Two lists side-by-side with drag & drop"
source: "fgl-topics/c_fgl_DragDrop_example_1.html"
breadcrumb: "User interface > User interface programming > Drag & drop > Examples > Example 1: Two lists side-by-side with drag & drop"
type: "concept"
description: "Form file: LAYOUT GRID { <t left > <t right> Left Right [a1 ] [a2 ] [a1 ] [a2 ] [a1 ] [a2 ] [a1 ] [a2 ] [s1 ][s2 ] } END END ATTRIBUTES a1 = FORMONLY.a1; a2 = FORMONLY.a2; s1 = FORMONLY.s1; s2 = ..."
---

# Example 1: Two lists side-by-side with drag & drop

Form
file:

```
LAYOUT
GRID
{
 <t left     >  <t     right>
 Left           Right
 [a1         ]  [a2         ]
 [a1         ]  [a2         ]
 [a1         ]  [a2         ]
 [a1         ]  [a2         ]
[s1           ][s2           ] 

}
END
END
ATTRIBUTES
a1 = FORMONLY.a1;
a2 = FORMONLY.a2;
s1 = FORMONLY.s1;
s2 = FORMONLY.s2;
END
INSTRUCTIONS
SCREEN RECORD sr_left(a1);
SCREEN RECORD sr_right(a2);
END
```

Program code:

```
MAIN
    DEFINE drag_index, drop_index, i INT
    DEFINE drag_source, drag_value STRING
    DEFINE arr_left, arr_right DYNAMIC ARRAY OF STRING
    DEFINE dnd ui.DragDrop 
    CONSTANT S_LEFT="sr_left"
    CONSTANT S_RIGHT="sr_right"
 
    OPEN FORM f FROM "dnd"
    DISPLAY FORM f 
 
    FOR i = 1 TO 10
       LET arr_left[i] = "left " || i 
       LET arr_right[i] = "right" || i 
    END FOR
 
    INITIALIZE drag_index TO NULL
 
    DIALOG ATTRIBUTES(UNBUFFERED)
 
        DISPLAY ARRAY arr_left TO sr_left.*
        ON DRAG_START(dnd)
            LET drag_source = S_LEFT
            LET drag_index = arr_curr()
            LET drag_value = arr_left[drag_index]
        ON DRAG_FINISHED(dnd)
            INITIALIZE drag_source TO NULL
        ON DRAG_ENTER(dnd)
            IF drag_source IS NULL THEN
                CALL dnd.setOperation(NULL)
            END IF
        ON DROP(dnd)
            IF drag_source == S_LEFT THEN
                CALL dnd.dropInternal()
            ELSE
                LET drop_index = dnd.getLocationRow()
                CALL DIALOG.insertRow(S_LEFT, drop_index)
                CALL DIALOG.setCurrentRow(S_LEFT, drop_index)
                LET arr_left[drop_index] = drag_value 
                CALL DIALOG.deleteRow(S_RIGHT, drag_index)
            END IF
        END DISPLAY
 
        DISPLAY ARRAY arr_right TO sr_right.*
        ON DRAG_START(dnd)
            LET drag_source = S_RIGHT
            LET drag_index = arr_curr()
            LET drag_value = arr_right[drag_index]
        ON DRAG_FINISHED(dnd)
            INITIALIZE drag_source TO NULL
        ON DRAG_ENTER(dnd)
            IF drag_source IS NULL THEN
                CALL dnd.setOperation(NULL)
            END IF
        ON DROP(dnd)
            IF drag_source == S_RIGHT THEN
                CALL dnd.dropInternal()
            ELSE
                LET drop_index = dnd.getLocationRow()
                CALL DIALOG.insertRow(S_RIGHT, drop_index)
                CALL DIALOG.setCurrentRow(S_RIGHT, drop_index)
                LET arr_right[drop_index] = drag_value 
                CALL DIALOG.deleteRow(S_LEFT, drag_index)
            END IF
        END DISPLAY
 
    ON ACTION cancel 
        EXIT DIALOG
 
    END DIALOG
END MAIN
```
