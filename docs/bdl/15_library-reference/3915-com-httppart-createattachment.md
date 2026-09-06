---
title: "com.HttpPart.CreateAttachment"
source: "fgl-topics/c_gws_ComHTTPPart_CreateAttachment.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.CreateAttachment"
type: "concept"
---

# com.HttpPart.CreateAttachment

> Creates a new HttpPart object based on a given filename located on disk.

## Syntax

```
com.HttpPart.CreateAttachment(
   filename STRING )
  RETURNS com.HttpPart
```

1. filename specifies the name of a file.

## Usage

Creates a new HttpPart object based on a given filename located on disk. To be
used via the [addPart()](3849-com-httprequest-addpart.md "Adds a new part to the HTTP root part request.") method.

The `com.HttpPart.CreateAttachment()` method automatically sets the following
headers for the created HttpPart object:

- `Content-Type` is defined based on the filename extension. If the file extension
  is not recognized, `Content-Type` defaults to
  `application/octet-stream`.
  > **Note:**
  >
  > File extensions to `Content-Type`
  > mapping can be customized in the file $FGLDIR/lib/wse/mime.cfg.
- `Content-Transfer-Encoding` is set to "`binary`".
- `Content-Disposition` is set with the base name of the given
  filename as follows: `attachment;
  filename="basename"`.

For example, when calling the method as
follows:

```
LET part = com.HttpPart.CreateAttachment( "/opt/myapp/resources/logo.jpg" )
```

The resulting HTTP part headers will look
like:

```
Content-Type: image/jpeg
Content-Transfer-Encoding: binary 
Content-Disposition: attachment; filename="logo.jpg"
```

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
