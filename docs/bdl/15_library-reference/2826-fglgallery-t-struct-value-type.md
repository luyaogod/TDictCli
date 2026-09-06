---
title: "fglgallery.t_struct_value type"
source: "fgl-topics/c_fgl_utility_functions_fglgallery_t_struct_value.html"
breadcrumb: "Library reference > Utility modules > fglgallery: Image gallery module > fglgallery.t_struct_value type"
type: "concept"
---

# fglgallery.t_struct_value type

> The t_struct_value type holds image selection data.

## Syntax

```
TYPE t_struct_value RECORD
    current INTEGER,
    selected DYNAMIC ARRAY OF INTEGER
  END RECORD
```

## Usage

This user-defined type defines a record structure to hold image selection information, to be used
in fglgallery web components.

Each picture is identified by its ordinal position in the image gallery.

- The `current` member defines the current picture in the gallery.
- The `selected` member is a dynamic array of integers, defining the selected
  images. This member is used, when multiple image selection is enabled with [`setMultipleSelection()`](2839-fglgallery-setmultipleselection.md "Enables/disables multiple picture selection in an fglgallery."). The current image can be different from a selected
  image. The current image may not be selected.

Image selection information is stored in the `WEBCOMPONENT` field value, as a JSON
formatted string similar to: `{"current":5,"selected":[4,5]}`. The variable bound to
the fglgallery `WEBCOMPONENT` form field must be defined as `VARCHAR`
or `STRING`.

The `t_struct_value` structure is typically used to get the selected image ids,
when the image selection action is fired by the fglgallery web component. The web component field
value can be parsed with the [`util.JSON.parse()`](3582-util-json-parse.md "Parses a JSON string and fills program variables with the values.") method, to fill a variable defined with the
`t_struct_value` type.

The web component field value can also be used to control the image selection from the program
code. To select specific images in the gallery, fill the `t_struct_value` variable
with image ids, then stringify the structure to JSON with the [`util.JSON.stringify()`](3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") method, and
set the web component field value with this JSON string.

## Example

In this code example, `rec.gallery_wc` is the name of the fglgallery web component
field:

```
DEFINE struct_value fglgallery.t_struct_value
...
   ON ACTION image_selection
      CALL util.JSON.parse( rec.gallery_wc, struct_value )
      DISPLAY struct_value.current,
              struct_value.selected.getLength()
...
   LET struct_value.current = 3
   LET struct_value.selected.clear()
   LET struct_value.selected[1] = 2
   LET struct_value.selected[2] = 7
   LET struct_value.selected[3] = 15
   LET rec.gallery_wc = util.JSON.stringify(struct_value)
...
```
