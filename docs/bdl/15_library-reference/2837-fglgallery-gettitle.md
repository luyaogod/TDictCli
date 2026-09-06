---
title: "fglgallery.getTitle()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_gettitle.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.getTitle()"
type: "concept"
---

# fglgallery.getTitle()

> Returns the description of a picture in an fglgallery.

## Syntax

```
FUNCTION getTitle(
   id SMALLINT,
   index STRING )
  RETURNS STRING
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.
2. index is the picture index in the image gallery.

## Usage

This function returns the title of a picture in the fglgallery.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The index parameter identifies the picture by its ordinal position in the
image gallery.

## Example

```
DISPLAY fglgallery.getTitle( id, 10 )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
