---
title: "WSAttachment"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSAttachment.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes set on both input and output > WSAttachment"
type: "concept"
---

# WSAttachment

> Defines file attachments in the REST message.

## Syntax

```
WSAttachment [= "regexp_pattern"]
```

Where:

1. `WSAttachment` is an attribute in the `ATTRIBUTES()` clause of an
   input parameter or return value, used to declare a file that is attached to the REST message.
2. regexp\_pattern is a regular expression value that the filename must match. It
   is used to limit filenames to a set of valid characters to protect against code injection. Using a
   regex pattern is optional.

`WSAttachment` is an optional attribute.

## Usage

You can attach files using the `WSAttachment` attribute. Files can be sent in the
request or the response. To receive files in a request, include one or more input parameters with
the `WSAttachment` attribute. The client must provide an absolute path for each file
parameter when calling the function.

To reduce the risk of code injection, you can provide a regular expression that restricts allowed
filenames. For example, the following pattern permits a simple filename with an extension, allowing
only letters, digits, and underscores:

```
WSAttachment = "[a-zA-Z0-9_]*\.[a-zA-Z0-9_]*"
```

The GWS validates the incoming filename against the specified pattern. It raises [error-42](../15_library-reference/3805-error-codes-of-com-webservicesengine.md) if the regular expression is invalid or if the filename does not match the expected
format. The `WSAttachment` attribute uses regular expressions that follow
libxml2/POSIX syntax. For more information about regular expressions, refer to the [XML Schema](http://www.w3.org/TR/xmlschema-2/#regexs)
(external link) specification.

You can control the MIME type of the attached file with the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute. If `WSMedia` is
not specified, the engine accepts all file types, following the OpenAPI wildcard MIME type
"`*/*`".

## Example WSAttachment with WSMedia

In this sample REST function a text file is returned as an attachment. In the return clause of
the function a `STRING` is defined with the `WSAttachment` attribute.
The `WSMedia` attribute specifies the MIME type as "text/plain". The absolute path to
the file is set on the return string.

The REST engine copies the file to the temporary directory defined by the TMP
environment variable. The file is removed from the temporary directory at the end of the REST
operation to avoid a build-up of files on your disk.

```
IMPORT com
IMPORT os

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION GetReadme() 
   ATTRIBUTES(WSGet,
              WSPath = "/file/README",
              WSDescription = "Returns a text file",
              WSThrows = "404:@userError") 
   RETURNS (STRING ATTRIBUTES(WSAttachment, WSMedia = "text/plain") )
     DEFINE ret, fname STRING
     DEFINE ok INTEGER

     LET fname = "/myservice/files/README"
     LET ok = os.Path.exists(fname)
     IF ok THEN
       LET ret = fname
     ELSE
       LET userError.message = SFMT("File (%1) does not exist", fname)
       CALL com.WebServiceEngine.SetRestError(404,userError)
     END IF
     RETURN ret
END FUNCTION
```

## Attaching files in request and response

In this sample REST function, a client sends an image file to the server, and the server returns
another image in the response. The wildcard media type (`image/*`) in
the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute lets the service
accept or return any image format. If the file can be of any type (not only images), omit the
`WSMedia` attribute from `WSAttachment`.

The actual media type used for the request or response depends
on the `Accept` or `Content-Type` headers. Your code is responsible
for replacing the `image/*` placeholder with the concrete media type you expect or
receive. You can also use the [WSContext](4806-wscontext.md) attribute to set these headers explicitly.

```
IMPORT os

PUBLIC FUNCTION EchoFile( input STRING ATTRIBUTES (WSAttachment,WSMedia = "image/*") )
  ATTRIBUTES(WSPost)
  RETURNS STRING ATTRIBUTES (WSAttachment, WSMedia = "image/*")
    DEFINE ok INTEGER
    LET ok = os.path.rename(input, "MyFile.png")
    RETURN "/usr/local/MyOtherFile.jpg"
END FUNCTION
```

## Attaching a file and validating the filename with a regular expression

In this REST sample, the input parameter uses
`WSAttachment="[a-zA-Z0-9_]*\.[a-zA-Z0-9_]*"` to declare an attachment and to
validate the filename. When a file is received, the GWS checks the filename against this regular
expression. Only alphanumeric characters and underscores are allowed in both the name and the
extension. Any character outside the class `[a-zA-Z0-9_]` causes validation to
fail.

> **Note:**
>
> Because the dot has a special meaning in regular expressions (it matches any character), it must
> be escaped (`\.`) to match a literal period between the name and the extension.

**Possible errors when validating the filename**

- Invalid regular expression in the `WSAttachment` attribute

  - Server error: [error-42](../15_library-reference/3805-error-codes-of-com-webservicesengine.md)
  - HTTP response:

    ```
    400 Regex syntax error
    ```
- Filename does not match the defined pattern

  - Server error: [error-42](../15_library-reference/3805-error-codes-of-com-webservicesengine.md)
  - HTTP response:

    ```
    400 File name does not match template
    ```

```
PUBLIC FUNCTION simpleFile(
    in STRING ATTRIBUTES(WSAttachment = "[a-zA-Z0-9_]*\.[a-zA-Z0-9_]*"))
    ATTRIBUTES(WSPost, WSPath = "/simple/txt")
    RETURNS()
END FUNCTION
```

## Related links

**Related concepts**  

[Handling file attachments with REST](4734-handling-file-attachments-and-data-transfer.md "The Genero REST high-level framework provides two mechanisms for handling attachments.")

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")
