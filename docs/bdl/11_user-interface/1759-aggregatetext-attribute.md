---
title: "AGGREGATETEXT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_AGGREGATETEXT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > AGGREGATETEXT attribute"
type: "concept"
---

# AGGREGATETEXT attribute

> The AGGREGATETEXT attribute defines a label to be displayed for aggregate fields.

## Syntax

```
AGGREGATETEXT = [%]"string"
```

1. string defines the label to be associated with the aggregate cell, with the %
   prefix it is a localized string.

## Usage

The `AGGREGATETEXT` attribute can be specified at the [`AGGREGATE` field level](1728-aggregate-item-definition.md "Defines screen-record fields that hold computed values to be displayed as footer cells in a TABLE container."),
or globally at the `TABLE` level, to define a label for the whole summary line.

When defining the `AGGREGATETEXT` attribute at the aggregate field level, the text
will be anchored to the value cell.

If the `AGGREGATETEXT` attribute is specified at the [`TABLE`](1746-table-item-definition.md "Defines attributes for a table layout tag.") level, the label will
appear on the left in the summary line.

When an aggregate text is defined at both levels, the global aggregate text of the table will be
ignored.

For more details, see [Summary lines in tables](2329-summary-lines-in-tables.md "Table views can display a summary line, to show aggregate values for columns.").

## Example

```
AGGREGATE tot = FORMONLY.total, AGGREGATETEXT="Total:";
```

## Related links

**Related concepts**  

[Aggregate fields](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
