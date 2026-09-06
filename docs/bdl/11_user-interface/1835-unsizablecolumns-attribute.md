---
title: "UNSIZABLECOLUMNS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UNSIZABLECOLUMNS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UNSIZABLECOLUMNS attribute"
type: "concept"
---

# UNSIZABLECOLUMNS attribute

> The UNSIZABLECOLUMNS attribute indicates that the columns of the table cannot be resized by the user.

## Syntax

```
UNSIZABLECOLUMNS
```

## Usage

When using this attribute in a `TABLE` or `TREE` definition,
the end user will not be allowed to resize the columns.

## Example

```
TABLE t1 ( UNSIZABLECOLUMNS )
```

## Related links

**Related concepts**  

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[TREE container](1725-tree-container.md "The TREE container defines the presentation of a list of ordered records in a tree-view widget.")

[UNSIZABLE attribute](1834-unsizable-attribute.md "The UNSIZABLE attribute indicates that the element cannot be resized by the user.")
