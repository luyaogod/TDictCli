---
title: "FORMAT EVERY ROW"
source: "fgl-topics/c_fgl_reports_FORMAT_EVERY_ROW.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > FORMAT EVERY ROW"
type: "concept"
---

# FORMAT EVERY ROW

> Default format specification of a report.

A report routine written with `FORMAT EVERY ROW` formats the report in a simple
default format, containing only the values that are passed to the `REPORT` program
block through its arguments, and the names of the arguments. You cannot modify the `EVERY
ROW` statement with any of the statements listed in report execution statements, and
neither can you include any control blocks in the `FORMAT` section.

The report engine uses as column headings the names of the variables that the report driver
passes as arguments at runtime. If all fields of each input record can fit horizontally on a
single line, the default report prints the names across the top of each page and the values
beneath. Otherwise, it formats the report with the names down the left side of the page and the
values to the right, as in the previous example. When a variable contains a null value, the
default report prints only the name of the variable, with nothing for the value.

The following example is a brief report specification that uses `FORMAT EVERY ROW`.
We assume here that the cursor that retrieved the input records for this report was declared with
an `ORDER BY` clause, so that no `ORDER BY` section is needed in this
report definition:

```
DATABASE stores

REPORT simple( order_num, customer_num, order_date )

  DEFINE order_num LIKE orders.order_num,
        customer_num LIKE orders.customer_num,
        order_date LIKE orders.order_date 

  FORMAT EVERY ROW

END REPORT
```

The example would produce the following output:

```
order_num customer_num order_date
     1001         104  01/20/1993
     1002         101  06/01/1993
     1003         104  10/12/1993
     1004         106  04/12/1993
     1005         116  12/04/1993
     1006         112  09/19/1993
     1007         117  03/25/1993
     1008         110  11/17/1993
     1009         111  02/14/1993
     1010         115  05/29/1993
     1011         104  03/23/1993
     1012         117  06/05/1993
```
