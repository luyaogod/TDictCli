---
title: "Example 1: Function fetching customer number"
source: "fgl-topics/c_fgl_Functions_example_1.html"
breadcrumb: "Language basics > Functions > Examples > Example 1: Function fetching customer number"
type: "concept"
description: "The example shows a typical function definition returning an INTEGER . FUNCTION findCustomerNumber(name VARCHAR(50)) RETURNS INTEGER DEFINE num INTEGER CONSTANT sqltxt = \"SELECT cust_num FROM customer ..."
---

# Example 1: Function fetching customer number

The example shows a typical function definition returning an
`INTEGER`.

```
FUNCTION findCustomerNumber(name VARCHAR(50)) RETURNS INTEGER
  DEFINE num INTEGER
  CONSTANT sqltxt = "SELECT cust_num FROM customer WHERE cust_name = ?"
  PREPARE stmt FROM sqltxt 
  EXECUTE stmt INTO num USING name 
  IF sqlca.sqlcode = 100 THEN
     LET num = -1
  END IF
  RETURN num 
END FUNCTION
```
