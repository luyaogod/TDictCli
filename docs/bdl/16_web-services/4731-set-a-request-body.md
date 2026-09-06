---
title: "Set a request body"
source: "fgl-topics/c_gws_restful_high_level_request_body.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set a request body"
type: "concept"
---

# Set a request body

> Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.

Typically, you specify an input body parameter when you
perform an HTTP PUT, POST, or PATCH request to a resource, otherwise you get [error-9106](../15_library-reference/4483-genero-bdl-errors.md).

In your BDL function, you specify a request body when you define an input parameter that does not
have any `WSHeader`, `WSQuery`, `WSCookie`, or
`WSParam` attributes.

If you have more than one of such input parameters, it means your payload is delivered in a
[multipart](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")
body request.

## Optional request body

You can set [WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") on the
`ATTRIBUTES()` clause of an input body parameter if a body is not always
required. If you have more than one input bodies and any one input body is optional, then
all input body parameters must set the `WSOptional` attribute; otherwise
fglcomp will raise [error-9138](../15_library-reference/4483-genero-bdl-errors.md).

When an input body is set as optional, the fglrestful tool generates two
Genero BDL functions for the same service in the stub file:

- The standard function, where an input body parameter is defined for which the client
  must provide a value.

  ```
  PUBLIC FUNCTION TestRecord(p_body testRecordRequestBodyType) RETURNS (INTEGER, INTEGER)
  ```
- A second function of the same name suffixed with "NoRequestBody", where there is no
  input body parameter, and the client does not need to provide a value.

  ```
  PUBLIC FUNCTION TestRecordNoRequestBody() RETURNS (INTEGER, INTEGER)
  ```

As a result, the service is able to handle all client requests, with or without an
input body, depending on which function is called.

For an example of a function using an optional input body, see Example: Optional input body.

## Example specifying a request body

In this sample REST function a user record is updated with a value sent in the request
body. The function's `id` parameter specifies the user to update. The
lname parameter of type `STRING` provides the user
details to update.

The `lname` data is passed in the message body in either JSON or XML
format.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION updateUsersField(
    id INTEGER ATTRIBUTES(WSParam, WSDescription = "User id"),
    lname STRING ATTRIBUTES(WSDescription = "User's last name"))
  ATTRIBUTES(WSPut,
             WSPath = "/users/{id}",
             WSDescription = "Update name of a given user",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE ret STRING
      TRY
        UPDATE users
        SET name = lname WHERE @id = id
        IF sqlca.sqlerrd[3] = 1 THEN # sqlerrd[3] = processed rows
          LET ret = SFMT("Updated user with ID: %1",id)
        ELSE
          LET ret = SFMT("No user with ID: %1",id)
        END IF
      CATCH
        LET ret = SFMT("Error updating user with ID: %1",id)     
        LET userError.message = SFMT("SQL error:%1 [%2]",
                                     sqlca.sqlcode, SQLERRMESSAGE)
        CALL com.WebServiceEngine.SetRestError(400,userError)
      END TRY
    RETURN ret
END FUNCTION
```

![Sample output of the HTTP request](../_images/rest_request_header_output.png)

*Output of the HTTP request*

## Example: Optional input body

In this sample REST function the value returned is dependent on whether data has been
received in an input body. The function has an input parameter `rec` that is
set with the `WSOptional` attribute. Therefore, data may or may not be
received in the request body.

A user-defined variable `Context` set with the [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") attribute is defined in the
module. This is used to detect if the client has sent a body at runtime. If the request does
not have an input body, the GWS sets the "`NoRequestBody`" entry in the
`Context` variable to true.

In the function the `Context` is checked for the status of the
"`NoRequestBody`" entry, and the value the function returns is based on
this.

```
PRIVATE DEFINE Context DICTIONARY ATTRIBUTES(WSContext) OF STRING

PUBLIC
FUNCTION testRecord(rec RECORD ATTRIBUTES(WSOptional) a INTEGER, b INTEGER END RECORD)
  ATTRIBUTES (WSPut, 
             WSPath = "/record",
             WSDescription = "Check for request body with WSContext variable")
  RETURNS (INTEGER)
  IF Context.contains("NoRequestBody") THEN
    DISPLAY Context["NoRequestBody"]
    RETURN 0
  ELSE
     RETURN rec.a + rec.b
  END IF
END FUNCTION
```

## Related links

**Related concepts**  

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")
