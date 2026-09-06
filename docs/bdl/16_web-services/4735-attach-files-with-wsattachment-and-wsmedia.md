---
title: "Attach files with WSAttachment and WSMedia"
source: "fgl-topics/c_gws_restful_high_level_wsattachments_with_wsmedia.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Attach files with WSAttachment and WSMedia"
type: "concept"
---

# Attach files with WSAttachment and WSMedia

> In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.

If a parameter has a [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") attribute, the REST engine treats
the parameter value as a path to a file to be attached, while the data format of the file can be
specified in the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute.

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

## Related links

**Related concepts**  

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")

[Transfer data in large objects](4738-transfer-data-in-large-objects.md "Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.")

## Child topics

- [Download a file as an attachment to a response](4736-download-file-in-response.md): This example demonstrates how to return a file as an attachment using the WSAttachment attribute
- [Example: upload a file as attachment in request body](4737-upload-file-in-request.md): Upload files by adding WSAttachment to an input parameter.
