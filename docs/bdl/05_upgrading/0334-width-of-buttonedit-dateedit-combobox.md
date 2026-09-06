---
title: "Width of ButtonEdit/DateEdit/ComboBox"
source: "fgl-topics/c_fgl_Migrate_to_130_fieldbutton_sizes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 1.30 upgrade guide > Width of ButtonEdit/DateEdit/ComboBox"
type: "concept"
---

# Width of ButtonEdit/DateEdit/ComboBox

> When using BUTTONEDIT/COMBOBOX/DATEEDIT fields, it is recommended that you account for the width of the widget button in addition to the input area.

The problem with BUTTONEDIT, DATEEDIT and COMBOBOX in versions prior to 1.30 is that a field
`[b ]` got the width 3, the same width as an edit field with the same
layout.

For example:

```
LAYOUT 
GRID
{
  [e  ]
  [b  ]
} 
END 
END 
ATTRIBUTES 
EDIT e=formonly.e; 
BUTTONEDIT b=formonly.b; 
END
```

In this example, the outer (visual) width of both elements was the same, but the edit portion of
"`b`" was much smaller, because the button did not count at all. (In practice
this meant that on average only one and a half characters of "`b`" was
visible). However, you could input 3 characters! This resulted in a
`BUTTONEDIT` where only one character was visible and inputting more than one
character was possible.

Starting with version 1.30, for the Button, the Form Compiler subtracts two character positions
from the width of
`BUTTONEDIT`/`COMBOBOX`/`DATEEDIT`. This is
possible because now the form compiler differentiates the width of the widget from the width
of the entry part.

In fact, there is no visual difference between version 1.20 and 1.30 regarding this example, but
in version 1.30 you can only enter one character, which is visually more correct.

In the example the `BUTTONEDIT` aligns with the
Edit; that's why the Edit part of the `BUTTONEDIT` is
usually still a bit bigger than one character (this depends
on the button size, but if a button edit is contained by an HBox,
it will get the exact size of "width" multiplied by the average
character pixel width.

To express the `BUTTONEDIT`/`COMBOBOX`/`DATEEDIT` layout
more visually, it is possible to specify:

```
  [e  ]
  [b- ]
```

the "`-`" sign marks the end of the edit portion
and the beginning of the button portion ( edit width ="1", widget
width ="3" ).

The two characters are also subtracted for a `BUTTONEDIT` which
is child of an HBox.

```
   [b  :]
```

gets also width="1" , but no widget width, because the HBox stacks
the elements horizontally without needing widget width definition.

The two extra characters are only used to show the real size relations
more WYSIWYG, and to have the same calculation as in a field without
an HBox parent.

```
  [e1:e2:e3:     ]
  [b1  :b2  :b3  ]
```

shows that three `BUTTONEDIT` fields are much larger
than three `EDIT` fields with the same width.

You can even write:

```
  [e1:e2:e3:     ]
  [b1- :b2- :b3- ]
```

or:

```
  [e1:e2:e3:  ]
  [b1-:b2-:b3-]
```

to use slim buttons and

```
  [e1:e2:e3:        ]
  [b1-  :b2-  :b3-  ]
```

if one uses large buttons to get the maximum WYSIWYG effect.

Please note that buttons do not grow if two characters "`-`  " are expanded to
three characters "`-` "; the button always computes its size from the image
used, it's just reserves more space in the form to match the real size.

## Related links

**Related concepts**  

[Form fields default sample](0335-form-fields-default-sample.md "An algorithm is used to compute the field width when no SAMPLE attribute is specified.")
