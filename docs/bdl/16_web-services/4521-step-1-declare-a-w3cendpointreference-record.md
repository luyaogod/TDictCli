---
title: "Step 1: Declare a W3CEndpointReference record to be used as state variable"
source: "fgl-topics/c_gws_stateful_services_011.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Server side > Step 1: Declare a W3CEndpointReference record"
type: "concept"
description: "This record MUST have: A mandatory member of type STRING, where you can define a different service end point URL, otherwise the current server URL will be used. A sub record to contain one or more BDL ..."
---

# Step 1: Declare a W3CEndpointReference record to be used as state variable

This record **MUST** have:

- A mandatory member of type STRING, where you can define a different service end point URL,
  otherwise the current server URL will be used.
- A sub record to contain one or more BDL variables used as state variables and defined as
  reference parameter in the WS-Addressing 1.0 specification.

For
example:

```
DEFINE EndpointReferenceState RECORD ATTRIBUTES(W3CEndpointReference)
  address STRING, # Mandatory
  ref RECORD # Sub-record Reference parameters containing one 
             # or more state variables
    OpaqueID STRING ATTRIBUTES(XMLName="OpaqueID"), # Unique ID to
               #identify the service state in the database
    Expiration DATE ATTRIBUTES(XMLName="Expiration",
      XMLNamespace="http://tempuri.org") # Session state expiration date
  END RECORD
END RECORD
```

You can use a unique ID from a database table to manage the Web services sessions in place of
`OpaqueID`.

## Related links

**Related concepts**  

[Step 2: Create a stateful WS-Addressing enabled web service with W3CEndpointReference record as a parameter](4522-step-2-create-a-stateful-ws-addressing-enabled-web-service.md "Step 2: Create a stateful WS-Addressing enabled web service with W3CEndpointReference record as a parameter")
