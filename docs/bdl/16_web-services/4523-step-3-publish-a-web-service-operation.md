---
title: "Step 3: Publish a web service operation returning the W3CEndpointReference state variable and set it as session initiator"
source: "fgl-topics/c_gws_stateful_services_013.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Server side > Step 3: Publish a Web service operation"
type: "concept"
description: "You must define which Web service operation will initiate the session on your service and return the W3CEndpointReference state variable. All other Web service operations (not defined as session ..."
---

# Step 3: Publish a web service operation returning the W3CEndpointReference state variable and set it as session initiator

You must define which Web service operation will initiate the session on your service and return
the W3CEndpointReference state variable.

All other Web service operations (not defined as session initiator) will return an error if they
don't get reference parameters defined in the W3CEndpointReference state variable as WS-Addressing
1.0 headers.

For
example:

```
DEFINE op com.WebOperation
LET op = com.WebOperation.CreateDocStyle("GetInstance",
  "GetInstance",NULL,EndpointReferenceState)
CALL op.initiateSession(TRUE)
CALL serv.publishOperation(op,NULL)
```

There is no restriction regarding the input parameter of the Web service initiator function, but
the output parameter must be the same W3CEndpointReference record passed to the service creation
constructor.

It is not required to have a Web operation which initiates the session in the same service, but
then you have to return the same W3CEndpointReference record in another Web service to instantiate
the session, such as a Factory service that instantiates all sessions for other stateful
services.

## Related links

**Related concepts**  

[Step 4: Create the BDL session initiator function and instantiate a new session](4524-step-4-create-the-bdl-session-initiator-function.md "Step 4: Create the BDL session initiator function and instantiate a new session")
