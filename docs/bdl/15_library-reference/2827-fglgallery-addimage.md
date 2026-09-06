---
title: "fglgallery.addImage()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_addimage.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.addImage()"
type: "concept"
---

# fglgallery.addImage()

> Adds a picture resource to an fglgallery.

## Syntax

```
FUNCTION addImage(
   id SMALLINT,
   path STRING,
   title STRING )
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.
2. path is the path to the image resource (URL).
3. title is the text to be displayed beside the picture.

## Usage

This function adds a new image to the fglgallery web component.

The function requires the gallery id, the path to the image file, which can be a URL or a local
relative path, and a title/description of the picture.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The function only registers the image resource for the gallery. In order to display the added
images, you must call the [`flush()`](2834-fglgallery-flush.md "Displays new added images to the end user.") or the [`display()`](2832-fglgallery-display.md "Displays an fglgallery to the end user.") function.

When displaying regular image resources available from a URL, you can directly specify this
URL.

When displaying an application image located on the server, you must use the [ui.Interface.filenameToURI()](3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.") method.

Leave title `NULL`, if you don't want to add a description.

## Example

Displaying an image from a URL:

```
CALL fglgallery.addImage(id,
   "http://freebigpictures.com/wp-content/uploads/2009/09/mountain-ridge.jpg",
   "Mountain ridge")
```

Displaying an image resource that is located on the application
server:

```
CALL fglgallery.addImage(id,
   Interface.filenameToURI("big_smiley.jpg"),
   "The big smiley.")
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
