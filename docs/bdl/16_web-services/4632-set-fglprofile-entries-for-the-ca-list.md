---
title: "Set FGLPROFILE entries for the CA list"
source: "fgl-topics/t_gws_ssl_config_cert_auth_soap.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Configure a WS client to access an HTTPS server > Set FGLPROFILE entries for the CA list"
type: "task"
---

# Set FGLPROFILE entries for the CA list

> Clients need to check to see if the server's certificate is trusted. This is done using a certificate authority list.

In this task you create the certificate authority list using the OpenSSL command line tool, and
set the global certificate authority entry in your FGLPROFILE file.

1. Create the certificate authority list.
   1. Access the URL of the HTTPS server and save its certificates to disk.

      Type the server's URL in your browser. When prompted, save **all** the certificates from
      the Certificate Hierarchy. For more information see, [Selecting the certificate to add](4589-selecting-the-certificate-to-add.md "The certificate authority (CA) is the authority that validates the server. The certificate to add to the CA list is the authority certificate, not the server certificate.") and
      [Missing certificates](4577-missing-certificates.md "Identifying missing certificates.").
   2. Create the Certificate Authority List by running the following command for each of the
      certificates that you saved to disk. 

      ```
      $ openssl x509 -in ServerCertificate.crt -text >> ClientCAList.pem
      ```

      All trusted certificate authorities are listed. These are checked following a chain of child to
      parent certificates until a certificate is reached which is trusted. All other certificates that
      were trusted by the Root Certificate Authority will also be considered as trusted by the client. For
      more information see [Certificate authorities](4566-encryption-and-authentication.md).
2. Set the entry for the global certificate authority list in your FGLPROFILE file.

   The global certificate authority list entry defines the file containing the certificate authority
   list used by the client's Genero Web Services to validate all certificates coming from the different
   servers it will connect
   to.

   ```
   security.global.ca = "ClientCAList.pem"
   ```

   The file is located based on the current execution directory. If you use Genero Studio, for
   instance, fglrun may not be executed in the same directory as when you use the
   command line. The recommended practice therefore is to specify an absolute path for the
   .pem, for
   example:

   ```
   security.global.ca = "/opt/usr/certs/ClientCAList.pem"
   ```

   If `security.global.ca` is not defined, Genero Web Services will look to see
   whether the operating system has a keystore; otherwise,
   `security.global.ca.lookuppath` will be used. For further information, see [Certificate
   authorities](4569-https-configuration.md) in [HTTPS configuration](4569-https-configuration.md "If no client certificate is provided, Genero Web Services (GWS) does the HTTPS request transparently.").

The client application is configured to use the appropriate certificate authority list to
validate a server's certificate.

**What to do next**

In your FGLPROFILE file ensure there are
configuration entries (`ws.*` ) for the HTTPS server URL and for the HTTP
authentication when accessing the HTTPS server. See [Set FGLPROFILE entries for the server URL](4573-set-fglprofile-entries-for-the-server-url.md "Add a set of configuration entries that specify the URL and the identity for the HTTPS server.").

## Related links

**Related concepts**  

[Error: Peer certificate is issued by a company not in our CA list](4576-error-peer-certificate-is-issued-by-a-company-not-in-our-ca.md "When a client connects to a server using HTTPS, the client needs to trust the server it is in communication with. So the client needs to add the server's CAs (certificate authorities lists) to its trusted CAs.")

[Enable OCSP](4574-enable-ocsp.md "To enable Online Certificate Status Protocol (OCSP), set the security.global.ocsp.enable and security.global.ocsp.url entries in FGLPROFILE.")

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")

[FGLPROFILE: HTTP(S) Proxy Authentication](4920-fglprofile-http-s-proxy-authentication.md "FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.")

[Accessing secured services](4568-accessing-secured-services.md "Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.")
