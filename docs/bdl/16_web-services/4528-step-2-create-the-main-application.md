---
title: "Step 2: Create the MAIN application"
source: "fgl-topics/c_gws_stateful_services_022.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Client side > Step 2: Create the MAIN application"
type: "concept"
description: "In your main application: Import the XML library. This is due to the support of WS-Addressing 1.0 with IMPORT XML . Import the generated .inc file with GLOBALS \"ws_stub.inc\" Manage the WS-Addressing ..."
---

# Step 2: Create the MAIN application

In your main application:

1. Import the XML library. This is due to the support of WS-Addressing 1.0 with `IMPORT
   XML`.
2. Import the generated .inc file with `GLOBALS
   "ws_stub.inc"`
3. Manage the WS-Addressing 1.0 reference parameters representing the session state (if your client
   has to handle several instances of the same service).

For
example:

```
IMPORT XML # Import the XML library required for WS-Addressing 1.0

GLOBALS "ws_stub.inc" # Import service global definition

TYPE InstanceType DYNAMIC ARRAY OF xml.DomDocument
 	# End point WSA reference parameters

DEFINE instance1,instance2,instance3 InstanceType
   # Store the different sessions the client will have to manage

MAIN
  ...
END MAIN
```

## Related links

**Related concepts**  

[Step 3: Instantiate a new session by calling the web service operation set as session initiator](4529-step-3-instantiate-a-new-session.md "Step 3: Instantiate a new session by calling the web service operation set as session initiator")
