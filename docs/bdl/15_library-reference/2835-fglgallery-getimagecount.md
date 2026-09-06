---
title: "fglgallery.getImageCount()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_getimagecount.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.getImageCount()"
type: "concept"
---

# fglgallery.getImageCount()

> Returns the number of pictures in an fglgallery.

## Syntax

```
FUNCTION getImageCount( id SMALLINT )
  RETURNS INTEGER
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.

## Usage

This function counts the number of pictures in the specified fglgallery web component.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

In an fglgallery web component, pictures are identified by their ordinal position. This function
can be used to loop through the picture indexes.

## Example

```
FOR idx = 1 TO fglgallery.getImageCount(id)
    DISPLAY fglgallery.getPath(id, idx)
END FOR
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
