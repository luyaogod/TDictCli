---
title: "Set FGLPROFILE entries for the server URL"
source: "fgl-topics/t_gws_ssl_config_cert_add_to_fglprofile.html"
breadcrumb: "Web services > Security > Configure a WS client to access an HTTPS server > Set FGLPROFILE entries for the server URL"
type: "task"
---

# Set FGLPROFILE entries for the server URL

> Add a set of configuration entries that specify the URL and the identity for the HTTPS server.

In this task you add configuration entries (`ws.*` ) in your FGLPROFILE file for
the HTTPS server URL and for HTTP authentication when accessing the HTTPS server. For an example,
see [FGLPROFILE: HTTP(S) Proxy Authentication](4920-fglprofile-http-s-proxy-authentication.md "FGLPROFILE entries can be used to define a connection to an HTTPS server via a proxy, and with HTTP and Proxy Authentication.").

Add configuration entries for the server. 

The following entries must be defined with an unique identifier (such as "myserver") :

1. `ws.myserver.url =
   "https://www.MyServer.com/gas/ws/r/MyWebService"`

   In a production environment, the Genero Application Server (GAS) provides the base URL from the
   server where the web service is deployed. This follows the
   format:

   ```
   http[s]://host:port/gas/ws/r/service-name
   ```

   1. host:port/gas/ws/r is
      the base URL from the server. For more information, refer to the Application Web
      Address page in the Genero Application Server User Guide.
   2. service-name is the name of the Web service.
2. `ws.myserver.security = "id1"`

   Where the value (`id1` in this example) must match the unique identifier
   defined by the client security entry created in [Set FGLPROFILE entries for the client certificate](4571-set-fglprofile-entries-for-the-client-certificate.md "Configure your application to use the certificate and the associated private key used by the client's Genero Web Services during HTTPS communication. For production systems, you add the configuration details to your FGLPROFILE file.").

> **Tip:**
>
> The unique identifier "myserver" can be used in the BDL client code in place of the
> actual URL.

## Related links

**Related concepts**  

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")

[Accessing secured services](4568-accessing-secured-services.md "Security and authentication are important. Genero Web Services provides various communications options for a client to connect to a Web service.")

[Troubleshoot certificate issues](4575-troubleshoot-certificate-issues.md "You may encounter known (and common) issues when completing the Genero Web Services tutorials or when adding Web services of your own. These issues and their solutions are presented in the following topics.")
