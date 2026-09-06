---
title: "Step 4: Call any web service operation with previously returned HTTP cookie"
source: "fgl-topics/c_gws_stateful_services_054.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Client side > Step 4: Call any Web service operation"
type: "concept"
description: "Before calling any Web service operation, set the HTTP cookie returned by a session initiator function to identify the session to the server. For example: # use instance1 LET ..."
---

# Step 4: Call any web service operation with previously returned HTTP cookie

Before calling any Web service operation, set the HTTP cookie returned by a session initiator
function to identify the session to the server.

For
example:

```
# use instance1
LET StatefulCookieService_StatefulCookieServicePortTypeEndpoint.Binding.Cookie =
 instance1 
# Call web operation MyFunction of instance 1
CALL MyFunction("Hello") RETURNING wsstatus,ret
```

## Related links

**Related concepts**  

[Step 5: Troubleshooting](4544-step-5-troubleshooting.md "Step 5: Troubleshooting")
