---
title: "Import a CA into the Windows key store"
source: "fgl-topics/t_gws_ssl_openssl_009.html"
breadcrumb: "Web services > Security > Certificates in practice > Import a CA into the Windows® key store"
type: "task"
---

# Import a CA into the Windows key store

> Import a Certificate Authority (CA) into the Windows key store so that Windows trusts certificates signed by that CA.

When using a local root CA for testing, as created in [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates."),
Windows does not trust it by default. Importing the CA certificate into the Windows key store
instructs Windows to trust any certificate signed by that CA.

1. Open the .crt certificate file.
2. Click **Install Certificate** and follow the instructions provided.

   Windows automatically places the certificate in the certificate authority list of the key
   store.

## Related links

**Related tasks**  

[Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates.")

[Create a certificate authority list](4582-create-a-certificate-authority-list.md "Create a CA list file containing the trusted certificate authorities used to verify certificates.")

[Import a certificate into the Windows key store](4583-import-a-certificate-into-the-windows-key-store.md "Import a certificate and its private key into the Windows key store as a PKCS12 file.")
