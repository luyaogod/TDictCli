---
title: "Concept"
source: "fgl-topics/c_gws_stateful_services_003.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Concept"
type: "concept"
---

# Concept

> A stateful service is a service that maintains a context between a Web services client and server.

It enables the service to keep track of previous requests from that context, in order to manage
different states in the Web service server.

## SOAP stateful Web service

Genero Web Services supports two kinds of stateful services:

- Based on the [WS-Addressing 1.0](4519-ws-addressing-1-0-stateful-services.md "WS-Addressing 1.0 uses the WS-Addressing EndpointReference type as a state variable to maintain a stateful service.") specification
  to define the XML format used to convey the context from the client to the server.
- Based on an [HTTP session cookie](4531-stateful-services-based-on-http-cookies.md "A stateful service based on HTTP cookies uses the HTTP transport protocol and its ability to transmit cookies, used as session context.") to convey
  the context from the client to the server.

The Genero Web Service engine uses a BDL variable defined at stateful service creation via [createStatefulWebService()](../15_library-reference/3759-com-webservice-createstatefulwebservice.md "Creates a new object to implement a stateful Web service.") as
service context. Use that variable to hold a service state in a database.

It is up to the BDL programmer to create, store and remove the service state in the database.

The SOAP engine is responsible for:

- Deserializing the state variable when getting a new incoming request. The programmer can then
  read the state variable for any published BDL Web service operation and restore the service state
  corresponding to that variable.
- Serializing a new instance of the state variable in a Web service response for all BDL Web
  service operations set as session initiator via [initiateSession()](../15_library-reference/3770-the-weboperation-class.md "The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service."). The programmer must instantiate a new state by filling the state variable
  and storing it in a database for future use.
