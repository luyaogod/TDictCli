---
title: "INITIALIZER attribute"
source: "fgl-topics/c_fgl_FSFAttributes_INITIALIZER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > INITIALIZER attribute"
type: "concept"
---

# INITIALIZER attribute

> The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item.

## Syntax

```
INITIALIZER = [module.]function
```

1. module is the name of the module implementing the function. If no module is
   specified, the function must be defined in a module that is already loaded by the runtime system. If
   a module is specified, it will be automatically loaded on demand. Module name is case
   sensitive.
2. function is an identifier defining the program function to be called to
   initialize the combobox item list. Function name is case insensitive.

## Usage

The `INITIALIZER` attribute defines the function to be called to fill the item
list of the `COMBOBOX` field.

The initialization function name is case insensitive.

When using a function qualified with a module prefix, the runtime system will automatically load
the module if it is not yet loaded. In such case, the module specification in the
`INITIALIZER` attribute is sufficient.

The module prefix of the initialization function name is case sensitive
(unlike the function name, which is case insensitive). If the module is not yet loaded, it will be
loaded automatically when the initializer function is needed.

When specifying a un-qualified function name (without the module as prefix), the module defining
the initialization function must have been loaded before the form is loaded.

> **Tip:**
>
> To make sure that the module is loaded, define other functions in the module, to
> be invoked with a regular `CALL` instruction before displaying the form.

The initialization function must be defined with an [`ui.ComboBox`](../15_library-reference/3247-the-combobox-class.md "The ui.ComboBox class provides an interface to the COMBOBOX form field view in the abstract user interface tree.")
parameter:

```
FUNCTION func-name(combobox ui.ComboBox) RETURNS ()
```

For more details about combobox initializer programming, see [Filling a COMBOBOX item list](2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.").

## Related links

**Related concepts**  

[ui.ComboBox.setDefaultInitializerFunction](../15_library-reference/3250-ui-combobox-setdefaultinitializerfunction.md "Define the default initializer for combobox form items.")
