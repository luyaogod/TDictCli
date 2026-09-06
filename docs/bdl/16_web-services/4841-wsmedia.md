---
title: "WSMedia"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSMedia.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes related to function parameters > WSMedia"
type: "concept"
---

# WSMedia

> Defines the supported media (MIME) types for a parameter or return value.

## Syntax

```
WSMedia = " MIME-type [,...]"
```

Where `WSMedia` is a comma-separated list of MIME types:

1. MIME-type is a [supported media
   type](4846-default-media-types.md "For RESTful Web services developed using the high-level framework, you can specify the MIME type or you can accept the default. The default MIME type is based on the data type.").

   A media type (also known as Multipurpose Internet Mail
   Extensions (MIME) type) is an identifier used by the HTTP protocol to denote message
   content. The format is based on standards from the Internet Assigned Numbers Authority
   (IANA).

`WSMedia` is an optional attribute.

## Usage

You use this attribute to define the supported data format of the message. You set the
`WSMedia` attribute in the `ATTRIBUTES()` clause of
variables and/or on input and output function parameters.

The GWS chooses the format based on whether the `Accept` or
`Content-Type` HTTP headers are set in requests and responses:

- If the `Accept` or `Content-Type` header **is** set, the GWS
  sets the data format accordingly provided the MIME type is listed in `WSMedia`. If
  the MIME type is not listed, the GWS returns a [400 No matching Rest operation found](4791-troubleshooting.md "Understanding how the GWS finds an operation to use for a request can help in troubleshooting unexpected errors.") error.
- If the `Accept` or `Content-Type` header is **not** set, the
  GWS defaults to:
  - Selecting the first MIME type listed in `WSMedia`.
  - Or, depending on the data type, selecting the appropriate [default
    format](4846-default-media-types.md "For RESTful Web services developed using the high-level framework, you can specify the MIME type or you can accept the default. The default MIME type is based on the data type.") for the type.
- If `WSMedia` is **not** set, the GWS defaults to selecting the appropriate
  default format for denoting type of message content.
  > **Tip:**
  >
  > You can specify the format at runtime by calling the REST service engine [com.WebServiceEngine.SetOption](../15_library-reference/3799-com-webserviceengine-setoption.md "Sets an option for the Web Service engine.") option, for details see [Setting MIME type at runtime](4772-setting-mime-type-at-runtime.md "Override the GWS default media format for messages.").

## Example: WSMedia with record variable

In this sample record definition, the `WSMedia` attribute is set to the list of
different MIME types it supports.

```
PUBLIC DEFINE myRecord RECORD
  ATTRIBUTES (WSMedia =
    "application/json, application/xml, application/x-www-form-urlencoded")
    a INTEGER,
    b FLOAT,
    c STRING
END RECORD
```

## WSMedia and file attachments

If a file is to be attached, you can handle this through a parameter with a [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") and the `WSMedia`
attribute. For details, see [Attach files with WSAttachment and WSMedia](4735-attach-files-with-wsattachment-and-wsmedia.md "In GWS REST attachments are handled via the WSAttachment and WSMedia attributes.").

## Related links

**Related concepts**  

[Handling file attachments with REST](4734-handling-file-attachments-and-data-transfer.md "The Genero REST high-level framework provides two mechanisms for handling attachments.")

[Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.")
