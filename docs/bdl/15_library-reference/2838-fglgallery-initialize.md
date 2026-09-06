---
title: "fglgallery.initialize()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_initialize.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.initialize()"
type: "concept"
---

# fglgallery.initialize()

> Prepares the fglgallery library for use.

## Syntax

```
FUNCTION initialize( )
```

## Usage

This function initializes the fglgallery module usage. When the fglgallery library is no longer
needed, call the finalization function [`fglgallery.finalize()`](2833-fglgallery-finalize.md "Releases the fglgallery library.").

Initialization and finalization functions can be called several times by different modules using
the fglgallery library.

## Example

```
IMPORT FGL fglgallery
FUNCTION show_image_gallery()
  ...
  CALL fglgallery.initialize()
  ...
  CALL fglgallery.finalize()
END FUNCTION
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
