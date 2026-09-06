---
title: "Step 1: Create the server's certificate authority list"
source: "fgl-topics/c_gws_ssl_deployment_server_004.html"
breadcrumb: "Web services > Deploy a Web Service > Configuring the Apache web server for HTTPS > Step 1: Create the server's certificate authority list"
type: "concept"
---

# Step 1: Create the server's certificate authority list

> Create the certificate authority list and add the root certificate authority certificate.

- Create the server's Certificate Authority
  List:

  ```
  $ openssl x509 -in MyRootCA.crt -text >> ServerCAList.pem
  ```

As the server trusts only the Root Certificate Authority, the list contains only that one
certificate authority; all other certificates that were trusted by the Root Certificate Authority
will also be considered as trusted by the server.

In the next step we register the server's certificate authority list, [Step 2: Register the server as a Web service in the GAS](4911-step-2-register-the-server-as-a-web-service-in-the-gas.md "Web services registered on the GAS are started automatically when the GAS starts.").

## Related links

**Related concepts**  

[The OpenSSL tool](4579-the-openssl-tool.md "The openssl command line tool creates certificates for the configuration of secured communications.")

[Certificates in practice](4578-certificates-in-practice.md "Procedures and tools for creating, importing, and viewing certificates and keys.")
