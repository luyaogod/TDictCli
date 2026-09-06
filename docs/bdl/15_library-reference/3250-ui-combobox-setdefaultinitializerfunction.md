---
title: "ui.ComboBox.setDefaultInitializerFunction"
source: "fgl-topics/c_fgl_ClassCombo_setDefaultInitializerFunction.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.setDefaultInitializerFunction"
type: "concept"
---

# ui.ComboBox.setDefaultInitializerFunction

> Define the default initializer for combobox form items.

## Syntax

```
ui.ComboBox.setDefaultInitializerFunction(
      initializer FUNCTION(
      comboBox ui.ComboBox
      ) RETURNS ()
 )
```

1. initializer is a [function
   reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") that implements the combobox initialization code.

## Usage

The `ui.ComboBox.setDefaultInitializerFunction()` class method specifies a default
initialization function to be called each time a `COMBOBOX` form field is created
when loading forms.

Use this method to define a global/default initialization function for all comboboxes of the
program. For individual comboboxes, consider using the [`INITIALIZER`](../11_user-interface/1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item.") form field attribute
instead.

The method takes a function reference as a parameter. The initialization function must be defined
with the [`ui.ComboBox`](3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree.") object as the
parameter.

## Example

The form file form.per:

```
LAYOUT 
GRID
{
[cb1             ]
[cb2             ]
[cb3             ]
}   
END
END
ATTRIBUTES
COMBOBOX cb1 = FORMONLY.city1, NOT NULL; -- Gets default initializer
COMBOBOX cb2 = FORMONLY.city2, NOT NULL; -- Gets default initializer
COMBOBOX cb3 = FORMONLY.city3, NOT NULL, INITIALIZER = setup.cb_init_2;
END
```

The main module:

```
IMPORT FGL setup
MAIN
    DEFINE rec RECORD city1, city2, city3 STRING END RECORD
    CALL ui.ComboBox.setDefaultInitializerFunction(FUNCTION setup.cb_init_1)
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1 -- initialization function is called
    INPUT BY NAME rec.*
END MAIN
```

The imported module
setup.4gl:

```
PUBLIC FUNCTION cb_init_1(comboBox ui.ComboBox) RETURNS ()
  CALL comboBox.clear()
  CALL comboBox.addItem(0,"<undef>")
  CALL comboBox.addItem(101,"London")
  CALL comboBox.addItem(102,"Paris")
  CALL comboBox.addItem(103,"Rome")
  CALL comboBox.addItem(104,"Copenhague")
END FUNCTION

PUBLIC FUNCTION cb_init_2(comboBox ui.ComboBox) RETURNS ()
  CALL comboBox.clear()
  CALL comboBox.addItem(0,"<undef>")
  CALL comboBox.addItem(201,"Berlin")
  CALL comboBox.addItem(202,"Berne")
  CALL comboBox.addItem(203,"Madrid")
  CALL comboBox.addItem(204,"Vienne")
END FUNCTION
```
