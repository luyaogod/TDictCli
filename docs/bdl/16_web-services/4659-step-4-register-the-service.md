---
title: "Step 4: Register the service"
source: "fgl-topics/c_gws_server_tutorial_008.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 1: Writing the entire server application > Step 4: Register the service"
type: "concept"
---

# Step 4: Register the service

> Register the service with the Genero Web Services (GWS) server.

Once the Service and operations are defined and the operations are published, the
`WebService` and `WebOperation` objects have completed their work.
Registering a service puts the Genero DVM in charge of the execution of all the operations of that
service - dispatching the incoming message to the right service, returning the correct output, and
so on. The same service may be registered at different locations on the Web.

The [WebServiceEngine](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.") is a global built-in
object that manages the Server part of the Genero DVM. Use the `RegisterService`
class method of the `WebServiceEngine` class. The parameter is:

1. The name of the WebService object

To register the "Calculator" service, for example, created in [Step 3: Create the service and operations](4658-step-3-create-the-service-and-operations.md "Describes how you provide your Web service and its operations to users who can access it on the net."):

```
CALL com.WebServiceEngine.RegisterService(serv)
```

If you want to create a single GWS Server DVM containing multiple Web Services, define additional
[input and output](4656-step-1-define-input-and-output-records.md "Define records for the input and output messages of the Web function.") records and repeat steps 2
through 6 for each Web Service. In [Step 5: Start the GWS server and process requests](4660-step-5-start-the-gws-server-and-process-requests.md "Code to start the Genero Web Services (GWS) Server."), a GWS Server DVM is started,
containing as many Web Services as you have defined. See [Web services server program deployment](4908-web-services-server-program-deployment.md "The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.") for additional discussion of GWS Services
and GWS Servers.

## Related links

**Related concepts**  

[Step 5: Start the GWS server and process requests](4660-step-5-start-the-gws-server-and-process-requests.md "Code to start the Genero Web Services (GWS) Server.")
