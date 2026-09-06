---
title: "Configuring the Apache web server for HTTPS"
source: "fgl-topics/c_gws_ssl_deployment_server_001.html"
breadcrumb: "Web services > Deploy a Web Service > Configuring the Apache web server for HTTPS"
type: "concept"
---

# Configuring the Apache web server for HTTPS

> Configuration steps to secure a web service for Apache server in HTTPS.

To secure your web server and applications, you must send a
Certificate Signing Request to one of the trusted Certificate Authority companies on the Internet
that will provide you with a certificate you can trust.

For information on generating a Certificate Signing Request, go to [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.").

For testing purposes, you may need to create a self-signed certificate.
To do so, start with the task to [Create a root certificate authority](4580-create-a-root-certificate-authority.md "Create a local root certificate authority for signing test certificates.") and follow with the
procedure to create the self-signed certificate in [Create a certificate](4581-create-a-certificate.md "Create a server or client certificate for use with SSL/TLS, and optionally a self-signed certificate for testing.").

Once you have a certificate, either issued and signed by a trusted Certificate Authority or self
signed for testing, follow the steps in this section to configure the Apache server for HTTPS.

## Related links

1. [Step 1: Create the server's certificate authority list](4910-step-1-create-the-server-s-certificate-authority-list.md)

   Create the certificate authority list and add the root certificate authority certificate.
2. [Step 2: Register the server as a Web service in the GAS](4911-step-2-register-the-server-as-a-web-service-in-the-gas.md)

   Web services registered on the GAS are started automatically when the GAS starts.
3. [Step 3: Configure Apache for HTTPS](4912-step-3-configure-apache-for-https.md)

   Add the locations of your certificates to the Apache configuration file.
4. [Step 4: Configure Apache for HTTP basic authentication](4913-step-4-configure-apache-for-http-basic-authentication.md)

   Create login details for an authenticated user for the Apache web server, and add the location of your authentication file to the Apache configuration file.

**Related concepts**  

[Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")

[The OpenSSL tool](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.")

**Related tasks**  

[View a certificate](4585-view-a-certificate.md "View the details of a certificate using the openssl command.")
