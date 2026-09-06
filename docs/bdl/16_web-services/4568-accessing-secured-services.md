---
title: "Accessing secured services"
source: "fgl-topics/c_gws_access_001.html"
breadcrumb: "Web services > Security > Accessing secured services"
type: "concept"
---

# Accessing secured services

> Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.

![Image shows the various communications options for a client to connect to a Web Service](../_images/communications.jpg)

*Communications options for a client to connect to a Web Service*

**HTTP**

Client connects to a Web Server (or a Web Service) using HTTP as the communication protocol.
(**No security** , **No authentication**).

**HTTP with Basic Authentication**

Client connects to a Web Server using HTTP as the communication protocol, but a valid login and
password are required from the Web server to grant access to the Web Service. (**No security** ,
**Weak Authentication)**. The login and password are sent in clear text on the communication
layer.

**HTTP with Digest Authentication**

Client connects to the Web Server using HTTP as the communication protocol, but a valid login
and password are required from the Web server to grant access to the Web Service. (**No
security** , **Authentication**). The login and password are encoded using a digest algorithm,
requiring additional information from the Web server. This means that the first connection will
always fail, but it is necessary in order to return Web server additional information back to the
client.

**HTTPS**

Client connects to a Web server using HTTPS as the communication protocol. (**Security** ,
**No authentication**). The communication channel is encrypted by SSL/TLS.

**HTTPS with Basic Authentication**

Client connects to a Web server using HTTPS as the communication protocol, but a valid login and
password are required from the Web server to grant access to the Web Service. (**Security** ,
**Weak Authentication**). The login and password are sent in clear text on the communication
layer, but the communication channel is encrypted by SSL/TLS.

**HTTPS with Digest Authentication**

Client connects to the Web Server using HTTPS as the communication protocol, but a valid login
and password are required from the Web server to grant access to the Web Service. (**Security** ,
**Authentication**). The login and password are encoded using a digest algorithm, requiring
additional information from the Web server. This means that the first connection will always fail,
but it is necessary in order to return Web server additional information back to the client. The
communication channel is encrypted by SSL/TLS.

## Accessing secured services via a proxy

To improve communication speed with the cache mechanism, or to restrict internet access to
specific clients, Genero Web Services allows a client to connect via proxies. The proxy is in charge
of dispatching the client request to the server, and uses the same protocol as that used by the
server. So, when a client connects via a proxy to access a HTTP server, the configuration of the
HTTP proxy is used, and when the client communicates in HTTPS, the HTTPS proxy configuration is used.

**HTTP proxy**

Client connects via a proxy using HTTP as the communication protocol.

**HTTP proxy with Basic Authentication**

Client connects via a proxy using HTTP as the communication protocol, but a valid login and
password are required from the proxy to dispatch the request to the Web Service. The login and
password are sent in clear text on the communication layer between client and proxy.

**HTTP proxy with Digest Authentication**

Client connects via a proxy using HTTP as the communication protocol, but a valid login and
password are required from the proxy to dispatch the request to the Web Service. The login and
password are encoded using a digest algorithm, requiring additional information from the proxy. This
means that the first connection will always fail, but it is necessary in order to return proxy
additional information back to the client.

**HTTPS proxy**

Client connects via a proxy using HTTPS as the communication protocol. The communication channel
is encrypted by SSL/TLS.

**HTTPS proxy with Basic Authentication**

Client connects via a proxy using HTTPS as the communication protocol, but a valid login and
password are required from the proxy to dispatch the request to the Web Service. The login and
password are sent in clear text on the communication layer between client and proxy, but the
communication channel is encrypted by SSL/TLS.

**HTTPS proxy with Digest Authentication**

Client connects via a proxy using HTTPS as the communication protocol, but a valid login and
password are required from the proxy to dispatch the request to the Web Service. The login and
password are encoded using a digest algorithm, requiring additional information from the proxy. This
means that the first connection will always fail, but it is necessary in order to return proxy
additional information back to the client. The communication channel between client and proxy is
encrypted by SSL/TLS.

## Related links

**Related reference**  

[Basic or digest HTTP authentication](4916-fglprofile-entries-for-web-services.md "Basic or digest HTTP authentication")

[Proxy configuration](4916-fglprofile-entries-for-web-services.md "Proxy configuration")
