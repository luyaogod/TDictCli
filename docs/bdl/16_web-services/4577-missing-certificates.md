---
title: "Missing certificates"
source: "fgl-topics/c_gws_certificates_007.html"
breadcrumb: "Web services > Security > Troubleshoot certificate issues > Missing certificates"
type: "concept"
---

# Missing certificates

> Identifying missing certificates.

Sometimes the CA hierarchy described in the server certificate is incomplete or needs another
certificate (default ones used by browsers or private ones).

![Screen shot of server certificate with incomplete hierarchy](../_images/FedExCert.jpg)

*Certificate Viewer in Firefox Web Browser; Details Tab*

When this occurs, you will get this kind of error message when you set
FGLWSDEBUG:

```
WS-DEBUG (Security error)
Error with certificate at depth: 3
 issuer = /C=US/O=VeriSign, Inc./OU=Class 3 Public Primary Certification Authority
 subject = /C=US/O=VeriSign, Inc./OU=Class 3 Public Primary Certification Authority
 err 19:self signed certificate in certificate chain
WS-DEBUG END
```

This means OpenSSL is looking for a third ancestor that is not listed in the hierarchy above. In
this example, gatewaybeta.fedex.com only has two ancestors, and none are named
"Class 3 Public Primary Certification Authority". You need to download the root certificates from
VeriSign and add "Class 3 Public Primary Certification Authority" in your CA list.

If the [certificate authorities](4569-https-configuration.md) are not found in the operating system keystore, you need to download
them:

- If you store the CA certificates in $FGLDIR/web\_utilities/certs directory,
  ensure they are named with the .crt extension.
- If you have configured the FGLPROFILE file using the global certificate authority list entry
  `security.global.ca`, you need to create the CA list with the
  .pem extension.

  The Firefox® browser allows
  you to download the CA list as a pem file; a PEM (cert) or
  PEM (chain). Use the PEM (chain) file in this case.
  Otherwise, you will need to download the certs individually and create the pem
  chain file as described in [Create a certificate authority list](4582-create-a-certificate-authority-list.md "Create a CA list file containing the trusted certificate authorities used to verify certificates.").

## Related links

**Related concepts**  

[Error: Peer certificate is issued by a company not in our CA list](4576-error-peer-certificate-is-issued-by-a-company-not-in-our-ca.md "When a client connects to a server using HTTPS, the client needs to trust the server it is in communication with. So the client needs to add the server's CAs (certificate authorities lists) to its trusted CAs.")

[The OpenSSL tool](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.")
