---
title: "Summary lines in tables"
source: "fgl-topics/c_fgl_ui_tables_aggregates.html"
breadcrumb: "User interface > User interface programming > Table views > Summary lines in tables"
type: "concept"
---

# Summary lines in tables

> Table views can display a summary line, to show aggregate values for columns.

To get a summary line in a table, define aggregate field item tags at the bottom of the
`TABLE` container, with the corresponding `AGGREGATE` form item
definitions in the `ATTRIBUTES` section.

Define the type of the aggregate field with the [`AGGREGATETYPE`](1760-aggregatetype-attribute.md "The AGGREGATETYPE attribute defines how the aggregate field value is computed.") attribute: The
aggregate value can be automatically computed, or set by program.

To get a global label for the summary line, specify the [`AGGREGATETEXT`](1759-aggregatetext-attribute.md "The AGGREGATETEXT attribute defines a label to be displayed for aggregate fields.") attribute at the
`TABLE` level. This aggregate label will appear on the left in the summary line, if
no aggregate text is defined at the aggregate field level.

Aggregate fields are not rendered when using a [`FLIPPED`](2319-controlling-table-rendering.md) table.

The following form definition file defines two aggregate fields for the second and third columns
of the table:

```
LAYOUT
TABLE
{
[c1   |c2                 |c3          ]
[c1   |c2                 |c3          ]
[c1   |c2                 |c3          ]
[c1   |c2                 |c3          ]
      [grt                |tot         ]
}
END
END

ATTRIBUTES
EDIT c1 = FORMONLY.pkey, TITLE="Num";
EDIT c2 = FORMONLY.name, TITLE="Name";
EDIT c3 = FORMONLY.price, TITLE="Price";
AGGREGATE grt = FORMONLY.o_great,
     AGGREGATETEXT = "Greatest: ",
     AGGREGATETYPE = MAX;
AGGREGATE tot = FORMONLY.o_total,
     AGGREGATETEXT = "Total:   ",
     AGGREGATETYPE = SUM;
END

INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.pkey THRU FORMONLY.price);
END
```

The program code implements a `DISPLAY ARRAY`:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
               pkey INT,
               name VARCHAR(20),
               price DECIMAL(10,2)
           END RECORD
    DEFINE x INT

    FOR x=1 TO 50
        LET arr[x].pkey = 100+x
        LET arr[x].name = SFMT("item%1",x)
        LET arr[x].price = (3.45 * x) + (1.23 * (x MOD 2))
    END FOR

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    DISPLAY ARRAY arr TO sr.* ATTRIBUTES(UNBUFFERED)

END MAIN
```

Visual result of the above sample:

![Table rendering aggregate fields for columns.](../_images/table_aggregate_1.jpg)

*Table with AGGREGATE fields*

For more details, see [Aggregate fields](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.").
