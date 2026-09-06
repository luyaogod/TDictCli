---
title: "Type attributes"
source: "fgl-topics/c_fgl_user_types_attributes.html"
breadcrumb: "Language basics > Types > Type attributes"
type: "concept"
---

# Type attributes

> Types can be defined with meta-data information.

## Syntax

In type specifications, the attributes-list clause is:

```
{ ATTRIBUTE | ATTRIBUTES } ( attribute [ = "value" ] [,...] )
```

1. attribute is the name of a definition
   attribute.
2. value is the value for the definition attribute, it is
   optional for boolean attributes.

## Basics

Types can be defined with the `ATTRIBUTES()` clause, to specify meta-data
information for the elements used in the type.

To specify metadata information when defining a type, use the `ATTRIBUTES` in the
type specification used to defined the
`TYPE`:

```
TYPE t_item_id INTEGER ATTRIBUTES(json_name="Item Id")
```

Type attributes can be specified on primitive types like `INTEGER`, and on complex
structures like [records](0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information."), [records](0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description."), [arrays](0736-attributes-on-array-definitions.md "Arrays can be defined with attributes, to complete the type description."), [dictionaries](0746-attributes-on-dictionary-definitions.md "Dictionaries can be defined with attributes, to complete the type description.").

When defining attributes for complex types, the `ATTRIBUTES` clause must be
specified right after the complex type
keywords:

```
TYPE t_uuids DYNAMIC ARRAY ATTRIBUTES(json_name="uuid-list") OF STRING
TYPE t_customer RECORD ATTRIBUTES(json_name = "customer-record")
          ...
       END RECORD
```

## JSON serialization attributes

To define JSON serialization options, use variable definition attributes such as
`json_null` and `json_name`:

```
TYPE t_rec RECORD
   cust_id INTEGER ATTRIBUTES(json_null="null"),
   cust_name INTEGER ATTRIBUTES(json_name="Customer Name"),
   ...
   orderlist DYNAMIC ARRAY ATTRIBUTES(json_null="undefined") OF RECORD
      ...
      END RECORD,
   ...
 END RECORD
```

List of supported JSON attributes for variable definitions:

- `json_null` (values can be `"null"` or
  `"undefined"`)
- `json_name`

For more details about JSON serialization attributes see [BDL names and JSON element names](../09_advanced-features/0958-bdl-names-and-json-element-names.md "To identify elements, JSON standards use different format as Genero BDL variable names.").

## XML serialization attributes

Type attributes are also used when defining variables for XML-based Web Services:

```
TYPE t_data RECORD
     val1 INTEGER ATTRIBUTES(XMLName="Value1"),
     val2 STRING ATTRIBUTES(XMLName="Value2"),
     attr INTEGER ATTRIBUTES(XMLAttribute,XMLName="MyAttr")
END RECORD
```

For more details about XML attributes, see [XML serialization rules and customization](../16_web-services/4958-xml-serialization-rules-and-customization.md).

## Related links

**Related concepts**  

[Attributes on variable definitions](0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.")

[Attributes on record definitions](0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description.")

[Attributes on array definitions](0736-attributes-on-array-definitions.md "Arrays can be defined with attributes, to complete the type description.")

[Attributes on dictionary definitions](0746-attributes-on-dictionary-definitions.md "Dictionaries can be defined with attributes, to complete the type description.")

[Function attributes](0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.")
