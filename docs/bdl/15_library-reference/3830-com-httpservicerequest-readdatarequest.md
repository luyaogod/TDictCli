---
title: "com.HttpServiceRequest.readDataRequest"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_readDataRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.readDataRequest"
type: "concept"
---

# com.HttpServiceRequest.readDataRequest

> Returns the body of a request in a BYTE.

## Syntax

```
readDataRequest(
  b BYTE)
```

1. b defines a variable of type
   `BYTE` that will be filled with the request body.

## Usage

The `readDataRequest()` method returns the body of the request in a [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.").

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a DELETE
> request, but clients may ignore it if they do not support it.

The `BYTE` variable must be located in memory, and will be filled with the request
body. The existing content of the `BYTE` will be discarded.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
