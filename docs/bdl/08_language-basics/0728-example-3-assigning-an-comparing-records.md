---
title: "Example 3: Assigning an comparing records"
source: "fgl-topics/c_fgl_records_006.html"
breadcrumb: "Language basics > Records > Examples > Example 3: Assigning an comparing records"
type: "concept"
---

# Example 3: Assigning an comparing records

> This example shows how to use records with comparison and assignment operators.

```
TYPE t_cust RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
     END RECORD
MAIN
  DEFINE cust1, cust2 t_cust

  LET cust1.cust_id = 999
  LET cust1.cust_name = "Scott Spinley"

  LET cust2 = cust1 -- Legacy syntax: LET cust2.* = cust1.*

  IF cust1.* = cust2.* THEN
     DISPLAY "1: Records are equal."
  END IF

  LET cust1.cust_name = NULL

  IF cust1.* != cust2.* THEN
     DISPLAY "2: Records are different."
  END IF

END MAIN
```
