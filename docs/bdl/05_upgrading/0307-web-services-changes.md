---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_200_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of Web services in Genero 2.00.

## Using the 1.3x Web Services API

There is no need to create a special runner for Genero Web Services 2.x. Instead, the GWS 2.x
library is imported into your applications. If you want to migrate your existing 1.x GWS Server
application to 2.x, to avoid the need for a special runner, as well as to take advantage of any bug
fixes, take the following steps:

1. Add the following statement at the top of any .4gl module where before you
   used GWS 1.3x functions:

   ```
   import com
   ```
2. Compile and re-link your GWS Server application (.42r).

This imports the new [GWS **com**
library](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services."), and ensures that any GWS 1.3x functions that you have used will be compatible. Your
existing Genero 1.3x Client applications, as well as third-party Client applications, will continue
to work.

## Using the new 2.00 Web Services API

If you want to take advantage of the new features and simplify future migrations, you can migrate
your Genero Web Services (GWS) Server runner and also use the new GWS 2.x APIs. All the 1.3x
publishing functions for all the operations in your application must be replaced with 2.x publishing
functions. Since this does not change the interface, all existing Genero 1.3x client applications,
as well as third-party client applications, will continue to work.

As 1.3x only supports RPC-Encoded style services, you must use the RPC style functions of the new
2.x APIs as the replacement functions, with `setInputEncoded` and`setOutputEncoded` set to true. And, you cannot add XML attributes to the records used as Web
Service function parameters.

To replace the `fgl_ws_server_publishfunction()` statement in an existing GWS
Server application; for
example:

```
CALL fgl_ws_server_publishfunction(
  "EchoInteger",
  "http://tempuri.org/webservices/types/in", "echoInteger_in",
	"http://tempuri.org/webservices/types/out", "echoInteger_out",
	"echoInteger")
```

1. Add this statement at the top of each module:

   ```
   import com
   ```
2. Define variables for the WebService and WebOperation
   objects:

   ```
   DEFINE serv  com.WebService
   DEFINE op    com.WebOperation  -- Operation of a WebService
   ```
3. Create the GWS Server
   object:

   ```
   LET serv = com.WebService.CreateWebService(
     "EchoInteger",
     "http://tempuri.org/webservices")
   ```
4. Use the 2.x publishing functions for each
   operation:

   ```
   LET op = com.WebOperation.CreateRPCStyle(
     "echoInteger",
     "EchoInteger",
     echoInteger_in,
     echoInteger_out)
   CALL op.setInputEncoded(true)
   CALL op.setOutputEncoded(true)
   CALL serv.publishOperation(op,NULL)
   ```
5. Compile and re-link your GWS Server application (.42r)

GWS 2.x also allows your Server application (.42r) to contain multiple
services. If you would like 2.x and 1.3x GWS to coexist in the same .42r
executable, replace the existing publishing 1.3x functions.

## Enhance the GWS server application to be WS-I compliant (recommended)

> **Important:**
>
> You must be able to change all the client applications that access your
> migrated Genero Web Services (GWS) Server.

If you use the Literal styles now available in GWS 2.x for your Web service, your application
will be WS-I compliant. However, the migration techniques still use the RPC/Encoded style (Only
RPC/Encoded was supported in GWS 1.3x.). If you can change all the client applications that access
your migrated GWS Server, we recommend that you enhance the GWS Server application to be WS-I
compliant.

1. Replace
   the publishing functions in the GWS Server application but omit the
   `setInputEncoded` and `setOutputEncoded` lines. The resulting style
   will be Literal.
2. The enhanced GWS Server will have a new RPC/Literal WSDL that must be used to regenerate the
   client stub with the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).")
   tool:

   ```
   fglwsdl -o NewClientstub http://localhost:8090/MyCalculator?WSDL
   ```
3. Compile that new client stub, and re-link it with the GWS client application. This operation
   must be repeated for each client application accessing that service.
4. Third party client applications must also be changed to use the new WSDL.
