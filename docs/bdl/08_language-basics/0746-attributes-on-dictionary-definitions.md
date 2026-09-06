---
title: "Attributes on dictionary definitions"
source: "fgl-topics/c_fgl_Dictionary_attributes.html"
breadcrumb: "Language basics > Dictionaries > Attributes on dictionary definitions"
type: "concept"
---

# Attributes on dictionary definitions

> Dictionaries can be defined with attributes, to complete the type description.

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

Dictionaries can be defined with the `ATTRIBUTES()` clause, to specify meta-data
information for the dictionary type.

In the next code example, the dictionary type definition gets definition attributes, to specify
the corresponding field names for JSON serialization:

```
IMPORT util
  
TYPE t_items DICTIONARY ATTRIBUTES(json_name="items") OF
     RECORD
         item_num INTEGER ATTRIBUTES(json_name="id"),
         item_desc VARCHAR(50)
     END RECORD

DEFINE stock RECORD
         itemlist t_items
     END RECORD

MAIN

    LET stock.itemlist["cf311"].item_num = 998
    LET stock.itemlist["cf311"].item_desc = "Hand gloves"
    LET stock.itemlist["cx242"].item_num = 999
    LET stock.itemlist["cx242"].item_desc = "Blue hat"

    DISPLAY util.JSON.format(util.JSON.stringify(stock))

END MAIN
```

Output:

```
{
    "items": {
        "cf311": {
            "id": 998,
            "item_desc": "Hand gloves"
        },
        "cx242": {
            "id": 999,
            "item_desc": "Blue hat"
        }
    }
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
