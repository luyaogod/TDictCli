---
title: "Example 2: Scrollgrid as listview"
source: "fgl-topics/c_fgl_ui_scrollgrid_example_2.html"
breadcrumb: "User interface > User interface programming > Scrollgrid views > Examples > Example 2: Scrollgrid as listview"
type: "concept"
description: "Figure: Form with scrollgrid using two data columns and an image The form file listview2bi.per : LAYOUT (TEXT=\"List view with 2 columns\") SCROLLGRID (STYLE=\"list\", WANTFIXEDPAGESIZE=NO) { [i01][f01 ] ..."
---

# Example 2: Scrollgrid as listview

![Screenshot of form with scrollgrid with listview layout](../_images/scrollgrid_listview_1.jpg)

*Form with scrollgrid using two data columns and an image*

The form file
listview2bi.per:

```
LAYOUT (TEXT="List view with 2 columns")
SCROLLGRID (STYLE="list", WANTFIXEDPAGESIZE=NO)
{
[i01][f01                       ]
[   ][f02                      ]
}
END
END
ATTRIBUTES
EDIT f01 = FORMONLY.name, STRETCH=X, STYLE="field1";
LABEL f02 = FORMONLY.id, STRETCH=X, STYLE="field2";
IMAGE i01 = FORMONLY.img, AUTOSCALE, STYLE="noborder";
END
INSTRUCTIONS
SCREEN RECORD scr(FORMONLY.*);
END
```

The presentation style file listview2bi.4st
contains:

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
    <Style name="Window">
        <StyleAttribute name="windowType" value="normal"/>
    </Style>
    <Style name="ScrollGrid.list">
        <StyleAttribute name="highlightCurrentRow" value="no"/>
    </Style>
    <Style name=".field1">
        <StyleAttribute name="fontSize" value="1.2em"/>
        <StyleAttribute name="fontWeight" value="bold"/>
    </Style>
    <Style name=".field2">
        <StyleAttribute name="fontSize" value="0.8em"/>
    </Style>
    <Style name="Image.noborder">
        <StyleAttribute name="border" value="none"/>
    </Style>
</StyleList>
```

The program listview2bi.4gl:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
        name STRING,
        id INTEGER,
        img STRING
    END RECORD
    DEFINE x INTEGER
    CALL ui.Interface.loadStyles("listview2bi")
    FOR x = 1 TO 100
        LET arr[x].name = SFMT("Item #%1", x)
        LET arr[x].id = 1000 + x
        LET arr[x].img = IIF(x MOD 2,"folder","file")
    END FOR
    OPEN FORM f1 FROM "listview2bi"
    DISPLAY FORM f1
    DISPLAY ARRAY arr TO scr.*
END MAIN
```
