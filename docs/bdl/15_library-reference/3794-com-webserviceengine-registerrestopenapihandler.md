---
title: "com.WebServiceEngine.RegisterRestOpenAPIHandler"
source: "fgl-topics/c_gws_ComWebServiceEngine_RegisterRestOpenAPIHandler.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.RegisterRestOpenAPIHandler"
type: "concept"
---

# com.WebServiceEngine.RegisterRestOpenAPIHandler

> Registers a callback function that dynamically modifies OpenAPI documentation for a REST service.

## Syntax

```
com.WebServiceEngine.RegisterRestOpenAPIHandler(
   service_name STRING, 
   FUNCTION my_callback )
```

1. service\_name specifies the name of the REST service for which the callback
   function will be invoked.
2. `FUNCTION`
   my\_callback is the callback function that will be invoked when a client requests
   the OpenAPI documentation for the specified service. For more information on using a
   `FUNCTION` reference, go to [Function references](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.").

## Callback Function Signature

The callback function must follow this signature:

```
FUNCTION my_callback(openapi util.JSONObject) RETURNS util.JSONObject
```

1. The function name is not fixed and may be chosen as needed.
2. openapi: The parameter name is strict and cannot vary. The generated OpenAPI
   document is a JSON object ([util.JSONObject](3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.")).
3. Return value:
   - `util.JSONObject`: The modified OpenAPI documentation
   - `NULL`: Returns an HTTP 500 error to the client

## Usage

Use this method to register a callback function that will be invoked when a client requests the
OpenAPI documentation for a REST service. This allows you to dynamically modify the OpenAPI JSON
specification before it is returned to the client.

The callback function receives the generated OpenAPI documentation as a [util.JSONObject](3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.") and returns a [util.JSONObject](3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.") with the updated OpenAPI
specification. If the callback function returns `NULL`, an HTTP 500 error is returned
to the client. If the service is not registered before running
`RegisterRestOpenAPIHandler`, [error-15584](4483-genero-bdl-errors.md) is thrown.

For example: This callback updates the `info.version` field and adds two custom
fields:

```
FUNCTION my_callback(openapi util.JSONObject) RETURNS util.JSONObject
    DEFINE info util.JSONObject
    # Get the 'info' section
    LET info = openapi.get("info")

    IF info IS NOT NULL THEN
        # Modify existing field
        CALL info.put("version", "35.0")
        # Add custom fields
        CALL info.put("x-custom-field", "Modified by handler")
        CALL info.put("x-version-custom", "2.0-custom")
    END IF
    RETURN openapi
END FUNCTION

#...
CALL com.WebServiceEngine.RegisterRestService("myservice", "callback_modified")
CALL com.WebServiceEngine.RegisterRestOpenAPIHandler("callback_modified", FUNCTION my_callback)
```

## OpenAPI generation errors

If the OpenAPI generation fails, the web service returns an HTTP 500 error to the client with the
message `Failed to generate OpenAPI documentation`. The GWS engine also returns [error-38](3805-error-codes-of-com-webservicesengine.md), and the developer can retrieve technical details of the error in
`sqlca.sqlerrm`.

## GWS engine

When you [retrieve the OpenAPI description](../16_web-services/4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")
(appending `?openapi.json` or `?openapi.yaml` to the service URL), the
GWS engine applies the callback to the OpenAPI document before returning it to the client. The
example below shows the modified JSON output.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image shows OpenAPI documentation with custom fields](../_images/callback_modified_openapi.png)

*Example OpenAPI excerpt modified by the callback*

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Function references](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")

**Related tasks**  

[Get the OpenAPI service description](../16_web-services/4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")
