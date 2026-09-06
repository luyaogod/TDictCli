---
title: "The fglgallery web component"
source: "fgl-topics/c_fgl_fgl_web_components_fglgallery.html"
breadcrumb: "User interface > User interface programming > Web components > Built-in web components > Built-in web components reference > The fglgallery web component"
type: "concept"
---

# The fglgallery web component

> The fglgallery built-in web component shows a collection of pictures the end user can choose from.

The `fglgallery` built-in web component is a [gICAPI web component](2391-using-a-gicapi-web-component.md "This section describes how to add a gICAPI-based web component to your application.").

Several rendering modes are supported: Mosaic, List and Thumbnails for basic rendering.

![Screenshot of a program using the fglgallery web component](../_images/fglgallery01_gbc.jpg)

*fglgallery web component*

## Defining the `fglgallery` web component in the form file

In your .per form definition file, define the image gallery as a [`WEBCOMPONENT`](1750-webcomponent-item-definition.md "Defines attributes for a generic form field that can receive an external widget.") form
item with the [`COMPONENTTYPE`](1771-componenttype-attribute.md "The COMPONENTTYPE attribute defines a name identifying the external widget for WEBCOMPONENT fields.") attribute set to the `"fglgallery"` value.

As the image gallery implements its own navigation controls or scrollbars, the [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item.") attribute can be set
to `NONE`.

Use [`SIZEPOLICY=FIXED`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") and
[`STRETCH=BOTH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), to get an image
gallery that resizes with the parent window.

For example:

```
LAYOUT
GRID
{
[ig                     ]
[                       ]
[                       ]
...
ATTRIBUTES
WEBCOMPONENT ig = FORMONLY.gallery_wc,
         COMPONENTTYPE = "fglgallery",
         PROPERTIES=(selection="image_selection"),
         SIZEPOLICY = FIXED,
         STRETCH = BOTH;
...
```

## Introducing the fglgallery.4gl library

The [fglgallery.4gl](../15_library-reference/2825-fglgallery-image-gallery-module.md) module implements a set of functions to control an
fglgallery web component.

The source can be found in $FGLDIR/src/webcomponents/fglgallery.

## Customizing the image gallery

The HTML elements created by fglgallery can be customized with CSS styling.

Default CSS styles for fglgallery are provided in
$FGLDIR/webcomponents/fglgallery/css/fglgallery.css.

Overwrite or define new style attributes for each fglgallery element type, in a secondary file
named $FGLDIR/webcomponents/fglgallery/css/fglgallery-custom.css, and uncomment
the link in $FGLDIR/webcomponents/fglgallery/fglgallery.html to load this custom
CSS file after fglgallery.css.

For example, to define a background color for the list display mode, add the following lines in
fglgallery-custom.css:

```
.list-main {
    background-color: black;
}
```

Do not modify the default CSS styles in fglgallery.css: This file will be
overwritten by a new installation. Use the fglgallery-custom.css file name
instead.

## fglgallery library initialization and finalization

Library initialization and finalization functions are provided to prepare the library before
usage, and free resources when the library is not longer
needed:

```
CALL fglgallery.initialize()
...
CALL fglgallery.finalize()
```

For more details see the [initialize()](../15_library-reference/2838-fglgallery-initialize.md "Prepares the fglgallery library for use.") and [`finalize()`](../15_library-reference/2833-fglgallery-finalize.md "Releases the fglgallery library.") functions.

## Creating and destroying an image gallery handle

Before using the image gallery web component, you must create a handle with the [`create()`](../15_library-reference/2829-fglgallery-create.md "Creates a new fglgallery handle.")
function:

```
DEFINE id INTEGER
...
LET id = fglgallery.create("formonly.gallery_wc")
```

The `create()` function takes the `WEBCOMPONENT` form item name as
parameter, and returns the id of the image gallery handle. The image gallery id will be used in
subsequent library calls. This technique allows for the manipulation of several image gallery items
in the same form, and the use of the library by different callers.

When the image gallery web component is no longer used, destroy the handle with the [`destroy()`](../15_library-reference/2831-fglgallery-destroy.md "Frees resources allocated for an fglgallery.")
function:

```
CALL fglgallery.destroy(id)
```

## The `fglgallery.t_struct_value` type

The fglgallery library provides a user-type to define a variable that will be used to send and
receive data from the image gallery web component.

The [`fglgallery.t_struct_value`](../15_library-reference/2826-fglgallery-t-struct-value-type.md "The t_struct_value type holds image selection data.") type is a `RECORD` structure, that
holds image selection
information:

```
PUBLIC TYPE t_struct_value RECORD
                 current INTEGER,
                 selected DYNAMIC ARRAY OF INTEGER
             END RECORD
```

In this structure, each picture is identified by its ordinal position in the image gallery.

## Data type for the fglgallery `WEBCOMPONENT` field

The fglgallery `WEBCOMPONENT` field value will be used to exchange image selection
information in JSON format.

Use `STRING` as data type for the variable bound to the
`WEBCOMPONENT` field.

The string variable can then be used to map the JSON to the [`fglgallery.t_struct_value`](../15_library-reference/2826-fglgallery-t-struct-value-type.md "The t_struct_value type holds image selection data.")
record:

```
DEFINE rec RECORD
         ...
         gallery_wc STRING,
         ...
       END RECORD
DEFINE 
...
INPUT BY NAME rec.* ...
    ...
    ON ACTION image_selection
        CALL util.JSON.parse( rec.gallery_wc, struct_value )
        DISPLAY "Current: ", fglgallery.getPath(id, struct_value.current)
...
```

## Filling the image gallery with pictures

In order to add a picture resource to the image gallery, use the [`addImage()`](../15_library-reference/2827-fglgallery-addimage.md "Adds a picture resource to an fglgallery.") function.
This function takes three parameters:

1. The id of the image gallery handle,
2. The URL of the picture resource,
3. The text/comment for the picture.

> **Important:**
>
> When using local application pictures, you need to understand how these
> pictures are transmitted to the front-end for display inside the `WEBCOMPONENT`. For
> more details, see [Using image resources with the gICAPI web component](2409-using-image-resources-with-the-gicapi-web-component.md "This section explains how to use image resources in a gICAPI web component.").

Pictures are identified by their ordinal position in the image gallery.

To add a picture resource that resides on the application server where the program executes, use
the `ui.Interface.filenameToURI()` method:

```
CALL fglgallery.addImage(id,
   ui.Interface.filenameToURI("big_smiley.jpg"),
   "The big smiley.")
```

To add a picture resources that is available as a URL, from the local network or the internet,
provide the URL
directly:

```
CALL fglgallery.addImage(id,
   "http://freebigpictures.com/wp-content/uploads/2009/09/mountain-ridge.jpg",
   "Mountain ridge")
```

After adding pictures, display the image gallery with [`display()`](../15_library-reference/2832-fglgallery-display.md "Displays an fglgallery to the end user."):

```
CALL fglgallery.display(id, FGLGALLERY_TYPE_MOSAIC, FGLGALLERY_SIZE_NORMAL)
```

Or call the [`flush()`](../15_library-reference/2834-fglgallery-flush.md "Displays new added images to the end user.") function, when the image gallery is already
displayed:

```
CALL fglgallery.flush(id)
```

This flush mechanism allows you to add a set of pictures in one front-end exchange.

## Get the number of pictures in the gallery

To get the number of pictures that are currently displayed in the gallery, call the [`getImageCount()`](../15_library-reference/2835-fglgallery-getimagecount.md "Returns the number of pictures in an fglgallery.")
function:

```
DISPLAY fglgallery.getImageCount(id)
```

## Cleaning the image gallery

Use the [`clean()`](../15_library-reference/2828-fglgallery-clean.md "Removes all pictures from an fglgallery.") function to remove all pictures from the image
gallery:

```
CALL fglgallery.clean(id)
```

## Displaying the image gallery

To display the image gallery, use the [`display()`](../15_library-reference/2832-fglgallery-display.md "Displays an fglgallery to the end user.") function.
This function takes the image gallery id as parameter, followed by the gallery type and size
options:

```
CALL fglgallery.display(id, FGLGALLERY_TYPE_MOSAIC, FGLGALLERY_SIZE_NORMAL)
```

This function can be called several times, in order to change the type and size of the image
gallery.

## Image gallery type (FGLGALLERY\_TYPE\_\*)

The type of the image gallery defines the layout / rendering of the gallery. This type is defined
by the second parameter of the `display()` function.

For possible values, see the [`display()`](../15_library-reference/2832-fglgallery-display.md "Displays an fglgallery to the end user.") function reference topic.

## Image gallery size (FGLGALLERY\_SIZE\_\*)

The size of the images displayed in the gallery is defined by the third parameter of the
`display()` function. The size is specified with an abstract option (small, normal,
large, etc). This size is related to the current font (the "em" unit is used in HTML).

For possible values, see the [`display()`](../15_library-reference/2832-fglgallery-display.md "Displays an fglgallery to the end user.") function reference topic.

## Image elements aspect ratio

All image elements in the fglgallery are displayed with the same size, to align properly when the
image resources have different sizes.

By default, image elements are displayed with a square (1:1) aspect ratio.

The aspect ratio can be configured with the [`setImageAspectRatio()`](../15_library-reference/2840-fglgallery-setimageaspectratio.md "Defines the aspect ratio for gallery images.") function.

## Get the picture resource path from a position

To get the picture resource name from an image position, use the [`getPath()`](../15_library-reference/2836-fglgallery-getpath.md "Returns the URL of a picture in an fglgallery.")
function:

```
LET path = fglgallery.getPath(id, struct_value.current)
```

## Get the picture title from a position

To get the title of a picture from an image position, use the [`getTitle()`](../15_library-reference/2837-fglgallery-gettitle.md "Returns the description of a picture in an fglgallery.")
function:

```
LET title = fglgallery.getTitle(id, struct_value.current)
```

## Delete pictures from the image gallery

To remove pictures from the image gallery, use the [`deleteImages()`](../15_library-reference/2830-fglgallery-deleteimages.md "Deletes pictures used in an fglgallery.")
function. This function takes a `DYNAMIC ARRAY OF INTEGER` as parameter, defining the
ordinal positions of the pictures to be deleted:

```
DEFINE todel DYNAMIC ARRAY OF INTEGER
LET todel[1] = 3
LET todel[2] = 6
CALL fglgallery.deleteImages(id, todel)
```

## Single and multiple picture selection

By default, the end user can only select a single picture from the image gallery. With the
Mosaic, List and Thumbnails display modes, you can enable multiple picture selection.

Multiple picture selection is not supported with the Light Gallery display mode.

To enable multiple picture selection, use the [`setMultipleSelection()`](../15_library-reference/2839-fglgallery-setmultipleselection.md "Enables/disables multiple picture selection in an fglgallery.")
function:

```
CALL fglgallery.setMultipleSelection(id, TRUE)
```

Picture selection can then be detected with an `ON ACTION` handler.

## Detecting image selection

To implement picture selection in the image gallery component, define the
`WEBCOMPONENT` form field with a `"selection"` property in the
`PROPERTIES`
attribute:

```
WEBCOMPONENT ig = FORMONLY.gallery_wc,
   ...
   PROPERTIES = (selection="image_selection"),
   ...
```

Picture selection is built-in to the fglgallery web component.

In the program code, detect picture selection with an `ON ACTION` handler defined
with an action name that matches the "selection" property of the `WEBCOMPONENT`
field.

When the action is fired, selection information is provided in the `WEBCOMPONENT`
field value as a JSON formatted string similar to:
`{"current":5,"selected":[4,5]}`

> **Tip:**
>
> In order to get the selection information directly in the program variable bound to
> the `WEBCOMPONENT` field, consider using the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") dialog
> attribute.

This JSON string can be parsed with the [`util.JSON.parse()`](../15_library-reference/3582-util-json-parse.md "Parses a JSON string and fills program variables with the values.") method, to fill a variable defined with the [`fglgallery.t_struct_value`](../15_library-reference/2826-fglgallery-t-struct-value-type.md "The t_struct_value type holds image selection data.") type:

```
DEFINE struct_value fglgallery.t_struct_value
...
   INPUT BY NAME rec.* ATTRIBUTES (UNBUFFERED, WITHOUT DEFAULTS)
   ...
   ON ACTION image_selection
      CALL util.JSON.parse( rec.gallery_wc, struct_value )
      DISPLAY struct_value.current,
              struct_value.selected.getLength()
...
```

The selection structure contains a "current" field, to identify the current selected picture, and
when the multiple-selection option is enabled, the "selection" field contains the list of picture
ids that are selected.

## Define selected pictures by program

In order to set the current picture, and select or de-select pictures from the program code, use
the variable defined with the `fglgallery.t_struct_value` type and assign it as a
JSON string to the web component field:

1. Fill the [`fglgallery.t_struct_value`](../15_library-reference/2826-fglgallery-t-struct-value-type.md "The t_struct_value type holds image selection data.") variable with the picture ordinal positions,
2. Convert to a JSON string with [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays."),
3. Set the `WEBCOMPONENT` field value with the resulting string.

For example, to set picture #3 as current, and select pictures #2, #7 et #15:

```
DEFINE struct_value fglgallery.t_struct_value
...
LET struct_value.current = 3
CALL struct_value.selected.clear()
LET struct_value.selected[1] = 2
LET struct_value.selected[2] = 7
LET struct_value.selected[3] = 15
LET rec.gallery_wc = util.JSON.stringify(struct_value)
```

## Complete fglgallery API reference

Check the $FGLDIR/src/webcomponents/fglgallery/fglgallery.4gl source for
more details about the image gallery API.

## Related links

**Related reference**  

[fglgallery: Image gallery module](../15_library-reference/2825-fglgallery-image-gallery-module.md "fglgallery: Image gallery module")

## Child topics

- [Examples](2419-examples.md): fglgallery built-in web component usage examples.
