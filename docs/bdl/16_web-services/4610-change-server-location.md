---
title: "Change server location"
source: "fgl-topics/c_gws_client_behavior_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Change server location"
type: "concept"
---

# Change server location

> If the Web service server location changes, you can update the address in the global endpoint record.

To change the server location at runtime, set the record `Uri` member with a valid
URL of another service. All services must respect the same WSDL contract. If you leave the variable
unset, the client will connect to the server URL defined in the WSDL at code generation time.

Example:

```
LET Calculator_CalculatorPortTypeEndpoint.Address.Uri = 
  "http://zeus:1111/mydomain/Calculator"
```

You can assign this variable with a URL set in the FGLPROFILE (see [Logical Service location](4629-use-logical-names-for-service-locations.md "Using a logical reference for the Web service, instead of the real URL, in your client application URL binding has advantages for working with and deploying applications.")).

If you are migrating from a version prior to 2.40, see also [Web Services changes](../05_upgrading/0244-web-services-changes.md "There are changes in support of web services in Genero 2.40.").
