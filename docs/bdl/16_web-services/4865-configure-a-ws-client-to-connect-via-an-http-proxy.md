---
title: "Configure a WS client to connect via an HTTP Proxy"
source: "fgl-topics/t_gws_ssl_proxy_tutorial_rest.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Configure a WS client to connect via an HTTP Proxy"
type: "task"
---

# Configure a WS client to connect via an HTTP Proxy

> Configuration steps to connect via a HTTP proxy.

1. Add the location of the proxy to your FGLPROFILE file with the
   `proxy.http.location` entry.

   Add the entry `proxy.http.location` to your fglprofile.
   For the value, provide the IP address of the HTTP proxy and the port number
   where the HTTP proxy is listening, separated by a colon. For example, to have a
   client connect via a HTTP proxy located at the IP address "10.0.0.170" and
   listening on port number "8080", add this entry to your
   fglprofile:

   `proxy.http.location =
   "10.0.0.170:8080"`

   To configure the client to connect via
   an HTTPS proxy, replace `http` with
   `https`.
2. Define the list of host names the client will **not** have to connect to via a proxy with
   the `proxy.http.list` entry.

   Add the entry `proxy.http.list` to your FGLPROFILE file. For the value, provide
   a semi-colon separated list of clients. For example, to exclude all hosts beginning with
   "www.mycompany.com" or "www.google." from connecting via a HTTP proxy, add this entry to your
   fglprofile:

   `proxy.http.list =
   "www.mycompany.com;www.google."`

## Related links

**Related reference**  

[Proxy configuration](4916-fglprofile-entries-for-web-services.md "Proxy configuration")
