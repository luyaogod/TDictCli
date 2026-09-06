---
title: "Define the input parameters"
source: "fgl-topics/c_gws_function_declaration_002.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Define the input parameters"
type: "concept"
---

# Define the input parameters

> Define a record for the input message of the Web function.

Even though input parameters are not allowed in Genero Web Service operations, each Web function
can have one public or module variable that defines the input message of the function. This variable
must be a record in which each field represents one of the input parameters of the Web function.

The name of each field corresponds to the name used in the SOAP request. These fields are filled
with the contents of the SOAP request by the Web Services engine just before executing the
corresponding BDL function.

## Example input record

```
# VARIABLE : Add
PUBLIC DEFINE Add RECORD ATTRIBUTES(XMLSequence,XMLName="Add",XMLNamespace="http://tempuri.org/",XMLNillable)
         a INTEGER ATTRIBUTES(XMLNamespace="",XMLNillable),
         b INTEGER ATTRIBUTES(XMLNamespace="",XMLNillable)
END RECORD
```

Genero version 2.0 allows you to add optional attributes to the definition of data types.
You can use attributes to map the BDL data types in a Genero application to their corresponding XML
data types. See [Attributes to Customize XML
Mapping](4958-xml-serialization-rules-and-customization.md) for additional information.
