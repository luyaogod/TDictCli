---
title: "Multipart requests or responses"
source: "fgl-topics/c_gws_restful_high_level_multipart_request_response.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Multipart requests or responses"
type: "concept"
---

# Multipart requests or responses

> In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.

You can specify a form-data type HTTP multipart request or response for transferring data of
several MIME types, such as JSON, XML, simple string, and to upload or download files.

## Multipart/form-data support

In the `multipart/form-data` type, entire files can be included in the data
transfer without encoding. This is ideally suited for uploading and downloading image or text files
to a Web service.

Separate parts are identified by the GWS REST engine's naming
convention sequence for headers and body parts that start with "rv0" and goes to the
"rvnth" number of parts. You can change default header naming by adding a [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attribute to your parameter and
setting a name.

The GWS REST engine also creates a "boundary" string in the "Content-Type: " header. This
boundary, is placed between the various parts, and at the beginning and end of the body of the
message.

From GWS version 4.00 onwards you can set `Content-Type` headers for each data
part in a multipart message via the `WSContext` dictionary variable. For more
information about the `WSContext` attribute and an example using multipart, see [Example WSContext with multipart Content-Type set at runtime](4806-wscontext.md).

## Define multipart request body

You define a multipart request body when you have **more than one** input parameters that do
not have `WSHeader`, `WSQuery`, `WSCookie`, or
`WSParam` attributes. The GWS handles this request as a multipart of
`form-data` type.
> **Important:**
>
> Typically, you specify an input body parameter when you
> perform an HTTP PUT, POST, or PATCH request to a resource, otherwise you get [error-9106](../15_library-reference/4483-genero-bdl-errors.md).

## Example multipart request

In this sample function there is an example of a multipart request. In the function's
`id` parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to update,
and the input parameters "addr" and "dtm" contain
values to update. These values will be sent in separate parts of the request body.

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

PUBLIC FUNCTION updateUsersAddress(
    id INTEGER ATTRIBUTES(WSParam),
    addr STRING,
    dtm DATETIME YEAR TO SECOND 
         ATTRIBUTES(WSDescription = "Must be datetime from year to second"))
  ATTRIBUTES(WSPut,
             WSPath = "/updatedUsers/{id}",
             WSDescription = "Multipart msg, update user address and date field",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE ret STRING
    TRY
      UPDATE users SET address = addr, last_notice = dtm
          WHERE @id = id
      IF sqlca.sqlerrd[3] = 1 THEN # sqlerrd[3] = processed rows
        LET ret = SFMT(" - Updated user with ID: %1",id)
      ELSE
        LET ret = SFMT(" - No user with ID: %1",id)
      END IF
    CATCH
      LET userError.message = SFMT(" - SQL error:%1 [%2]",
                                   sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
    RETURN ret
END FUNCTION
```

## Define multipart response body

You define a multipart response body if you have **more than one** return values without
`WSHeader` attributes. The GWS handles this as a multipart response of type
`form-data`.
> **Important:**
>
> A message body in the response is required when you
> perform an HTTP GET, POST, PUT, DELETE operation on a resource, otherwise the response
> results in the [error-9106](../15_library-reference/4483-genero-bdl-errors.md).

## Example multipart response

In this sample function there is an example of a multipart response. There are two string values
defined in the `RETURNS` clause of the function. These do not have
`WSHeader` attributes, so values are sent in separate parts in the response
body.

```
PUBLIC FUNCTION help()
  ATTRIBUTES (WSGet,
              WSPath = "/help")
  RETURNS (INTEGER ATTRIBUTES(WSHeader), STRING, STRING)
    RETURN 3, "Hello world", "Have a nice day."
END FUNCTION
```

## Related links

**Related concepts**  

[Transfer data in large objects](4738-transfer-data-in-large-objects.md "Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.")

[Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md "In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.")

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")

## Child topics

- [Download a file in a multipart response](4744-download-file-and-data-response.md): This example demonstrates how to return a file in a form-data type HTTP multipart response while also transferring data of other types in the same message response.
- [Example: upload a file in a multipart request](4745-upload-file-and-data-request.md): Shows an example of a method you might use from a client to upload an image along with other data in a function using a form-data type HTTP multipart request.
