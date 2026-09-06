---
title: "fglgallery.display()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_display.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.display()"
type: "concept"
---

# fglgallery.display()

> Displays an fglgallery to the end user.

## Syntax

```
FUNCTION display(
   id SMALLINT,
   type INTEGER,
   size INTEGER )
```

1. id is the gallery identifier, as returned by the `create()`
   function.
2. type is the type of rendering to be used for the gallery
   (`FGLGALLERY_TYPE_*`).
3. size is the size hint for the gallery rendering
   (`FGLGALLERY_SIZE_*`).

## Usage

This function displays the fglgallery in the corresponding web component field.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The function can be called several times. It is typically used to refresh the image gallery, when
the type or the size of the gallery needs to be changed.

The gallery type can be one off:

- `FGLGALLERY_TYPE_MOSAIC`
- `FGLGALLERY_TYPE_LIST`
- `FGLGALLERY_TYPE_THUMBNAILS`
- `FGLGALLERY_TYPE_SLIDESHOW`

The gallery size can be one off (ignored for
`FGLGALLERY_TYPE_SLIDESHOW`):

- `FGLGALLERY_SIZE_XSMALL`
- `FGLGALLERY_SIZE_SMALL`
- `FGLGALLERY_SIZE_NORMAL`
- `FGLGALLERY_SIZE_LARGE`
- `FGLGALLERY_SIZE_XLARGE`

If you only need to add images to the gallery without changing the gallery type or size, there is
no need to re-display the gallery: Use the [`addImage()`](2827-fglgallery-addimage.md "Adds a picture resource to an fglgallery.") and [`flush()`](2834-fglgallery-flush.md "Displays new added images to the end user.") functions
instead.

## Example

```
CALL fglgallery.display( id, FGLGALLERY_TYPE_MOSAIC, FGLGALLERY_SIZE_NORMAL )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
