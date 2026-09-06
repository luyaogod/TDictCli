---
title: "ui.Form.setDefaultInitializerFunction"
source: "fgl-topics/c_fgl_ClassForm_setDefaultInitializerFunction.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setDefaultInitializerFunction"
type: "concept"
---

# ui.Form.setDefaultInitializerFunction

> Define the default initializer for all forms.

## Syntax

```
ui.Form.setDefaultInitializerFunction(
   initializer FUNCTION(
      form ui.Form
      ) RETURNS ()
 )
```

1. initializer is a [function
   reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") that implements the form initialization code.

## Usage

Specify a default initialization function with the
`ui.Form.setDefaultInitializerFunction()` method, to implement global processing
when a form is opened with `OPEN FORM / DISPLAY FORM` or with `OPEN WINDOW
... WITH FORM`.

Use this method to define a global/default initialization function for all forms of the
program.

The method takes a function reference as a parameter. The initialization function must be defined
with the [`ui.Form`](3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.") object as the
parameter.

## Example

The form file form.per:

```
LAYOUT 
GRID
{   
[f1             ]
}   
END 
END 
ATTRIBUTES
EDIT f1 = FORMONLY.cust_name;
END
```

The main module:

```
IMPORT FGL setup
MAIN
    DEFINE cust_name STRING
    CALL ui.Form.setDefaultInitializerFunction(FUNCTION setup.form_init)
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1 -- initialization function is called
    INPUT BY NAME cust_name
END MAIN
```

The imported module
setup.4gl:

```
PUBLIC FUNCTION form_init(form ui.Form) RETURNS ()
    CALL form.loadToolBar("common_toolbar")
END FUNCTION
```

## Related links

**Related concepts**  

[OPEN WINDOW](../11_user-interface/1572-open-window.md "Creates and displays a new window.")

[OPEN FORM](../11_user-interface/1578-open-form.md "Declares a compiled form in the program.")

[DISPLAY FORM](../11_user-interface/1579-display-form.md "Displays and associates a form with the current window.")
