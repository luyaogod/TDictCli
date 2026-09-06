---
title: "FGLAPPSERVER"
source: "fgl-topics/c_fgl_EnvVariables_FGLAPPSERVER.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLAPPSERVER"
type: "concept"
---

# FGLAPPSERVER

> Defines the listening TCP port of the Web service in development context.

The FGLAPPSERVER environment variable defines the TCP port on which the web service
server will be started.

If the FGLAPPSERVER environment variable is not set, the default TCP port is 80.

During development, define this environment variable before starting the web service
server program, to let web service clients connect directly to the runtime system. You typically
defined FGLAPPSERVER to the port 8090.

In production, Genero Application Server (GAS) is used to deploy web services servers.
The GAS will automatically set FGLAPPSERVER. Do not manually set FGLAPPSERVER when GAS
is involved.

## Related links

**Related concepts**  

[The WebServiceEngine class](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.")

[Step 6: Create the server](../16_web-services/4692-step-6-create-the-server.md "Provide a file with a BDL function that starts your Web service with the Genero Web Service server instead of Axis.")

[fgl\_ws\_server\_start() (version 1.3)](../16_web-services/5068-fgl-ws-server-start-version-1-3.md "Creates and starts the Web services server.")

[Testing the GWS service in stand-alone mode](../16_web-services/4669-testing-the-gws-service-in-stand-alone-mode.md "Test that your service is reachable and that it can generate the WSDL.")

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
