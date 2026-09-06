---
title: "Create a certificate"
source: "fgl-topics/t_gws_ssl_openssl_006.html"
breadcrumb: "Web services > Security > Certificates in practice > Create a certificate"
type: "task"
---

# Create a certificate

> Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.

A **server certificate** proves the server's identity to clients. A **client
certificate** proves the client's identity to the server, and is used for web services or web
applications requiring mutual authentication.

To secure your web server and applications, you must send a
Certificate Signing Request to one of the trusted Certificate Authority companies on the Internet
that will provide you with a certificate you can trust.

The CSR must be created on the server where the certificate is to be deployed. Use the [openssl](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.") tool to create the CSR.

The steps in this topic mirror the real-world process of obtaining a certificate from a
Certificate Authority (CA). The root CA created in [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates.") acts
as the CA, and the certificate created here is the equivalent of one issued by a trusted CA on the
internet.

1. Create a Certificate Signing Request and private key:

   ```
   $ openssl req -new -out MyCert.csr
   ```

   You are prompted to provide a Distinguishing Name (DN) for the
   certificate, including fields such as Country Name, State or Province Name, Organization Name, and
   Common Name (CN).

   Set the CN according to the intended use of the certificate:

   - **Server certificate:** The CN must match the server's domain name (DNS), otherwise
     clients will not trust it. For example, for
     https://www.MyServer.com/fastcgi/ws/r/MyWebService, set the CN to
     www.MyServer.com. Use the `subjectAltName` parameter to
     cover variations of the domain name (for example, MyServer.com,
     mail.MyServer.com).
     > **Tip:**
     >
     > You can also add the IP address in
     > `SubjectAltName` as an alternative name for the CN. For details, go to
     > [OpenSSL](https://www.openssl.org).
   - **Client certificate:** The CN does not need to match the client's DNS, provided the
     certificate is issued by a trusted CA. You can leave the CN field open.

   Fill in the remaining fields with your organization details, required if you are purchasing the
   certificate from a certificate authority.

   Two files are created: MyCert.csr and a private key file
   privkey.pem (the default OpenSSL name for private key files).
   > **Note:**
   >
   > **About the CSR and its private key:** 
   >
   > - If you want an official Certificate Authority, you must send the CSR file to one of the
   >   self-established Certificate Authority companies on the Internet instead of creating it with
   >   openssl. See [Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.").
   > - The CSR file is also used to encrypt messages that only its corresponding private key can
   >   decrypt.

To create a test certificate signed by a local certificate authority, perform the
following steps:

2. Remove the private key password (Optional):

   ```
   $ openssl rsa -in privkey.pem -out MyCert-nopass.pem
   ```

   The unprotected private key is output in MyCert-nopass.pem.
3. Create a certificate from the CSR signed by the certificate created in [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates."):

   (line breaks added for readability)

   ```
   $ openssl x509 -in MyCert.csr -out MyCert.crt -req
      -CA MyRootCA.crt -CAkey MyRootCA.pem -days 365
   ```

   The certificate is output in MyCert.crt.

Once you have a certificate, configure it according to its intended use:

| Use | Action |
| --- | --- |
| Server | Configure in your web server (for example, IIS or Apache). For an example, go to [Configuring the Apache web server for HTTPS](4909-configuring-the-apache-web-server-for-https.md "Configuration steps to secure a web service for Apache server in HTTPS."). |
| Client (web service) | Configure entries in the FGLPROFILE file. For details, go to [Set FGLPROFILE entries for the client certificate](4571-set-fglprofile-entries-for-the-client-certificate.md "Configure your application to use the certificate and the associated private key used by the client's Genero Web Services during HTTPS communication. For production systems, you add the configuration details to your FGLPROFILE file."). |
| Client (web application, HTTPS) | Install the certificate in the browser or OS keystore/keychain. |

For client certificates, you also need to declare the server's CA List:

| Use | Action |
| --- | --- |
| Client (web service) | Configure entries in the FGLPROFILE file. For details, go to [Set FGLPROFILE entries for the CA list](4572-set-fglprofile-entries-for-the-ca-list.md "Clients need to check to see if the server's certificate is trusted. This is done using a certificate authority list."). |
| Client (web application, HTTPS) | Install the CA List in the browser or OS keystore/keychain. For an example using a Windows keystore, go to [Import a CA into the Windows key store](4584-import-a-ca-into-the-windows-key-store.md "Import a Certificate Authority (CA) into the Windows key store so that Windows trusts certificates signed by that CA."). For details on creating a CA list, go to [Create a certificate authority list](4582-create-a-certificate-authority-list.md "Create a CA list file containing the trusted certificate authorities used to verify certificates."). |

## Related links

**Related tasks**  

[Configure a WS client to access an HTTPS server](4570-configure-a-ws-client-to-access-an-https-server.md "Configuration steps to access a server in HTTPS.")
