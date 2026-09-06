---
title: "Transfer data in large objects"
source: "fgl-topics/c_gws_restful_high_level_handle_lobs.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer > Transfer data in large objects"
type: "concept"
---

# Transfer data in large objects

> Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.

Specify an input or return parameter of a Web service function as a `TEXT` or `BYTE` type in order to transfer a large object (LOB) to a web service server or client.

## Which object to use

| Situation | Data type to use |
| --- | --- |
| Your web service transfers binary file contents.For example: an image, PDF, zip file, etc. | BYTE |
| You have a binary large object (BLOB) stored in a table and your web service must write it out to a file.For example, an image. | BYTE |
| Your web service transfers ASCII content.For example: text, HTML, etc. | TEXT |

This method transfers data in the message body.

> **Important:**
>
> To send a LOB from a REST client function you must, `LOCATE` it in memory or file to load the data
> before calling the server function. On the server REST function, you do not need to locate the
> input parameters `TEXT`/`BYTE` types as they are implicitly located
> in memory when the server function is called. See code samples in [Upload a large object in the request body](4740-upload-large-object-in-request-body.md "This example demonstrates how to upload a large object in the message request body.").

## Related links

**Related concepts**  

[Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md "In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.")

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")

## Child topics

- [Download a large object in the response body](4739-download-large-object-in-response-body.md): This example demonstrates how to send a large object in the message response body.
- [Upload a large object in the request body](4740-upload-large-object-in-request-body.md): This example demonstrates how to upload a large object in the message request body.
