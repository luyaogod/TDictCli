---
title: "Planning a Web service"
source: "fgl-topics/c_gws_concepts_008.html"
breadcrumb: "Web services > General > Introduction to Web services > Planning a Web service"
type: "concept"
---

# Planning a Web service

> Creating a Web service application requires planning for the future use and reuse of the service.

When creating a Web service, you not only have to think of the task at hand, but you must also
consider growth. You likely want the Web service to be flexible; to be able to handle different
types of input. Prepare the Web service for what is *probable*. Developers should think bigger
than the needs of a single application. Consider how you might reuse existing services, and how your
services might be reused by others.

Security will likely play a larger role than it did previously
with existing in-house application infrastructures using programmed
links between systems; you will need to become versed in security
issues.

Decide which protocol is right for your requirements, Simple Object Access Protocol (SOAP) or
Representational State Transfer (REST).

While both are similar, it is worth researching the advantages of one over the other in terms of
the requirements and scope of your project. Keep in mind when designing and coding web services that
you need to have: flexibility, reusability, and interoperability.

## Related links

**Related concepts**  

[SOAP](4495-soap.md "Simple Object Access Protocol (SOAP) is a communication protocol that defines an XML data flow between a server and a client.")

[REST](4496-rest.md "Representational State Transfer (REST) is a Web standard architecture that provides a method for communication between a Web service and a client over HTTP.")

[Designing REST Web services](4704-designing-rest-web-services.md "Identifying the resources to expose to a RESTful Web service client is an essential step in designing a RESTful Web service server.")
