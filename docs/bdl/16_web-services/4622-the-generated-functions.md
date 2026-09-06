---
title: "The generated functions"
source: "fgl-topics/c_gws_handlers_client_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > The generated functions"
type: "concept"
---

# The generated functions

> Defines the requirements for Genero Web Services (GWS) client functions.

Genero Web Services (GWS) client functions have the following requirements:

- The function cannot have input parameters.
- The function cannot have return values.
- The function's input message must be defined as a public RECORD.
- The function's output message must be defined as a public RECORD.

As a result, two types of GWS functions are generated in the stub file for the Web Service
operation that you requested:

- One function type uses public records for the input and output. The names of these functions end
  in "\_g". Before calling the function in your GWS Client application, you must set the values in the
  public input record. After the function call, the status of the request is returned from the server,
  and the output message is stored in the public output record. In addition to performing the desired
  operation, this function handles the communication for the SOAP request and response, and sets the
  values in the `wsError` record as needed.
- The other function type serves as a "wrapper" for the "\_g" function. It passes the values of
  input parameters to the "\_g" function, and returns the output values and status received from the
  "\_g" function. Your client application does not need to directly access the public records. This
  function can only be used if the parameters are simple variables (no records or arrays).

## Backward compatibility for globals

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your GWS client app prior to version 4.00, you will find
prototypes of the functions for the GWS operation, and the definitions of the global
`INPUT` and `OUTPUT` records in the [globals file](4619-example-globals-file.md "The example WSDL file for the Calculator Web Service provides information about the service.") (.inc).
