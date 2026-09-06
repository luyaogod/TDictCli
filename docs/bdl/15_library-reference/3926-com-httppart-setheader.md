---
title: "com.HttpPart.setHeader"
source: "fgl-topics/c_gws_ComHTTPPart_setHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.setHeader"
type: "concept"
---

# com.HttpPart.setHeader

> Sets a named HTTP multipart header using a string value.

## Syntax

```
setHeader(
   name STRING,
   value STRING )
```

1. name specifies the multipart header
   name.
2. value specifies the multipart header
   value (such as HTTP headers).

## Usage

Use this method to set HTTP multipart headers.

For instance, when you send a multipart image, it is recommended that you specify the image mime type with this
header method. If the image is a png, you have to do
`part.setHeader("Content-Type","image/png")`, which allows the peer to know the
format of the attached file it has to process.

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
CALL req.setHeader("MyClientHeader","Hello")
```

## Related links

**Related concepts**  

[com.HttpPart.getHeader](3921-com-httppart-getheader.md "Returns a named HTTP multipart header as a string.")
