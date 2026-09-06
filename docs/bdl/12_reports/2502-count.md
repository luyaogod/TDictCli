---
title: "COUNT"
source: "fgl-topics/c_fgl_reports_COUNT.html"
breadcrumb: "Reports > Report aggregate functions > COUNT"
type: "concept"
---

# COUNT

> Counts a number of rows in a report based on a condition.

## Syntax

```
[GROUP] COUNT(*) [ WHERE condition ]
```

1. condition is a boolean expression evaluated to compute the aggregate
   value.

## Usage

This aggregate report instruction returns the total number of records
qualified by the optional `WHERE` condition.

The `WHERE` condition is evaluated after any `OUTPUT TO REPORT` execution. Even
if it is typically used in `AFTER
GROUP OF` blocks, the aggregate expression is not evaluated in that block.
Changing values of the `WHERE` clause in the `AFTER GROUP`
context will not have an immediate effect.

Using the `GROUP` keyword causes the aggregate instructions to include only
data of the current group of records that have the same value for the variable that you specify
in the `AFTER GROUP OF` control block.

## Example

The following fragment of a report definition uses the `AFTER GROUP OF`
control block and `GROUP` keyword to form sets of records depending on how many
items are in each order. The last `PRINT`
statement calculates the total price of each order, adds a shipping charge, and prints the
result. Because no `WHERE` clause is specified here, `GROUP SUM()`
combines the *total\_price* of every item in the group included in the order.

```
  AFTER GROUP OF number 
    SKIP 1 LINE
    PRINT 4 SPACES, "Shipping charges for the order: ",
            ship_charge USING "$$$$.&&"
    PRINT 4 SPACES, "Count of small orders: ",
         GROUP COUNT(*) WHERE total_price < 200.00 USING "##,###"
    SKIP 1 LINE
    PRINT 5 SPACES, "Total amount for the order: ",
          ship_charge + GROUP SUM(total_price) USING "$$,$$$,$$$.&&"
```

## Related links

**Related concepts**  

[Report engine configuration](2508-report-engine-configuration.md "Report engine behavior can be controlled with FGLPROFILE settings.")

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")
