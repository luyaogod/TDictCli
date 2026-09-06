---
title: "Step 4: Call any web service operation with previously returned WS-Addressing 1.0 reference parameters"
source: "fgl-topics/c_gws_stateful_services_024.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Client side > Step 4: Call any web service operation"
type: "concept"
description: "Before calling any Web service operation, you must set the WS-Addressing 1.0 reference parameters returned by a session initiator function to identify the session to the server. For example: LET ..."
---

# Step 4: Call any web service operation with previously returned WS-Addressing 1.0 reference parameters

Before calling any Web service operation, you must set the WS-Addressing 1.0 reference parameters
returned by a session initiator function to identify the session to the server.

For
example:

```
LET StatefulWSAddressingService_StatefulWSAddressingServicePortTypeEndpoint.
Address.Parameters.* = instance1.* 
# assign WS-Addressing 1.0 reference parameters dynamic array by reference
CALL MyFunction("Hello") RETURNING wsstatus,ret  
# Call web operation MyFunction of instance 1
```
