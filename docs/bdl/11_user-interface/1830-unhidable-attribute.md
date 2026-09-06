---
title: "UNHIDABLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNHIDABLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNHIDABLE attribute"
type: "concept"
---

# UNHIDABLE attribute

> The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu.

## Syntax

```
UNHIDABLE
```

## Usage

By default, a `TABLE` or `TREE` container allows the user to
hide the columns by a right-click on the column header.

Use the `UNHIDABLE` attribute to prevent the user from hiding a specific
column.

The end user is also not allowed to show columns that are hidden by default with
`HIDDEN=USER`.

Makes sense only for a field that is used for the definition of a column in a
`TABLE` or `TREE` container.

## Example

```
EDIT c01 = item.comment, UNHIDABLE;
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[HIDDEN attribute](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.")

[UNHIDABLECOLUMNS attribute](1831-unhidablecolumns-attribute.md "The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu.")
