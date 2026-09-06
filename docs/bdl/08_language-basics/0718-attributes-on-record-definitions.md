---
title: "Attributes on record definitions"
source: "fgl-topics/c_fgl_records_012.html"
breadcrumb: "Language basics > Records > Attributes on record definitions"
type: "concept"
---

# Attributes on record definitions

> Records can be defined with attributes, to complete the type description.

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

Records can be defined with the `ATTRIBUTES()` clause, to specify meta-data
information for the record type.

In the next code example, the `t_cust` record type gets definition attributes to
specify the corresponding field names for JSON serialization:

```
IMPORT util
  
TYPE t_cust RECORD ATTRIBUTES(json_name="customer")
         cust_id INTEGER ATTRIBUTES(json_name="id"),
         cust_name VARCHAR(50)
     END RECORD

TYPE t_custinfo RECORD
         cust t_cust
     END RECORD

MAIN
    DEFINE ci t_custinfo

    LET ci.cust.cust_id = 998
    LET ci.cust.cust_name = "Scott Finley"

    DISPLAY util.JSON.format(util.JSON.stringify(ci))

END MAIN
```

Output:

```
{
    "customer": {
        "id": 998,
        "cust_name": "Scott Finley"
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
