---
title: "Example 1: User type with a record structure"
source: "fgl-topics/c_fgl_user_types_example_1.html"
breadcrumb: "Language basics > Types > Examples > Example 1: User type with a record structure"
type: "concept"
description: "The example shows how to define a user type as a RECORD . TYPE t_customer RECORD cust_num INTEGER, cust_name VARCHAR(50), cust_addr VARCHAR(200) END RECORD MAIN DEFINE custrec t_customer DEFINE ..."
---

# Example 1: User type with a record structure

The example shows how to define a user type as a `RECORD`.

```
TYPE t_customer RECORD
            cust_num INTEGER,
            cust_name VARCHAR(50),
            cust_addr VARCHAR(200)
    END RECORD

MAIN
    DEFINE custrec t_customer
    DEFINE custarr DYNAMIC ARRAY OF t_customer

    LET custrec.cust_num = 123
    LET custrec.cust_name = "Mike Pantock"

    LET custarr[1] = custrec
    LET custarr[2] = custrec
    LET custarr[3] = custrec

END MAIN
```
