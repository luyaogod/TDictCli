---
title: "Record initializers"
source: "fgl-topics/c_fgl_records_007.html"
breadcrumb: "Language basics > Records > Record initializers"
type: "concept"
---

# Record initializers

> Records can be initialized in their definition.

To initialize a record variable in its definition, use the equal sign followed by a record
initializer.

A record initializer is specified with parentheses, where record member values must be specified
as a `key:value` pair, separated by a
comma:

```
DEFINE rec RECORD
           cust_id INT,
           cust_name VARCHAR(50)
       END RECORD = ( cust_id: 101, cust_name: "Mike TORN" )
```

Below is a complex record example, using a sub-record and a dynamic
array:

```
DEFINE reader RECORD
            id INTEGER,
            name VARCHAR(100),
            birth DATE,
            address RECORD
                num VARCHAR(20),
                street VARCHAR(200),
                city_id INTEGER,
                state_id VARCHAR(5)
            END RECORD,
            book_ids DYNAMIC ARRAY OF INTEGER
       END RECORD = (
            id : 123,
            name : "Scott Spinley",
            birth : MDY(12,24,1998),
            address : (
                num : "2A",
                street : "Sunset Bld",
                city_id : 9834,
                state_id : "CA"
            ),
            book_ids : [ 234, 34, 458 ]
       )
```

For more details, see [Variable initializers](0691-variable-initializers.md "Variables can be initialized in their definition.").
