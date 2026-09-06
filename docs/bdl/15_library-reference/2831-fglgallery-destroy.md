---
title: "fglgallery.destroy()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_destroy.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.destroy()"
type: "concept"
---

# fglgallery.destroy()

> Frees resources allocated for an fglgallery.

## Syntax

```
FUNCTION destroy( id SMALLINT )
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.

## Usage

This function releases the fglgallery identified by the specified handler.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

Call this method after closing the window/form displaying the gallery web component.

## Example

```
CALL fglgallery.destroy( id )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
