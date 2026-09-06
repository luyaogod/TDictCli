---
title: "Error: Peer certificate is issued by a company not in our CA list"
source: "fgl-topics/c_gws_troubleshooting_003.html"
breadcrumb: "Web services > Security > Troubleshoot certificate issues > Error: Peer certificate is issued by a company not in our CA list"
type: "concept"
---

# Error: Peer certificate is issued by a company not in our CA list

> When a client connects to a server using HTTPS, the client needs to trust the server it is in communication with. So the client needs to add the server's CAs (certificate authorities lists) to its trusted CAs.

This error means the client CA list is missing a certificate authority in its CA list.

To display the client CA list, use the following
command:

```
openssl x509 -in ClientCAList.pem -noout -text
```

Solution:

1. Add the missing CA list to the client CA
   list.

   ```
   openssl x509 -in MyCompanyCA.crt -text >> ClientCAList.pem
   ```

## Theory

Usually certificates work in pairs: a public key and a private key.

![This image illustrates the hierarchy of certificate authority for the client and server. Each has a CA root authority at the root of the hierarchy that signs certificates. The client CA hierarchy needs to have all the server CA, including the server CA root.](../_images/Certificate_Theory.jpg)

*Certificates working in pairs: a public key and a private key*

The certificate which appears at the top of the list on the client and server side is the
Certificate Authority (CA) root. This means that the client has a certificate that can be signed by
an authority signed itself by a root authority. Likewise, the server has a certificate that can be
signed by an authority signed itself by a root authority. In some instances, a certificate can be
signed by itself.

Things to note:

- The server certificate is expected to have its host name set as CN (Common Name). For example,
  if you want to access the server https://www.mycompany.com the CN is expected
  to be "www.mycompany.com".
- In the client CA list it is recommended that you have all the CA of the server. In this example
  you need the server CA (5) and the server CA Root (4). If the server is self-signed then add the
  server certificate (6) to the client CA list.
- Sometimes, the needed CAs are not listed in the certificates hierarchy. Setting environment
  variable FGLWSDEBUG=3, will give you information about the missing CA.

## Related links

**Related concepts**  

[Missing certificates](4577-missing-certificates.md "Identifying missing certificates.")

[Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")

[FGLWSDEBUG](../07_configuration/0537-fglwsdebug.md "The FGLWSDEBUG environment variable enables web services library debugging.")

**Related tasks**  

[Set FGLPROFILE entries for the CA list](4572-set-fglprofile-entries-for-the-ca-list.md "Clients need to check to see if the server's certificate is trusted. This is done using a certificate authority list.")
