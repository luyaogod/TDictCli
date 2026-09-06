---
title: "com.HttpServiceRequest.getUrlQuery"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_getUrlQuery.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.getUrlQuery"
type: "concept"
---

# com.HttpServiceRequest.getUrlQuery

> Takes a dynamic array of RECORD of two strings and fills that array with the decoded query string of an HTTP service request.

## Syntax

```
getUrlQuery(
   query RECORD)
```

1. query defines a record with the following
   structure:

   ```
   DEFINE query DYNAMIC ARRAY OF RECORD
         name STRING,
         value  STRING
   END RECORD
   ```

   The
   WSHelper library provides the `WSHelper.WSQueryType` variable for your use. See [WSHelper.WSQueryType](../16_web-services/4925-wshelper-wsquerytype.md "The WSQueryType defines a dynamic array of key-value pairs that stores the query string of a URL.").

## Usage

Takes a dynamic array of RECORD of two strings and fills that array
with the decoded query string of an HTTP service request. If there is no query string,
the dynamic array size will be zero (0).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

If the given array is not of the expected type, it raises exception [-15535](4483-genero-bdl-errors.md).

If there is an encoding issue, the status is set to [-15552](4483-genero-bdl-errors.md). No exception is raised. You
can process the query parts that do not have any UTF-8 conversion issues.

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
