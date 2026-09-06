---
title: "Attributes on array definitions"
source: "fgl-topics/c_fgl_Arrays_attributes.html"
breadcrumb: "Language basics > Arrays > Attributes on array definitions"
type: "concept"
---

# Attributes on array definitions

> Arrays can be defined with attributes, to complete the type description.

## Syntax

In type specifications, the attributes-list clause is:

```
{ ATTRIBUTE | ATTRIBUTES } ( attribute [ = "value" ] [,...] )
```

1. attribute is the name of a definition
   attribute.
2. value is the value for the definition attribute, it is
   optional for boolean attributes.

## Usage

Arrays can be defined with the `ATTRIBUTES()` clause, to specify meta-data
information for the array type.

In the next code example, the array type definition gets definition attributes, to specify the
corresponding field names for JSON serialization:

```
IMPORT util
  
TYPE t_items DYNAMIC ARRAY ATTRIBUTES(json_name="items") OF
     RECORD
         item_num INTEGER ATTRIBUTES(json_name="id"),
         item_desc VARCHAR(50)
     END RECORD

DEFINE stock RECORD
         itemlist t_items
     END RECORD

MAIN

    LET stock.itemlist[1].item_num = 998
    LET stock.itemlist[1].item_desc = "Hand gloves"
    LET stock.itemlist[2].item_num = 999
    LET stock.itemlist[2].item_desc = "Blue hat"

    DISPLAY util.JSON.format(util.JSON.stringify(stock))

END MAIN
```

Output:

```
{
    "items": [
        {
            "id": 998,
            "item_desc": "Hand gloves"
        },
        {
            "id": 999,
            "item_desc": "Blue hat"
        }
    ]
}
```

For more details, see [Type attributes](0756-type-attributes.md "Types can be defined with meta-data information.").

## Attributes meta-data belong to the type

When not using a user-defined [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."), a variable definition with a primitive type or complex type such as a
`RECORD`, `DYNAMIC ARRAY` or `DICTIONARY`, creates an
[anonymous type](0755-anonymous-types.md "Anonymous types are created from DEFINE instructions.").

If the `ATTRIBUTES` clause is used, this meta-data information belongs to the type
definition, it does not belong to the variable.

## Related links

**Related concepts**  

[Type attributes](0756-type-attributes.md "Types can be defined with meta-data information.")
