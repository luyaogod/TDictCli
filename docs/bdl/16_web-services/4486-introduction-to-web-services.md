---
title: "Introduction to Web services"
source: "fgl-topics/c_gws_concepts_001.html"
breadcrumb: "Web services > General > Introduction to Web services"
type: "concept"
---

# Introduction to Web services

> Web services are a standard way of communicating between applications over an intranet or Internet.

Web services can be invoked via the HyperText Transfer Protocol (HTTP), by requests for services,
and information is returned in either JSON documents or XML documents. It does not matter if the
platform that runs the web service is different to the platform that receives the JSON or XML
document.

## SOAP and REST

Genero Web Services supports web services implemented by the Simple Object Access Protocol (SOAP)
and Representational State Transfer (REST) architecture, which are defined standards for
communicating with Web services. A web service defines how to communicate between two entities:

- A server that exposes services
- A client that consumes services

## Server usage example

A server exposes, for example, a "StockQuotation" service that responds to an operation
"getQuote". For the "getQuote" operation, the input message is a stock symbol as a string, and the
output message is a stock value as a decimal number.

The "getQuote" operation is a function written in Genero BDL, and it is published on the server.
This function retrieves the stock value for the stock symbol passed in, and returns it.

## Client usage example

The Web service client application calls the function as if it were a local function. It passes
the stock symbol to the function, and stores the returned value in a variable.

## SOAP

If the Web service operation is named `WebService_StockQuotation_getQuote` and the
local variable is svalue, the Web service is called as
follows:

```
LET svalue = WebService_StockQuotation_getQuote( "MyStockSymbol" )
```

## REST

The Web service operation is requested using the URI with the required parameter that matches the
function defined in the "StockQuotation" Web service server side. Therefore, `quotes`
becomes the resource name part of the URI included in our example code. `req` is
defined as an object of the `com.HttpRequest` class.

```
LET req = com.HttpRequest.Create("http://localhost:8090/ws/r/quotes?symbol=FJS)
```

## Related links

**Related concepts**  

[RESTful web services](4703-restful-web-services.md "Create RESTful Web service applications (server and/or client) with Genero Web Services. RESTful Web services conform to the REST architectural style.")

## Child topics

- [Service Oriented Architecture (SOA) and Web services](4487-service-oriented-architecture-soa.md): Service Oriented Architecture (SOA) is based on a philosophy of how to connect systems and exchange data to solve business problems.
- [Migrating to SOA and Web services](4488-migrating-to-soa-and-web-services.md): To migrate your application from an existing integration method to a Service Oriented Architecture (SOA) one and move to Web services requires an iterative and evolutionary approach.
- [Planning a Web service](4489-planning-a-web-service.md): Creating a Web service application requires planning for the future use and reuse of the service.
- [Genero Web Services extension](4490-genero-web-services-extension.md): Applications providing Web services use special libraries of the Genero Business Development Language.
- [Web services standards](4491-web-services-standards.md): Web services standards are defined by the World Wide Web and other organizations.
- [Web services style options](4499-web-services-style-options.md): Information on Web services Style options available for SOAP Genero Web services. There is no style concept in REST.
