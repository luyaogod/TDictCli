---
title: "IMAGE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_IMAGE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > IMAGE attribute"
type: "concept"
---

# IMAGE attribute

> The IMAGE attribute defines the image resource to be displayed for the form item.

## Syntax

```
IMAGE = "resource"
```

1. resource defines the file name, path or URL to the image source.

## Usage

The `IMAGE` attribute is used to define the image resource to be displayed
for form items such as a `BUTTON`, `BUTTONEDIT`, a
`TOOLBAR` button, or a static `IMAGE` item.

For more details about image resource specification, see
[Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

This attribute is also an action attribute that can be defined in the
`ACTION DEFAULTS` section of a form or directly in an action view
(`BUTTON`), see [IMAGE action attribute](2274-image-action-attribute.md "The IMAGE attribute defines the image resource to be displayed for the action.")
for more details.

## Example

```
-- As action default
ACTION DEFAULTS
  ACTION print (IMAGE="printer")
END

-- In a form buttonedit or button
BUTTONEDIT f001 = FORMONLY.field01, IMAGE = "zoom";
BUTTON b01: open_file, IMAGE = "buttons/fileopen";
BUTTON b02: accept, IMAGE = "http://myserver/images/accept.png";

-- In a static image form item
IMAGE: img1, IMAGE = "mylogo.png"
```

## Related links

**Related concepts**  

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")

[IMAGE item type](1695-image-item-type.md "Defines an area that can display an image resource.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
