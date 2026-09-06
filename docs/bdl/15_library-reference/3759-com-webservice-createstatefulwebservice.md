---
title: "com.WebService.CreateStatefulWebService"
source: "fgl-topics/c_gws_ComWebService_createStatefulWebService.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.CreateStatefulWebService"
type: "concept"
---

# com.WebService.CreateStatefulWebService

> Creates a new object to implement a stateful Web service.

## Syntax

```
com.WebService.CreateStatefulWebService(
   name STRING,
   ns STRING ,
   state RECORD  )
  RETURNS com.WebService
```

1. name defines the Web service identifier.
2. ns defines the namespace for the Web
   service.
3. state defines the state between
   the client and server.

## Usage

The `com.WebService.CreateStatefulWebService()` class method creates a new
`com.WebService` object implementing a Web service that is stateful.

The name and ns parameters must
uniquely identify the Web service across the entire application, when multiple Web service programs
run on the same server. In theory, the value of ns+name must
be unique on the internet.

The state variable is used to identify the state between the client and the
server:

- For a WS-Addressing stateful service, the state variable must be a
  `RECORD` with the following structure, with the `W3CEndpointReference`
  variable
  attribute:

  ```
  RECORD ATTRIBUTES(W3CEndpointReference)
    address STRING, -- The location of the Web Service (for ex: URL)
    ref RECORD
      ... (other members defining the state)
    END RECORD
  END RECORD
  ```
- For a stateful service based on HTTP cookies, the state variable must be a
  simple variable defined with a basic [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").

It is up to the programmer to manage the state variable and to restore the
service state from a database.

When creating a stateful Web service, all published Web operations require a session in the
client request except those defined as '`initiateSession`'.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Stateful SOAP Web services](../16_web-services/4517-stateful-soap-web-services.md "The GWS provides support for stateful service. Different options for implementing SOAP stateful services are described.")
