---
title: "UNSIZABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNSIZABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNSIZABLE attribute"
type: "concept"
---

# UNSIZABLE attribute

> The UNSIZABLE attribute indicates that the element cannot be resized by the user.

## Syntax

```
UNSIZABLE
```

## Usage

By default, a `TABLE` or `TREE` container allows the user to
resize the columns by a drag-click on the column header.

Use the `UNSIZABLE` attribute to prevent a resize on a specific column.

Makes sense only for a field that is used for the definition of a column in a
`TABLE` or `TREE` container.

## Example

```
EDIT c01 = item.comment, UNSIZABLE;
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNSIZABLECOLUMNS attribute](1835-unsizablecolumns-attribute.md "The UNSIZABLECOLUMNS attribute indicates that the columns of the table cannot be resized by the user.")
