---
title: "Global Endpoint type definition"
source: "fgl-topics/r_gws_soap_global_endpoint_record.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Global Endpoint type definition"
type: "reference"
---

# Global Endpoint type definition

> The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.

The `tGlobalEndpointType` global type is used by any generated client stub to
allow the programmer to change the client behavior at
runtime.

```
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

Description of variables:

- `Address.Uri`: Represents the location of the server.
  > **Important:**
  >
  > It replaces the global variable of type STRING generated prior to version
  > 2.40, therefore it is mandatory to regenerate the client stub and to modify the location assignation
  > in your application.
- `Binding.Version`: Represents the HTTP version to use for communication (only 1.0
  or 1.1 allowed, default is 1.1).
- `Binding.Cookie`: Represents the HTTP cookie to use for communication (or NULL if
  there is no cookie to send).
- `Binding.ConnectionTimeout` : Represents the maximum time in seconds to wait for
  the establishment of the connection to the server.
- `Binding.ReadWriteTimeout`: Represents the maximum time in seconds to wait for a
  connection read or write operation before breaking the connection.

## Related links

**Related concepts**  

[Global Endpoint user-defined type definition](4621-global-endpoint-user-defined-type-definition.md "Bindings defined for the Web service can be changed at runtime.")

**Related reference**  

[WS-Addressing 1.0 Global Endpoint type definition](4609-ws-addressing-1-0-global-endpoint-type-definition.md "The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.")
