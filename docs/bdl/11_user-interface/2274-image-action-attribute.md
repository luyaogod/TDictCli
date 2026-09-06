---
title: "IMAGE action attribute"
source: "fgl-topics/c_fgl_action_attribute_IMAGE.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes list > IMAGE action attribute"
type: "concept"
---

# IMAGE action attribute

> The IMAGE attribute defines the image resource to be displayed for the action.

## Syntax 1 (Dialog action handlers and form action defaults):

```
IMAGE = "resource"
```

## Syntax 2 (Global .4ad action defaults file):

```
image = "resource"
```

1. resource defines the file name, path or URL to the image source.

## Usage

The `IMAGE` attribute is used
to define the image resource for action views such as `BUTTON`,
`BUTTONEDIT`, or a `TOOLBAR` button.

For more details about image resource specification, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

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
ON ACTION print ATTRIBUTES(IMAGE="printer")

-- As action default
ACTION DEFAULTS
  ACTION print (IMAGE="printer")
END

-- In a form buttonedit or button
BUTTONEDIT f001 = FORMONLY.field01, IMAGE = "zoom";
BUTTON b01: open_file, IMAGE = "buttons/fileopen";
BUTTON b02: accept, IMAGE = "http://myserver/images/accept.png";
```

## Related links

**Related concepts**  

[IMAGE attribute](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.")

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")
