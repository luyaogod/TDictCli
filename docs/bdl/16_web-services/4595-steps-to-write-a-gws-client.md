---
title: "Steps for writing a client for a SOAP Web service"
source: "fgl-topics/c_gws_client_tutorial_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client"
type: "concept"
---

# Steps for writing a client for a SOAP Web service

> A SOAP Web service contains functions that you need to call. Create a Genero BDL application that requests the functions of the service.

Genero Web Services (GWS) allows a Genero BDL program to access Web services found on the
Internet. GWS supports the WSDL1.1 specification of March 15, 2002.

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

## Related links

1. [Generate a GWS client stub file from the WSDL](4596-generate-the-client-stub-file.md)

   To access a SOAP Web service, retrieve the WSDL from the service provider to learn about the functions provided by the Web service. The fglwsdl command-line tool can use the WSDL to generate much of the client code for you.
2. [Call the web service](4598-call-the-web-service.md)

   Import the required libraries and write the program to call a SOAP Web Service.
3. [Set a time period for the response](4603-set-a-time-period-for-the-response.md)

   To protect against remote server failure or unavailability, set a timeout value that indicates how long you are willing to wait for the server to respond to your request.
4. [Handle GWS server errors](4604-handle-gws-server-errors.md)

   When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.
5. [Compile the client application](4606-compile-the-client-application.md)

   Compiling the client and the stub file.
