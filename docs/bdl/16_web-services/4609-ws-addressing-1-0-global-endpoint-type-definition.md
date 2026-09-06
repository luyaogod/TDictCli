---
title: "WS-Addressing 1.0 Global Endpoint type definition"
source: "fgl-topics/r_gws_soap_WSAglobal_endpoint_record.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > WS-Addressing 1.0 Global Endpoint type definition"
type: "reference"
---

# WS-Addressing 1.0 Global Endpoint type definition

> The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.

The `tGlobalWSAEndpointType` global type is used by any generated client stub
where support for WS-Addressing 1.0 is enabled. It allows the programmer to change the client
behavior at runtime, and to send additional WS-Addressing 1.0 reference parameters to a server.

If this global type is used in your main application, you must add the `IMPORT
xml` instruction.

```
PUBLIC TYPE tGlobalWSAEndpointType RECORD # End point
	  Address RECORD # Address
	     Uri STRING, # URI
         Parameters DYNAMIC ARRAY OF xml.DomDocument # End point WSA reference parameters
	  END RECORD,
	  Binding RECORD  # Binding
		Version STRING, # HTTP Version (1.0 or 1.1)
		Cookie STRING,  # Cookie to be set
		Request RECORD  # HTTP request
		   Headers DYNAMIC ARRAY OF RECORD  # HTTP Headers
		      Name  STRING,
		      Value STRING
		   END RECORD
		END RECORD,
		Response RECORD  # HTTP response
		   Headers DYNAMIC ARRAY OF RECORD  # HTTP Headers
		       Name STRING,
		       Value STRING
		   END RECORD
		END RECORD,
		ConnectionTimeout INTEGER, # Connection timeout
		ReadWriteTimeout INTEGER,  # Read write timeout
		CompressRequest STRING     # HTTP request compression mode (gzip or deflate)
	   END RECORD
END RECORD
```

Description of variables:

- `Address.Parameters`: Represents the WS-Addressing 1.0 reference parameter to send
  to a WS-Addressing 1.0 compliant server.

  For a description of other parameters, see [Global Endpoint type definition](4608-global-endpoint-type-definition.md "The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.").

## Related links

**Related concepts**  

[Global Endpoint user-defined type definition](4621-global-endpoint-user-defined-type-definition.md "Bindings defined for the Web service can be changed at runtime.")
