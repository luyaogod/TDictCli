---
title: "com.HttpResponse methods"
source: "fgl-topics/c_gws_ComHTTPResponse_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods"
type: "concept"
---

# com.HttpResponse methods

> Methods for the com.HttpResponse class.

| Name | Description |
| --- | --- |
| beginJSONResponse() RETURNS json.JSONReader | Starts a HTTP response streaming JSON. |
| beginXmlResponse() RETURNS xml.StaxReader | Starts a streaming HTTP response. |
| endJSONResponse( json json.JSONReader ) | Performs the HTTP request. |
| endXmlResponse( stax xml.StaxReader ) | Performs the HTTP request. |
| getDataResponse( b BYTE ) | Returns the entire HTTP response in a BYTE. |
| getFileResponse( ) RETURNS STRING | Returns the entire HTTP response to a file on the disk. |
| getHeader( name STRING ) RETURNS STRING | Returns the value of an HTTP header. |
| getHeaderCount() RETURNS INTEGER | Returns the number of headers. |
| getHeaderName( pos INTEGER ) RETURNS STRING | Returns the name of a header by position. |
| getHeaderValue( pos INTEGER ) RETURNS STRING | Returns the value of a header by position. |
| getStatusCode() RETURNS INTEGER | Returns the HTTP status code. |
| getStatusDescription() RETURNS STRING | Returns the HTTP status description. |
| getTextResponse() RETURNS STRING | Returns the entire HTTP response in a string. |
| getXmlResponse() RETURNS xml.DomDocument | Returns the entire HTTP response in a DOM document. |

| Name | Description |
| --- | --- |
| getMultipartType() RETURNS STRING | Returns whether a response is multipart or not, and the kind of multipart if any. |
| getPart( pos INTEGER ) RETURNS com.HttpPart | Returns the HTTP part object at the specified index of the current HTTP response. |
| getPartCount() RETURNS INTEGER | Returns the number of additional parts in the HTTP response. |
| getPartFromID( id STRING ) RETURNS com.HttpPart | Returns the HTTP part object marked with the given Content-ID value as identifier, or NULL if none. |

| Name | Description |
| --- | --- |
| getServerCookies( cookies RECORD ) | Returns all cookies set as response from a server. |
