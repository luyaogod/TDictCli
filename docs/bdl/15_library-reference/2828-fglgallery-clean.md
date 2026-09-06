---
title: "fglgallery.clean()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_clean.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.clean()"
type: "concept"
---

# fglgallery.clean()

> Removes all pictures from an fglgallery.

## Syntax

```
FUNCTION clean( id SMALLINT )
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.

## Usage

This function removes all pictures from the fglgallery.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

The cleaned gallery is automatically displayed to the end user.

## Example

```
CALL fglgallery.clean( id )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
