---
title: "fglgallery.create()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_create.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.create()"
type: "concept"
---

# fglgallery.create()

> Creates a new fglgallery handle.

## Syntax

```
FUNCTION create( name STRING )
  RETURNS SMALLINT
```

1. name is the name of the `WEBCOMPONENT` form field.

## Usage

This function creates a new fglgallery handler.

The window/form containing the fglgallery web component field must be created before calling this
function.

The id returned by this function identifies the web component field in subsequent fglgallery API
calls.

When the fglgallery web component is no longer required, use the [`destroy()`](2831-fglgallery-destroy.md "Frees resources allocated for an fglgallery.") function,
to free the resources allocated for this fglgallery object.

## Example

```
DEFINE id SMALLINT
LET id = fglgallery.create("formonly.gallery")
...
CALL fglgallery.destroy(id)
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
