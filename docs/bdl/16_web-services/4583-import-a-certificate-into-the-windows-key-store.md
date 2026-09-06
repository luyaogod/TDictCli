---
title: "Import a certificate into the Windows key store"
source: "fgl-topics/t_gws_ssl_openssl_008.html"
breadcrumb: "Web services > Security > Certificates in practice > Import a certificate into the Windows® key store"
type: "task"
---

# Import a certificate into the Windows key store

> Import a certificate and its private key into the Windows key store as a PKCS12 file.

To use a certificate in Windows applications, you must import it into the Windows key store.
The certificate and its private key are first combined into a PKCS12 (.p12)
file, which is password-protected and can be safely transported.

1. Create a PKCS12 file containing the certificate and its private key:

   ```
   $ openssl pkcs12 -export -inkey MyCert-nopass.pem -in MyCert.crt -out MyCert.p12
   ```

   The .p12 file is protected by a password and can be transported without
   any risk.
2. On a Windows® system, open the
   .p12 file and follow the instructions provided.

   If you select strong verification during the import process, Windows displays a dialog box each
   time an application accesses the private key, asking whether the application is allowed to use
   it.

## Related links

**Related tasks**  

[Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.")

[Import a CA into the Windows key store](4584-import-a-ca-into-the-windows-key-store.md "Import a Certificate Authority (CA) into the Windows key store so that Windows trusts certificates signed by that CA.")
