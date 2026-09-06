---
title: "Migrate an I4GL web service provider to Genero"
source: "fgl-topics/c_gws_i4gl_migration_guide_002.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero"
type: "concept"
---

# Migrate an I4GL web service provider to Genero

> Migrate a Web service from I4GL to Genero BDL using the SOAP protocol.

This section explains how to migrate an I4GL Web service provider to a Genero application
providing the same Web service; allowing access to all clients already accessing that service -
unmodified (excepted for the host name of course).

The migration is based on the SOA zipcode demo in the I4GL package.

## Related links

1. [Step 1: Use the I4GL function and the I4GL.4cf configuration file](4687-step-1-use-the-i4gl-function-and-the-i4gl-4cf-configuration.md)

   The I4GL .4cf configuration file has all the information you need about the I4GL Web service.
2. [Step 2: Create a BDL RECORD for the input parameters](4688-step-2-create-a-bdl-record-for-the-input-parameters.md)

   Define a BDL record for the input message of the Web function.
3. [Step 3: Create a BDL RECORD for the output parameters](4689-step-3-create-a-bdl-record-for-the-output-parameters.md)

   Define a BDL record for the output message of the Web function.
4. [Step 4: Create a BDL wrapper function](4690-step-4-create-a-bdl-wrapper-function.md)

   Create the wrapper function that uses BDL records to call the I4GL function.
5. [Step 5: Publish the wrapper function as a Genero web service](4691-step-5-publish-the-wrapper-function-as-a-genero-web-service.md)

   Based on the details in the configuration file (.4cf) file, create a function that registers your Web service with the Genero Web Service server.
6. [Step 6: Create the server](4692-step-6-create-the-server.md)

   Provide a file with a BDL function that starts your Web service with the Genero Web Service server instead of Axis.
7. [Step 7: Configure the database](4693-step-7-configure-the-database.md)

   Based on the DATABASE entry in the I4GL .4cf configuration file, use the Genero instruction to connect to the Informix® database at server startup.
8. [Step 8: Compile and run the Genero service](4694-step-8-compile-and-run-the-genero-service.md)

   Describes how to compile and run the service to test it.
