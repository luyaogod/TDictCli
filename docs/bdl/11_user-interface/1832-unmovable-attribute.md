---
title: "UNMOVABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNMOVABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNMOVABLE attribute"
type: "concept"
---

# UNMOVABLE attribute

> The UNMOVABLE attribute prevents the user from moving a defined column of a table.

## Syntax

```
UNMOVABLE
```

## Usage

By default, a `TABLE` or `TREE` container allows the user
to move the columns by dragging and dropping the column header.

Use the `UNMOVABLE` attribute to prevent the user from changing the order of a
specific column.

Makes sense only for a field that is used for the definition of a column in a
`TABLE` or `TREE` container.

Typically, `UNMOVABLE` is used on at least two columns, to prevent the user from
changing the order of the input on these columns.

## Example

```
EDIT c01 = item.comment, UNMOVABLE;
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNMOVABLECOLUMNS attribute](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.")
