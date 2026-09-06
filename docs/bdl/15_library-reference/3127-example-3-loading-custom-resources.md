---
title: "Example 3: Loading custom resources"
source: "fgl-topics/c_fgl_ClassInterface_example_4.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > Examples > Example 3: Loading custom resources"
type: "concept"
description: "Form file ( customer.per ): LAYOUT GRID { Num: [f1 ] Name: [f2 ] } END END ATTRIBUTES EDIT f1 = FORMONLY.cust_id; EDIT f2 = FORMONLY.cust_name, STYLE=\"mandatory\"; END Program file: MAIN DEFINE rec ..."
---

# Example 3: Loading custom resources

Form file
(customer.per):

```
LAYOUT
GRID
{
Num:  [f1       ]
Name: [f2                 ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.cust_id;
EDIT f2 = FORMONLY.cust_name, STYLE="mandatory";
END
```

Program file:

```
MAIN
    DEFINE rec RECORD
               cust_id INT,
               cust_name VARCHAR(50)
           END RECORD
    CALL ui.Interface.loadActionDefaults("myactdefs")
    CALL ui.Interface.loadStyles("mystyles")
    OPEN FORM f1 FROM "customer"
    DISPLAY FORM f1
    INPUT BY NAME rec.*
END MAIN
```

Styles file
(mystyles.4st):

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
  <Style name="Window">
     <StyleAttribute name="windowType" value="normal" />
  </Style>
  <Style name="Edit.mandatory">
     <StyleAttribute name="backgroundColor" value="lightRed" />
  </Style>
</StyleList>
```

Action Defaults file
(myactdefs.4ad):

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<ActionDefaultList>
  <ActionDefault name="accept" text="Accept" acceleratorName="Return" acceleratorName2="Enter" />
  <ActionDefault name="cancel" validate="no" text="Cancel" acceleratorName="Escape" />
  <ActionDefault name="dialogtouched" validate="no" defaultView="no" contextMenu="no"/>
</ActionDefaultList>
```
