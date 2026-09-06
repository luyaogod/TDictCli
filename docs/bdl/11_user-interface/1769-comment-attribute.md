---
title: "COMMENT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_COMMENT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > COMMENT attribute"
type: "concept"
---

# COMMENT attribute

> The COMMENT attribute defines a hint for the user about the form element.

## Syntax

```
COMMENT = [%]"string"
```

1. string is the text to display. With the % prefix, it is a localized
   string.

## Usage

The most common use of the `COMMENT` attribute is to give information or
instructions to the user.

The `COMMENT` attribute can be used for different types of form elements:

- with form field definitions, to show a message when the field gets the focus.
- with action views, to give a hint to the user about the action.

With form fields, the `COMMENT` attribute is particularly appropriate when the
field accepts only a limited set of values.

The screen location where the comment message displays depends on configuration settings:

- In TUI mode, the comment message will be displayed in the [`COMMENT LINE`](1573-open-window-attributes.md) defined globally or for the current window.
- In GUI mode, the comment message will be displayed in a tooltip hint when hovering the mouse
  over the form element.

The `COMMENT` attribute can, in specific situations,
define a value placeholder message when the field is empty. If the `COMMENT` cannot
be rendered in a dedicated area (as is the case with mobile devices), the `COMMENT`
is used as the default [`PLACEHOLDER`](1808-placeholder-attribute.md "The PLACEHOLDER attribute defines a hint for the user when the field contains no value.").

With actions, this attribute can be defined in the `ACTION DEFAULTS` section of a
form or directly in an action view (`BUTTON`). See [COMMENT action attribute](2271-comment-action-attribute.md "The COMMENT attribute defines hint for the user about the action.") for more details.

The `COMMENT` attribute may be used for accessibility, if a screen reader is
available on the front-end side. For more details, see [Screen readers](1591-screen-readers.md "How to integrate with platform screen readers?").

## Example

```
-- In a form field definition
EDIT f1 = customer.name, COMMENT = "The customer name";

-- As action default
ACTION DEFAULTS
  ACTION print (COMMENT="Print current order information")
END

-- In a form button
BUTTON b1: print, COMMENT = "Print customer details";
```

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[OPEN WINDOW](1572-open-window.md "Creates and displays a new window.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
