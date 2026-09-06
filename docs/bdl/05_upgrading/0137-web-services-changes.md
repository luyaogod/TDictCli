---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_400_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 4.00.

## Security Note: sameSite HTTP cookie attribute

When using HTTP cookies, make sure to check the `sameSite` attribute usage.

For more details, see [Changes to how GWS handles cookies](0158-web-services-changes.md) and [Single sign-on (OpenID Connect, SAML, and GIP) sameSite security](0158-web-services-changes.md).

## Validation checks on high-level REST

Starting with version 4.00, the GWS performs the following validation checks on the high-level
REST function and the compiler raises errors if validation fails.

- If you use `WSPost`, `WSPut`, or `WSPatch` as REST
  resource operation, the FUNCTION must have at least one REST input body. Not specifying this will
  cause the [error-9106](../15_library-reference/4483-genero-bdl-errors.md). However,
  you can set `WSOptional` if
  a message body is not always required.
- If you use `WSGet`, `WSHead`, `WSDelete`,
  `WSOptions`, or `WSTrace` as the REST resource operation, the FUNCTION
  must **not** have a request (input) body. Specifying these requests with a message body will
  cause the [error-9106](../15_library-reference/4483-genero-bdl-errors.md) or [error-9128](../15_library-reference/4483-genero-bdl-errors.md).
- If you use `WSHead`, `WSOptions`, `WSTrace`, or
  `WSPatch` as the REST resource operation, the FUNCTION must **not** have a REST
  response (output) body. Specifying these responses with a message body will cause the [error-9129](../15_library-reference/4483-genero-bdl-errors.md).
