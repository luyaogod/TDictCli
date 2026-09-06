---
title: "Global Endpoint user-defined type definition"
source: "fgl-topics/c_gws_handlers_client_global_endpoint_user_defined_type_definition.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Global Endpoint user-defined type definition"
type: "concept"
---

# Global Endpoint user-defined type definition

> Bindings defined for the Web service can be changed at runtime.

A public variable for the bindings of the Web Service is generated in the stub file generated
by the fglwsdl. For
example:

```
# Location of the SOAP endpoint.
# You can reassign this value at run-time.
PUBLIC Calculator_CalculatorPortTypeSoap12Endpoint WSHelper.tGlobalEndpointType
```

This references the `WSHelper.tGlobalEndpointType` defined in the [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module, included in the
$FGLDIR/lib directory of the Genero Web Services package.

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

The `CompressRequest` entry is of type string. It is `NULL` by
default, meaning no request is compressed. To compress a request, set this variable to **gzip**
or **deflate**. The server must support compression; otherwise, the request will be rejected.

## Related links

**Related reference**  

[Global Endpoint type definition](4608-global-endpoint-type-definition.md "The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.")

[WS-Addressing 1.0 Global Endpoint type definition](4609-ws-addressing-1-0-global-endpoint-type-definition.md "The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.")
