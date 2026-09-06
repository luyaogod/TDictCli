---
title: "Example 2: User type defined in a module"
source: "fgl-topics/c_fgl_user_types_example_2.html"
breadcrumb: "Language basics > Types > Examples > Example 2: User type defined in a module"
type: "concept"
description: "This example shows how to use a type defined in another module. The module \" type_order.4gl \": PUBLIC TYPE t_order RECORD order_num INTEGER, store_num INTEGER, order_date DATE, cust_num INTEGER, ..."
---

# Example 2: User type defined in a module

This example shows how to use a type defined in another module.

The module "type\_order.4gl":

```
PUBLIC TYPE t_order RECORD
                 order_num INTEGER,
                 store_num INTEGER,
                 order_date DATE,
                 cust_num INTEGER,
                 fac_code CHAR(3)
             END RECORD
```

The main
program:

```
IMPORT FGL type_order
  
MAIN
  DEFINE o type_order.t_order

  CONNECT TO "custdemo"

  DECLARE order_c CURSOR FOR
     SELECT orders.*
        FROM orders ORDER BY cust_num
  START REPORT order_list
  FOREACH order_c INTO o.*
    OUTPUT TO REPORT order_list(o)
  END FOREACH
  FINISH REPORT order_list

END MAIN

REPORT order_list(ro t_order)
  FORMAT
    ON EVERY ROW
     PRINT ro.order_num, ro.order_date
END REPORT
```
