---
title: "com.HttpPart methods"
source: "fgl-topics/c_gws_ComHTTPPart_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods"
type: "concept"
---

# com.HttpPart methods

> Methods for the com.HttpPart class.

| Name | Description |
| --- | --- |
| com.HttpPart.CreateAttachment( filename STRING ) RETURNS com.HttpPart | Creates a new HttpPart object based on a given filename located on disk. |
| com.HttpPart.CreateFromData( b BYTE ) RETURNS com.HttpPart | Creates a new HttpPart object based on given BYTE located in memory. |
| com.HttpPart.CreateFromDomDocument( doc xml.DomDocument ) RETURNS com.HttpPart | Creates a new HttpPart object based on given XML document. |
| com.HttpPart.CreateFromString( str STRING ) RETURNS com.HttpPart | Creates a new HttpPart object based on given string. |

| Name | Description |
| --- | --- |
| clearHeaders() | Remove all headers from the HTTP part. |
| getAttachment() RETURNS STRING | Returns the absolute path to the HTTP part. |
| getContentAsData( b BYTE ) | Returns the HTTP part as a BYTE. |
| getContentAsDomDocument() RETURNS xml.DomDocument | Returns the HTTP part as an XML document. |
| getHeader( name STRING ) RETURNS STRING | Returns a named HTTP multipart header as a string. |
| getHeaderCount() RETURNS INTEGER | Retrieve the number of headers for the current HTTP part. |
| getHeaderName( pos INTEGER ) RETURNS STRING | Retrieve the name of an HTTP multipart header as a string, where the multipart header is specified by its position. |
| getHeaderValue( pos INTEGER ) RETURNS STRING | Retrieve the value of an HTTP multipart header as a string, where the multipart header is specified by its position. |
| getContentAsString() RETURNS STRING | Returns the HTTP part as a string. |
| removeHeader( name STRING ) | Remove the header of given name from the current HttpPart object. |
| setHeader( name STRING, value STRING ) | Sets a named HTTP multipart header using a string value. |

## Child topics

- [com.HttpPart.clearHeaders](3916-com-httppart-clearheaders.md): Remove all headers from the HTTP part.
- [com.HttpPart.getHeaderCount](3922-com-httppart-getheadercount.md): Retrieve the number of headers for the current HTTP part.
- [com.HttpPart.getHeaderName](3923-com-httppart-getheadername.md): Retrieve the name of an HTTP multipart header as a string, where the multipart header is specified by its position.
- [com.HttpPart.getHeaderValue](3924-com-httppart-getheadervalue.md): Retrieve the value of an HTTP multipart header as a string, where the multipart header is specified by its position.
- [com.HttpPart.removeHeader](3925-com-httppart-removeheader.md): Remove the header of given name from the current HttpPart object.
