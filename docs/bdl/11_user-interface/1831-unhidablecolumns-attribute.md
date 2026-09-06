---
title: "UNHIDABLECOLUMNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNHIDABLECOLUMNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNHIDABLECOLUMNS attribute"
type: "concept"
---

# UNHIDABLECOLUMNS attribute

> The UNHIDABLECOLUMNS attribute indicates that the columns of the table cannot be hidden or shown by the user with the context menu.

## Syntax

```
UNHIDABLECOLUMNS
```

## Usage

When using the `UNHIDABLECOLUMNS` attribute in a `TABLE` or
`TREE` definition, the end user will not be allowed to hide columns.

The end user is also not allowed to show columns that are hidden by default with
`HIDDEN=USER`.

## Example

```
TABLE t1 ( UNHIDABLECOLUMNS )
```

## Related links

**Related concepts**  

[HIDDEN attribute](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.")

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNHIDABLE attribute](1830-unhidable-attribute.md "The UNHIDABLE attribute indicates that the element cannot be hidden or shown by the user with the context menu.")
