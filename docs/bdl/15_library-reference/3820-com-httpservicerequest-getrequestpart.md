---
title: "com.HttpServiceRequest.getRequestPart"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_getRequestPart.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.getRequestPart"
type: "concept"
---

# com.HttpServiceRequest.getRequestPart

> Returns the HttpPart object at the specified index position.

## Syntax

```
getRequestPart(
   pos INTEGER)
  RETURNS com.HttpPart
```

1. pos defines the index position.

## Usage

Returns the HttpPart object at the specified index position. The file part is stored in the
temporary directory as defined by the DVM environment or the system. If the DVM [DBTEMP](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files.") environment variable is set, this is used.
Otherwise the temporary directory as defined by the system is used. On Windows® the system temporary
directory is set by TEMP. On UNIX® the
/tmp directory is used.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

Can raise error [-15554](4483-genero-bdl-errors.md)
(Index is out of bounds).

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
