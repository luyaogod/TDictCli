---
title: "Step 5: Restore the session in any BDL web operation from the W3CEndpointReference record"
source: "fgl-topics/c_gws_stateful_services_015.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Server side > Step 5: Restore the session"
type: "concept"
description: "In a published BDL Web function, the SOAP engine deserializes the WS-Addressing 1.0 reference parameter headers in the W3CEndpointReference sub-record so that you can retrieve the session from the ..."
---

# Step 5: Restore the session in any BDL web operation from the W3CEndpointReference record

In a published BDL Web function, the SOAP engine deserializes the WS-Addressing 1.0 reference
parameter headers in the W3CEndpointReference sub-record so that you can retrieve the session from
the state variable.

For
example:

```
FUNCTION MyFunction()
  IF EndpointReferenceState.ref.OpaqueID IS NULL THEN
    CALL com.WebServiceEngine.SetFaultString("Invalid session id")
    RETURN
  ELSE
    ... Restore the service session based on the OpaqueID state 
    ... variable from the database
  END IF
  ... Process the operation
END FUNCTION
```

## Related links

**Related concepts**  

[Client side](4526-client-side.md "Perform these steps to communicate with a stateful Web service based on WS-Addressing 1.0.")
