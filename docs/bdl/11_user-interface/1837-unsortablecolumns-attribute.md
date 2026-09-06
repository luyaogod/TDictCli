---
title: "UNSORTABLECOLUMNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNSORTABLECOLUMNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNSORTABLECOLUMNS attribute"
type: "concept"
---

# UNSORTABLECOLUMNS attribute

> The UNSORTABLECOLUMNS attribute indicates that the columns of the table cannot be selected by the user for sorting.

## Syntax

```
UNSORTABLECOLUMNS
```

## Usage

When using this attribute in a `TABLE` or `TREE` definition, the
end user will not be allowed to sort rows.

## Example

```
TABLE t1 ( UNSORTABLECOLUMNS )
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNSORTABLE attribute](1836-unsortable-attribute.md "The UNSORTABLE attribute indicates that the element cannot be selected by the user for sorting.")
