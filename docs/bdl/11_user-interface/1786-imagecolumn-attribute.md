---
title: "IMAGECOLUMN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_IMAGECOLUMN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > IMAGECOLUMN attribute"
type: "concept"
---

# IMAGECOLUMN attribute

> The IMAGECOLUMN attribute defines the form field containing the image for the current field.

## Syntax

```
IMAGECOLUMN = column-name
```

1. column-name is a form field name.

## Usage

The `IMAGECOLUMN` attribute allows an image to be displayed on the left of the
value of the column value. The image can be different for each row.

A typical usage is the `TREE` container: `IMAGECOLUMN` will allow a
row-specific image to be displayed left of the tree node text. You defined only one image column for
a tree node decoration.

When used in the definition of a `TABLE` column, the image and the column
will be displayed in the same table cell. There can be several `TABLE`
columns using an `IMAGECOLUMN`.

For `TREE` containers, the images defined by the [`IMAGECOLLAPSED`](1787-imagecollapsed-attribute.md "The IMAGECOLLAPSED attribute sets the global icon to be used when a tree node is collapsed."), [`IMAGEEXPANDED`](1788-imageexpanded-attribute.md "The IMAGEEXPANDED attribute sets the global icon to be used when a tree node is expanded.") and [`IMAGELEAF`](1789-imageleaf-attribute.md "The IMAGELEAF attribute defines the global icon for leaf nodes of a TREE container.") attributes take
precedence over the images defined by the `IMAGECOLUMN` cell.

This attribute references form field that contains the name of an image. This form field
must be defined as a `PHANTOM` form field, that will be part of the screen
record definition in the `INSTRUCTIONS` section.

For more details about image resource specification in the `PHANTOM`
column, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

## Example

```
...
ATTRIBUTES
PHANTOM FORMONLY.icon;
EDIT FORMONLY.file_name, IMAGECOLUMN=icon;
...
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.icon, FORMONLY.file_name, ...);
...
```

## Related links

**Related concepts**  

[Phantom fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).")

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")

[Tree views](2352-tree-views.md "Describes how to implement tree views.")
