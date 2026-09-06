---
title: "Set data format with WSMedia"
source: "fgl-topics/c_gws_restful_high_level_media_formats.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set data format with WSMedia"
type: "concept"
---

# Set data format with WSMedia

> It is important to set the correct MIME type for a Web service request or response. You can specify the data format via the WSMedia attribute.

The default MIME type for records and arrays in the message
response or request is JSON (`application/json`). If you do not set the MIME
type with the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute,
the response or request will be in JSON.

`WSMedia` can contain a comma-separated list of MIME types. If you
specify more than one type, the Accept header received from the client determines the payload format
chosen.

For an example of `WSMedia` used to output records to the client in XML
format, see [Customize XML serialization with WSName and XMLName](4742-customize-xml-data-output.md "Use XML serialization attributes to customize serialization at runtime and improve the readability of the XML output.")

For `WSMedia` examples using images or files, see [Example: upload a file in a multipart request](4745-upload-file-and-data-request.md "Shows an example of a method you might use from a client to upload an image along with other data in a function using a form-data type HTTP multipart request.").

## Child topics

- [Customize XML serialization with WSName and XMLName](4742-customize-xml-data-output.md): Use XML serialization attributes to customize serialization at runtime and improve the readability of the XML output.
