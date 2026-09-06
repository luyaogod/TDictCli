---
title: "ui.Form.setDefaultInitializer"
source: "fgl-topics/c_fgl_ClassForm_setDefaultInitializer.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setDefaultInitializer"
type: "concept"
---

# ui.Form.setDefaultInitializer

> Define the default initializer for all forms.

## Syntax

> **Note:**
>
> The `ui.Form.setDefaultInitializer()` method is deprecated, use [`ui.Form.setDefaultInitializerFunction()`](3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.") instead.

```
ui.Form.setDefaultInitializer(
   initializer STRING )
```

1. initializer is the name of a function in the program. This can be a simple
   function name, or a function name prefixed by a module name in the form
   `"module-name.function-name"`.

## Usage

Specify a default initialization function with the
`ui.Form.setDefaultInitializer()` method, to implement global processing when a form
is opened with `OPEN FORM / DISPLAY FORM` or with `OPEN WINDOW ... WITH
FORM`.

The method takes the name of the initialization function as a parameter. It can be prefixed by
the module name followed by a dot.

The initialization function name is case insensitive.

The module prefix of the initialization function name is case sensitive
(unlike the function name, which is case insensitive). If the module is not yet loaded, it will be
loaded automatically when the initializer function is needed.

The initialization function is called with the [`ui.Form`](3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.") object as the parameter.

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
    CALL setup.init_form_setup(FALSE)
    CALL ui.Form.setDefaultInitializer("form_init")
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1 -- initialization function is called
    INPUT BY NAME cust_name
END MAIN
```

The imported module
setup.4gl:

```
PRIVATE DEFINE with_toolbar BOOLEAN

PUBLIC FUNCTION init_form_setup(tb)
    DEFINE tb BOOLEAN
    LET with_toolbar = tb
END FUNCTION

PUBLIC FUNCTION form_init(form ui.Form)
    IF with_toolbar THEN
       CALL form.loadToolBar("common_toolbar")
    END IF
END FUNCTION
```

## Related links

**Related concepts**  

[OPEN WINDOW](../11_user-interface/1572-open-window.md "Creates and displays a new window.")

[OPEN FORM](../11_user-interface/1578-open-form.md "Declares a compiled form in the program.")

[DISPLAY FORM](../11_user-interface/1579-display-form.md "Displays and associates a form with the current window.")
