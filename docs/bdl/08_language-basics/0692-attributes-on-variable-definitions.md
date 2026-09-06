---
title: "Attributes on variable definitions"
source: "fgl-topics/c_fgl_variables_attributes.html"
breadcrumb: "Language basics > Variables > Attributes on variable definitions"
type: "concept"
---

# Attributes on variable definitions

> Variables can be defined with meta-data information.

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

Variables can be defined with the `ATTRIBUTES()` clause, to specify meta-data
information for the anonymous type created for this variable.

To specify metadata information when defining a variable, use the `ATTRIBUTES`
clause right after the
type:

```
DEFINE myvar INTEGER ATTRIBUTES(json_name="my variable")
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
