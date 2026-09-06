---
title: "fglgallery.flush()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_flush.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.flush()"
type: "concept"
---

# fglgallery.flush()

> Displays new added images to the end user.

## Syntax

```
FUNCTION flush( id SMALLINT )
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.

## Usage

This function flushes the new images to the fglgallery web component.

The `flush()` method is used after adding a set of new images with the [`addImage()`](2827-fglgallery-addimage.md "Adds a picture resource to an fglgallery.") function,
to show the images when the image gallery is already displayed. A `flush()` call is
not required before calling the [`display()`](2832-fglgallery-display.md "Displays an fglgallery to the end user.") method.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

## Example

```
CALL fglgallery.addImage(id, "http://...", "Picture 1")
CALL fglgallery.addImage(id, "http://...", "Picture 2")
CALL fglgallery.addImage(id, "http://...", "Picture 3")
...
CALL fglgallery.flush()
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
