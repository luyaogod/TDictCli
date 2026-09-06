---
title: "Function attributes"
source: "fgl-topics/c_fgl_Functions_attributes.html"
breadcrumb: "Language basics > Functions > Function attributes"
type: "concept"
---

# Function attributes

> Function attributes can be used to add definition information about the function, its parameters and its return values.

## Syntax

In type specifications, the attributes-list clause is:

```
{ ATTRIBUTE | ATTRIBUTES } ( attribute [ = "value" ] [,...] )
```

1. attribute is the name of a definition
   attribute.
2. value is the value for the definition attribute, it is
   optional for boolean attributes.

## Purpose of function attributes

Function attributes can be used to define properties for the function itself, for the parameters
accepted by the function, and for the returned values/types.

Function attributes are typically used to define [RESTful Web services (high-level framework)](../16_web-services/4705-restful-web-services-high-level-framework.md "These topics give you the information you need to begin working with RESTful Web services applications using BDL function with support for attributes.").

## Usage

Function, parameter and return value attributes are specified with the
`ATTRIBUTES()` clause. The `ATTRIBUTES()` clause contains a
comma-separated list of attributes. An attribute can be a
`name=value` pair, or single
`name` for a boolean
attribute:

```
ATTRIBUTES( WSScope='book.write', WSPut, WSPath='{id}', ... )
```

The `ATTRIBUTES()` clause can be specified:

- after a type of parameter in the ( ) parameter list, to define parameter attributes.
- after the closing brace of the parameter list, to define function attributes.
- after a type in the `RETURNS` clause, to define return value attributes.

## Example

```
TYPE profileType RECORD ... END RECORD

PUBLIC FUNCTION UpdateUserProfile(
    id STRING ATTRIBUTE(WSParam),
    thisUser profileType )
  ATTRIBUTES (WSPut,
              WSPath="/users/{id}",
              WSDescription="Update user with the given id",
              WSThrows='404:user not found')
  RETURNS STRING
    DEFINE ret STRING
    # ... function code  ...
    LET ret = SFMT("Updated user with ID: %1",id)
    RETURN ret
END FUNCTION
```

## Related links

**Related concepts**  

[REST](../16_web-services/4496-rest.md "Representational State Transfer (REST) is a Web standard architecture that provides a method for communication between a Web service and a client over HTTP.")

**Related reference**  

[High-level RESTful Web service attributes](../16_web-services/4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
