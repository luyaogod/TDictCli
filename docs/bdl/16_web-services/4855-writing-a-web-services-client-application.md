---
title: "Writing a Web services client application"
source: "fgl-topics/c_gws_rest_section_client.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application"
type: "concept"
---

# Writing a Web services client application

> Create, configure and deploy a RESTful Genero Web Services (GWS) client.

In this section, you will learn about the basic steps of creating a client application to access
the "Add" function in the calculator RESTful GWS Web Service that is detailed in the [Calculator RESTFul Web services server application](4877-calculator-server-source.md "The source code for the server-side application included in the RESTful Web services calculator demo.").
The topics explain:

1. What functions are available on the Server, and which query parameters/arguments to use with the
   URIS (resource). See [Step 1: Obtain information about Web service resources](4856-step-1-obtain-information-about-resources.md "To access a RESTful Web services server, you must first get information about the services offered by a server from the service provider.").
2. What GWS packages you need, see [Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md "The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages.").
3. What Genero Business Development Language records you need to define, see [Step 3: Define the records](4858-step-3-define-the-records.md "In this step you define the records you need for the HTTP Request and Response and the processing of the data.").
4. How to request the Web service functions. [Step 4: Build the HTTP request](4859-step-4-build-the-http-request.md "In this step you code the HTTP request and perform the request to the server to add two numbers.")
5. How to get the result and handle errors.[Step 5: Process the HTTP response](4860-step-5-process-the-http-response.md "In the final step you process the HTTP response; check for errors, and process the data.").

In discussing the above five steps, the [Calculator
RESTFul Web services client application](4878-calculator-client-source.md "The source code for the client-side application included in the RESTful Web services calculator demo.") is used to provide code examples.

## Child topics

- [Step 1: Obtain information about Web service resources](4856-step-1-obtain-information-about-resources.md): To access a RESTful Web services server, you must first get information about the services offered by a server from the service provider.
- [Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md): The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages.
- [Step 3: Define the records](4858-step-3-define-the-records.md): In this step you define the records you need for the HTTP Request and Response and the processing of the data.
- [Step 4: Build the HTTP request](4859-step-4-build-the-http-request.md): In this step you code the HTTP request and perform the request to the server to add two numbers.
- [Step 5: Process the HTTP response](4860-step-5-process-the-http-response.md): In the final step you process the HTTP response; check for errors, and process the data.
- [Configure a WS client to access an HTTPS server](4861-configure-a-ws-client-to-access-an-https-server.md): Configuration steps to access a server in HTTPS.
- [Configure a WS client to connect via an HTTP Proxy](4865-configure-a-ws-client-to-connect-via-an-http-proxy.md): Configuration steps to connect via a HTTP proxy.
- [Configure a WS client to use IPv6](4866-configure-a-ws-client-to-use-ipv6.md): Configuration steps to customize IPv6 for a WS client.
- [Authenticate the WS client to a server (HTTP basic authentication)](4867-authenticate-the-ws-client-to-a-server.md): Configuration steps to authenticate the client to a server.
- [Authenticate the WS client to a proxy](4868-authenticate-the-ws-client-to-a-proxy.md): Configuration steps to authenticate the client to a proxy (proxy authentication).
