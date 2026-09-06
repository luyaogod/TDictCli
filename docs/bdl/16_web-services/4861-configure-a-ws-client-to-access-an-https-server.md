---
title: "Configure a WS client to access an HTTPS server"
source: "fgl-topics/t_gws_ssl_client_tutorial_rest.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Configure a WS client to access an HTTPS server"
type: "task"
---

# Configure a WS client to access an HTTPS server

> Configuration steps to access a server in HTTPS.

To configure access to an HTTPS server, you need a client certificate.

Before you begin, there are options to consider depending on how you wish to use the client certificate:

- If you do not have the certificate information in your FGLPROFILE file, Genero Web Services
  creates a certificate for you. This is an implicit or temporary certificate that is valid for a
  session only. For more information, go to [HTTPS configuration](4569-https-configuration.md "If no client certificate is provided, Genero Web Services (GWS) does the HTTPS request transparently.").

  For the
  implicit certificate, no configuration is required.
- Alternatively, for stronger security you should have your own certificate signed by a trusted
  certificate authority that the web server may verify before granting access. You should configure
  your application to use the certificate by adding the configuration details to the FGLPROFILE file.
  Follow the steps outlined in this section.
  > **Important:**
  >
  > In a production environment, it is not recommended to use self-signed certificates.

  In a production environment, some servers provide a client certificate and you use the
  certificate as provided, and add the configuration details to the FGLPROFILE file.

  Most servers
  do not check the identity of the clients. For these servers, the client's certificate does not
  necessarily need to be trusted; it is only used for data encryption purpose. If, however, the server
  performs client identification, you must trust a Certificate Authority in which it has total
  confidence concerning the validity of the client's certificates.

To get a certificate signed by a trusted certificate authority, you must send a Certificate
Signing Request to one of the trusted Certificate Authority companies on the Internet that will
provide you with a certificate you can trust. For information on generating a Certificate Signing
Request, go to [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.").

For testing purposes, you may need to create a self-signed certificate.
To do so, start with the task to [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates.") and follow with the
procedure to create the self-signed certificate in [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.").

Once you have a certificate, either issued and signed by a trusted Certificate Authority or self
signed for testing, follow the three steps in this section to configure the FGLPROFILE security
entries used by the client's Genero Web Services during HTTPS communication.

## Related links

1. [Set FGLPROFILE entries for the client certificate](4862-set-fglprofile-entries-for-the-client-certificate.md)

   Configure your application to use the certificate and the associated private key used by the client's Genero Web Services during HTTPS communication. For production systems, you add the configuration details to your FGLPROFILE file.
2. [Set FGLPROFILE entries for the CA list](4863-set-fglprofile-entries-for-the-ca-list.md)

   Clients need to check to see if the server's certificate is trusted. This is done using a certificate authority list.
3. [Set FGLPROFILE entries for the server URL](4864-set-fglprofile-entries-for-the-server-url.md)

   Add a set of configuration entries that specify the URL and the identity for the HTTPS server.

**Related concepts**  

[Configuring the Apache web server for HTTPS](4909-configuring-the-apache-web-server-for-https.md "Configuration steps to secure a web service for Apache server in HTTPS.")

[Certificates in practice](4578-certificates-in-practice.md "Procedures and tools for creating, importing, and viewing certificates and keys.")

[Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")

[Accessing secured services](4568-accessing-secured-services.md "Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.")
