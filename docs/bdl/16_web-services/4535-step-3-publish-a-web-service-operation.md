---
title: "Step 3: Publish a Web service operation defined as session initiator"
source: "fgl-topics/c_gws_stateful_services_043.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Server side > Step 3: Publish a Web service operation"
type: "concept"
description: "Define which Web service operation will initiate the session on your service and instantiate a new session. All other Web service operations (not defined as session initiator) will return an error if ..."
---

# Step 3: Publish a Web service operation defined as session initiator

Define which Web service operation will initiate the session on your service and instantiate a
new session. All other Web service operations (not defined as session initiator) will return an
error if they don't get an HTTP cookie called **GSESSIONID**.

For
example:

```
DEFINE op com.WebOperation
LET op = com.WebOperation.CreateDocStyle("GetInstance","GetInstance",NULL,NULL)
CALL op.initiateSession(true)
CALL serv.publishOperation(op,NULL)
```

There is no restriction on the Web service session initiator function regarding the input and
output parameters.

## Related links

**Related concepts**  

[Step 4: Create the BDL session initiator function and instantiate a new session](4536-step-4-create-the-bdl-session-initiator-function.md "Step 4: Create the BDL session initiator function and instantiate a new session")
