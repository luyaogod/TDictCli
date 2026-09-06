---
title: "com.HttpServiceRequest.getUrl"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_getUrl.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.getUrl"
type: "concept"
---

# com.HttpServiceRequest.getUrl

> Returns the URL of the HTTP service request. The URL may consists of a host name, a port number, a path, and a query.

## Syntax

```
getUrl()
  RETURNS STRING
```

## Usage

The `getUrl()` method returns the entire URL request containing the host, port,
document and query string.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

URLs are sent in UTF-8 on the network. If the query part of the URL cannot be converted from
UTF-8 to the fglrun locale charset, status will be set to [-15552](4483-genero-bdl-errors.md). Is this case, the document
part of the URL is available, but the query string must be retrieved through [HttpServiceRequest.readFormEncodedRequest()](3832-com-httpservicerequest-readformencodedrequest.md "Returns the string of a GET request with UTF-8 conversion option."). As a general advice, run your WS server program
in UTF-8.

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
