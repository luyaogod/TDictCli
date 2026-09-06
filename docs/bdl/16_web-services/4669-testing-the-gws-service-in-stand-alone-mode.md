---
title: "Testing the GWS service in stand-alone mode"
source: "fgl-topics/c_gws_server_tutorial_015.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Testing the GWS service in stand-alone mode"
type: "concept"
---

# Testing the GWS service in stand-alone mode

> Test that your service is reachable and that it can generate the WSDL.

For testing and development purposes only, the Genero Web Services Server application can be
executed directly, without using the Genero Application Server (GAS).

1. Use the Genero fglrun command to execute the GWS Server application. The
   application must reside on the same machine:

   ```
   fglrun <gws-application>
   ```

   This will start the GWS Server on the port specified by the FGLAPPSERVER environment variable.
   For example, if FGLAPPSERVER is set to 8090, the server will be started on that port. If this
   environment variable is not set, port number 80 is used.

   > **Warning:**
   >
   > Do not set the FGLAPPSERVER variable in production environments, since the
   > port number is selected by the Genero Application Server.
2. Obtain the WSDL information for your Service and write a test Client application. For example,
   if the GWS Server in step 1 was started on your local machine and FGLAPPSERVER was set to 8090, the
   command to get the WSDL information would be:

   ```
   fglwsdl -o <test-client> http://localhost:8090/<service-name>?WSDL
   ```
