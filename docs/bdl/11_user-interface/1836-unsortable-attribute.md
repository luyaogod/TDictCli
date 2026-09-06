---
title: "UNSORTABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNSORTABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNSORTABLE attribute"
type: "concept"
---

# UNSORTABLE attribute

> The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting.

## Syntax

```
UNSORTABLE
```

## Usage

By default, a `TABLE` or `TREE` container allows the
user to sort the columns by a left-click on the column header.

Use the `UNSORTABLE` attribute to prevent a sort on a specific
column.

Makes sense only for a field that is used for the definition of a column in a
`TABLE` or `TREE` container.

## Example

```
EDIT c01 = item.comment, UNSORTABLE;
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNSORTABLECOLUMNS attribute](1837-unsortablecolumns-attribute.md "The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting.")
