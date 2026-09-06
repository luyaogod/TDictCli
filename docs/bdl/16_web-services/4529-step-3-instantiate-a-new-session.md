---
title: "Step 3: Instantiate a new session by calling the web service operation set as session initiator"
source: "fgl-topics/c_gws_stateful_services_023.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Client side > Step 3: Instantiate a new session"
type: "concept"
description: "Call the BDL function generated from the WSDL that is defined as session initiator on the server. This function returns a W3CEndpointReference parameter that contains the WS-Addressing 1.0 reference ..."
---

# Step 3: Instantiate a new session by calling the web service operation set as session initiator

Call the BDL function generated from the WSDL that is defined as session initiator on the server.
This function returns a `W3CEndpointReference` parameter that contains the
WS-Addressing 1.0 reference parameters representing the new instance created on server side.

If your application handles several instances, you will have to copy and store those parameters
in your application to identify a service instance for further requests.

As the WS-Addressing 1.0 reference parameters are defined as an XML document, they are
represented as a dynamic list of `xml.DomDocument` in BDL.

For example:

```
DISPLAY "Creating a new instance ..."
LET wsstatus = GetInstance_g() # call the service session initiator 
                               # web function
IF wsstatus == 0 THEN
   FOR ind=1 TO
     ns1GetInstanceResponse.return.ReferenceParameters._LIST_0.getLength()
      LET instance1[ind]=
        ns1GetInstanceResponse.return.ReferenceParameters._LIST_0[ind].clone() 
        # copy the service returned WS-Addressing 1.0 reference parameters
   END FOR
ELSE
     ... handle soap errors
END IF
```

> **Important:**
>
> When creating a new instance, ensure that the Parameters member of the
> generated global variable of type [tWSAGlobalEndpointType](4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.") has
> been set to `NULL`, otherwise the server will throw an error.

## Related links

**Related concepts**  

[Step 4: Call any web service operation with previously returned WS-Addressing 1.0 reference parameters](4530-step-4-call-any-web-service-operation.md "Step 4: Call any web service operation with previously returned WS-Addressing 1.0 reference parameters")
