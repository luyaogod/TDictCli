---
title: "Implicit creation of certificates for HTTPS"
source: "fgl-topics/c_fgl_Migrate_to_250_https_certif.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > Implicit creation of certificates for HTTPS"
type: "concept"
---

# Implicit creation of certificates for HTTPS

> Certificates for HTTPS are now created implicitly when nothing is specified in FGLPROFILE.

Before version 2.50, certificates for HTTPS had to be specified
explicitly in FGLPROFILE.

Starting with 2.50, no HTTPS certificate is defined in FGLPROFILE, when a web services program
starts, the creation is implicit.

## Related links

**Related concepts**  

[HTTPS configuration](../16_web-services/4569-https-configuration.md "If no client certificate is provided, Genero Web Services (GWS) does the HTTPS request transparently.")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
