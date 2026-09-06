---
title: "Download a file as an attachment to a response"
source: "fgl-topics/c_gws_restful_high_level_download_file_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Attach files with WSAttachment and WSMedia > Download file in response"
type: "concept"
---

# Download a file as an attachment to a response

> This example demonstrates how to return a file as an attachment using the WSAttachment attribute

In order to return a file as an attachment to a client, set the [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") attribute on an output
parameter.

In this sample REST function a file is returned as attachment. The function has one
return parameter:

- A return paramater is defined as `STRING` type with a
  `WSAttachment` attribute. The REST engine treats the parameter value as a path to a
  file to be attached.
- A `WSMedia` attribute is
  added to handle the data format for images. The wildcard media type (`image/*`) in
  the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute lets the service
  accept or return any image format. If the file can be of any type (not only images), omit the
  `WSMedia` attribute from `WSAttachment`.

The actual media type used for the request or response depends
on the `Accept` or `Content-Type` headers. Your code is responsible
for replacing the `image/*` placeholder with the concrete media type you expect or
receive. You can also use the [WSContext](4806-wscontext.md) attribute to set these headers explicitly.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") – is set to handle errors. If
the file does not exist, the `WSError` attributed variable is set with a description
of the error and the `SetRestError()` method is called to return it.

```
IMPORT com
IMPORT os

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION downloadImageFile()
  ATTRIBUTES (WSGet, 
              WSPath = "/files3/images",
              WSDescription = "download image file to the client with WSAttachment",
              WSThrows = "400:@userError")
  RETURNS (STRING ATTRIBUTES(WSAttachment, WSMedia = "image/*") )
    DEFINE ret, fname STRING
    DEFINE ok INTEGER

    LET fname = "favicon.ico"
    LET ok = os.Path.exists(fname)
    IF ok THEN
        LET ret = fname
    ELSE
      LET userError.message = SFMT("File (%1) does not exist", fname)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END IF
  RETURN ret
END FUNCTION
```

![Image shows](../_images/rest_response_wsattachment.png)

*Output of HTTP response to download image file*

In Figure 1 the `Content-Disposition` response header in the output indicates
that the content is expected as an attachment. Therefore, when the function is called, the file is
downloaded and saved to the client locally in its TMP directory.

## Related links

**Related concepts**  

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")

[Download a large object in the response body](4739-download-large-object-in-response-body.md "This example demonstrates how to send a large object in the message response body.")
