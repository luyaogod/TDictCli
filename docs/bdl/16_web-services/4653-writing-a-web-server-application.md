---
title: "Writing a Web server application"
source: "fgl-topics/c_gws_server_tutorial_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application"
type: "concept"
---

# Writing a Web server application

> Follow examples showing you how to write a complete Web service application for the SOAP protocol. .

This tutorial guides you through the steps to create a server application for a Genero Web
Service that can be accessed over the web by client applications. A complete example is provided at
$FGLDIR/demo/WebServices.

You can write your server application based on input/output records that you have defined. Or,
you can use the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool to include third-party
WSDL information in your server application.

## Related links

1. [Including the web services library](4654-including-the-web-services-library.md)

   Import the com class.
2. [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md)

   Design a simple Web service.
3. [Example 2: Writing a server using third-party WSDL (the fglwsdl tool)](4661-example-2-writing-a-server-using-third-party-wsdl-the-fglwsd.md)

   Describes using a server stub from a third-party Web service in your GWS server application.
4. [Enabling MTOM on the server side](4667-enabling-mtom-on-the-server-side.md)

   Enable the Message Transmission Optimization Mechanism (MTOM) feature to efficiently send binary data to and from Web services.
5. [Compiling GWS server applications](4668-compiling-gws-server-applications.md)

   When compiling, remember to include the WSHelper library.
6. [Testing the GWS service in stand-alone mode](4669-testing-the-gws-service-in-stand-alone-mode.md)

   Test that your service is reachable and that it can generate the WSDL.
7. [Configuring the Genero Application Server for the GWS application](4670-configuring-the-gas-for-the-gws-application.md)

   Prepare for a production environment.
8. [Making the GWS service available](4671-making-the-gws-service-available.md)

   Provide your Web service to users who can access it on the net.
