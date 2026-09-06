---
title: "Step 4: Create the BDL session initiator function and instantiate a new session"
source: "fgl-topics/c_gws_stateful_services_044.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Server side > Step 4: Create the BDL session initiator function"
type: "concept"
description: "In your BDL function declared as session initiator, you must: Handle the creation of the session. Fill the state variable before returning from the function. Store the state variable in a database ..."
---

# Step 4: Create the BDL session initiator function and instantiate a new session

In your BDL function declared as session initiator, you must:

- Handle the creation of the session.
- Fill the state variable before returning from the function.
- Store the state variable in a database based on the state variable (in order to keep the session
  across consecutive requests from the same client).

For example:

```
FUNCTION GetInstance()
  # Generate an unique string (can come from a database table id)
  LET ServiceState = security.RandomGenerator.CreateUUIDString()
      ... Store ServiceState value into database or use directly a 
          database table entry to hold the session
END FUNCTION
```

Ensure that `IMPORT security` is called at the beginning of the file when using
[security.RandomGenerator.CreateUUIDString()](../15_library-reference/4412-security-randomgenerator-createuuidstring.md "Creates a new universal unique identifier (UUID).").

## Related links

**Related concepts**  

[Step 5: Restore the session in any BDL web operation from the state variable](4537-step-5-restore-the-session.md "Step 5: Restore the session in any BDL web operation from the state variable")
