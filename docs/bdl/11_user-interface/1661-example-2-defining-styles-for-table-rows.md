---
title: "Example 2: Defining styles for table rows"
source: "fgl-topics/c_fgl_presentation_styles_example_2.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Examples > Example 2: Defining styles for table rows"
type: "concept"
description: "This example shows how to define styles for tables and table rows. The presentation style definition file: <?xml version=\"1.0\" encoding=\"ANSI_X3.4-1968\"?> <StyleList> <!-- Applies to all type of ..."
---

# Example 2: Defining styles for table rows

This example shows how to define styles for tables and table rows.

The presentation style definition
file:

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
  <!-- Applies to all type of elements -->
  <Style name=".bigfont">
     <StyleAttribute name="fontSize" value="large" />
  </Style>
  <!-- Background color form odd rows in tables -->
  <Style name="Table:odd">
     <StyleAttribute name="backgroundColor" value="yellow" />
  </Style>
</StyleList>
```

The form definition
file:

```
LAYOUT
TABLE
{
[c1   |c2             |c3                  ]
[c1   |c2             |c3                  ]
[c1   |c2             |c3                  ]
[c1   |c2             |c3                  ]
}
END
ATTRIBUTES
EDIT c1 = FORMONLY.col1, TITLE="C1";
EDIT c2 = FORMONLY.col2, TITLE="C2";
EDIT c3 = FORMONLY.col3, TITLE="C3", STYLE="bigfont";
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

Program source file:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
                col1 INTEGER,
                col2 STRING,
                col3 STRING
           END RECORD,
           i INTEGER

    FOR i=1 TO 20
        LET arr[i].col1 = i
        LET arr[i].col2 = "Item #"||i
        LET arr[i].col3 = IIF(i MOD 2, "odd", "even")
    END FOR

    CALL ui.Interface.loadStyles("styles")

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    DISPLAY ARRAY arr TO sr.*

END MAIN
```

Graphical result:

![Screenshot of form displayed with styles applied. Odd rows are yellow. Column 3 has larger font.](../_images/present_styles_ex_2_gbc.jpg)

*Form displayed based on styles applied*

How the styles were applied:

1. The odd rows get a yellow background because of the `name="Table:odd"` style
   (using the `odd` pseudo-selector).
2. Column 3 defined with the `bigfont` style name gets a large font because of
   the `name=".bigfont"` style.

## Related links

**Related concepts**  

[Pseudo selectors](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.")

[Syntax of presentation styles file (.4st)](1609-syntax-of-presentation-styles-file-4st.md "A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.")
