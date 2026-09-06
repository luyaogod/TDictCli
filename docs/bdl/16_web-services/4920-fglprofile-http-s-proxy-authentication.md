---
title: "FGLPROFILE: HTTP(S) Proxy Authentication"
source: "fgl-topics/c_gws_ssl_configuration_010.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > Examples > FGLPROFILE: HTTP(S) Proxy Authentication"
type: "concept"
---

# FGLPROFILE: HTTP(S) Proxy Authentication

> FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.

```
# Security configuration
security.global.script   =  "Cert/password.sh"
security.global.ca       =  "Cert/CAList.pem"
security.global.cipher   =  "HIGH" # Use only HIGH encryption ciphers
security.mykey.certificate =  "Cert/MyCertificateA.crt"
security.mykey.privatekey  =  "Cert/MyPrivateKeyA.pem"

# Proxy HTTP Authentication
authenticate.proxyauth.login    =  "myapplication"
authenticate.proxyauth.password =  "mypswd"
authenticate.proxyauth.scheme   =  "Basic"

# HTTPS Proxy configuration
proxy.https.location     =  "10.0.0.170"
proxy.https.list         =  "www.mycompany.com;www.mycompany.com"
proxy.https.authenticate =  "proxyauth"

# Server HTTP Authentication
authenticate.serverauth.login    =  "mylogin"
authenticate.serverauth.password =  "password"

# Server configuration
ws.myserver.url =  "https://www.MyServer.com/gas/ws/r/MyWebService"
ws.myserver.authenticate =  "serverauth"
ws.myserver.security     =  "mykey"
```

## Related links

**Related concepts**  

[Encryption, BASE64 and password agent with fglpass tool](4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent.")

[Accessing secured services](4568-accessing-secured-services.md "Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.")

**Related reference**  

[Basic or digest HTTP authentication](4916-fglprofile-entries-for-web-services.md "Basic or digest HTTP authentication")

[Proxy configuration](4916-fglprofile-entries-for-web-services.md "Proxy configuration")
