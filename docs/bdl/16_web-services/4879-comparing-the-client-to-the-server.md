---
title: "Comparing the client to the server"
source: "fgl-topics/c_gws_rest_calc_cs_comp.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > The RESTful calculator demo source > Comparing the client to the server"
type: "concept"
---

# Comparing the client to the server

> Comparing the code between the RESTful calculator client and server.

The purpose of this side-by-side comparison is to provide a glimpse of how the client code and
the server code relate. In some cases, only the initial line of server code has been provided; you
can look at the full source for either the client or the server for the complete code.

- [Calculator client source](4878-calculator-client-source.md "The source code for the client-side application included in the RESTful Web services calculator demo.")
- [Calculator server source](4877-calculator-server-source.md "The source code for the server-side application included in the RESTful Web services calculator demo.")

| Calculator client | Calculator server |
| --- | --- |
| LET req = com.HttpRequest.Create("http://localhost:8090/add?a=" \|\| add_in.a \|\| "&b=" \|\| add_in.b)The client creates the HTTP request, which is a URL consisting of a path and parameters. | LET m_reqInfo.path = req.getUrlPath() CALL req.getUrlQuery(m_reqInfo.items)The server parses out the path and the query from the HTTP service request. |
| CALL req.setMethod("GET")The client sets the method with the verb "GET". | LET m_reqInfo.method = req.getMethod()The server parses out the method verb from the HTTP service request. |
| CALL req.setHeader("Content-Type", "application/json")The client sets the Content-Type header request. | LET m_reqInfo.ctype = getHeaderByName(req,"Content-Type") IF m_reqInfo.ctype.getIndexOf("/xml",1) THEN LET m_reqInfo.informat = "XML" ELSE LET m_reqInfo.informat = "JSON" END IFThe server parses out the value of the Content-Type header request. |
| CALL req.setHeader("Accept", "application/json")The client sets the Accept header request. | LET m_reqInfo.caccept = getHeaderByName(req,"Accept") IF m_reqInfo.caccept.getIndexOf("/xml",1) THEN LET m_reqInfo.outformat = "XML" ELSE LET m_reqInfo.outformat = "JSON" END IFThe server parses out the value of the Accept header request. |
| CALL req.doRequest()The client submits the request. | LET req = com.WebServiceEngine.getHTTPServiceRequest(-1)The server receives the HTTP service request. |
| LET resp = req.getResponse() LET info.response = resp.getTextResponse() CALL util.JSON.parse(info.response, add_out)The client:retrieves the HTTP response using a method of the `com.HttpRequest` class,returns it as a string using methods of the `com.HttpResponse` class, andparses out the desired value using methods of the `util.JSON` class.This example is for JSON; if the data is XML instead, then the methods `doXmlRequest()` and `getXmlResponse()` would be used instead of getting a text response and JSON parsing. | CALL req.setResponseHeader("Content-Type","application/json") CALL req.sendTextResponse(200, "OK", util.JSON.stringify(output_variable))The server returns a response to the client. |
