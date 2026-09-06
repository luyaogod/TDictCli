---
title: "Enabling MTOM on the server side"
source: "fgl-topics/c_gws_srvr_enable_mtom.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Enabling MTOM on the server side"
type: "concept"
---

# Enabling MTOM on the server side

> Enable the Message Transmission Optimization Mechanism (MTOM) feature to efficiently send binary data to and from Web services.

By default, MTOM is not enabled. As a result, any BYTE represented as `xsd:base64`
or `xsd:hexbinary` will be inlined in the SOAP request or response.

To enable MTOM on server side, set the MTOM feature to `TRUE`.

```
# Create Web Service
LET service = com.WebService.CreateWebService("MTOMStaxService","https://mywebsite.com/services/stax/mtom")
CALL service.setFeature("MTOM",true)
```

With MTOM enabled on the server side, the generated WSDL includes the MTOM policy.
