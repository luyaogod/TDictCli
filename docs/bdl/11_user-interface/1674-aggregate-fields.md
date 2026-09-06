---
title: "Aggregate fields"
source: "fgl-topics/c_fgl_FormSpecFiles_AGGREGATE_Fields.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form fields > Aggregate fields"
type: "concept"
---

# Aggregate fields

> An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.

## Syntax

```
AGGREGATE item-tag = field-name [ , attribute-list ] ;
```

1. item-tag is an identifier that defines the name of the item tag in the layout section.
2. field-name identifies the name of the screen record field.
3. attribute-list defines the aspect and behavior of the form item.

## Usage

An `AGGREGATE` field defines a form field that is used to display a
summary cell for a given column of a `TABLE` container. The aggregate
fields are displayed after the last data line of the table. Such fields are
typically used to show computed values for the corresponding column which appears
above the aggregate cell.

An aggregate field can be based on a database column defined in a schema file, or as
`FORMONLY` field.

The `AGGREGATETYPE` attribute defines how the value of the field will
be computed. For example, the `SUM` keyword (the default) can be used
to instruct the runtime system to automatically compute the total of the associated
column. By using the `PROGRAM` keyword, you indicate that the value
of the aggregate field will be computed and displayed by program code. A simple
`DISPLAY BY NAME` or `DISPLAY TO` can be used to show
the summary value.

The value displayed in the `AGGREGATE` field follows the
`FORMAT` attribute of the corresponding column, if defined. The
`FORMAT` attribute is applied for automatically computed values, as well as for
values displayed by user code with `DISPLAY BY NAME` or `DISPLAY
TO`.

The label of an aggregate field can be specified with the `AGGREGATETEXT` attribute. The text
defined with this attribute will be displayed on the left of the aggregate value (in the aggregate
cell), except if there is no room to display the label (for example if the aggregate value is too
large or if the column values are aligned to the left). An aggregate label can be a localized string
with the `%"..."` string syntax. You can also specify an `AGGREGATETEXT`
attribute at the `TABLE` level, to get a global label for the summary line. If no
text is defined for an aggregate field, the global aggregate text will appear on the left in the
summary line.

Table aggregate decoration can be modified with a presentation style. Use the
`summaryLine` [pseudo-selector](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.")
to change the font type and color, as well as the background of the summary line.

The item tag of an aggregate field must appear in the last line in the layout block of the [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container, and must be
aligned vertically with a table column item tag. You can specify several aggregate item tags for the
same table:

```
TABLE
{
[c1      |c2       |c3           |c4       |c5      ]
[c1      |c2       |c3           |c4       |c5      ]
[c1      |c2       |c3           |c4       |c5      ]
[c1      |c2       |c3           |c4       |c5      ]
         [cnt      ]             [tot_c4   |tot_c5  ]
}
END
```

## Example

```
SCHEMA stores
LAYOUT( TEXT = "Orders" )
TABLE table1
{
 Num      Date            Order total
[c1      |c2             |c3           ]
[c1      |c2             |c3           ]
[c1      |c2             |c3           ]
[c1      |c2             |c3           ]
                         [total        ]
}
END
END
TABLES
orders, items
END
ATTRIBUTES
  EDIT c1 = orders.order_num;
  EDIT c2 = orders.order_date;
  EDIT c3 = items.total_price;
  AGGREGATE total = FORMONLY.order_total,
                       AGGREGATETEXT = "Total:",
                       AGGREGATETYPE = SUM;
END
INSTRUCTIONS
  SCREEN RECORD sr(
     orders.order_num,
     orders.order_date,
     items.total_price
  );
END
```

## Related links

**Related concepts**  

[AGGREGATE item definition](1728-aggregate-item-definition.md "Defines screen-record fields that hold computed values to be displayed as footer cells in a TABLE container.")
