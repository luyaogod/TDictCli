---
title: "Step 5: Restore the session in any BDL web operation from the state variable"
source: "fgl-topics/c_gws_stateful_services_045.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Server side > Step 5: Restore the session"
type: "concept"
description: "In a published BDL web function, the SOAP engine deserializes the HTTP Cookie called GSESSIONID from the HTTP layer into the state variable. You can then retrieve the session in BDL via that state ..."
---

# Step 5: Restore the session in any BDL web operation from the state variable

In a published BDL web function, the SOAP engine deserializes the HTTP Cookie called
**GSESSIONID** from the HTTP layer into the state variable. You can then retrieve the session in
BDL via that state variable.

For
example:

```
FUNCTION MyFunction()
  IF ServiceState IS NULL THEN
    CALL com.WebServiceEngine.SetFaultString("Invalid session id")
    RETURN
  ELSE
    ... Restore the service session based on the ServiceState
        variable from the database
  END IF
  ... Process the operation
END FUNCTION
```

## Related links

**Related concepts**  

[Step 6: Deployment recommendation](4538-step-6-deployment-recommendation.md "Step 6: Deployment recommendation")
