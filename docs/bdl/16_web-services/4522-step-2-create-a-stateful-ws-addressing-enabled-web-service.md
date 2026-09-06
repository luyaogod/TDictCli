---
title: "Step 2: Create a stateful WS-Addressing enabled web service with W3CEndpointReference record as a parameter"
source: "fgl-topics/c_gws_stateful_services_012.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Server side > Step 2: Create a stateful WS-Addressing enabled Web service"
type: "concept"
description: "The Genero Web Service extension provides a new Web service constructor called createStatefulWebService() to perform stateful services. This function works as the stateless constructor, but expects a ..."
---

# Step 2: Create a stateful WS-Addressing enabled web service with W3CEndpointReference record as a parameter

The Genero Web Service extension provides a new Web service constructor called [createStatefulWebService()](../15_library-reference/3759-com-webservice-createstatefulwebservice.md "Creates a new object to implement a stateful Web service.") to
perform stateful services. This function works as the stateless constructor, but expects a
W3CEndpointReference record as parameter.

For
example:

```
DEFINE serv com.WebService

# Create a stateful service with a W3CEndpointReference state variable
LET serv = com.WebService.CreateStatefulWebService(
  "StatefulWSAddressingService","http://mywebsite.com/services",
  EndpointReferenceState)

# Enable support of WS-Addressing 1.0
CALL serv.setFeature("WS-Addressing1.0","REQUIRED")
```

## Related links

**Related concepts**  

[Step 3: Publish a web service operation returning the W3CEndpointReference state variable and set it as session initiator](4523-step-3-publish-a-web-service-operation.md "Step 3: Publish a web service operation returning the W3CEndpointReference state variable and set it as session initiator")
