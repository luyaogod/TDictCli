---
title: "com.WebServiceEngine.SetRestError"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetRestError.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetRestError"
type: "concept"
---

# com.WebServiceEngine.SetRestError

> Manages error handling for a REST high-level Web Service function.

## Syntax

```
com.WebServiceEngine.SetRestError(
   code INTEGER
   error any-type )
```

1. code defines a standard HTTP response status code (400 - 599) as defined in
   [RFC 2616
   Section 10.4](http://tools.ietf.org/html/rfc2616#section-10.4) and [RFC 2616 Section 10.5](http://tools.ietf.org/html/rfc2616#section-10.5)
2. error defines the program variable that has details of the error. The any-type can be any Genero BDL variable
   type: `RECORD`, `ARRAY`, or simple data type.
   It can return a `NULL` value.

## Usage

You use this method to set the HTTP status code and details of the error to be returned to the
client when the request to the resource has finished. `SetRestError` accepts the HTTP
code, a Genero BDL variable (if any), or a text description to be returned. This method must be
called inside a Web Service REST high-level function.

If the `SetRestError()` method error parameter references a
variable, the variable must be specified in the [WSThrows](../16_web-services/4829-wsthrows.md "Defines the list of error codes the REST function may return.") declaration of the function, otherwise
the method fails and returns an error code of [-15570](4483-genero-bdl-errors.md).

Variables referenced by `WSThrows`, must be defined as `PUBLIC` at
the modular level and specified with the `WSError` attribute.

If `WSThrows` does not reference a variable, just a message string (for example
"`404:my error message`"), you must call the method with the `NULL`
value (`SetRestError(404,NULL)`) to send the "my error message" as the HTTP reason
phrase. Ultimately, this will depend on whether the web server allows you to change the standard
protocol message or not.

It is recommended that if you need to ensure an error description other than the standard
protocol message is returned to the client, you must use `WSThrows` with
"`code:@variable`" option and call the method
with a reference to the variable, for example `SetRestError(400,myError)`.

In the examples Example 1: managing expected errors and Example 2: managing unexpected errors different uses of the method are shown.

## Example 1: managing expected errors

In this example, common application level type errors, like resource not found, are expected.
Instead of the standard HTTP status code (`400 bad request`), the call to
`SetRestError(400,myerror)` will send the description set on the
myError.message variable as the reason for the error.

```
PUBLIC DEFINE myError RECORD ATTRIBUTE(WSError="user error")
  message STRING
END RECORD

PUBLIC
FUNCTION queryAccountById(id VARCHAR(10) ATTRIBUTE(WSParam)
    )
    ATTRIBUTES(WSGet,
        WSPath="/accounts/{id}",
        WSThrows="400:@myError,500:Internal Server Error")
    RETURNS accountType ATTRIBUTES(WSName="body")

  DEFINE thisAccount RECORD LIKE accounts.*
  # ... retrieve user id from database ...

    CASE sqlca.sqlcode
    WHEN 0
        EXIT CASE
    WHEN NOTFOUND
        LET myError.message = SFMT("Could not find account id :%1",id)
        CALL com.WebServiceEngine.SetRestError(400,myError)
    OTHERWISE
        CALL com.WebServiceEngine.SetRestError(500,NULL)
    END CASE
    RETURN thisAccount.*
END FUNCTION
```

## Example 2: managing unexpected errors

> **Tip:**
>
> In general, the recommended option is to list all the expected errors as shown
> in Example 1, in order
> to get them generated in the OpenAPI specification file and trapped on the client side.

In this example you code to trap a status code not defined in the `WSThrows` list.
In the call to the `SetRestError()` method the status code and optionally the error
variable is returned.

```
PUBLIC DEFINE myError RECORD ATTRIBUTE(WSError="user error")
  message STRING
END RECORD

PUBLIC
FUNCTION queryAccountById(id VARCHAR(10) ATTRIBUTE(WSParam)
    )
    ATTRIBUTES(WSGet,
        WSPath="/accounts/{id}",
        WSThrows="400:@myError,500:Internal Server Error")
    RETURNS accountType ATTRIBUTES(WSName="body")
  
  DEFINE thisAccount RECORD LIKE accounts.*
  # ... retrieve user id from database ...

    CASE sqlca.sqlcode
    WHEN 0 # success
        EXIT CASE
    WHEN NOTFOUND
        LET myError.message = SFMT("Could not find account id :%1",id)
        CALL com.WebServiceEngine.SetRestError(400,myError)
    WHEN < 0 
        CALL com.WebServiceEngine.SetRestError(500,NULL)
    OTHERWISE
       CALL com.WebServiceEngine.SetRestError(505,NULL)
    END CASE
    RETURN thisAccount.*
END FUNCTION
```

## Related links

**Related concepts**  

[Handling application level errors](../16_web-services/4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")

[Define functions in a module](../16_web-services/4723-define-functions-in-a-module.md "A GWS REST service is defined in a module.")

**Related reference**  

[High-level RESTful Web service attributes](../16_web-services/4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
