---
title: "Step 2: Create the MAIN application"
source: "fgl-topics/c_gws_stateful_services_052.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Client side > Step 2: Create the MAIN application"
type: "concept"
description: "In your main application: Import the generated .inc file with GLOBALS \"ws_stub.inc\" . Manage the HTTP cookies representing the session state (if your client has to handle several instances of the same ..."
---

# Step 2: Create the MAIN application

In your main application:

- Import the generated .inc file with `GLOBALS
  "ws_stub.inc"`.
- Manage the HTTP cookies representing the session state (if your client has to handle several
  instances of the same service).

For
example:

```
GLOBALS "ws_stub.inc" # Import service global definition

# Store the different sessions the client will have to manage 
# in a string
DEFINE instance1,instance2,instance3 String 

MAIN
  ...
END MAIN
```

## Related links

**Related concepts**  

[Step 3: Instantiate a new session by calling the web service operation set as session initiator](4542-step-3-instantiate-a-new-session.md "Step 3: Instantiate a new session by calling the web service operation set as session initiator")
