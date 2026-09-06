---
title: "Enable OCSP"
source: "fgl-topics/c_gws_certificates_oscp.html"
breadcrumb: "Web services > Security > Enable OCSP"
type: "concept"
---

# Enable OCSP

> To enable Online Certificate Status Protocol (OCSP), set the security.global.ocsp.enable and security.global.ocsp.url entries in FGLPROFILE.

When these options are set, for each HTTPS connection, once the X509 certificate has been
validated, the Web service will check whether all certificates used for that validation are still
valid and have not been revoked at the time of the connection.

## Related links

**Related concepts**  

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
