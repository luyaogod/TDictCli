---
title: "com.WebServiceEngine.RegisterRestResources"
source: "fgl-topics/c_gws_ComWebServiceEngine_RegisterRestResources.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.RegisterRestResources"
type: "concept"
---

# com.WebServiceEngine.RegisterRestResources

> Registers the resources of a REST service in the engine.

## Syntax

```
com.WebServiceEngine.RegisterRestResources(
   basePath STRING,
   info RECORD,
   scope STRING,
   { module | package } STRING [,...]
)
```

1. basePath defines the path to the REST web service as it appears in the URL.
   It may also be used as the name of the Web service if there is no `WSInfo` information.
2. info defines the record used to generate the [service information](../16_web-services/4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). If the record
   is `NULL`, or information is not specified, the service name will be used by default.
3. scope defines the global scope (can be `NULL`)
4. Provides the option to specify one or more resources either in modules or packages.The method
   takes a variable number of arguments. Several module names can be specified in a comma-separated
   list with each module (42m) in quotes: `"module1",
   "module2","module3"` or if the modules are in packages:
   `"package.module1","package/subpackage/module2","package.subpackage.module3"`.

   If service modules are in a package, the package name must have a
   dot-separated (`"package.module1"`) or slash-separated
   (`"package/module1"`) list corresponding to the path from a top-directory identifying
   the root of the package tree down to where service modules (42m) are located.
   For more information on working with packages, refer to the [PACKAGE](../09_advanced-features/0816-package.md "Defines the package the module belongs to.") page.

## Usage

Use this method to register resources in modules for your web service with the Genero Web Service
(GWS) engine and publish the service.
> **Note:**
>
> If your service has only one module, use the [com.WebServiceEngine.RegisterRestService](3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine.") method instead.

For instance: (Line breaks have
been added to improve readability.)

- If you have a web service called "fruits" that contains several resources (one per module);
  where one module is called "apple", and another "peach", import the modules with `IMPORT
  FGL` statements in your web service application and then call the function to register the
  web service as shown:

  ```
  IMPORT FGL apple
  IMPORT FGL peach
  #...
   CALL com.WebServiceEngine.RegisterRestResources("fruits", serviceInfo, "scope", 
                                                      "apple", "peach")
  ```
- If the "fruits" web service that contains several resources (one per module) is inside a FGL
  `PACKAGE`; where one module is called "apple", and another "peach", you can import
  these modules from the package with an `IMPORT FGL` statement in your web service
  application, using either dot notation or slashes in the path and a wildcard (`.*`)
  as shown. Then a call to the function as shown registers the web
  service:

  ```
  IMPORT FGL myWSPackage.season.*
  #...
  # using dot notation
    CALL com.WebServiceEngine.RegisterRestResources("fruits", serviceInfo, "scope", 
                                                     "myWSPackage.season.apple",
                                                     "myWSPackage.season.peach")
  # or using slashes
    CALL com.WebServiceEngine.RegisterRestResources("fruits", serviceInfo, "scope", 
                                                     "myWSPackage/season/apple",
                                                     "myWSPackage/season/peach")
  ```

  > **Important:**
  >
  > **Aliasing not supported**
  >
  > Using an alias on the
  > package name **is not** supported. For example, this will
  > fail:
  >
  > ```
  > IMPORT FGL myWSPackage.season.peach AS peach
  >   #...
  >   CALL com.WebServiceEngine.RegisterRestResources("fruits", serviceInfo, "scope", 
  >                                                    "peach") -- WILL NOT WORK
  > ```

You typically deploy the service behind a Genero Application Server (GAS). In the context of the
GAS, you therefore may have the group-name (if not a member of the default group)
and the service configuration file (xcf-file ) name in the URL before the service
name. For example, the URL for the service's OpenAPI specification is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/service-name?openapi.json
```

For example, a URI to access the "apple" resource "orders" operation is at:
http://myhost/gas/ws/r/myGroup/myXcf/fruits/apple/orders

Where:

- myhost/gas/ws/r is the base URL from the server.
- myGroup is the name of the group on the GAS where the Web service is
  deployed.
- myXcf is the name of the Web service configuration
  .xcf file.
- fruits is the web service name
- apple is the resource module name

  The resource name will always appear in
  the URI, even if there is only one module.
- orders is the resource-path (defined with [WSPath](../16_web-services/4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")) for the "apple"
  resource.

For more information on resource URI, see [REST resource URI naming practice](../16_web-services/4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Publish REST service with multiple resources](../16_web-services/4757-rest-service-with-several-modules.md "Publish a service with several resources.")
