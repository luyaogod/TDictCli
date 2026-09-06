---
title: "fglgallery.setMultipleSelection()"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_setmultipleselection.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.setMultipleSelection()"
type: "concept"
---

# fglgallery.setMultipleSelection()

> Enables/disables multiple picture selection in an fglgallery.

## Syntax

```
FUNCTION setMultipleSelection(
   id SMALLINT,
   on BOOLEAN )
```

1. id is the fglgallery identifier, as returned by the `create()`
   function.
2. on can be `TRUE` (to enable multi-selection) or
   `FALSE` (to disable multi-selection).

## Usage

This function controls multiple picture selection in the specified image gallery.

The id parameter is the fglgallery handler returned by the [`create()`](2829-fglgallery-create.md "Creates a new fglgallery handle.") function.

When multiple selection is enabled, the end user can mark pictures as selected.

Indexes of selected images are provided in the web component field value, using the following
JSON notation: `{"current":5,"selected":[4,5]}`. Convert this JSON string to a [fglgallery.t\_struct\_value](2826-fglgallery-t-struct-value-type.md "The t_struct_value type holds image selection data.")
record to get the list of picture indexes in the `"selected"` member.

## Example

```
CALL fglgallery.setMultipleSelection( id, TRUE )
```

## Related links

**Related concepts**  

[The fglgallery web component](../11_user-interface/2418-the-fglgallery-web-component.md "The fglgallery built-in web component shows a collection of pictures the end user can choose from.")
