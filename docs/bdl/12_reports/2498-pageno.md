---
title: "PAGENO"
source: "fgl-topics/c_fgl_reports_PAGENO.html"
breadcrumb: "Reports > Report operators > PAGENO"
type: "concept"
---

# PAGENO

> Contains the current page number in a report.

## Syntax

```
PAGENO
```

## Usage

This operator takes no operand but returns the number of the page the report engine is
currently printing.

You can use `PAGENO` in the `PAGE HEADER` or `PAGE TRAILER` block, or in other
control blocks to number the pages of a report sequentially.

## Example

If you use the SQL aggregate `COUNT(*)` in the `SELECT`
statement to find how many records are returned by the query, and if the number of records
that appear on each page of output is both fixed and known, you can calculate the total
number of pages, as in this example:

```
  FIRST PAGE HEADER
    SELECT COUNT(*) INTO cnt FROM customer 
    LET y = cnt/50 -- Assumes 50 records per page 
  ON EVERY ROW
    PRINT COLUMN 10, r.customer_num, ...
  PAGE TRAILER
    PRINT PAGE PAGENO USING "<<" OF cnt USING "<<"
```

If the calculated number of pages was 20, the first page trailer would be:

```
Page 1 of 20
```

`PAGENO` increments with each page, so the last page trailer would be:

```
Page 20 of 20
```

## Related links

**Related concepts**  

[START REPORT](2470-start-report.md "The START REPORT instruction initializes a report execution.")
