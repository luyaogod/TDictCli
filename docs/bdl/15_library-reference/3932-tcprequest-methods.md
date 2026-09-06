---
title: "com.TcpRequest methods"
source: "fgl-topics/c_gws_ComTCPRequest_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods"
type: "concept"
---

# com.TcpRequest methods

> Methods of the com.TcpRequest class.

| Name | Description |
| --- | --- |
| com.TcpRequest.Create( uri STRING ) RETURNS com.TcpRequest | Creates a new TCP request object. |

| Name | Description |
| --- | --- |
| beginXmlRequest() RETURNS xml.StaxWriter | Starts a streaming XML request. |
| doDataRequest( data BYTE ) | Performs the request by sending binary data. |
| doRequest() | Performs a TCP request. |
| doTextRequest( str STRING ) | Performs a request with a string. |
| doXmlRequest( doc xml.DomDocument ) | Performs a request with a DOM document. |
| endXmlRequest( stax xml.StaxWriter ) | Terminates a streaming TCP request. |
| getAsyncResponse() RETURNS com.TcpResponse | Returns the response after performing a TCP request, asynchronously. |
| getResponse() RETURNS com.TcpResponse | Returns the response after performing a TCP request. |
| setConnectionTimeOut( timeout INTEGER ) | Defines the connection time out. |
| setKeepConnection( keep INTEGER ) | Defines if the TCP connection is kept open after sending a request. |
| setMaximumResponseLength( length INTEGER ) | Defines the maximum size in Kbyte of the response. |
| setTimeOut( timeout INTEGER ) | Defines the time out for read/write operations. |