- [WSPath](../16_web-services/4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") cannot have the same verb more
  than once for the same path, otherwise the [error-9127](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

For more information on use of HTTP verbs and their corresponding RESTful high-level API
attributes, see [HTTP verbs and attributes](../16_web-services/4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.").

## Changes to the high-level REST Web services URL format

Starting with version 4.00, there are changes to RESTful Web service URL using the high-level
API. It is a change to the format of the URL you use to access services; it does not involve any
application configuration changes. The new format enhances the way Genero Web Services (GWS) finds
resources.

The previous URL
was:

```
http[s]://host:port/gas/ws/r/resource-name/endpoint
```

The new URL
is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/resource-name/endpoint
```

The new REST URL includes two new path elements for when Web services are deployed on the Genero
Application Server (GAS):

**group-name (optional)**

If your Web services is part of a group of Web services configured on the GAS, you must provide
the group name. If it does not belong to a group, group-name can be omitted; in
this case, the Web service belongs to the default group, which does not have to be named in the URL.

**xcf-file**

You must provide the Web services configuration filename. For example, if the configuration
file is myServices.xcf, you specify "`myServices`" for the
xcf-file path element.

For more information, see [REST resource URI naming practice](../16_web-services/4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.").

## Default REST media type JSON for records and arrays

Starting with version 4.00, the REST server generates JSON (`application/json`) as
the default MIME type for records and arrays. If you do not set the data format with a `WSMedia` attribute, the GWS defaults to
serializing the record or array as JSON.

In version 3.20, the default was either JSON (`application/json`) or XML
(`application/xml`). You can restore this default behavior by calling the REST
service engine to set the option at
runtime:

```
com.WebServiceEngine.SetOption("server_restdefaultformat","both")
```

For examples using `WSMedia`, see [REST API reference (high-level framework)](../16_web-services/4795-reference.md "The reference information provided in this section allow you to create RESTful Web service server and client applications using the Genero Web Services high-level framework.").

## Changes to how fglrestful handles errors in the stub

When upgrading your RESTful Web services from version 3.20 to version 4.00, be aware of the
changes to how Genero Web Server (GWS) errors are handled in the client stub.

The [wsError](../16_web-services/4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.") record now handles
unexpected errors in the client stub. This has implications for how you handle server errors in your
client application.

- With version 3.20, the GWS returned a "`-1`" in the return status code
  (`wsstatus`) of the client-stub function for a GWS runtime error, such as the service
  not running. No further details were available to the client.
- With version 4.00, the client application now gets details of the error (code and description)
  from the `wsError`
  record.

There is also a change to how you can use the status code (`wsstatus`) to trap
errors. You must now code in the following way:

- If `wsstatus < 0`, check `wsError` for details. The
  `wsError` allows you to get details of the error. Its `code` field
  will have a value of either `< 0` (if it is a GWS runtime error) or the HTTP
  status code. Its `description` field will have a message.
  > **Note:**
  >
  > If the status code is less than zero, it indicates a GWS runtime error, such as -15333,
  > or it indicates an HTTP error that the stub could not handle because it was not specified in the
  > openapi.json.
- If `wsstatus == 0`, there is no error.
- If `wsstatus > 0`, check the related error record for details. This error is
  expected, for example, it is specified in the openapi.json, and
  fglrestful has generated the `userError` record in the stub that
  can return details.

For an example and more details about handling server errors in the client, see [Handle GWS REST server errors](../16_web-services/4790-handle-rest-server-errors.md "Handle the status code returned by a REST service call and process expected and unexpected errors in the client application.").

## Changes to how fglwsdl generates stubs and uses linking

Prior to version 4.00, The fglwsdl tool generated two files from a WSDL:

- a stub file (.4gl).
- a Genero BDL include source file (.inc), also known as a globals file, that
  was [linked](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.") with the stub file
  (.4gl).

Starting with version 4.00, fglwsdl no longer creates an
.inc file. It now generates one stub file that defines all types and variables,
considered global, as `PUBLIC` modular types and variables that you can access using
the `IMPORT FGL` statement.

This may impact the migration of your SOAP Web services from version 3.20 to version 4.00.
See:

- Changes to how fglwsdl generates the client stub
- Changes to how fglwsdl generates the server stub
- Changes to how fglwsdl generates the client stub with DOM API

For GWS services or applications written prior to version 4.00, you can still use the
.inc file and traditional linking by running the fglwsdl
tool with the [legacy
option](../13_programming-tools/2523-fglwsdl.md).

If you are generating stub files for a new GWS service or application, where you previously
included (for example) a `GLOBALS "ws_myGlobals.inc` " statement,
you must now use `IMPORT FGL`
statements instead:

- Import the stub file with an `IMPORT FGL ws_myStub` statement
  in your client app.
- If you are using the [wsError](../16_web-services/4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.")
  record, it is now defined in the WSHelper module. You must import this module with an `IMPORT
  FGL WSHelper` statement.
- It is recommended to call functions in the stub using their qualified names. For example, use
  `CALL ws_myStub.Add(1,2)` instead of `CALL
  Add(1,2)` or `CALL callback_module.HandleRequest()`
  instead of `CALL HandleRequest()`.

> **Note:**
>
> [Linking](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.") is no longer required
> when compiling.

For further details, see [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).").

## Changes to how fglwsdl generates the client stub

Starting with version 4.00, there are changes  to
how [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") generates stubs and uses linking. The `fglwsdl
-c` command generates one stub file where it defines all types and variables, considered
global, as `PUBLIC` modular types and variables that you can access using the
`IMPORT FGL` statement. For example of the command and its use, see [Generate client stub file](../13_programming-tools/2523-fglwsdl.md) and [WS client stubs and handlers](../16_web-services/4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.").

For GWS applications written prior to version 4.00, you can still use the
.inc file and traditional linking by creating it with the fglwsdl -c
-legacy command. For further details, see  [legacy option](../13_programming-tools/2523-fglwsdl.md).

## Changes to how fglwsdl generates the server stub

Starting with version 4.00, there are changes  to
how [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") generates stubs and uses linking. The `fglwsdl -s`command generates one stub file where it defines all types and variables, considered
global, as `PUBLIC` modular types and variables that you can access using the
`IMPORT FGL` statement. The command requires a mandatory filename as the name of the
module where you write the functions for the service operations. For example of the command and its
use, see [Generate server stub file](../13_programming-tools/2523-fglwsdl.md) and [Example 2: Writing a server using third-party WSDL (the fglwsdl tool)](../16_web-services/4661-example-2-writing-a-server-using-third-party-wsdl-the-fglwsd.md "Describes using a server stub from a third-party Web service in your GWS server application.").

For GWS services written prior to version 4.00, you can still use the .inc
file and traditional linking by creating it with the fglwsdl -s -legacy command.
For further details, see [legacy option](../13_programming-tools/2523-fglwsdl.md).

## Changes to how fglwsdl generates the client stub with DOM API

Starting with version 4.00, there are changes  to
how [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") generates stubs and uses linking. The fglwsdl
-domhandler command tool option (to create a client stub file from a WSDL based on the DOM
API) generates one stub file where it defines all types and variables, considered global, as
`PUBLIC` modular types and variables that you can access using the `IMPORT
FGL` statement.

You must give the command a mandatory filename, as the name of the module where you write the
callback functions for the service operations. For
example:

```
fglwsdl -domHandler callback_module -o myDomStub service.wsdl
```

For more information on its use, see [The generated callback handlers](../16_web-services/4623-the-generated-callback-handlers.md "Understand the function of callback handlers, and how to generate them in your client stub.").

For GWS applications written prior to version 4.00, you can still create the
.inc file and traditional linking by running the fglwsdl -domHandler
-legacy command.

## WSContext supports multipart content

Starting with version 4.00, the high-level REST API supports setting multipart content-type via
the `WSContext` dictionary variable. For example, in FGLGWS 3.20 you could set the
`Content-Type` for data, such as an image file, in the body of the message at runtime
using the `WSContext` dictionary. From GWS version 4.00 onwards you can set the
`Content-Type` for each data part in a multipart message. For more information, see
[WSContext](../16_web-services/4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.").

## REST API supports versioning

Starting with FGLGWS version 4.00.03, the high-level REST API supports versioning via the [WSVersion (function)](../16_web-services/4825-wsversion-function.md "Specifies the version or versions in which the REST function is available.") attribute. In FGLGWS 3.20 you
could set the version of operations by hand setting either the URI path or a custom header in the
operations. Now you can set a default version for the entire Web service with the [WSVersion (module)](../16_web-services/4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.") attribute, and you can specify
whether versioning is done through URI Path, query parameters, or custom header with the [WSVersionMode](../16_web-services/4809-wsversionmode.md "Specifies how the service version is selected for OpenAPI generation (URI, header, or query).") attribute. For more
information, see [Version a REST web service API](../16_web-services/4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.").

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Web
services changes in BDL 3.21](0155-web-services-changes.md "There are changes in support of web services in Genero 3.21.").

Notable changes introduced in maintenance releases:

- [Support for OpenSSL 3](0155-web-services-changes.md), also available in FGLGWS 4.01.05.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
