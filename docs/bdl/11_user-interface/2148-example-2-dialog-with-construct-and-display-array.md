---
title: "Example 2: DIALOG with CONSTRUCT and DISPLAY ARRAY"
source: "fgl-topics/c_fgl_multiple_dialogs_example_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Examples > Example 2: DIALOG with CONSTRUCT and DISPLAY ARRAY"
type: "concept"
description: "Form file \" form1.per \": LAYOUT GRID { <g g1 > Name: [f1 ] State: [f2 ] City: [f3 ] Zip-code: [f4 ] [ :cc :sr ] < > <g g2 > <t t1 > Id Name [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] < > < > [ :cw ] } END END ..."
---

# Example 2: DIALOG with CONSTRUCT and DISPLAY ARRAY

Form file
"form1.per":

```
LAYOUT
GRID
{
<g g1                                >
 Name:     [f1                       ]
 State:    [f2                       ]
 City:     [f3                       ]
 Zip-code: [f4                       ]
[                      :cc    :sr    ]
<                                    >
<g g2                                >
 <t t1                              >
  Id        Name
 [c1       |c2                      ]
 [c1       |c2                      ]
 [c1       |c2                      ]
 <                                  >
<                                    >
[                              :cw    ]
}
END
END

ATTRIBUTES
GROUP g1: TEXT = "Search criteria";
EDIT f1 = FORMONLY.cust_name TYPE VARCHAR;
EDIT f2 = FORMONLY.cust_state TYPE VARCHAR;
EDIT f3 = FORMONLY.cust_city TYPE VARCHAR;
EDIT f4 = FORMONLY.cust_zipcode TYPE VARCHAR;
BUTTON cc: clear, TEXT="Clear";
BUTTON sr: fetch, TEXT="Fetch";
GROUP g2: TEXT = "Customer list";
EDIT c1 = FORMONLY.c_id TYPE INTEGER;
EDIT c2 = FORMONLY.c_name TYPE VARCHAR;
BUTTON cw: close;
END

INSTRUCTIONS
SCREEN RECORD sr (FORMONLY.c_id, FORMONLY.c_name);
END
```

Program source code:

```
MAIN
  DEFINE custarr DYNAMIC ARRAY OF RECORD
            c_id INTEGER,
            c_name VARCHAR(50)
          END RECORD
  DEFINE where_clause STRING

  OPTIONS INPUT WRAP 

  OPEN FORM f1 FROM "form1"
  DISPLAY FORM f1

  DIALOG ATTRIBUTES(FIELD ORDER FORM, UNBUFFERED)

     CONSTRUCT BY NAME where_clause 
          ON cust_name, cust_state, cust_city, cust_zipcode 
        ON ACTION clear 
          CLEAR cust_name, cust_state, cust_city, cust_zipcode 
     END CONSTRUCT

     DISPLAY ARRAY custarr TO sr.*
       BEFORE ROW
         MESSAGE SFMT("Row: %1/%2", DIALOG.getCurrentRow("sr"),
                                DIALOG.getArrayLength("sr"))
     END DISPLAY

     ON ACTION fetch 
        MESSAGE "Where:", where_clause 
        -- Execute SQL query here to fill custarr ...

     ON ACTION close 
        EXIT DIALOG

  END DIALOG

END MAIN
```
