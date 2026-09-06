---
title: "Troubleshooting stub file creation"
source: "fgl-topics/c_gws_rest_client_stub_file_troubleshooting.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Before generating the stub file > Troubleshooting"
type: "concept"
---

# Troubleshooting stub file creation

> Errors or warnings reported during stub file creation indicate issues in the OpenAPI description or in the data types defined by the service. This topic describes how to interpret and resolve these messages.

## Errors during stub file creation

The `fglrestful` tool reports errors when the [OpenAPI description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") contains
unsupported structures, incompatible types, or missing definitions. Errors prevent the stub file
from being generated.

Error messages include the JSON path of the OpenAPI element where the problem occurred and the
reason the tool cannot process it. For example:

```
Error in /paths/Property/UpdateProperties/{applicationID}/post/requestbody  
Reason : ComplexType='UpdatePropertiesRequestBodyType' not supported on binary operation
```

Review the OpenAPI description at the path shown in the message to determine whether the issue
can be corrected in the JSON description or if it requires assistance from your Four Js support
center.

## Warnings during stub file creation

Warnings do not stop stub file creation. However, enabling warnings may help identify problems in
the OpenAPI description that could affect client behavior. To display warnings, run the
`fglrestful` tool with the warning option set to `yes`:

```
fglrestful -W yes -o stub_filename http://localhost:8090/ws/r/myResource?openapi.json
```

The following warning indicates that an operation name appears more than once:

```
Warning in /paths/Liability/GetReasonTypes/{subscriberId}  
Reason : OperationId=GetReasonTypes has already been defined, append _1
```

Warnings help highlight inconsistencies in the OpenAPI description that you may want to correct,
even if the stub file can be generated successfully.

## Related links

**Related concepts**  

[fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.")

[REST stub file overview](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.")

**Related tasks**  

[Generate a REST stub file](4783-generate-stub-file.md "Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.")

**Related reference**  

[Error codes of com.WebServicesEngine](../15_library-reference/3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods.")
