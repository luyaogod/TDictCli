---
title: "fglgallery.deleteImages()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_deleteimages.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.deleteImages()"
type: "concept"
---

# fglgallery.deleteImages()

> Deletes pictures used in an fglgallery.

## Syntax

```
FUNCTION deleteImages(
   id SMALLINT,
   indexes DYNAMIC ARRAY OF INTEGER )
  RETURNS STRING
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.
2. indexes is the dynamic array of indexes of the pictures to remove from the
   gallery.

## Usage

This function deletes the specified pictures from the fglgallery.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The indexes parameter is a `DYNAMIC ARRAY OF INTEGER` that
defines a set of picture indexes. Each corresponding picture reference will be removed from the
image gallery.

The new gallery is automatically displayed to the end user.

## Example

```
DEFINE todel DYNAMIC ARRAY OF INTEGER
LET todel[1] = 3
LET todel[2] = 6
CALL fglgallery.deleteImages(id, todel)
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
