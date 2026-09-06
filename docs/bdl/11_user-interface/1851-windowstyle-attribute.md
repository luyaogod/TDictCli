---
title: "WINDOWSTYLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WINDOWSTYLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WINDOWSTYLE attribute"
type: "concept"
---

# WINDOWSTYLE attribute

> The WINDOWSTYLE attribute defines the style to be used by the parent window of a form.

## Syntax

```
WINDOWSTYLE = "string"
```

1. string is a user-defined style name.

## Usage

The `WINDOWSTYLE` attribute can be used to specify the style of the parent window
that will hold the form. This attribute is specific to the `LAYOUT` element. Do not confuse this
attribute with the `STYLE` attribute, which is used to specify the decoration style
of the form elements.

When a form is loaded by the [`OPEN
WINDOW`](1572-open-window.md "Creates and displays a new window.") or [`DISPLAY
FORM`](1579-display-form.md "Displays and associates a form with the current window.") instruction, the runtime system automatically assigns the
`WINDOWSTYLE` to the `STYLE` attribute of the parent window
element.

## Example

```
LAYOUT ( STYLE="BigFont", WINDOWSTYLE="dialog" )
```

## Related links

**Related concepts**  

[STYLE attribute](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")
