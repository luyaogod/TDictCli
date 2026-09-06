---
title: "Writing a Web Services client application"
source: "fgl-topics/c_gws_section_client.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application"
type: "concept"
---

# Writing a Web Services client application

> Create, configure and deploy a Genero Web Services client using the SOAP protocol.

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

## Child topics

- [Steps for writing a client for a SOAP Web service](4595-steps-to-write-a-gws-client.md): A SOAP Web service contains functions that you need to call. Create a Genero BDL application that requests the functions of the service.
- [Change WS client behavior at runtime](4607-change-ws-client-behavior-at-runtime.md): Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.
- [WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md): To access a remote Web Service, you first must get the WSDL information from the service provider.
- [Use logical names for service locations](4629-use-logical-names-for-service-locations.md): Using a logical reference for the Web service, instead of the real URL, in your client application URL binding has advantages for working with and deploying applications.
- [Configure a WS client to access an HTTPS server](4630-configure-a-ws-client-to-access-an-https-server.md): Configuration steps to access a server in HTTPS.
- [Configure a WS client to connect via an HTTP Proxy](4634-configure-a-ws-client-to-connect-via-an-http-proxy.md): Configuration steps to connect via a HTTP proxy.
- [Configure a WS client to use IPv6](4635-configure-a-ws-client-to-use-ipv6.md): Configuration steps to customize IPv6 for a WS client.
- [Authenticate the WS client to a proxy](4636-authenticate-the-ws-client-to-a-proxy.md): Configuration steps to authenticate the client to a proxy (proxy authentication).
- [Authenticate the WS client to a server (HTTP basic authentication)](4637-authenticate-the-ws-client-to-a-server.md): Configuration steps to authenticate the client to a server.
- [How to create a SOAP WSSecurityUserName token](4638-create-a-wssecurityusername-token.md): Creating a WSSecurityUserName token requires you to generate a client stub using the --domHandler option and using methods from the securityHelper module to create a WSSecurityUserName token.
