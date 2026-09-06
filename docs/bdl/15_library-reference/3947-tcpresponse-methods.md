---
title: "com.TcpResponse methods"
source: "fgl-topics/c_gws_ComTCPResponse_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpResponse class > TcpResponse methods"
type: "concept"
---

# com.TcpResponse methods

> Methods of the com.TcpResponse class.

| Name | Description |
| --- | --- |
| beginXmlResponse() RETURNS xml.StaxReader | Starts a streaming TCP response. |
| endXmlResponse( stax xml.StaxReader ) | Ends a streaming TCP response. |
| getDataResponse( data BYTE ) | Returns a TCP response in binary format. |
| getTextResponse() RETURNS STRING | Returns a TCP response in string format. |
| getXmlResponse() RETURNS xml.DomDocument | Returns an entire DOM document as TCP response. |
