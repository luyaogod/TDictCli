---
title: "AGGREGATETYPE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_AGGREGATETYPE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > AGGREGATETYPE attribute"
type: "concept"
---

# AGGREGATETYPE attribute

> The AGGREGATETYPE attribute defines how the aggregate field value is computed.

## Syntax

```
AGGREGATETYPE = { PROGRAM | SUM | AVG | MIN | MAX | COUNT }
```

## Usage

`PROGRAM` specifies that the aggregate value will be computed and displayed
by the program code.

An aggregate type different from `PROGRAM` specifies that the aggregate value
is computed automatically:

- `SUM` computes the total of all values of the corresponding
  numeric column.
- `AVG` computes the average of all values of the
  corresponding numeric column.
- `MIN` displays the minimum value of the corresponding
  numeric column.
- `MAX` displays the maximum value of the corresponding
  numeric column.
- `COUNT` computes the number of rows.

The `SUM` and `AVG` aggregate types apply to data types that
can be used as operand for an addition, such as `INTEGER`, `DECIMAL`,
`INTERVAL`.

The `MIN` and `MAX` aggregate types apply to data types that
can be compared, such as `INTEGER`, `DECIMAL`, `INTERVAL`,
`CHAR`, `DATETIME`.

When all values of a column are `NULL`, the value of a computed aggregate is
`NULL` and show up empty.

With `SUM` and `AVG` aggregate types, the `NULL`
values are equivalent to zero. With `MIN` and `MAX` aggregate types,
`NULL` values are ignored.

Aggregate fields are not rendered when using a [`FLIPPED`](1780-flipped-attribute.md "The FLIPPED attribute flips TABLE columns into rows.") table.

For more details, see [Summary lines in tables](2329-summary-lines-in-tables.md "Table views can display a summary line, to show aggregate values for columns.").

## Example

```
AGGREGATE tot = FORMONLY.total, AGGREGATETYPE=PROGRAM;
```

## Related links

**Related concepts**  

[AGGREGATE item definition](1728-aggregate-item-definition.md "Defines screen-record fields that hold computed values to be displayed as footer cells in a TABLE container.")

[TABLE item definition](1746-table-item-definition.md "Defines attributes for a table layout tag.")
