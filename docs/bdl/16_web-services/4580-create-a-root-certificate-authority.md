---
title: "Create a root certificate authority"
source: "fgl-topics/t_gws_ssl_openssl_004.html"
breadcrumb: "Web services > Security > Certificates in practice > Create a root certificate authority"
type: "task"
---

# Create a root certificate authority

> Create a local root certificate authority for signing test certificates.

> **Important:**
>
> This task is optional because a root certificate authority is only needed if you are creating
> test certificates.

For details about creating a Certificate Signing Request, go to [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing."). For more information about certificate authorities, go to [Certificate authorities](4566-encryption-and-authentication.md).

Use the [openssl](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.") tool to create a root
certificate authority.

1. Create the root certificate authority serial file:

   ```
   $ echo 01 > MyRootCA.srl
   ```

   This command creates a serial file with an initial HEX value 01. OpenSSL uses this file to
   track the serial numbers of certificates it creates. The serial file is typically given the
   same name as the root CA with the extension .srl.
2. Create a certificate signing request and a private key:

   ```
   $ openssl req -new -out MyRootCA.csr
   ```

   You are prompted to provide a Distinguishing Name (DN) for the
   certificate, including fields such as Country Name, State or Province Name, Organization Name, and
   Common Name (CN).

   Two files are created, MyRootCA.csr and a file called
   privkey.pem. The private key file of a root certificate authority must be
   handled with care because it validates certificates it has signed and it is used in creating future
   certificates. As a result, it must not be accessible by other users.

Create a self-signed root CA certificate:

3. Remove the password of the private key (Optional):

   ```
   $ openssl rsa -in privkey.pem -out MyRootCA.pem
   ```

   You are prompted for the
   passphrase.
   > **Warning:**
   >
   > Removing the password of a certificate authority's private key is not recommended.

   The unprotected private key is output in MyRootCA.pem.
4. Create a certificate from the CSR that is valid for 730 days, and that is signed by the
   unprotected private key:

   ```
   $ openssl x509 -trustout -in MyRootCA.csr -out MyRootCA.crt
    -req -signkey MyRootCA.pem -days 730
   ```

   The root certificate authority certificate is output in
   MyRootCA.crt.

You can use MyRootCA.crt to encrypt data as a self-signed certificate, but
users will be shown a warning that the certificate is not trusted. To have it trusted, create your
own certificate signed by this certificate authority and install it as a trusted certificate in the
browser or in the keystore/keychain of the machine. For details, go to [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.").

## Related links

**Related concepts**  

[Configuring the Apache web server for HTTPS](4909-configuring-the-apache-web-server-for-https.md "Configuration steps to secure a web service for Apache server in HTTPS.")

**Related tasks**  

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")
