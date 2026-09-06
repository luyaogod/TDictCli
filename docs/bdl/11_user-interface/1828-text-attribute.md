---
title: "TEXT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_TEXT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > TEXT attribute"
type: "concept"
---

# TEXT attribute

> The TEXT attribute defines the label associated with a form item.

## Syntax

```
TEXT = [%]"string"
```

1. string defines the label to be associated with the form item, with the %
   prefix it is a localized string.

## Usage

The `TEXT` attribute is used to define the label of a
form item, for example for a `CHECKBOX` form field or a
`BUTTON` action view.

Consider using localized strings with the
`%"string-id"` syntax, if you plan to internationalize your
application.

This attribute is also an action attribute that can be defined in the
`ACTION DEFAULTS` section of a form or directly in an action view
(`BUTTON`), see [TEXT action attribute](2276-text-action-attribute.md "The TEXT attribute defines the label associated to the action.") for more
details.

The `TEXT` attribute may be used for accessibility, if a screen reader is
available on the front-end side. For more details, see [Screen readers](1591-screen-readers.md "How to integrate with platform screen readers?").

## Example

```
-- As form action default
ACTION DEFAULTS
  ACTION print (TEXT="Print")
END

-- As a CHECKBOX label
CHECKBOX cb01 = FORMONLY.checkbox01, 
                TEXT="OK" ... ;

-- As a BUTTON label
BUTTON b1: print, TEXT="Print";
```

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Form items](1665-form-items.md "The concept of form item includes all elements used in the definition of a form.")
