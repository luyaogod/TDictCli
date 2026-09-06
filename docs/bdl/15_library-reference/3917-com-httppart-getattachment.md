---
title: "com.HttpPart.getAttachment"
source: "fgl-topics/c_gws_ComHTTPPart_getAttachment.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.getAttachment"
type: "concept"
---

# com.HttpPart.getAttachment

> Returns the absolute path to the HTTP part.

## Syntax

```
getAttachment()
  RETURNS STRING
```

## Usage

Returns the absolute path location of the received part file.

The file is created in the [temporary directory used by the runtime system
(DBTEMP)](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files."). The name of the file is the basename found in the HTTP Content-Disposition Header.
If this basename is not specified, the filename is created with a UUID. If a file with the same
name already exists in the temporary directory, the API prefixes the new file with a number. It is
then of the form: /tmp/ABC/filename\_index.ext, where
index represents the number of files with the same name on disk.

If the file is encoded in base64, you can use the Genero Web Services fglpass
-dec64 command to convert it back to binary.

It is up to the programmer to remove the file from the disk when it is no longer needed.

To be used via methods: [com.HttpResponse.getPart](3901-com-httpresponse-getpart.md "Returns the HTTP part object at the specified index of the current HTTP response."),
[com.HttpResponse.getPartCount](3902-com-httpresponse-getpartcount.md "Returns the number of additional parts in the HTTP response."), and [com.HttpResponse.getPartFromID](3903-com-httpresponse-getpartfromid.md "Returns the HTTP part object marked with the given Content-ID value as identifier, or NULL if none.")

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
