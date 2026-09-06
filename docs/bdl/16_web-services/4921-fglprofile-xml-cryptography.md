---
title: "FGLPROFILE: XML cryptography"
source: "fgl-topics/c_gws_ssl_configuration_011.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > Examples > FGLPROFILE: XML cryptography"
type: "concept"
---

# FGLPROFILE: XML cryptography

> Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.

```
# Security configuration
security.global.agent     = "4444"

# Crypto configuration
xml.keystore.calist     = "RSARootCertificate.crt;DSARootCertificate.crt"
xml.keystore.cadir      = "/opt/usr/caCerts"
xml.keystore.x509list   = "RSA1024Certificate.crt;DSA1024Certificate.crt"
xml.id1.x509          = "RSA1024Certificate.crt"
xml.id2.x509          = "DSA1024Certificate.crt"
xml.id3.key           = "RSA1024Key.pem"
xml.id4.key           = "DSA1024Key.der"
xml.id5.key           = "HMAC.bin"
```

## Related links

**Related concepts**  

[Encryption, BASE64 and password agent with fglpass tool](4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent.")

**Related reference**  

[XML configuration](4916-fglprofile-entries-for-web-services.md "XML configuration")
