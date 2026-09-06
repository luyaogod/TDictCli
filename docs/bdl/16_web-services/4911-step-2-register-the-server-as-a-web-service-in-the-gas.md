---
title: "Step 2: Register the server as a Web service in the GAS"
source: "fgl-topics/c_gws_ssl_deployment_server_005.html"
breadcrumb: "Web services > Deploy a Web Service > Configuring the Apache web server for HTTPS > Step 2: Register the server as a Web service in the GAS"
type: "concept"
---

# Step 2: Register the server as a Web service in the GAS

> Web services registered on the GAS are started automatically when the GAS starts.

As the Web server is in charge of the complete HTTPS protocol with all the clients, there is no
additional GAS configuration needed to add security. Simply register the BDL server to the list of
Web Services of the GAS. For more information, refer to the Genero Application Server User Guide.

For more details, see [Web services server program deployment](4908-web-services-server-program-deployment.md "The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.").

In the next step we configure Apache for HTTPS, [Step 3: Configure Apache for HTTPS](4912-step-3-configure-apache-for-https.md "Add the locations of your certificates to the Apache configuration file.").
