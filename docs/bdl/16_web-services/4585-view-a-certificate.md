---
title: "View a certificate"
source: "fgl-topics/t_gws_ssl_openssl_010.html"
breadcrumb: "Web services > Security > Certificates in practice > View a certificate"
type: "task"
---

# View a certificate

> View the details of a certificate using the openssl command.

Use this command to verify certificate details such as the expiry date, issuer, and Common
Name (CN).

Run the following command:

```
openssl x509 -in MyCert.crt -noout -text
```

## Related links

**Related concepts**  

[Troubleshoot certificate issues](4575-troubleshoot-certificate-issues.md "You may encounter known (and common) issues when completing the Genero Web Services tutorials or when adding Web services of your own. These issues and their solutions are presented in the following topics.")
