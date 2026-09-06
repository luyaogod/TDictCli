---
title: "Set FGLPROFILE entries for the client certificate"
source: "fgl-topics/t_gws_ssl_config_client_cert_soap.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Configure a WS client to access an HTTPS server > Set FGLPROFILE entries for the client certificate"
type: "task"
---

# Set FGLPROFILE entries for the client certificate

> Configure your application to use the certificate and the associated private key used by the client's Genero Web Services during HTTPS communication. For production systems, you add the configuration details to your FGLPROFILE file.

**Before you begin:**

You have generated a certificate. For
more details on how to create a self-signed certificate, go to the [Certificates in practice](4578-certificates-in-practice.md "Procedures and tools for creating, importing, and viewing certificates and keys.") pages.

In this task you add the certificate information to your FGLPROFILE file.

Set the configuration entries for your security certificates in your FGLPROFILE file.

The security entry must be defined with an unique identifier (**id1** in this
example).

`security.id1.certificate = "MyServer.crt"`  
`security.id1.privatekey  = "MyServer.pem"`

If the private key is protected with a password, the password must be removed using the OpenSSL
command. For example,

```
$ openssl rsa -in privkey.pem -out MyServer.pem
```

You
are prompted for the passphrase. The unprotected private key is output in
MyServer.pem.

> **Tip:**
>
> If you use the [setCertificateAndKey()](../15_library-reference/3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.") method, you can
> override this configuration dynamically at runtime. While the [clearCertificateAndKey](../15_library-reference/3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") resets the
> value to the one configured in FGLPROFILE.

Your applications are configured to use the client certificate.

**What to do next**

Create the certificate authority list
from the server and configure the global certificate authority list
(`security.global.ca`) in your FGLPROFILE file. See [Set FGLPROFILE entries for the CA list](4572-set-fglprofile-entries-for-the-ca-list.md "Clients need to check to see if the server's certificate is trusted. This is done using a certificate authority list.").

## Related links

**Related concepts**  

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")

[FGLPROFILE: HTTP(S) Proxy Authentication](4920-fglprofile-http-s-proxy-authentication.md "FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.")

[Encryption and authentication](4566-encryption-and-authentication.md "A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.")

[The OpenSSL tool](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.")

**Related tasks**  

[View a certificate](4585-view-a-certificate.md "View the details of a certificate using the openssl command.")
