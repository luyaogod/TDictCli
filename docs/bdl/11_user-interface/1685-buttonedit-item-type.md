---
title: "BUTTONEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_BUTTONEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > BUTTONEDIT item type"
type: "concept"
---

# BUTTONEDIT item type

> Defines a line-edit with a push-button that can trigger an action.

## BUTTONEDIT item basics

The `BUTTONEDIT` form item defines an edit field that gets user input, with an
additional push button that can fire an action.

![BUTTONEDIT rendering](../_images/FormItemType_BUTTONEDIT.jpg)

*BUTTONEDIT form item type*

This type of form field is typically used to open a secondary window, to let the user choose from
a large list of items and set the field value.

## Defining a BUTTONEDIT

The [`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.") attribute of a
`BUTTONEDIT` form item defines the picture to be displayed on the button.

By default, the text editor of a `BUTTONEDIT` allows the user to change the field
value. Use the [`NOTEDITABLE`](1804-noteditable-attribute.md "The NOTEDITABLE attribute disables the text editor.")
attribute to deny text modification. The field still gets the focus, and the action button remains
active, if there is a corresponding action handler in the current
dialog.

```
BUTTONEDIT ...
   IMAGE = "zoom",
   NOTEDITABLE;
```

The button of `BUTTONEDIT` form items can inherit action default attributes, to
avoid having to specify the `IMAGE` attributes in all elements bound to the same
action. For more details, see [Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.").

Most of the attributes described in the [`EDIT`](1690-edit-item-type.md "Defines a simple line-edit field.") item type can also be used with the `BUTTONEDIT`.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [ButtonEdit style attributes](1632-buttonedit-style-attributes.md "ButtonEdit presentation style attributes apply to BUTTONEDIT elements.").

## Detecting BUTTONEDIT button action

The button of a `BUTTONEDIT` form element acts as an action view for a dialog
action, and is bound to the `ON ACTION` handler by the [`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item.") attribute.

> **Important:**
>
> The `BUTTONEDIT` button will be ghosted, if the field is not
> managed by a current dialog, or when the field is disabled. For example, if a `MENU`
> defines the `ON ACTION lookup` handler and the `BUTTONEDIT` has
> `ACTION=lookup`, the button will NOT be active, because the field is not active
> during a `MENU` statement. The button is also ghosted when the field is controlled by
> an non-menu dialog, but is disabled with `DIALOG.setFieldActive()` or defined as
> `NOENTRY`.

In the form, the `ACTION` attribute of `BUTTONEDIT` defines the
name of the action to be sent to the program when the user clicks on the button. In the program, the
action handler is defined with an `ON ACTION`
block:

```
-- Form file:
BUTTONEDIT ...
   ACTION = open_city_list;

-- Program file:
ON ACTION open_city_list
   -- Open the city pick-list
```

The button action of a `BUTTONEDIT` can also be a prefixed with a sub-dialog
identifier and/or field name, to define a qualified action view to be used in conjunction with
`ON ACTION action-name INFIELD field-name`.
However, for convenience, even if the action is not qualified with a field name, the runtime system
considers that action as field-qualified, to make the `BUTTONEDIT` button always
active.

For more details, see [Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?"), [Field-specific actions (INFIELD clause)](2285-field-specific-actions-infield-clause.md "Using the INFIELD clause of ON ACTION provides automatic action activation when a field gets the focus.").

## Where to use a BUTTONEDIT

A `BUTTONEDIT` can be defined with an item tag and a [BUTTONEDIT item definition](1731-buttonedit-item-definition.md "Defines attributes for a line-edit field with a push-button that can trigger an action.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of a `BUTTONEDIT` widget is computed by the layout rules as described in
[Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md).

## Field input length

In grid-based layout, the input length in a `BUTTONEDIT` fields is defined by the
item tag and may need to get the [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field.") attribute. For more details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Related links

**Related concepts**  

[GRID container](1722-grid-container.md "Defines a layout area based on a grid of cells.")

[Defining action views in forms](2278-defining-action-views-in-forms.md "How to define action views that will fire action events.")
