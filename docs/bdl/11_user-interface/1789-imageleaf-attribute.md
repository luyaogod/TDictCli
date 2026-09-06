---
title: "IMAGELEAF attribute"
source: "fgl-topics/c_fgl_FSFAttributes_IMAGELEAF.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > IMAGELEAF attribute"
type: "concept"
---

# IMAGELEAF attribute

> The IMAGELEAF attribute defines the global icon for leaf nodes of a TREE container.

## Syntax

```
IMAGELEAF = "image-name"
```

1. image-name is an image resource.

## Usage

This attribute is used in the definition of a `TREE` container, to specify the
name of the icon that must be used for leaf nodes.

It overwrites the program array image defined by `IMAGECOLUMN`, if both are used.

This attribute is optional.

For more details about image resource specification, see [Providing the image resource](1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

For more details about treeview programming, see [Tree views](2352-tree-views.md "Describes how to implement tree views.").

## Related links

**Related concepts**  

[Using images](1583-using-images.md "Describes how to use pictures in the forms of your application.")
