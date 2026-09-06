---
title: "Create a certificate authority list"
source: "fgl-topics/t_gws_ssl_openssl_007.html"
breadcrumb: "Web services > Security > Certificates in practice > Create a certificate authority list"
type: "task"
---

# Create a certificate authority list

> Create a CA list file containing the trusted certificate authorities used to verify certificates.

A CA list is a file containing one or more trusted Certificate Authority (CA) certificates in
PEM format. Clients use it to verify that a server or client certificate was signed by a trusted CA.
If you are using the local root CA created in [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates."), your CA
list contains a single entry.

Concatenate the certificate authority certificates into a single CA list file, in certificate
chain order:

```
$ openssl x509 -in MyCA1.crt -text >> CAList.pem
$ openssl x509 -in MyCA2.crt -text >> CAList.pem
$ openssl x509 -in MyCA3.crt -text >> CAList.pem
```

For the local test setup from [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates."), the CA list contains
a single entry:

```
$ openssl x509 -in MyRootCA.crt -text >> CAList.pem
```

## Related links

**Related tasks**  

[Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates.")

[Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.")

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")
