---
title: "TEXT action attribute"
source: "fgl-topics/c_fgl_action_attribute_TEXT.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > TEXT action attribute"
type: "concept"
---

# TEXT action attribute

> The TEXT attribute defines the label associated to the action.

## Syntax 1 (Dialog action handlers and form action defaults):

```
TEXT = [%]"string"
```

1. string defines the label for the action, with the % prefix for a localized
   string.

## Syntax 2 (Global .4ad action defaults file):

```
text = "string"
```

1. string is the text to display. Use `<LStr/>` child element
   for a localized string.

## Usage

The `TEXT` attribute is used to define the label associated to an action, for
example for a `CHECKBOX` form field or a `BUTTON` action view.

Consider using localized strings with the `%"string-id"` syntax, if you
plan to internationalize your application.

This action attribute can be specified as action default attribute
in a global .4ad file, in the `ACTION DEFAULTS` section of form
files, as dialog action attribute (using `ON ACTION name
ATTRIBUTES(...)`), or directly as action view attribute (in the form definition file).

Action view decoration attributes `TEXT`,
`IMAGE` and `COMMENT` defined at dialog level with `ON ACTION
name ATTRIBUTES(...)` apply only to implicit action views
such as [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), [context menu action views](2288-action-display-in-the-context-menu.md "The CONTEXTMENU action default attribute allows you to control action visibility in the context menu.") and [rowbound action views](2291-actions-bound-to-the-current-row.md "Actions can be configured with the ROWBOUND attribute depending on whether there is a current row."): The
explicit action views (such as buttons in the form layout) will not get dialog-level
attribute values. Note that the `COMMENT` attribute will not be rendered on context
menu and rowbound action views. See also [Dialog action handler attributes](2262-dialog-action-handler-attributes.md)
for more details.

> **Note:**
>
> When a `TEXT` (label) of an action view (`BUTTON` elements,
> `TOPMENU` commands, `TOOLBAR` buttons, etc) contains an &
> ampersand character, it will be hidden (you need to double the ampersand to get one). For example,
> `"Fish &Chips && Sauce"` will show as `"Fish Chips &
> Sauce"`. This is implemented to simplify migration from older front ends.

## Example

```
-- As action handler attribute
ON ACTION print ATTRIBUTES(TEXT="Print")

-- As form action default
ACTION DEFAULTS
  ACTION print (TEXT="Print")
END

-- As a CHECKBOX label
CHECKBOX cb01 = FORMONLY.checkbox01, 
                TEXT="OK" ... ;

-- As a BUTTON label, using a localized string id
BUTTON b1: print, TEXT=%"actions.print.label";

-- In a global action defaults file with a localized string id
<ActionDefault name="zoom" text="Zoom" ... >
   <LStr text="actions.zoom.label" />
</ActionDefault>
```

## Related links

**Related concepts**  

[TEXT attribute](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
