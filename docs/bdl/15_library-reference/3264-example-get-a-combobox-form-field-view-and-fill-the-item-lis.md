---
title: "Example Get a ComboBox form field view and fill the item list"
source: "fgl-topics/c_fgl_ClassCombo_example_1.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > Examples > Example Get a ComboBox form field view and fill the item list"
type: "concept"
description: "Form Specification File: LAYOUT GRID { Airport: [cb01 ] } END END ATTRIBUTES COMBOBOX cb01 = FORMONLY.airport TYPE CHAR; END Program File: MAIN DEFINE cb ui.ComboBox DEFINE airport CHAR(3) OPEN FORM ..."
---

# Example Get a ComboBox form field view and fill the item list

Form Specification
File:

```
LAYOUT
GRID
{
  Airport:  [cb01                 ]
}
END
END
ATTRIBUTES
COMBOBOX cb01 = FORMONLY.airport TYPE CHAR;
END
```

Program
File:

```
MAIN
  DEFINE cb ui.ComboBox
  DEFINE airport CHAR(3)
  OPEN FORM f1 FROM "combobox"
  DISPLAY FORM f1
  LET cb = ui.ComboBox.forName("formonly.airport")
  IF cb IS NULL THEN
     ERROR "Form field not found in current form"
     EXIT PROGRAM
  END IF
  CALL cb.clear()
  CALL cb.addItem("CDG", "Paris-Charles de Gaulle, France")
  CALL cb.addItem("LCY", "London-City Airport, UK")
  CALL cb.addItem("LHR", "London-Heathrow, UK")
  CALL cb.addItem("FRA", "Frankfurt Airport, Germany")
  IF cb.getIndexOf("SFO") == 0 THEN
     CALL cb.addItem("SFO", "San Francisco International Airport, CA" )
  END IF
  CALL cb.removeItem("FRA")
  INPUT BY NAME airport
END MAIN
```
