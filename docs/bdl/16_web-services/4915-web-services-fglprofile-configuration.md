---
title: "Web services FGLPROFILE configuration"
source: "fgl-topics/c_gws_ssl_configuration_001.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration"
type: "concept"
---

# Web services FGLPROFILE configuration

> The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.

When using BDL web services on server side, it is the web server that is in charge of the BDL web
services server security, not the BDL server application itself. You must refer to your web server
manual to secure the server part of the web services.

This is useful for deployment purposes, as no additional code modification is necessary, even
if the location of the different servers change, or if different cryptography keys or X509
certificates are necessary for the same application but intended for several customers with their
own needs.

The Genero web services secured communication and the support of XML-Security is based on the
[OpenSSL](http://www.openssl.org) engine. It allows
a BDL web services client, or a BDL application using the `com` or `xml` API, to communicate with any secured server
over HTTP or HTTPS, and to handle encrypted and/or signed XML document in BDL coming from any other
application.

## Child topics

- [FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md): The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.
- [Examples](4917-examples.md): SSL/TLS configuration examples.
