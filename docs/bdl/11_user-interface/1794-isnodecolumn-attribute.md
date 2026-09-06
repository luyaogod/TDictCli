---
title: "ISNODECOLUMN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ISNODECOLUMN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ISNODECOLUMN attribute"
type: "concept"
---

# ISNODECOLUMN attribute

> The ISNODECOLUMN attribute specifies the form field that indicates whether a tree node has children.

## Syntax

```
ISNODECOLUMN = column-name
```

1. column-name is a form field name.

## Usage

This attribute is used in the definition of a `TREE` container, to specify the
name of the form field indicating whether a tree node has children.

Even if the program node does not contain child nodes for this tree node, this
attribute may be used, to implement dynamic filling of tree views.

You must specify form field column names, not item tag identifiers.

This attribute is optional.

For more details about treeview programming, see [Tree views](2352-tree-views.md "Describes how to implement tree views.").

## Related links

**Related concepts**  

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
