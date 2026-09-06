---
title: "Example 5: Type attributes"
source: "fgl-topics/c_fgl_variables_example_5.html"
breadcrumb: "Language basics > Variables > Examples > Example 5: Type attributes"
type: "concept"
description: "This example shows how to use type attributes. IMPORT util TYPE t_good RECORD num INTEGER ATTRIBUTES(json_name=\"Item Id\"), name VARCHAR(50) ATTRIBUTES(json_name=\"Item Name\") END RECORD TYPE t_load ..."
---

# Example 5: Type attributes

This example shows how to use type attributes.

```
IMPORT util

TYPE t_good RECORD
         num INTEGER ATTRIBUTES(json_name="Item Id"),
         name VARCHAR(50) ATTRIBUTES(json_name="Item Name")
     END RECORD

TYPE t_load RECORD
         goods DYNAMIC ARRAY ATTRIBUTES(json_name="The Goods") OF t_good
     END RECORD

MAIN
    DEFINE load t_load
    LET load.goods[1].num = 101
    LET load.goods[1].name = "Apples"
    DISPLAY util.JSON.format( util.JSON.stringify(load) )
END MAIN
```

Ouput:

```
{
    "The Goods": [
        {
            "Item Id": 101,
            "Item Name": "Apples"
        }
    ]
}
```
