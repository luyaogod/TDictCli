---
title: "Access HTTP request and response headers for a service"
source: "fgl-topics/c_gws_client_behavior_010.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Access HTTP request and response headers for a service"
type: "concept"
---

# Access HTTP request and response headers for a service

> Configure additional headers for requests and responses by adding them to the global endpoint record.

To access HTTP headers exchanged between the Genero client and a web service, you must use the
following records in the binding section:

- one record called `Request` in order to customize HTTP headers to be sent to a
  web service
- one record called `Response` in order to retrieve all HTTP headers returned by a
  web service

```
TYPE tGlobalEndpointWithHttpLayerType RECORD # End point
  Address RECORD # Address
    Uri STRING   # URI
  END RECORD,
  Binding RECORD # Binding
    Version STRING, # HTTP Version (1.0 or 1.1)
    Cookie STRING,  # Cookie to be set
    Request RECORD
      Headers DYNAMIC ARRAY OF RECORD # HTTP Headers
        Name STRING,
        Value STRING
      END RECORD
    END RECORD,
    Response RECORD
      Headers DYNAMIC ARRAY OF RECORD # HTTP Headers
        Name STRING,
        Value STRING
      END RECORD
    END RECORD,       
    ConnectionTimeout INTEGER, # Connection timeout
    ReadWriteTimeout INTEGER   # Read write timeout
    CompressRequest  STRING    # HTTP compression mode (gzip or deflate)
  END RECORD
END RECORD
```

Description of additional Request and Response variables:

- `Binding.Request.Headers`: Represents the additional HTTP headers to be sent to
  the web service. (Notice that client stub headers will replace user ones with the same name).
- `Binding.Response.Headers`: Represents the HTTP headers returned by a web
  service.

## Related links

**Related reference**  

[Global Endpoint type definition](4608-global-endpoint-type-definition.md "The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.")

[WS-Addressing 1.0 Global Endpoint type definition](4609-ws-addressing-1-0-global-endpoint-type-definition.md "The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.")
