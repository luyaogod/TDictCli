---
title: "Example 1: DIALOG controlling two lists"
source: "fgl-topics/c_fgl_multiple_dialogs_example_1.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Examples > Example 1: DIALOG controlling two lists"
type: "concept"
description: "Form file \" lists.per \": LAYOUT GRID { <t t1 > [f11 |f12 ] < > <t t2 > [f21 |f22 ] < > } END END ATTRIBUTES EDIT f11 = FORMONLY.column_11; EDIT f12 = FORMONLY.column_12; EDIT f21 = FORMONLY.column_21; ..."
---

# Example 1: DIALOG controlling two lists

Form file
"lists.per":

```
LAYOUT
GRID
{
<t t1              >
[f11  |f12         ]
<                  >
<t t2              >
[f21  |f22         ]
<                  >
}
END
END
ATTRIBUTES
EDIT f11 = FORMONLY.column_11;
EDIT f12 = FORMONLY.column_12;
EDIT f21 = FORMONLY.column_21;
EDIT f22 = FORMONLY.column_22;
END
INSTRUCTIONS
SCREEN RECORD sr1(FORMONLY.column_11,FORMONLY.column_12);
SCREEN RECORD sr2(FORMONLY.column_21,FORMONLY.column_22);
END
```

Program source code:

```
DEFINE
   arr1 DYNAMIC ARRAY OF RECORD 
        column_11 INTEGER,
        column_12 VARCHAR(10)
   END RECORD,
   arr2 DYNAMIC ARRAY OF RECORD 
        column_21 INTEGER,
        column_22 VARCHAR(10)
   END RECORD

MAIN
  DEFINE i INTEGER 
  FOR i = 1 TO 20
    LET arr1[i].column_11 = i 
    LET arr1[i].column_12 = "aaa "||i 
    LET arr2[i].column_21 = i 
    LET arr2[i].column_22 = "aaa "||i 
  END FOR
  OPTIONS INPUT WRAP
  OPEN FORM f FROM "lists"
  DISPLAY FORM f 
  DIALOG ATTRIBUTES(UNBUFFERED)
    DISPLAY ARRAY arr1 TO sr1.*
      BEFORE DISPLAY
        MESSAGE "We are in list one"
    END DISPLAY
    DISPLAY ARRAY arr2 TO sr2.*
      BEFORE DISPLAY
        MESSAGE "We are in list two"
    END DISPLAY
    ON ACTION close 
      EXIT DIALOG
  END DIALOG
END MAIN
```
