---
title: "IMAGE item type"
source: "fgl-topics/c_fgl_FormSpecFiles_IMAGE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > IMAGE item type"
type: "concept"
---

# IMAGE item type

> Defines an area that can display an image resource.

## IMAGE item basics

The `IMAGE` item type defines an area where a picture resource can be
displayed.

![IMAGE rendering](../_images/FormItemType_IMAGE.jpg)

*IMAGE form item type*

## Defining an IMAGE

An `IMAGE` form item can be defined as a form field image or as a static
image. Use a form field image when the content of the image will change often during program
execution (for example, to display images from the database). Use a static image if the image
remains the same during program execution.

Front-ends support different presentation and behavior options, which can be controlled
by a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute.
For more details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.")
and [Image style attributes](1641-image-style-attributes.md "Image style presentation attributes apply to an IMAGE element.").

## Form field IMAGE item

Use a form field image item to display values that change often during program execution, for
example if the image is stored in the database.

The picture resource is defined by the value of the field.

The value can be changed by the program using the `DISPLAY BY NAME / DISPLAY TO`
instruction, or just by changing the value of the program variable bound to the image field when
using the `UNBUFFERED` mode in an interactive instruction.

When defining the `IMAGE` item in the form, use a field name to identify the
element in programs:

```
IMAGE f001 = cars.picture, SIZEPOLICY=FIXED, STRETCH=BOTH;
```

## Providing the image resource

To display an image, the front-end needs the image data, which can be provided in different ways.

For example, you can specify a URL, a mapped icon, or a plain image file (centralized on the
application server).

For more details about image resource specification, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

## Detecting IMAGE clicks

To inform the dialog immediately when an image was clicked, define the [`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item.") attribute in the `IMAGE`
item, and implement the corresponding `ON ACTION` handler in the dialog:

```
-- Form file (grid layout)
IMAGE : logo, IMAGE="fourjs.png",
   ACTION=show_about;

-- Program file:
ON ACTION show_about
   -- The image was clicked
```

The program can then react immediately when the user selects the image element.

## Where to use IMAGE

An `IMAGE` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and an [IMAGE item definition](1739-image-item-definition.md "Defines attributes for an area that can display an image resource.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of an `IMAGE` widget can be controlled by several attributes such as
`SIZEPOLICY`, `AUTOSCALE` and `STRETCH`.

For more details about image sizing, see [Controlling the image layout](1585-controlling-the-image-layout.md "Explains how image form items can be sized in different front-end layout systems.").

## Related links

**Related concepts**  

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")
