---
title: "Authenticate the WS client to a proxy"
source: "fgl-topics/t_gws_ssl_auth_tutorial_client2proxy.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Authenticate the WS client to a proxy"
type: "task"
---

# Authenticate the WS client to a proxy

> Configuration steps to authenticate the client to a proxy (proxy authentication).

For an example, see [FGLPROFILE: HTTP(S) Proxy Authentication](4920-fglprofile-http-s-proxy-authentication.md "FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.").

1. Add an HTTP authenticate entry to your FGLPROFILE file.

   To connect via a proxy with HTTP Proxy Authentication, it is necessary to define the client
   login and password as registered on the HTTP proxy.

   The following two entries must be defined with
   an unique identifier (**proxyauth** for our example) to define a HTTP Proxy Authentication with
   **myapplication** as login and **mypassword** as
   password:

   `authenticate.proxyauth.login     = "myapplication"`  
   `authenticate.proxyauth.password  = "mypassword"`

   See [[RFC2617]](http://www.ietf.org/rfc/rfc2617.txt) for more details.
2. For proxy authentication, an entry must be made to the HTTP proxy configuration in order to
   authenticate a client.

   To authenticate a client known as **myapplication** with **mypassword** as password by the
   HTTP Proxy, add the following entry to the HTTP proxy
   configuration:

   `proxy.http.authenticate = "proxyauth"`

   To authenticate the client to a HTTPS proxy, replace
   http with https.

## Related links

**Related reference**  

[Proxy configuration](4916-fglprofile-entries-for-web-services.md "Proxy configuration")
