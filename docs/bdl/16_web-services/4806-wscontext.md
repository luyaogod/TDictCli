---
title: "WSContext"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSContext.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSContext"
type: "concept"
---

# WSContext

> Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.

## Syntax

```
WSContext
```

`WSContext` is an optional attribute.

## Usage

Use WSContext to retrieve REST operation context such as BaseURL, Media, and Scope, or to
set a default context for the `Content-Type` header if the `WSMedia` attribute value
contains a wildcard.

You need to define a `DICTIONARY` variable in your REST module that
specifies `WSContext` in an `ATTRIBUTES()` clause. The context
variable needs to be a private modular variable, otherwise [error-9125](../15_library-reference/4483-genero-bdl-errors.md) is
thrown.

```
PRIVATE DEFINE Context DICTIONARY ATTRIBUTES(WSContext)OF STRING
```

The GWS engine sets the dictionary key/values before the function is executed. For details
of the available context keys and examples of their use, refer to Table 1.

In addition to the standard keys listed below, preprocessing and postprocessing callbacks may set
the `HTTPProcessingErrorMsg` key in the context dictionary. When a callback returns a
non-zero status, the value provided (if non-empty) is used as the HTTP reason phrase for the
response. See the dictionary entry for more details.
> **Important:**
>
> The WSContext dictionary is normally read-only. Your program can query all keys, but cannot
> modify their values. The only exception is the `Content-Type` entry, which may be set
> at runtime to specify the concrete MIME type when the `WSMedia` attribute allows a
> wildcard or multiple media types.

