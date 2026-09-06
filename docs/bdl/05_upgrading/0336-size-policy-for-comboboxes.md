---
title: "Size policy for ComboBoxes"
source: "fgl-topics/c_fgl_Migrate_to_130_combobox_sizepolicy.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 1.30 upgrade guide > Size policy for ComboBoxes"
type: "concept"
---

# Size policy for ComboBoxes

> You can use the SIZEPOLICY attribute for a COMBOBOX.

Starting with version 1.30 you can use the SIZEPOLICY attribute for COMBOBOXes.

`COMBOBOX` form items had a special behavior in versions prior to 1.30,
because they adapted their size to the maximum item of the value list. On one hand, this
is very convenient because the programmer doesn't have to find the biggest string in the
value list, and to estimate how large it will be on the screen (with proportional fonts
the string with the highest number of characters is not automatically the largest string).
On the other hand, this behavior often led to an unpredictable layout if the programmer
didn't reserve enough space for the `COMBOBOX`.

The `SIZEPOLICY` attribute gives better control
of the result.

```
<G "Combo makes edit2 too big"                     >
 [edit1]
 [combo   ]
 [edit2   ]
... 
ATTRIBUTES 
EDIT edit1=formonly.edit1; 
COMBOBOX combo=formonly.combo,
   ITEMS=((0,"Veeeeeeeery Loooooooooooooooong Item"),(1,"hallo")),
   DEFAULT=0; 
EDIT edit2=formonly.edit2; 
END
```

![SIZEPOLICY usage screenshot](../_images/NFCombo1.jpg)

*Use of SIZEPOLICY*

In this case, the "`combo"` field gets very large
as does "`edit2`", because it ends in the same grid
column. It will confuse the end user if he can input only eight characters
and the field is apparently much bigger. Two possibilities exist to
surround this:

Use an HBox to prevent the `edit2` from growing,
and use HBoxes for all fields which start together with `combo` and
are as large or bigger than `combo`:

```
 <G "Edit2 in HBox doesn't grow" >
 [edit1]
 [combo  :]
 [edit2  :]
 ...
```

![HBox usage screenshot](../_images/NFCombo2.jpg)

*Use of HBox*

Use the new `SIZEPOLICY` attribute, and set it to `fixed`
to prevent `combo` from getting bigger than the initial six characters
(6+Button):

```
 <G "Combo has a fixed size">
 ...
 [combo   ]
 [edit2   ]
 ... 
 ATTRIBUTES
 ...
 COMBOBOX combo=formonly.combo,
    ITEMS = ((0,"Veeeeeeeery Looooooooooooooooong Item"),(1,"hallo")),
    DEFAULT=0, SIZEPOLOCY=FIXED ;
....
```

![SIZEPOLICY as fixed screenshot](../_images/NFCombo3.jpg)

*Use of SIZEPOLICY as fixed*

In this example the `edit2` dictates the maximum size of `combo`,
because even if the SIZEPOLICY is *fixed*, the elements are aligned by the Grid.

To prevent this and have *exactly* six characters (numbers) in the ComboBox, you need
to de-couple `combo` from `edit2` by using an HBox.

```
<G "Combo has a fixed size,sample 0,in HBox"
... 
Combo [combo  :]
Edit2 [edit2  :]
... 
COMBOBOX combo=formonly.combo,
   ITEMS = ((0,"12345678 Looooooooooooooooong Item"),(1,"hallo")),
   DEFAULT=0, SIZEPOLICY=FIXED, SAMPLE="0";
```

![HBox usage screenshot](../_images/NFCombo4.jpg)

*Use of HBox*

Now the wanted six numbers are displayed and `combo` does not grow to the
size of `edit2`.

## Related links

**Related concepts**  

[SIZEPOLICY attribute](../11_user-interface/1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.")
