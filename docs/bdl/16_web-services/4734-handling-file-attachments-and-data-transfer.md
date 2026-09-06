---
title: "Handling file attachments with REST"
source: "fgl-topics/c_gws_rest_high_level_handle_file_attachments.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling file attachments and data transfer"
type: "concept"
---

# Handling file attachments with REST

> The Genero REST high-level framework provides two mechanisms for handling attachments.

- Send (or receive) the complete file as a simple attachment.
- Send (or receive) a file in a MIME multipart message body.

## Which method to use

| Situation | Method to choose |
| --- | --- |
| If your REST service needs one file, whether it is a zip file, image, or PDF. | Send the complete file as a simple attachment.See [Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md "In GWS REST attachments are handled via the WSAttachment and WSMedia attributes."). |
| If the design of your REST Web services requires multipart to attach files. | Send the file as multipart.See [Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body."). |
| If you have several attachments to send in a single HTTP request or response message. | Send the file as multipart.See [Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body."). |
| If you have a large data object, whether it is a binary, text file, or a BLOB in a database. | Send the data in a BYTE or TEXT object in the message body. See [Transfer data in large objects](4738-transfer-data-in-large-objects.md "Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.") |

## Child topics

- [Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md): In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.
- [Transfer data in large objects](4738-transfer-data-in-large-objects.md): Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.
