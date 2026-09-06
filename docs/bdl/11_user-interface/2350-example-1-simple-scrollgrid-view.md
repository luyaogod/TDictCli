---
title: "Example 1: Simple scrollgrid view"
source: "fgl-topics/c_fgl_ui_scrollgrid_example_1.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Examples > Example 1: Simple scrollgrid view"
type: "concept"
description: "Figure: Form with paged scrollgrid The form file form.per : LAYOUT SCROLLGRID (WANTFIXEDPAGESIZE=NO, INITIALPAGESIZE=4, STYLE=\"paged\") { [f1 ] Id: [f2 ] Name:[f3 ] [ ] Details: [ ] [f4 ] [ ] [ ] } END ..."
---

# Example 1: Simple scrollgrid view

![Screenshot of form with scrollgrid using the tilelist option](../_images/scrollgrid_paged_1.jpg)

*Form with paged scrollgrid*

The form file form.per:

```
LAYOUT
SCROLLGRID (WANTFIXEDPAGESIZE=NO, INITIALPAGESIZE=4, STYLE="paged")
{
[f1   ] Id: [f2   ] Name:[f3           ]
[     ] Details:
[     ] [f4                            ]
[     ] [                              ]
}
END 
END 
ATTRIBUTES
EDIT  f2 = FORMONLY.key;
EDIT  f3 = FORMONLY.name;
IMAGE f1 = FORMONLY.image, SIZEPOLICY=FIXED, AUTOSCALE;
EDIT  f4 = FORMONLY.detail;
END 
INSTRUCTIONS
SCREEN RECORD list1(FORMONLY.*);
END
```

The presentation style file
contains:

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
    <Style name="Window">
        <StyleAttribute name="windowType" value="normal"/>
    </Style>
    <Style name="ScrollGrid.paged">
        <StyleAttribute name="customWidget" value="pagedScrollGrid" />
    </Style>
</StyleList>
```

The program main.4gl:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name STRING,
               image STRING,
               detail STRING
           END RECORD,
           x INTEGER
    CALL ui.Interface.loadStyles("mystyles")
    FOR x=1 TO 60
        LET arr[x].key = x
        LET arr[x].name = SFMT("Item %1", x)
        LET arr[x].image = IIF(x MOD 2,"file","folder")
        LET arr[x].detail = SFMT("This is item %1", x)
    END FOR
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    DISPLAY ARRAY arr TO list1.*
            ATTRIBUTES(UNBUFFERED,DOUBLECLICK=myselect)
        ON ACTION myselect
           MESSAGE "myselect:", arr_curr()
    END DISPLAY
END MAIN
```
