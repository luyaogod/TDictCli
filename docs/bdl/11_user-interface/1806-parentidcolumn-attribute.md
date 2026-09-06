---
title: "PARENTIDCOLUMN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_PARENTIDCOLUMN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > PARENTIDCOLUMN attribute"
type: "concept"
---

# PARENTIDCOLUMN attribute

> The PARENTIDCOLUMN attribute specifies the form field that contains the identifier of the parent node of a tree node.

## Syntax

```
PARENTIDCOLUMN = column-name
```

1. column-name is a form field name.

## Usage

This attribute is used in the definition of a `TREE` container, to define the name
of the form field containing the identifier of the tree node that is the parent of the current node
in a tree view.

You must specify form field column names, not item tag identifiers.

This attribute is mandatory.

For more details about treeview programming, see [Tree views](2352-tree-views.md "Describes how to implement tree views.").

## Related links

**Related concepts**  

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")

[IDCOLUMN attribute](1784-idcolumn-attribute.md "The IDCOLUMN attribute specifies the form field that contains the identifier of a tree node.")
