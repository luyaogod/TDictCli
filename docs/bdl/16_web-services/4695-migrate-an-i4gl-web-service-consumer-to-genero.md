---
title: "Migrate an I4GL web service consumer to Genero"
source: "fgl-topics/c_gws_i4gl_migration_guide_012.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero"
type: "concept"
---

# Migrate an I4GL web service consumer to Genero

> Migrate a client application from I4GL to a Genero BDL application using the SOAP protocol.

This section explains how to migrate an I4GL Web service consumer to a Genero application
accessing the same Web service.

The migration is based on the SOA demo in the I4GL package.

## Child topics

- [Step 1: Generate the Genero web service stub from an I4GL WSDL](4696-step-1-generate-the-genero-web-service-stub-from-an-i4gl-wsd.md): Use the fglwsdl tool to get the WSDL information from the service provider.
- [Step 2: Modify the Genero .inc stubs to fix wrong I4GL WSDL](4697-step-2-modify-the-genero-inc-stubs-to-fix-wrong-i4gl-wsdl.md): Remove the namespace attributes in the .inc file stub.
- [Step 3: Include the generated stub in your I4GL application](4698-step-3-include-the-generated-stub-in-your-i4gl-application.md): Use a GLOBALS statement to specify the Web service your application uses.
- [Step 4: Modify the I4GL web service function call](4699-step-4-modify-the-i4gl-web-service-function-call.md): Rename I4GL function name names to Genero Web service function names.
- [Step 5: Handle Genero Web Services errors](4700-step-5-handle-genero-web-services-errors.md): Describes how to code to check the error status returned by the I4GL Web service, and get details of the error.
- [Step 6: Compile and run the Genero client](4701-step-6-compile-and-run-the-genero-client.md): Describes how to compile and run the client to test it.
- [Standalone Axis server is buggy](4702-standalone-axis-server-is-buggy.md): Describes a bug you can expect using the I4GL standalone axis server.
