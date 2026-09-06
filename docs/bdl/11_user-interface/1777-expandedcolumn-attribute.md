---
title: "EXPANDEDCOLUMN attribute"
source: "fgl-topics/c_fgl_FSFAttributes_EXPANDEDCOLUMN.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > EXPANDEDCOLUMN attribute"
type: "concept"
---

# EXPANDEDCOLUMN attribute

> The EXPANDEDCOLUMN attribute specifies the form field that indicates whether a tree node is expanded.

## Syntax

```
EXPANDEDCOLUMN = column-name
```

1. column-name is the name of the form field holding the flag indicating whether
   a tree node is expanded (opened).

## Usage

This attribute is used in the definition of a `TREE` container.

You must specify form field column names, not item tag identifiers.

This attribute is optional.

For more details about treeview programming, see [Tree views](2352-tree-views.md "Describes how to implement tree views.").

## Related links

**Related concepts**  

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[Item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container.")
