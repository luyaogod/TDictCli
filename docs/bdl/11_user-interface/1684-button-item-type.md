---
title: "BUTTON item type"
source: "fgl-topics/c_fgl_FormSpecFiles_BUTTON.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > BUTTON item type"
type: "concept"
---

# BUTTON item type

> Defines a push-button that can trigger an action.

## BUTTON item basics

The `BUTTON` form item type defines a standard push button with a label and/or
an icon.

![BUTTON rendering](../_images/FormItemType_BUTTON.jpg)

*BUTTON form item type*

## Defining a BUTTON

The label of a `BUTTON` form item is defined with the `TEXT`
attribute. The `COMMENT` attribute can be used to define a hint for the button.
Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") for these
attributes.

The picture is defined by the `IMAGE` attribute. Consider using [centralized icons](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.") for button images.

```
BUTTON ...
   TEXT = %"common.button.text.ok",
   IMAGE = "accept",
   COMMENT = %"common.button.hint.ok";
```

`BUTTON` form items can inherit action default attributes, to avoid having to
specify the `TEXT`, `COMMENT` and `IMAGE` attributes
in all elements bound to the same action. For more details, see [Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.").

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [Button style attributes](1631-button-style-attributes.md "Button presentation style attributes apply to BUTTON elements.").

## Detecting BUTTON action

A `BUTTON` form item acts as an action view for a dialog action, and is bound
to the `ON ACTION` handler by name. The action name can be prefixed with a sub-dialog
identifier and/or a field name, to define a qualified action view:

```
-- Form file (grid layout)
BUTTON b1: print;

-- Program file:
ON ACTION print
   -- Execute code related to the print action
```

When controlled by a `COMMAND` action handler in a `DIALOG`
interactive instruction, form buttons can get the focus and thus be part of the tabbing list ([`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") attribute).

For more details, see [Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?").

## Where to use a BUTTON

A `BUTTON` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [BUTTON item definition](1730-button-item-definition.md "Defines attributes for a push-button that can trigger an action.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Related links

**Related concepts**  

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
