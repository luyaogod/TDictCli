---
title: "Web services server program deployment"
source: "fgl-topics/c_gws_deployment_001.html"
breadcrumb: "Web services > Deploy a Web Service > Web services server program deployment"
type: "concept"
---

# Web services server program deployment

> The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.

## Introduction

In a production environment, Genero Web Services becomes a part of a global
application architecture handled by the Genero Application Server (GAS). The GWS DVMs are managed by
the GAS.

This architecture takes care of:

- Security issues
- Scalability
  - Load management
  - Balancing of the Web service requests amongst the available virtual
    machines
- Runtime monitoring

## GAS configuration

For deployment, the GWS Server application must be added to the GAS configuration.
See Configure applications for Web service in the Genero Application Server User Guide.

The web services application can be added to the GAS in different ways:

- GWS Server application implementing a single web service.

  This application can be deployed on various physical machines. A Genero Web
  Services VMProxy (GWSProxy) is started on each machine where the GWS Server application is executed,
  to manage the requests for a service and manage the DVMs that handle the requests. A single VMProxy
  can communicate with multiple GWS DVMs, and manage the load balancing.
- GWS Server application implementing multiple web services.

  The GWSProxy would manage the client requests, dispatching the request to the
  appropriate DVM and the appropriate web service.

A web wervice Server must be stateless; several instances of the same service can be
created to support load balancing.

The basic deployment strategy can be implemented in varying permutations, depending
on your business needs and the volume of requests.

![Diagram showing path from client to web server to GAS dispatcher to GWS Proxy to DVMs to Database server (and back)](../_images/GAS_GWS_Deploy_Arch.jpg)

*Deployment strategy*

- Using the World Wide Web, a web service client requests WSDL information for
  a particular SOAP web service from the web server or the OpenAPI specification file for a RESTful
  service.
- The web service client uses this information to make a web service request
  from the web server.
- The web server passes the request to the GAS dispatcher.
- The GAS dispatcher starts a GWSProxy, which will be in charge of the pool of
  DVMs that will serve the web service application.
- The GWSProxy will start the number of DVMs specified by the START element
  defined for the web service application.

For a more detailed explanation of the Services Pool for web services, refer to the
*GAS Architecture* topic in the Genero Application Server User Guide.

## Access the web services server from a client application

To reach the web service from the internet, client applications must use the
following URL form:

```
http[s]://host-name/connector.uri/ws/r/app_id
```

1. host-name defines the web server host name where the GAS is
   running.
2. connector.uri is provided by the GAS. Typically, it has the value "gas".
3. app\_id is the XCF file of the GAS web services
   application.

For a more detailed explanation of web service addresses, refer to the *Application web
addresses* section in the Genero Application Server User Guide.

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](../13_programming-tools/2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")
