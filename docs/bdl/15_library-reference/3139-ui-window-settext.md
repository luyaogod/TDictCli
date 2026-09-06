---
title: "ui.Window.setText"
source: "fgl-topics/c_fgl_ClassWindow_setText.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.setText"
type: "concept"
---

# ui.Window.setText

> Set the window title.

## Syntax

```
setText(
   title STRING )
```

1. title is the title of the window.

## Usage

The `setText()` method defines the title of the window.

By default, the title of a window is defined by the [`TEXT`](../11_user-interface/1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") attribute of the
[`LAYOUT`](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
definition in form files.

## Example

Definiting the title of the current window/form:

```
MAIN
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  CALL ui.Window.getCurrent().setText("Customer")
  MENU "Test"
     COMMAND "exit" EXIT MENU
  END MENU
END MAIN
```

Defining the title of the default `SCREEN` window when no form was displayed
yet:

```
MAIN
    MENU "My Menu"
        BEFORE MENU
            CALL ui.Window.getCurrent().setText("My window title")
        ON ACTION exit ATTRIBUTES(TEXT="Exit")
            EXIT MENU
    END MENU
END MAIN
```

## Related links

**Related concepts**  

[Windows and forms](../11_user-interface/1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
