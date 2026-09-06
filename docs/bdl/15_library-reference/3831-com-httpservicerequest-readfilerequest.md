---
title: "com.HttpServiceRequest.readFileRequest"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_readFileRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.readFileRequest"
type: "concept"
---

# com.HttpServiceRequest.readFileRequest

> Returns the body of the request to a file on disk and returns the filename.

## Syntax

```
readFileRequest( )
    RETURNS STRING
```

## Usage

The `readFileRequest()` method returns the body of the request to a file on disk
and the retrieved value is the filename.

The file is created in the [temporary directory used by the runtime system
(DBTEMP)](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files."). The name of the file is the basename found in the HTTP Content-Disposition Header.
If this basename is not specified, the filename is created with a UUID. If a file with the same
name already exists in the temporary directory, the API prefixes the new file with a number. It is
then of the form: /tmp/ABC/filename\_index.ext, where
index represents the number of files with the same name on disk.

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a DELETE
> request, but clients may ignore it if they do not support it.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
