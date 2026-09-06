---
title: "ui.ComboBox.setDefaultInitializer"
source: "fgl-topics/c_fgl_ClassCombo_setDefaultInitializer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.setDefaultInitializer"
type: "concept"
---

# ui.ComboBox.setDefaultInitializer

> Define the default initializer for combobox form items.

## Syntax

> **Note:**
>
> The `ui.ComboBox.setDefaultInitializer()` method is deprecated, use [`ui.ComboBox.setDefaultInitializerFunction()`](3249-ui-combobox-setdefaultinitializer.md "Define the default initializer for combobox form items.") instead.

```
ui.ComboBox.setDefaultInitializer(
   initializer STRING )
```

1. initializer is the name of the initialization function. This can be a simple
   function name, or a function name prefixed by a module name in the form
   `"module-name.function-name"`.

## Usage

The `ui.ComboBox.setDefaultInitializer()` class method specifies a default
initialization function to be called each time a `COMBOBOX` form field is created
when loading forms.

The method takes the name of the initialization function as a parameter. It can be prefixed by the
module name followed by a dot.

The initialization function name is case insensitive.

The module prefix of the initialization function name is case sensitive
(unlike the function name, which is case insensitive). If the module is not yet loaded, it will be
loaded automatically when the initializer function is needed.

The initialization function is called with the [`ui.ComboBox`](3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree.") object as the parameter.

A combobox initialization function typically fills the drop down list of `COMBOBOX`
fields with items.

For more details about combobox initializer programming, see [Filling a COMBOBOX item list](../11_user-interface/2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.").

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
COMBOBOX cb3 = FORMONLY.city3, NOT NULL, INITIALIZER=cb_init_2;
END
```

The main module:

```
IMPORT FGL setup
MAIN
    DEFINE rec RECORD city1, city2, city3 STRING END RECORD
    CALL setup.init_combo_setup(TRUE)
    CALL ui.ComboBox.setDefaultInitializer("cb_init_1")
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1 -- initialization function is called
    INPUT BY NAME rec.*
END MAIN
```

The imported module
setup.4gl:

```
PRIVATE DEFINE with_undef BOOLEAN

PUBLIC FUNCTION init_combo_setup(wu)
    DEFINE wu BOOLEAN
    LET with_undef = wu
END FUNCTION
    
PUBLIC FUNCTION cb_init_1(comboBox ui.ComboBox)
  CALL comboBox.clear()
  IF with_undef THEN
     CALL comboBox.addItem(0,"<undef>")
  END IF
  CALL comboBox.addItem(101,"London")
  CALL comboBox.addItem(102,"Paris")
  CALL comboBox.addItem(103,"Rome")
  CALL comboBox.addItem(104,"Copenhague")
END FUNCTION

PUBLIC FUNCTION cb_init_2(comboBox ui.ComboBox)
  CALL comboBox.clear()
  IF with_undef THEN
     CALL comboBox.addItem(0,"<undef>")
  END IF
  CALL comboBox.addItem(201,"Berlin")
  CALL comboBox.addItem(202,"Berne")
  CALL comboBox.addItem(203,"Madrid")
  CALL comboBox.addItem(204,"Vienne")
END FUNCTION
```
