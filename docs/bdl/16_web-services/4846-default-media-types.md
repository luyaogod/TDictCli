---
title: "Default media types"
source: "fgl-topics/c_gws_high_level_rest_supported_media_types.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Default media types"
type: "concept"
---

# Default media types

> For RESTful Web services developed using the high-level framework, you can specify the MIME type or you can accept the default. The default MIME type is based on the data type.

Table 1 lists
the default MIME types. These default MIME types are used when [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") is **not** set or when the
`Accept` or `Content-Type` HTTP headers are not set in
requests and responses. The data type determines the default MIME type for message
content.

A media type (also known as Multipurpose Internet Mail
Extensions (MIME) type) is an identifier used by the HTTP protocol to denote message
content. The format is based on standards from the Internet Assigned Numbers Authority
(IANA).

| Data type | Example | MIME type |
| --- | --- | --- |
| Simpletype | Simple types are data structures that represent a simple value like strings, numbers, and booleans.DEFINE x, y INTEGER DEFINE s STRING DEFINE b BOOLEAN | `text/plain` |
| Record (only) | A structure of type [`RECORD`](../08_language-basics/0717-record.md "The RECORD keyword defines a structured type or variable.") that defines a fixed number of elements.DEFINE rec RECORD x, y INTEGER END RECORD | `application/json` |
| Record and array | A structured type like a dynamic [`array`](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") or a record with dynamic elements like an array and all other types supported by the [BDL to JSON](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.") conversion class.DEFINE arr DYNAMIC ARRAY OF STRINGDEFINE rec DYNAMIC ARRAY OF RECORD x, y INTEGER END RECORDStatic arrays are not supported. | `application/json`The default MIME type for records and arrays in the message response or request is JSON (`application/json`). If you do not set the MIME type with the [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute, the response or request will be in JSON. |
| BYTE | A structured `BYTE` data type that defines a large object structure for any binary data, such as images or sounds.DEFINE b BYTE | `application/octet-stream` |

## The Image file type (`image/*` )

A special media type identifies an image file. It consists of an identifier and a wildcard
(`image/*` ) Set in `WSMedia`, this type allows the Web service to
handle all image file types: .jpeg, .png, and so on. For
an example, see [Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md "In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.").

## Related links

**Related concepts**  

[Set data format with WSMedia](4741-set-data-format-with-wsmedia.md "It is important to set the correct MIME type for a Web service request or response. You can specify the data format via the WSMedia attribute.")

[Setting MIME type at runtime](4772-setting-mime-type-at-runtime.md "Override the GWS default media format for messages.")

[Transfer data in large objects](4738-transfer-data-in-large-objects.md "Use the BYTE and TEXT data types to transfer large objects (LOBs) with a Genero RESTful web service.")
