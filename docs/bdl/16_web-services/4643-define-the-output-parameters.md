---
title: "Define the output parameters"
source: "fgl-topics/c_gws_function_declaration_003.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Define the output parameters"
type: "concept"
---

# Define the output parameters

> Define a record for the output message of the function.

Even though output parameters in Genero Web functions are not allowed, each Web function can have
one global variable or module variable that defines the output message of the function. This message
must be a record where each field represents one of the output parameters of the Web function.

The name of each field corresponds to the name used in the SOAP response. These fields are
retrieved from the Web Services engine immediately after executing the BDL function, and sent back
to the client.

## Example output record

```
# VARIABLE : AddResponse
PUBLIC DEFINE AddResponse RECORD ATTRIBUTE(XMLSequence,XMLName="AddResponse",XMLNamespace="http://tempuri.org/",XMLNillable)
         r INTEGER ATTRIBUTE(XMLNamespace="",XMLNillable)
END RECORD
```

GWS 2.0 allows you to add optional attributes to the definition of data types. You can use
attributes to map the BDL data types in a Genero application to their corresponding XML data types.
See [Attributes to Customize XML Mapping](4958-xml-serialization-rules-and-customization.md) for
additional information.
