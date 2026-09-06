---
title: "fglgallery.setImageAspectRatio()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_setimageaspectratio.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.setImageAspectRatio()"
type: "concept"
---

# fglgallery.setImageAspectRatio()

> Defines the aspect ratio for gallery images.

## Syntax

```
FUNCTION setImageAspectRatio(
   id SMALLINT,
   ratio DECIMAL(5,2) )
```

1. id is the gallery identifier, as returned by the `create()`
   function.
2. ratio is the aspect ratio to be used (for example, 1.77 or 16/9).

## Usage

For all display types except `FGLGALLERY_TYPE_SLIDESHOW`, in order to align image
elements properly when image resources have different sizes, the fglgallery uses the same size for
each image.

The size of image elements is defined by the `FGLGALLERY_SIZE_*` parameter passed
to the `display()` function. This size parameter is relative to the current font size
(`em` unit is used in HTML).

The `setImageAspectRatio()` function defines the aspect ratio for fglgallery image
elements.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The [`display()`](2832-fglgallery-display.md "Displays an fglgallery to the end user.") function has to used after `setImageAspectRatio()` to
have the changes in the fglgallery take effect.

The aspect ratio is ignored when using the `FGLGALLERY_TYPE_SLIDESHOW` display
type. For other display types, a square aspect ratio (`1:1`) is used by default.

## Example

```
CALL fglgallery.setImageAspectRatio( id, 16/9 )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
