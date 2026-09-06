---
title: "UNMOVABLECOLUMNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNMOVABLECOLUMNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNMOVABLECOLUMNS attribute"
type: "concept"
---

# UNMOVABLECOLUMNS attribute

> The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.

## Syntax

```
UNMOVABLECOLUMNS
```

## Usage

When using this attribute in a `TABLE` or `TREE` definition,
the end user will not be allowed to move columns around.

## Example

```
TABLE t1 ( UNMOVABLECOLUMNS )
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNMOVABLE attribute](1832-unmovable-attribute.md "The UNMOVABLE attribute prevents the user from moving a defined column of a table.")
