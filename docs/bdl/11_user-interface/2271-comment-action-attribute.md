---
title: "COMMENT action attribute"
source: "fgl-topics/c_fgl_action_attribute_COMMENT.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > COMMENT action attribute"
type: "concept"
---

# COMMENT action attribute

> The COMMENT attribute defines hint for the user about the action.

## Syntax 1 (Dialog action handlers and form action defaults):

```
COMMENT = [%]"string"
```

1. string is the comment to display, with the % prefix for a localized
   string.

## Syntax 2 (Global .4ad action defaults file):

```
comment = "string"
```

1. string is the comment to display. Use `<LStr/>` child
   element for a localized string.

## Usage

Use the `COMMENT` attribute to define a description for the action.
This text will typically be displayed as a hint for the corresponding action view.

Consider using localized strings with the `%"string-id"`
syntax, if you plan to internationalize your application.

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

## Example

```
-- As action handler attribute
ON ACTION print ATTRIBUTES(COMMENT="Prints current record")

-- As action default
ACTION DEFAULTS
  ACTION print (COMMENT="Print current order information")
END

-- In a form buttom, using a localized string id
BUTTON b1: print, COMMENT=%"actions.print.comment";

-- In a global action defaults file with a localized string id
<ActionDefault name="zoom" comment="Opens a zoom window" ... >
   <LStr comment="actions.zoom.comment" />
</ActionDefault>
```

## Related links

**Related concepts**  

[COMMENT attribute](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
