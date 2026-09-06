---
title: "Example 2: Writing a server using third-party WSDL (the fglwsdl tool)"
source: "fgl-topics/c_gws_server_tutorial_010.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 2: Writing a server using third-party WSDL (the fglwsdl tool)"
type: "concept"
---

# Example 2: Writing a server using third-party WSDL (the fglwsdl tool)

> Describes using a server stub from a third-party Web service in your GWS server application.

To write a Web Service that is compatible with the specification of the input and output records
defined by a third-party (for example, a vendor of manufacturing software, or a WSDL specialist in
your company), you can use the fglwsdl tool to obtain the WSDL information and
generate a part of the Server application. See [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") for a
complete description of the tool and its use.

## Related links

1. [Step 1: Get the WSDL description and generate server stub](4662-step-1-get-the-wsdl-description-and-generate-server-stub.md)

   Use the fglwsdl tool to generate the BDL server stub from a WSDL.
2. [Step 2: Write a BDL function for your service operation](4664-step-2-write-a-bdl-function-for-your-service-operation.md)

   Write functions that implement the functions in the stub file. This allows you to create your own version of the function.
3. [Step 3: Create service, start server and process requests](4666-step-3-create-service-start-server-and-process-requests.md)

   Code to start the Genero Web Services (GWS) Server.
