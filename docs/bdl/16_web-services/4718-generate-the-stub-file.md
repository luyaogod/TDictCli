---
title: "Generate the stub file"
source: "fgl-topics/t_gws_restful_high_level_quick_start_generate_stub.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 3: RESTful client application > Generate the stub file"
type: "task"
---

# Generate the stub file

> Use the fglrestful tool to generate the client stub from a REST Web service URL.

The generated [stub](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.") file contains
the complete code to manage calls made by client apps to the Web service resources. You must import
the stub file into the client app.

The "MyService" GWS REST service must be running on the specified port in order to provide the
service information.

## Steps

1. Use the [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") tool to generate the stub

   `fglrestful -o clientstub http://localhost:8090/MyService?openapi.json`

   Where:
   1. clientstub specifies the filename of the stub in the output
      (`-o`) option.
   2. The URL of the Web service is specified with the query string ?openapi.json
      to get the specification file.

   The clientstub.4gl is generated from the specification.
2. Compile the stub file.

   fglcomp clientstub.4gl

## Related links

**Related concepts**  

[REST stub file overview](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.")
