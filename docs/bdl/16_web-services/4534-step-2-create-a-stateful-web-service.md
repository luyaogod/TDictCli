---
title: "Step 2: Create a stateful web service with state variable as parameter"
source: "fgl-topics/c_gws_stateful_services_042.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Server side > Step 2: Create a stateful Web service"
type: "concept"
description: "The Genero Web Service extension provides a new Web service constructor called createStatefulWebService() to perform stateful services. This function works as the stateless constructor, but expects a ..."
---

# Step 2: Create a stateful web service with state variable as parameter

The Genero Web Service extension provides a new Web service constructor called [createStatefulWebService()](../15_library-reference/3759-com-webservice-createstatefulwebservice.md "Creates a new object to implement a stateful Web service.") to
perform stateful services. This function works as the stateless constructor, but expects a simple
state variable as parameter.

## Example

```
DEFINE serv com.WebService
# Create a stateful service with a simple BDL variable as state variable
LET serv = com.WebService.CreateStatefulWebService("StatefulCookieService",
           "http://mywebsite.com/services",ServiceState)
```

## Related links

**Related concepts**  

[Step 3: Publish a Web service operation defined as session initiator](4535-step-3-publish-a-web-service-operation.md "Step 3: Publish a Web service operation defined as session initiator")
