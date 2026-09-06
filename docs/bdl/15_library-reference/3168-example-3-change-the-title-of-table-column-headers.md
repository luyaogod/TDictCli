---
title: "Example 3: Change the title of table column headers"
source: "fgl-topics/c_fgl_ClassForm_example_3.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > Examples > Example 3: Change the title of table column headers"
type: "concept"
description: "The form file ( coltitle.per ): LAYOUT GRID { <TABLE t1 > Id Name [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] } END END ATTRIBUTES c1 = FORMONLY.col1; c2 = FORMONLY.col2; END INSTRUCTIONS SCREEN ..."
---

# Example 3: Change the title of table column headers

The form file (coltitle.per):

```
LAYOUT
GRID
{
<TABLE t1             >
 Id      Name
[c1     |c2           ]
[c1     |c2           ]
[c1     |c2           ]
[c1     |c2           ]
[c1     |c2           ]
}
END
END
ATTRIBUTES
c1 = FORMONLY.col1;
c2 = FORMONLY.col2;
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

The program file:

```
MAIN
    DEFINE f ui.Form, i INT
    DEFINE arr DYNAMIC ARRAY OF RECORD
                   id INT,
                   name VARCHAR(40)
               END RECORD
    OPEN FORM f1 FROM "coltitle"
    DISPLAY FORM f1
    FOR i=1 TO 10
        LET arr[i].id = i
        LET arr[i].name = "aaa"||i
    END FOR
    DISPLAY ARRAY arr TO sr.* ATTRIBUTES(UNBUFFERED)
        BEFORE DISPLAY
           let f = dialog.getForm()
        ON ACTION change_title
           CALL f.setElementText("formonly.col1","ID")
           CALL f.setElementText("formonly.col2","NAME")
    END DISPLAY
END MAIN
```
