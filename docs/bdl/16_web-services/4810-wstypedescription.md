---
title: "WSTypeDescription"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSTypeDescription.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSTypeDescription"
type: "concept"
---

# WSTypeDescription

> Provides a description for a user‑defined type in the REST service.

## Syntax

```
WSTypeDescription=" description "
```

Where:

1. description provides a description of the user-defined [type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") it is applied to.

`WSTypeDescription` is an optional attribute.

## Usage

You use this attribute to provide useful description for any user-defined [type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") declared at the module level of the high-level REST web
service.

Information you specify with this attribute, is generated in
the OpenAPI documentation file. You can control generation of descriptions in the stub file with the
`--comment` option of the [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") tool.

## Example using WSTypeDescription in a type definition

In this example, the `WSTypeDescription` attribute describes the
`profileType` record.

```
PUBLIC TYPE profileType RECORD ATTRIBUTE(WSTypeDescription = "profile of user")
     id INTEGER ATTRIBUTES (WSDescription="Internal Identifier"), 
     name VARCHAR(100) ATTRIBUTES (WSDescription="Lastname"),
     email VARCHAR(255),
     category VARCHAR(10) ATTRIBUTES (WSDescription="User Demographic"),
     status INTEGER,
     ccode VARCHAR(3) ATTRIBUTES (WSDescription="Country Code")
     # ...
   END RECORD
```

In the OpenAPI documentation, the `profileType` will be referenced in the
`#/components/schema` section. The GWS engine exposes the description property in the
`schema` of the JSON object.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the profileType JSON schema with a description property](../_images/rest_openapi_api_wstypedescription_json_schema.png)

*Sample JSON schema with description properties*

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
