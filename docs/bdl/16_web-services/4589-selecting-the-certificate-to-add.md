---
title: "Selecting the certificate to add"
source: "fgl-topics/c_gws_certificates_006.html"
breadcrumb: "Web services > Security > Examining certificates > Selecting the certificate to add"
type: "concept"
---

# Selecting the certificate to add

> The certificate authority (CA) is the authority that validates the server. The certificate to add to the CA list is the authority certificate, not the server certificate.

There are default certificates known by browsers like:

- VeriSign: <http://www.verisign.com/support/roots.html>
- Thawte: <http://www.thawte.com/roots/index.html>

Get the server issuer certificate (and all the parents, grandparents, and so on).

For example, if your server is validated by Thawte, add the Thawte certificate to the list.

To check whether your certificate is the CA certificate, search for the CNs (Common Names) in the
.cer files. The CA Subject entry should be the Issuer CN in the server
certificate. Running the openssl command as follows outputs the CN:

```
openssl x509 -in server.pem -noout -subject
```

subject= /C=ZA/ST=Western Cape/L=Cape Town/O=Thawte Consulting
cc/OU=CertificationServices Division/  **CN=Thawte Server CA** /emailAddress=server-certs@thawte.com

To convert a .cer certificate to the .pem format used
by Genero Web Services run the following openssl
command:

```
openssl x509 -inform DER -in server.cer -outform PEM -out server.crt
```
