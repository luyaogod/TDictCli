---
title: "Example 1: Simple table view"
source: "fgl-topics/c_fgl_ui_tables_example_1.html"
breadcrumb: "User interface > User interface programming > Table views > Examples > Example 1: Simple table view"
type: "concept"
description: "Figure: Form with simple table The form file form.per : LAYOUT TABLE { [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] [c1 |c2 ] } END END ATTRIBUTES PHANTOM FORMONLY.key; c1 = FORMONLY.name, TITLE=\"Name\", ..."
---

# Example 1: Simple table view

![Screenshot of form with table using two columns and image](../_images/table_simple_list_1.jpg)

*Form with simple table*

The form file form.per:

```
LAYOUT
TABLE
{
[c1         |c2                                      ]
[c1         |c2                                      ]
[c1         |c2                                      ]
[c1         |c2                                      ]
}
END
END
ATTRIBUTES
PHANTOM FORMONLY.key;
c1 = FORMONLY.name, TITLE="Name", IMAGECOLUMN=image;
PHANTOM FORMONLY.image;
c2 = FORMONLY.detail, TITLE="Detail";
END
INSTRUCTIONS
SCREEN RECORD list1(FORMONLY.*);
END
```

The program code main.4gl:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name STRING,
               image STRING,
               detail STRING
           END RECORD,
           x INTEGER
    FOR x=1 TO 60
        LET arr[x].key = x
        LET arr[x].name = SFMT("Item %1", x)
        LET arr[x].image = IIF(x MOD 2, "file", "folder")
        LET arr[x].detail = SFMT("This is item %1", x)
    END FOR
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    DISPLAY ARRAY arr TO list1.* ATTRIBUTES(UNBUFFERED,DOUBLECLICK=myselect)
        ON ACTION myselect
           MESSAGE "myselect:", arr_curr()
    END DISPLAY
END MAIN
```

## Related links

**Related concepts**  

[Runtime images](1588-runtime-images.md "Explains how to display pictures at runtime.")