| Key | Description | Example usage |
| --- | --- | --- |
| Media | The dictionary variable entry contains:one of the media values set in the `WSMedia` attribute of the input parameter; the variable returns the MIME type chosen when several MIME types are possible.the default MIME type when no `WSMedia` attribute has been defined.A media type (also known as Multipurpose Internet Mail Extensions (MIME) type) is an identifier used by the HTTP protocol to denote message content. The format is based on standards from the Internet Assigned Numbers Authority (IANA). | DISPLAY context["Media"] This example displays a MIME type. For instance, if WSMedia is set with a list of MIME types (`"application/json,application/xml"`), you can know the exact media type requested in the current execution of the function: `application/xml`. |
| BaseURL | Contains the base URL of the service when executed.The URL for an application being delivered by the Genero Application Server (GAS) and the URL for an application not behind a GAS differ. | DISPLAY context["BaseURL"]This example displays the service's BaseURL, such as: HTTPS://myhost/gas/ws/r/myXcf/Account if the service has been registered as "Account". This may be useful if your REST function has to query another function, and the exact URL is needed. |
| RequestURL | Contains the URL of the service request from the client app. | DISPLAY context["RequestURL"]The request URL may consist of a host name, a port number, a resource path, and a query string. The GWS rebuilds the URL from the context BaseURL value, [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL."), and the function parameter values.You may need the URL to perform a forward or a redirect to another service. |
| Scope | Contains the valid scope (if there is one) that grants access to the REST function. | DISPLAY context["Scope"]This may display "profile.read". For instance, if `WSScope (function)` has been set in the function with a list of scopes ("profile.read, profile.write, profile.mgr"), it may be useful to know which one really applies in the current function execution. |
| RequestPATH | Contains the `WSPath` attribute value of the current REST operation. | `RequestPath=/complexxml/{coeff}`The path will be needed by the REST engine to perform preprocessing or postprocessing callbacks. |
| RequestVERB | Contains the HTTP Verb of the current REST operation. | `RequestVERB=POST`The request verb will be needed by the REST engine to perform preprocessing or postprocessing callbacks. |
| RequestOPERATION | Contains the name of the current REST operation. | `RequestOPERATION=TestEchoXML`The name of the request operation will be needed by the REST engine to perform preprocessing or postprocessing callbacks. |
| Content-Type | Contains the real MIME type the REST function returns if the `WSMedia` attribute value contains "image/\*".**Tip:**You can set the `Content-Type` in a multipart response to transfer data of different MIME types in the same message. See Example WSContext with multipart Content-Type set at runtime. | You can set this context dictionary value at runtime.LET context["Content-Type"]="image/jpeg"This statement will set the returned `Content-Type` header to "image/jpeg". For more examples, go to Set WSContext with Content-Type at runtime |
| HTTPProcessingErrorMsg | When set by a preprocessing or postprocessing callback that returns a non‑zero status, this string value overrides the HTTP reason phrase sent to the client. The behavior is as follows:If the value is non‑empty, the string you provide is used as the reason phrase.If the value is an empty string or `NULL`, the engine falls back to the standard HTTP phrase (for example, "`Unauthorized`" for 401).If not set, a default message is used ("Preprocessing done" or "Postprocessing done").The value is ignored when the callback returns `0` (success). | LET Context["HTTPProcessingErrorMsg"] = "Postprocessing failed"Example of setting a custom reason phrase in a callback. For more examples, go to [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.") |

## Set WSContext with Content-Type at runtime

If the `WSMedia` return parameter attribute has a list of values or a
wildcard value, "image/\*", you code in a function to set the `Content-Type` header for the
actual value returned at runtime in the response. Table 2
outlines options for defining the `Content-Type` header for the image.

| Header type | Header Name | Description | Example |
| --- | --- | --- | --- |
| Default header | `Content-Type` | Setting the header with the format of the image when returning one file. | LET Context["Content-Type"]="image/png" |
| Custom header | name.`Content-Type` | Setting custom headers via GWS default naming. | LET Context["rv0.Content-Type"]="image/png" |
| Custom header | name.`Content-Type` | Setting custom headers with parameters set with [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message."). | LET Context["hello.Content-Type"]="image/jpg" |

## Example WSContext with Content-Type set at runtime

In this sample REST function an image file is requested and returned by the service if
available. The [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") attribute is
set on the function's `id` parameter as the path template for the image to
return. The client passes a value for the `id` parameter in the URL
(`124143` in the example):

http://myhost/gas/ws/r/myGroup/myXcf/MyService/photos/124143

The [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attribute value set in
the return parameter to define the header, "`image.Content-Type`" in the
`Context` variable. The header value is set to the actual image type,
depending on whether it is a jpg or a png type image is returned.

If the image is not found, the `message` field of the
`userError` variable is set, and a call to [SetRestError()](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function.") returns the
message to the client.

```
IMPORT com
IMPORT os

PRIVATE DEFINE Context DICTIONARY ATTRIBUTES(WSContext) OF STRING

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getImage( id INTEGER ATTRIBUTES(WSParam) )
  ATTRIBUTES(WSGET,
             WSPath = "/photos/{id}" )
  RETURNS STRING ATTRIBUTES (WSAttachment, WSName = "image", WSMedia = "image/*")
    DEFINE img STRING
    DEFINE ok BOOLEAN
    
    LET ok = FALSE
    IF os.Path.exists("usr/app/files/"|| id || ".png") THEN
       LET img = "usr/app/files/"|| id || ".png"
       LET Context["image.Content-Type"] = "image/png"
       LET ok = TRUE
    ELSE
      IF os.Path.exists("usr/app/files/"|| id || ".jpg") THEN 
        LET img = "usr/app/files/"|| id || ".jpg"
        LET Context["image.Content-Type"] = "image/jpg"
        LET ok = TRUE
      END IF
    END IF
    IF NOT ok THEN
      LET userError.message = "File does not exist"
      CALL com.WebServiceEngine.SetRestError(404,userError)
    END IF
    RETURN img
END FUNCTION
```

## Example WSContext with multipart Content-Type set at runtime

In this sample REST function a multipart response body is defined to return three images of
different media types. A `Content-Type` header for each image is set using
the `Context` dictionary and the value is set to the actual image type
returned.

In the example, the `Content-Type` header for the first parameter is set
with the GWS REST engine's default naming for headers that start with "rv0":
`rv0.Content-Type`. The second parameter, using the `WSName`
attribute, is given the name "`hello.Content-Type`". The third parameter
continues with the default sequence and is set to `rv2.Content-Type`.

```
PRIVATE DEFINE Context DICTIONARY ATTRIBUTES(WSContext) OF STRING

PUBLIC FUNCTION getMultiImage()
  ATTRIBUTES(WSGet,
            WSPath = "/photos") 
  RETURNS (BYTE, BYTE ATTRIBUTES(WSName = "hello"), BYTE)
  DEFINE b0,b1,b2 BYTE

  # ... function code ...
  
  LET Context["rv0.Content-Type"] = "image/png"
  LET Context["hello.Content-Type"] = "image/jpg"
  LET Context["rv2.Content-Type"] = "image/gif"
  # ...
  RETURN b0,b1,b2
END FUNCTION
```

## Related links

**Related concepts**  

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")
