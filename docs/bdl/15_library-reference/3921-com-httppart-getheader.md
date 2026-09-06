---
title: "com.HttpPart.getHeader"
source: "fgl-topics/c_gws_ComHTTPPart_getHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.getHeader"
type: "concept"
---

# com.HttpPart.getHeader

> Returns a named HTTP multipart header as a string.

## Syntax

```
getHeader(
   name STRING )
  RETURNS STRING
```

1. name specifies the name of the header part.

## Usage

Use this method to retrieve HTTP multipart headers.

The method returns the value for the header part specified by name.

In case of related multipart (i.e., the part is
multipart/related and set via the
`com.HttpRequest.setMultipartType("related",NULL,NULL)`), it is
mandatory to set a unique Content-ID header. To set up a unique Content-ID header, use
the [security.RandomGenerator.CreateUUIDString](4412-security-randomgenerator-createuuidstring.md "Creates a new universal unique identifier (UUID).") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
LET val = httppart.getHeader("MyClientHeader")
```

## Related links

**Related concepts**  

[com.HttpPart.setHeader](3926-com-httppart-setheader.md "Sets a named HTTP multipart header using a string value.")
