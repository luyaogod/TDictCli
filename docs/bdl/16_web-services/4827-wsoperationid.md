---
title: "WSOperationId"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSOperationId.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > WSOperationId"
type: "concept"
---

# WSOperationId

> Specifies a unique string to identify the operation. It corresponds to the operationId keyword in the OpenAPI specification.

## Syntax

```
WSOperationId=" operationId "
```

Where:

1. operationId is the custom identifier to publish in the OpenAPI document.

`WSOperationId` is optional. If you do not set it, the engine uses the 4GL
function name as the `operationId`.

This attribute applies only to functions that expose a REST operation.

## Usage

You use `WSOperationId` when you want to publish a specific `operationId` in the OpenAPI document instead of relying on the default name derived from the 4GL function.

It allows you to override the default value that GWS outputs at runtime. By specifying a more
recognizable name for the `operationId`, you improve the readability of the generated
OpenAPI definition.

The value of `WSOperationId` must be unique across all operations in the REST service. The GWS engine does not validate this, so you must ensure that each identifier is distinct.

## Example: specifying a custom operationId

The following function sets a custom `operationId` using
`WSOperationId`.

```
FUNCTION user()
  ATTRIBUTE(WSGet,
            WSPath = "/user",
            WSDescription = "Custom operationId example",
            WSOperationId = "getUser")
  RETURNS(STRING ATTRIBUTE(WSMedia = "application/json"))
  #... function code
  RETURN "Hello User"
END FUNCTION
```

## OpenAPI

```
paths:
  /user:
    get:
      description: Custom operationId example
      operationId: getUser
      responses:
        '200':
          description: successful operation
```

## Compilation rules

1. Leading and trailing whitespace characters in the string are removed.
2. If the resulting string is empty or contains only whitespace, the engine uses the 4GL function
   name as the default operationId.
3. All special characters are allowed and preserved in the `operationId`.

## GWS engine

The GWS engine exports the `WSOperationId` text to the OpenAPI document under the
`operationId` keyword.

> **Important:**
>
> The GWS engine does not validate the value of
> `WSOperationId`. It does not check correctness, format, or uniqueness at compile time
> or at runtime. You must ensure that each `WSOperationId` in your REST service is
> unique and suitable for inclusion in the generated OpenAPI document.

## fglrestful

fglrestful includes the `operationId` in the generated OpenAPI
document, enabling developers to maintain uniqueness across modules without altering function
names.

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
