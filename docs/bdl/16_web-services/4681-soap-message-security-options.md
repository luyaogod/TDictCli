---
title: "SOAP message security options"
source: "fgl-topics/c_gws_ssl_security_how_to_007.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > SOAP security standards > SOAP message security options"
type: "concept"
---

# SOAP message security options

> Describes the Wss10 SOAP Message Security 1.0 options that are supported.

```
<sp:Wss10 xmlns:sp="http://schemas.xmlsoap.org/ws/2005/07/securitypolicy">
  <sp:MustSupportRefKeyIdentifier />
  <sp:MustSupportRefIssuerSerial />
```

- `MustSupportRefKeyIdentifier` means that initiator and recipient are able to
  generate and process key identifier reference.
- `MustSupportRefIssuerSerial` means that initiator and recipient are able to
  generate and process issuer and token serial reference.
