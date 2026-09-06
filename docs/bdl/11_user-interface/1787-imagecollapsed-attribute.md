---
title: "IMAGECOLLAPSED attribute"
source: "fgl-topics/c_fgl_FSFAttributes_IMAGECOLLAPSED.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > IMAGECOLLAPSED attribute"
type: "concept"
---

# IMAGECOLLAPSED attribute

> The IMAGECOLLAPSED attribute sets the global icon to be used when a tree node is collapsed.

## Syntax

```
IMAGECOLLAPSED = "image-name"
```

1. image-name is an image resource.

## Usage

This attribute is used in the definition of a `TREE` container, to
define the icon to be used for nodes that are collapsed.

It overwrites the program array image defined by `IMAGECOLUMN`, if
both are used.

This attribute is optional.

For more details about image resource specification, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

For more details about treeview programming, see [Tree views](2352-tree-views.md "Describes how to implement tree views.").

## Related links

**Related concepts**  

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")

[IMAGECOLUMN attribute](1786-imagecolumn-attribute.md "The IMAGECOLUMN attribute defines the form field containing the image for the current field.")
