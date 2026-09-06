---
title: "AGGREGATE item definition"
source: "fgl-topics/c_fgl_FormSpecFiles_AGGREGATE_Item_Type.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section > AGGREGATE item definition"
type: "concept"
---

# AGGREGATE item definition

> Defines screen-record fields that hold computed values to be displayed as footer cells in a TABLE container.

## Syntax

```
AGGREGATE item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in
   the layout section.
2. field-name identifies the name of the screen record field.
3. attribute-list defines the aspect and behavior of the form item.

## Attributes

[`AGGREGATETEXT`](1759-aggregatetext-attribute.md "The AGGREGATETEXT attribute defines a label to be displayed for aggregate fields."), [`AGGREGATETYPE`](1760-aggregatetype-attribute.md "The AGGREGATETYPE attribute defines how the aggregate field value is computed.").

## Usage

Aggregate fields must be declared with an `AGGREGATE` element in the
`ATTRIBUTES` section.

For more details see [Aggregate fields](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.").

## Example

```
AGGREGATE total = FORMONLY.o_total,
                  AGGREGATETEXT = "Total:",
                  AGGREGATETYPE = SUM;
```

## Related links

**Related concepts**  

[Summary lines in tables](2329-summary-lines-in-tables.md "Table views can display a summary line, to show aggregate values for columns.")
