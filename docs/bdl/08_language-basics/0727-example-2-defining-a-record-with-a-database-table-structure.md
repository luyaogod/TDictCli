---
title: "Example 2: Defining a record with a database table structure"
source: "fgl-topics/c_fgl_records_005.html"
breadcrumb: "Language basics > Records > Examples > Example 2: Defining a record with a database table structure"
type: "concept"
---

# Example 2: Defining a record with a database table structure

> This example shows how to define a record with members using the same types as database columns.

```
SCHEMA stores 
DEFINE cust RECORD LIKE customer.*
MAIN
  DATABASE stores
  SELECT * INTO cust.* FROM customer WHERE customer_num=2
  DISPLAY cust.*
END MAIN
```
