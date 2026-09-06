---
title: "Step 3: Instantiate a new session by calling the web service operation set as session initiator"
source: "fgl-topics/c_gws_stateful_services_053.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Client side > Step 3: Instantiate a new session"
type: "concept"
description: "Call the BDL function generated from the WSDL that was defined as session initiator on the server. This function returns a new HTTP Cookie saved in the Binding.Cookie member of the global service ..."
---

# Step 3: Instantiate a new session by calling the web service operation set as session initiator

Call the BDL function generated from the WSDL that was defined as session initiator on the
server. This function returns a new HTTP Cookie saved in the `Binding.Cookie` member
of the global service variable of type [tGlobalEndpointType](4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records."). If your application handles several instances, you will have to copy
and store that cookie in your application to identify a service instance for further requests.

For
example:

```
DISPLAY "Creating a new instance ..."
LET wsstatus = GetInstance_g() # call the service session 
                               # initiator web function
IF wsstatus == 0 THEN
  # copy the service returned HTTP cookie
  LET instance1 = 
   StatefulCookieService_StatefulCookieServicePortTypeEndpoint.Binding.Cookie 
ELSE
  ... handle soap errors
END IF
```

When creating a new instance, ensure that the `Binding.Cookie` member of the
generated global variable of type [tGlobalEndpointType](4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.") has been set to NULL, otherwise the server will throw an error.

## Related links

**Related concepts**  

[Step 4: Call any web service operation with previously returned HTTP cookie](4543-step-4-call-any-web-service-operation.md "Step 4: Call any web service operation with previously returned HTTP cookie")
