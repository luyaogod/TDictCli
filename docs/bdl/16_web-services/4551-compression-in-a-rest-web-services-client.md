---
title: "Compression in a REST web services client"
source: "fgl-topics/c_gws_http_compression_rest_client.html"
breadcrumb: "Web services > Concepts > Data compression in GWS > HTTP compression in REST Web services > Compression in a REST web services client"
type: "concept"
---

# Compression in a REST web services client

> Send and receive compressed requests from a REST web services client.

When you create a low-level web service and do not have any stubs created by
fglrestful, you need to manage it by setting the HTTP headers.

## Send a compressed request

The method used to set up the client for sending a compressed request depends on whether the
Genero Web Services client is a high-level or low-level Web services client.
A high-level client is
a Genero Web services client using the GWS REST high-level framework that includes the stub
files created by the fglrestful tool.
A low-level client is a Genero Web services client that does
not utilize stub files created by the fglrestful tool.

Regardless of the type of client, the server must be set up to handle such compression;
otherwise, the request will be rejected.

## Send a compressed request from a high-level client

In the stub file, an `Endpoint` variable of type
`tGlobalEndpointType` is defined as public.

Set the variable `Binding.CompressRequest` to either "gzip" or
"deflate".

```
LET clientStub.Endpoint.Binding.CompressRequest = "gzip"
```

The `Binding.CompressRequest` variable is defined in the stub
file.

```
TYPE tGlobalEndpointType RECORD # REST Endpoint
      Address RECORD 
         Uri STRING 
      END RECORD,
      Binding RECORD 
         Version STRING, # HTTP Version (1.0 or 1.1)
         ConnectionTimeout INTEGER,
         ReadWriteTimeout INTEGER, 
         CompressRequest STRING 
      END RECORD
END RECORD
```

## Send a compressed request from a low-level client

Set the `Content-Encoding` field in the request header to either "gzip" or
"deflate".

This example sets the `Content-Encoding` field to "gzip", where the request is a
`com.HttpRequest` object.

```
CALL request.setHeader("Content-Encoding","gzip")
```

## Accept a compressed response

A Genero Web services client can accept a compressed request if it sets the
`Accept-Encoding` field in the header to "gzip, deflate". These values represent
supported compression schema names (called content-coding tokens) separated by commas.

This example sets the `Accept-Encoding` field with the `setHeader`
method, where the request is a `com.HttpRequest` object.

```
CALL request.setHeader("Accept-Encoding","gzip, deflate")
```

## Related links

**Related concepts**  

[Compression in a REST Web services server](4552-compression-in-a-rest-web-services-server.md "Send and receive compressed requests from a Web services server.")

[Information needed to generate a REST stub file](4781-before-generating-the-stub-file.md "You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.")

[REST Web services APIs](4511-rest-web-services-apis.md "Genero provides BDL function attributes to define high-level REST Web services. Low-level REST web services are implemented using standard HTTP classes.")
