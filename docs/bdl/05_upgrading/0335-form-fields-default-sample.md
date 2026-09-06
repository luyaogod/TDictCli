---
title: "Form fields default sample"
source: "fgl-topics/c_fgl_Migrate_to_130_default_sample.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 1.30 upgrade guide > Form fields default sample"
type: "concept"
---

# Form fields default sample

> An algorithm is used to compute the field width when no SAMPLE attribute is specified.

Starting with version 1.30, if no `SAMPLE` attribute is specified in the form
files, the client uses an algorithm to compute the field width. In this case, a very pessimistic
algorithm is used to compute the field widths: The client assumes a default `SAMPLE`
of "M" for the first six characters and then "0" for the subsequent characters and applies this
algorithm to all fields, with some exceptions like `DATEEDIT` fields.

The default algorithm tends to produce larger forms compared to forms used in BDS V3 and very
first versions of Genero. Do not hesitate to modify the `SAMPLE` attribute in the
form file, to make your fields shorter.

If you do not want to touch all your forms, a more tailored automatic solution would be to specify
a `ui.form.setDefaultInitializer()` function, to set the `SAMPLE`
depending on the AUI tag. In this example small `UPSHIFT` fields get a sample of "M";
all other fields get a sample of "0". This will preserve the original width for
`UPSHIFT` fields. However, numeric and normal String fields will get the sample of "0"
and make the overall width of the form smaller.

Program:

```
# this demo program shows how to affect the "sample" attribute in a 
#  ui.form.setDefaultInitializer function 
# the main concern is to set a default sample of "0" and to 
# correct the sample attribute for small UPSHIFT fields to "M" 
# to be able to display full uppercase letter for fields with a small width
 
MAIN
  DEFINE three_char_upshift CHAR(3)
  DEFINE three_digit_number Integer
  DEFINE longstring CHAR(100)
  CALL ui.form.setDefaultInitializer("myinit")
  OPEN form f from "sampletest2"
  DISPLAY form f
  INPUT BY NAME three_char_upshift,three_digit_number,longstring 
END MAIN
 
FUNCTION myInit(f)
  DEFINE f ui.Form
  CALL checkSampleRecursive(f.getNode()) 
END FUNCTION
 
FUNCTION checkSampleRecursive(node)
  DEFINE node,child om.DomNode
  LET child= node.getFirstChild()
  WHILE child IS NOT NULL
    CALL checkSampleRecursive(child)
    CALL setSample(child)
    LET child=child.getNext()
  END WHILE 
END FUNCTION
 
FUNCTION setSample(node)
  DEFINE node,parent om.DomNode
  LET parent=node.getParent()
  -- only set the "sample" for FormFields in this example
  IF parent.getTagName()<>"FormField" THEN
    RETURN
  END IF
  IF node.getAttribute("shift")="up"
    AND node.getAttribute("width")<=6 THEN
    CALL node.setAttribute("sample","M")
  ELSE
    CALL node.setAttribute("sample","0")
  END IF
  DISPLAY "set sample attribute of ",node.getId()," to \"",
    node.getAttribute("sample"),"\"" 
END FUNCTION
```

Form File:

```
LAYOUT(text="sampletest2") 
GRID
{
 <G sampletest                                                                 >
  3 Letter Code: [a  ] 3 digit code:[b  ]  Description:[longstring ]

 <G "What can be seen"                                                         >
  There is no default sample set in this form, but due to a
  ui.form.setDefaultInitializer function, small UPSHIFT fields
  are adjusted to a sample of "M", all other fields get the sample "0"
  -----------------------------------------------------------------------------
  1. The 3 letter code should show up exactly "MMM" because of the applied
     sample="M"
  2. The 3 letter digit code should show up exactly "123" without additional  
     spacing
} 
END 
END 
ATTRIBUTES 
EDIT a=formonly.three_char_upshift,UPSHIFT,default="MMM"; 
EDIT b=formonly.three_digit_number,default="123"; 
EDIT longstring=formonly.longstring,UPSHIFT,
     default="DESCRIPTION OF THE ITEM",SCROLL; 
END
```

![Sample usage in form screenshot](../_images/NFSampleTest.jpg)

*Sample usage in form*

## Related links

**Related concepts**  

[SAMPLE attribute](../11_user-interface/1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.")
