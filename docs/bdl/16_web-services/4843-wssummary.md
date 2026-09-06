---
title: "WSSummary"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSSummary.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes related to function parameters > WSSummary"
type: "concept"
---

# WSSummary

> Specifies a short, human-readable summary of a REST operation. It corresponds to the summary keyword in the OpenAPI specification.

## Syntax

```
WSSummary=" summary "
```

Where:

1. summary gives a concise summary of the function it is applied to.

`WSSummary` is an optional attribute. It corresponds to the
`summary` keyword in the OpenAPI specification.

## Usage

Provides a concise description of the operation’s purpose, helping you quickly understand the
role of each endpoint. Complements [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types."), which can be
used for a more detailed explanation.

## Example

```
FUNCTION getCustomer(
    id INTEGER ATTRIBUTE(WSParam))
    ATTRIBUTES(WSGet,
        WSPath = "/customer/{id}",
        WSTags = "Customer",
        WSSummary = "Retrieve a customer by ID",
        WSDescription = "Returns the customer details for the given ID")
    RETURNS(customerType)

END FUNCTION
```

## OpenAPI

```
paths:
  /customer/{id}:
    get:
      tags:
        - Customer
      summary: Retrieve a customer by ID
      description: Returns the customer details for the given ID
      responses:
        '200':
          description: successful operation
```

## Compilation rules

1. `WSSummary` applies only to operation objects (functions exposed via Web Services).
2. Accepted values: a single string.
3. Must contain a value.

## GWS engine

Length recommendation: keep the text short to ensure good display in Swagger UI and client documentation.

The GWS engine exports the `WSSummary` text to the OpenAPI document under the `summary` keyword. At runtime, the summary has no functional effect: it is used for documentation and tooling only.

## fglrestful

fglrestful includes the summary text in the generated OpenAPI file.

The summary may also appear in generated client documentation to help developers quickly identify the purpose of each client function.

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

[Set a request body](4731-set-a-request-body.md "Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.")

[Set a response body and header](4732-set-a-response-body-and-header.md "You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.")
