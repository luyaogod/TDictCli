---
title: "Example 1: Defining styles for grid elements"
source: "fgl-topics/c_fgl_presentation_styles_example_1.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Examples > Example 1: Defining styles for grid elements"
type: "concept"
description: "This example shows how to define styles for grid elements. The presentation style definition file: <?xml version=\"1.0\" encoding=\"ANSI_X3.4-1968\"?> <StyleList> <!-- Applies to all type of elements --> ..."
---

# Example 1: Defining styles for grid elements

This example shows how to define styles for grid elements.

The presentation style definition
file:

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
  <!-- Applies to all type of elements -->
  <Style name=".bigfont">
     <StyleAttribute name="fontSize" value="large" />
  </Style>
  <!-- Default text color and font family for all labels -->
  <Style name="Label">
     <StyleAttribute name="textColor" value="blue" />
     <StyleAttribute name="fontFamily" value="sans-serif" />
  </Style>
  <!-- Background color for Edits having focus -->
  <Style name="Edit:focus">
     <StyleAttribute name="backgroundColor" value="yellow" />
  </Style>
  <!-- Text color for Edits with STYLE="mandatory" -->
  <Style name="Edit.mandatory">
     <StyleAttribute name="textColor" value="red" />
  </Style>
</StyleList>
```

The form definition file:

```
LAYOUT
GRID
{
[l1    ][f1            ]
[l2    ][f2            ]
[l3    ][f3            ]
}
END
ATTRIBUTES
LABEL l1: TEXT="Label 1:";
EDIT  f1 = FORMONLY.field1;
LABEL l2: TEXT="Label 2:";
EDIT  f2 = FORMONLY.field2;
LABEL l3: TEXT="Label 3:", STYLE="bigfont";
EDIT  f3 = FORMONLY.field3, STYLE="bigfont mandatory";
END
```

Program source file:

```
MAIN
    DEFINE rec RECORD
                field1 STRING,
                field2 STRING,
                field3 STRING
           END RECORD

    LET rec.field1 = "Field 1"
    LET rec.field2 = "Field 2"
    LET rec.field3 = "Field 3"

    CALL ui.Interface.loadStyles("styles")

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    INPUT BY NAME rec.* WITHOUT DEFAULTS

END MAIN
```

Graphical result:

![Screenshot of form displayed with styles applied. Field 2 is yellow. Label 3 and Field 3 have larger font. Field 3 has red text.](../_images/present_styles_ex_1_gbc.jpg)

*Form displayed based on styles applied*

How the styles were applied:

1. All labels get a blue text color and sans-serif font family because of the
   `name="Label"` style.
2. Label 3 and Edit 3 defined with the `bigfont` style name get a large font because
   of the `name=".bigfont"` style.
3. The Edit field having the focus gets a yellow background color because of the
   `name="Edit:focus"` style (using the `focus` pseudo-selector).
4. Edit fields defined with the `mandatory` style name get a red text color because
   of the `name="Edit.mandatory"` style.

## Related links

**Related concepts**  

[Syntax of presentation styles file (.4st)](1609-syntax-of-presentation-styles-file-4st.md "A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.")
