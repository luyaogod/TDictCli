---
title: "Example Using the INITIALIZER attribute in the form file"
source: "fgl-topics/c_fgl_ClassCombo_example_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > Examples > Example Using the INITIALIZER attribute in the form file"
type: "concept"
description: "Form Specification File: LAYOUT GRID { Airport: [cb01 ] } END END ATTRIBUTES COMBOBOX cb01 = FORMONLY.airport TYPE CHAR, INITIALIZER = initcombobox; END Initialization function: FUNCTION ..."
---

# Example Using the INITIALIZER attribute in the form file

Form Specification File:

```
LAYOUT
GRID
{
  Airport:  [cb01                 ]
}
END
END
ATTRIBUTES
COMBOBOX cb01 = FORMONLY.airport TYPE CHAR, INITIALIZER = initcombobox;
END
```

Initialization function:

```
FUNCTION initcombobox(comboBox ui.ComboBox) RETURNS ()
  CALL comboBox.clear()
  CALL comboBox.addItem("CDG", "Paris-Charles de Gaulle, France")
  CALL comboBox.addItem("LCY", "London-City Airport, UK")
  CALL comboBox.addItem("LHR", "London-Heathrow, UK")
  CALL comboBox.addItem("FRA", "Frankfurt Airport, Germany")
  CALL comboBox.addItem("SFO", "San Francisco International Airport, CA" )
END FUNCTION
```

## Related links

**Related concepts**  

[Filling a COMBOBOX item list](../11_user-interface/2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.")
