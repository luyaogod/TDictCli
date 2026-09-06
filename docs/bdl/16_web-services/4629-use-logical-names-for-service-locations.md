---
title: "Use logical names for service locations"
source: "fgl-topics/c_gws_repository_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Use logical names for service locations"
type: "concept"
---

# Use logical names for service locations

> Using a logical reference for the Web service, instead of the real URL, in your client application URL binding has advantages for working with and deploying applications.

Genero Web Services provide a repository for Web service locations using FGLPROFILE. To achieve
maximum flexibility, you can map a logical reference used by your Web Services client application to
an actual URL. This is subject to the network configuration and access rights management of the
deployment site.

## Create logical reference with an FGLPROFILE entry

FGLPROFILE entries can be used to define server URLs by using URL patterns. The following entry
in the FGLPROFILE file maps the logical reference "myservice" to an actual URL:

```
ws.myservice.url = "http://www.MyServer.com/fastcgi/ws/r/MyWebService"
```

## Logical reference in the client application

When you [generate a Client
stub](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.") from WSDL information using the fglwsdl tool, a public variable for
the URL of the Web Service is contained in the stub file (.4gl) file, which
references the global endpoint type in the [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module.
For example:

```
# Location of the SOAP server.
# You can reassign this value at run-time.
PUBLIC DEFINE Calculator_CalculatorPortTypeSoap12Endpoint WSHelper.tGlobalEndpointType
```

You can use this public variable in your Web Services client application to assign a logical name
to the Web service URL:

```
LET Calculator_CalculatorPortTypeSoap12Endpoint.Address.Uri = "alias://myservice"
```

When the client application accesses the service, the location will be supplied
by the entry in the FGLPROFILE file on the client machine. This allows you to provide the same
compiled .42r application to different customers; the entries in the FGLPROFILE
file on each customer's machine would customize the Web Service location for that customer.

If you are migrating from a version prior to 2.40, see also [Web Services changes](../05_upgrading/0244-web-services-changes.md "There are changes in support of web services in Genero 2.40.").

## Logical reference in a legacy client application

Similarly, if you have [generated
stub files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your legacy GWS client application using
`GLOBALS`, a global variable for the URL of the Web Service is contained in the
.inc file. For
example:

```
# Location of the SOAP server.
# You can reassign this value at run-time.
DEFINE Calculator_CalculatorPortTypeEndpoint tGlobalEndpointType
```

You can use this global
variable in your Web Services client application to assign a logical name to the Web service URL:

```
LET Calculator_CalculatorPortTypeEndpoint.Address.Uri = "alias://myservice"
```

When the client application accesses the service, the location will be supplied
by the entry in the FGLPROFILE file on the client machine. This allows you to provide the same
compiled .42r application to different customers; the entries in the FGLPROFILE
file on each customer's machine would customize the Web Service location for that customer.

If you are migrating from a version prior to 2.40, see also [Web Services changes](../05_upgrading/0244-web-services-changes.md "There are changes in support of web services in Genero 2.40.").

## Examples of using logical reference in the URL

When you deploy a Genero Web Service with a GAS behind a Web Server, the service can be accessed
by two different URLs; for example:

- For internal clients: http://zeus/ws/r/MyWebService
- For clients using the Web:
  http://www.MyServer.com/fastcgi/ws/r/MyWebService

You use a logical name in the URL binding record and map the actual location of the Web Service
in your FGLPROFILE file, depending on the location of the client machine.

- For internal clients, the FGLPROFILE entry would
  be:

  ```
  ws.myservice.url = "http://zeus/ws/r/MyWebService"
  ```
- For clients using the Web, the FGLPROFILE entry would
  be:

  ```
  ws.myservice.url = "http://www.MyServer.com/fastcgi/ws/r/MyWebService"
  ```

## Related links

**Related concepts**  

[FGLPROFILE: Server URL patterns](4922-fglprofile-server-url-patterns.md "FGLPROFILE entries can be used to define multiple server URLs, by using URL patterns.")

[Global Endpoint user-defined type definition](4621-global-endpoint-user-defined-type-definition.md "Bindings defined for the Web service can be changed at runtime.")
