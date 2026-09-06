---
title: "Compression in a SOAP Web services client"
source: "fgl-topics/c_gws_http_compression_client.html"
breadcrumb: "Web services > Concepts > Data compression in GWS > HTTP compression in SOAP Web services > Compression in a SOAP Web services client"
type: "concept"
---

# Compression in a SOAP Web services client

> Send and receive compressed requests from a SOAP Web services client.

When you create a low-level Web service and do not have any stubs created by
fglwsdl, you need to manage it by setting the HTTP headers.

## Send a compressed request

The method used to set up the client for sending a compressed request depends on whether the
Genero Web Services client is a high-level or low-level Web services client.
A high-level client is a
Genero Web Services client that includes the stub files created by the fglwsdl
tool.
A low-level client is a Genero Web Services client that does not
utilize stub files created by the fglwsdl tool.

Regardless of the type of client, the server must be set up to handle such compression, otherwise
the request will be rejected.

## Send a compressed request from a high-level client

Set the variable `Binding.CompressRequest` to either "gzip" or
"deflate".

```
LET EchoDocStyle_EchoDocStylePortTypeEndpoint.Binding.CompressRequest = "gzip"
```

The `Binding.CompressRequest` variable is defined in the stub file, specifically
the client's global (inc)
file.

```
#
# Global Endpoint user-defined type definition in WSHelper
#
PUBLIC TYPE tGlobalEndpointType RECORD  # End point
	Address RECORD # Address
	   Uri STRING  # URI
	END RECORD,
	Binding RECORD # Binding
	   Version STRING,	# HTTP Version (1.0 or 1.1)
	   Cookie STRING,	 # Cookie to be set
	   Request RECORD	 # HTTP request
             Headers DYNAMIC ARRAY OF RECORD  # HTTP Headers
	        Name STRING,
	        Value STRING
	      END RECORD
	   END RECORD,
	   Response RECORD # HTTP response
             Headers DYNAMIC ARRAY OF RECORD # HTTP Headers
	         Name STRING,
                Value STRING
             END RECORD
          END RECORD,
          ConnectionTimeout INTEGER,  # Connection timeout
          ReadWriteTimeout INTEGER,   # Read write timeout
          CompressRequest STRING      # HTTP request compression mode (gzip or deflate)
	END RECORD
END RECORD
```

## Send a compressed request from a low-level client

A low-level client is a Genero Web Services client that does not
utilize stub files created by the fglwsdl tool.

Set the `Content-Encoding` field in the request header to either "gzip" or
"deflate".

This example sets the `Content-Encoding` field to "gzip", where the request is a
`com.HttpRequest` object.

```
CALL request.setHeader("Content-Encoding","gzip")
```

## Accept a compressed response

A Genero Web Services client can accept a compressed request if it sets the
`Accept-Encoding` field in the header to "gzip, deflate". These values represent
supported compression schema names (called content-coding tokens) separated by commas.

This example sets the `Accept-Encoding` field with the `setHeader`
method, where the request is a `com.HttpRequest` object.

```
CALL request.setHeader("Accept-Encoding","gzip, deflate")
```

## Related links

**Related concepts**  

[Global Endpoint user-defined type definition](4621-global-endpoint-user-defined-type-definition.md "Bindings defined for the Web service can be changed at runtime.")

[WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")

[Compression in a SOAP Web services server](4549-compression-in-a-soap-web-services-server.md "Send and receive compressed requests from a Web services server.")

[SOAP Web services APIs](4510-soap-web-services-apis.md "Genero provides high-level and low-level APIs for creating SOAP Web services.")
