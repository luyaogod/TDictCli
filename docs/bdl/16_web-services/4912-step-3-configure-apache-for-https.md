---
title: "Step 3: Configure Apache for HTTPS"
source: "fgl-topics/c_gws_ssl_deployment_server_006.html"
breadcrumb: "Web services > Deploy a Web Service > Configuring the Apache web server for HTTPS > Step 3: Configure Apache for HTTPS"
type: "concept"
---

# Step 3: Configure Apache for HTTPS

> Add the locations of your certificates to the Apache configuration file.

You must configure Apache to support HTTPS by adding the required modules. For more information,
refer to the [Apache Web
server documentation](https://httpd.apache.org/docs/) for your Apache version.

Once the Apache Web server supports HTTPS, you must change or add the following directives to the
Apache configuration file:

- Set the Apache Web server Certificate Authority List directive created in [Step 1](4910-step-1-create-the-server-s-certificate-authority-list.md "Create the certificate authority list and add the root certificate authority certificate.") :

  `SSLCACertificateFile
  D:/Apache-Server/conf/ssl/ServerCAList.pem`
- Set the Apache Web server Certificate and associated private key directives:

  `SSLCertificateFile
  D:/Apache-Server/conf/ssl/MyServer.crt`

  `SSLCertificateKeyFile
  D:/Apache-Server/conf/ssl/MyServer.pem`
- Require the Apache Web server to verify the validity of all client certificates:

  `SSLVerifyClient require`

The Apache Web server must be started on a machine where the host is the same as the
one defined in the subject of the server's certificate (www.MyServer.com
in our example).

In the next step we configure Apache for HTTP basic authentication, [Step 4: Configure Apache for HTTP basic authentication](4913-step-4-configure-apache-for-http-basic-authentication.md "Create login details for an authenticated user for the Apache web server, and add the location of your authentication file to the Apache configuration file.").
