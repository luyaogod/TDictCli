---
title: "com.WebService.CreateWebService"
source: "fgl-topics/c_gws_ComWebService_createWebService.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.CreateWebService"
type: "concept"
---

# com.WebService.CreateWebService

> Creates a new object to implement a Web Service.

## Syntax

```
com.WebService.CreateWebService(
   name STRING,
   ns STRING )
  RETURNS com.WebService
```

1. name defines the Web service identifier.
2. ns defines the namespace for the Web
   service.

## Usage

The `com.WebService.CreateWebService()` class method creates a new
`com.WebService` object implementing a Web Service.

The name and ns parameters must
uniquely identify the Web service across the entire application, when multiple Web service programs
run on the same server. In theory, the value of ns+name must
be unique on the internet.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
