---
title: "com.WebServiceEngine.RegisterRestService"
source: "fgl-topics/c_gws_ComWebServiceEngine_RegisterRestService.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.RegisterRestService"
type: "concept"
---

# com.WebServiceEngine.RegisterRestService

> Registers a REST service in the engine.

## Syntax

```
com.WebServiceEngine.RegisterRestService(
   { module | package } STRING, 
   basePath STRING )
```

1. Provides the option to specify the module or package name of the service or services.

   If service modules are in a package, the package name must have a
   dot-separated (`"package.module1"`) or slash-separated
   (`"package/module1"`) list corresponding to the path from a top-directory identifying
   the root of the package tree down to where service modules (42m) are located.
   For more information on working with packages, refer to the [PACKAGE](../09_advanced-features/0816-package.md "Defines the package the module belongs to.") page.
2. basePath defines the path to the web service as it appears in the URL.

## Usage

Use this method to register all the public functions in your web service with the Genero Web
Service engine and publish them for users on the internet.
> **Note:**
>
> If your service has more than one module, use the [com.WebServiceEngine.RegisterRestResources](3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.") method instead.

The basePath sets the path to the web service in the Genero Web Service engine
as it appears in the URL for the service.

For instance:

- If you have a web service and you want your Genero BDL module called "accounts" to be published
  as a REST service with the name "Accounts" (uppercase "A"), import the accounts module with an
  `IMPORT FGL` statement and then call the function as shown to register the web
  service:

  ```
  IMPORT FGL accounts
    #...
     CALL com.WebServiceEngine.RegisterRestService("accounts","Accounts")
  ```
- If your accounts web service is inside a FGL `PACKAGE`, you can import the module
  from the package with an `IMPORT FGL` statement in your web service application,
  using either dot notation or slashes and a wildcard (`.*`) as shown. Then a call to
  the function as shown, either using dot notation or slashes (`"/"`) in the path to
  the accounts module, registers the web service:

  ```
  IMPORT FGL myWSPackage.*
    #...
    # using dot notation
      CALL com.WebServiceEngine.RegisterRestService("myWSPackage.v1.accounts", "Accounts")

    # or using slashes
      CALL com.WebServiceEngine.RegisterRestService("myWSPackage/v1/accounts", "Accounts")
  ```

  > **Important:**
  >
  > **Aliasing not supported**
  >
  > Using an alias on
  > the package name **is not** supported. For example, this will
  > fail:
  >
  > ```
  > IMPORT FGL myWSPackage.v1.accounts AS Account
  >   #...
  >   CALL com.WebServiceEngine.RegisterRestService("Account",
  >               "Accounts") -- WILL NOT WORK
  > ```

You typically deploy the service behind a Genero Application Server (GAS). In the context of the
GAS, you therefore have the group-name name (if not a member of the default
group) and the service xcf-file filename in the URL. For example, the URL for
the service's OpenAPI specification is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/service-name?openapi.json
```

For instance, a URI to access the "users" operation is at:

http://myhost/gas/ws/r/myGroup/myXcf/accounts/users

- myhost/gas/ws/r is the base URL from the server.
- myGroup is the name of the group on the GAS where the web service is
  deployed.
- myXcf is the name of the web service configuration
  .xcf file.
- accounts is the name of the web service.
- users is a resource-path of a specific resource
  operation set with [WSPath](../16_web-services/4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.").

For more information on resource URI, see [REST resource URI naming practice](../16_web-services/4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
