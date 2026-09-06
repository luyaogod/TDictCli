---
title: "com.HttpServiceRequest methods"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods"
type: "concept"
---

# com.HttpServiceRequest methods

> Methods of the com.HttpServiceRequest class.

| Name | Description |
| --- | --- |
| getUrl() RETURNS STRING | Returns the URL of the HTTP service request. The URL may consists of a host name, a port number, a path, and a query. |
| getUrlHost() RETURNS STRING | Returns the host name contained in the URL of an HTTP service request. |
| getUrlPath() RETURNS STRING | Returns the path contained in the URL of an HTTP service request. |
| getUrlPort() RETURNS INTEGER | Returns the port number contained in the URL of an HTTP service request. |
| getUrlQuery( query RECORD) | Takes a dynamic array of RECORD of two strings and fills that array with the decoded query string of an HTTP service request. |

| Name | Description |
| --- | --- |
| getRequestHeader( name STRING ) RETURNS STRING | Returns the value of an HTTP header. |
| getRequestHeaderCount() RETURNS INTEGER | Returns the number of request headers. |
| getRequestHeaderName( ind INTEGER ) RETURNS STRING | Returns a request header name by position. |
| getRequestHeaderValue( ind INTEGER ) RETURNS STRING | Returns a request header value by position. |
| getRequestVersion() RETURNS STRING | Returns the HTTP version of the service request. |

| Name | Description |
| --- | --- |
| beginXmlRequest() RETURNS xml.StaxReader | Starts an HTTP streaming request. |
| endXmlRequest( stax xml.StaxReader ) | Terminates an HTTP streaming request. |
| getMethod() RETURNS STRING | Returns the HTTP method of the service request. |
| hasRequestKeepConnection() RETURNS INTEGER | Returns `TRUE` if the connection remains open after sending a response. |
| readDataRequest( b BYTE) | Returns the body of a request in a `BYTE`. |
| readFileRequest( ) RETURNS STRING | Returns the body of the request to a file on disk and returns the filename. |
| readFormEncodedRequest( utf8 INTEGER ) RETURNS STRING | Returns the string of a GET request with UTF-8 conversion option. |
| readTextRequest() RETURNS STRING | Returns the request body as a plain string. |
| readXmlRequest() RETURNS xml.DomDocument | Returns the request body as an XML document. |

| Name | Description |
| --- | --- |
| getRequestMultipartType() RETURNS STRING | Returns the multipart type of an incoming request. |
| getRequestPart( pos INTEGER) RETURNS com.HttpPart | Returns the HttpPart object at the specified index position. |
| getRequestPartCount() RETURNS INTEGER | Returns the number of additional multipart elements. |
| getRequestPartFromID( id STRING) RETURNS com.HttpPart | Returns the HttpPart object of the given Content-ID value. |

| Name | Description |
| --- | --- |
| setResponseCharset( charset STRING ) | Defines the HTTP response character set. |
| setResponseHeader( name STRING, value STRING ) | Defines a header for the HTTP response. |
| setResponseVersion( version STRING ) | Defines the HTTP response version. |

| Name | Description |
| --- | --- |
| beginXmlResponse( code INTEGER, description STRING ) RETURNS xml.StaxWriter | Starts an HTTP streaming response. |
| endXmlResponse( stax xml.StaxWriter ) | Terminates an HTTP streaming response. |
| sendResponse( code INTEGER, description STRING ) | Sends an HTTP response without body. |
| sendDataResponse( code INTEGER, description STRING, b BYTE ) | Sends an HTTP response with data of a `BYTE` variable. |
| sendFileResponse( code INTEGER, description STRING, filename STRING ) | Sends an HTTP response with the data contained in a file. |
| sendTextResponse( code INTEGER, description STRING, txt STRING ) | Sends an HTTP response with data from a plain string. |
| sendXmlResponse( code INTEGER, desc STRING, doc xml.DomDocument ) | Sends an HTTP response with data from a XML document object. |

| Name | Description |
| --- | --- |
| addResponsePart( part com.HttpPart) | Adds a new part to the HTTP root part response. |
| setResponseMultipartType( type STRING, start STRING, boundary STRING ) | Sets HTTP response in multipart mode of given type. |

| Name | Description |
| --- | --- |
| findRequestCookie( name STRING) RETURNS STRING | Enables the server to retrieve a cookie sent by the client by name. |
| setResponseCookies( cookies WSHelper.WSServerCookiesType ) | Allows the server to return cookies to be set on the client application sending the request. |
