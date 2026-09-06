---
title: "Step 1: Generate the Genero web service stub from an I4GL WSDL"
source: "fgl-topics/c_gws_i4gl_migration_guide_013.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Step 1: Generate the Genero web service stub from an I4GL WSDL"
type: "concept"
---

# Step 1: Generate the Genero web service stub from an I4GL WSDL

> Use the fglwsdl tool to get the WSDL information from the service provider.

Use the I4GL WSDL located on the Axis server to generate the Genero Web service client stub via
the tool called [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).")>.

For example, the WSDL file of the I4GL zipcode demo is located on
$INFORMIXDIR/AXIS2C/services/ws\_zipcode/zipcode\_details.wsdl.
So run the following command:

```
$ fglwsdl -noFacets zipcode_details.wsdl
```

It will generate these two Genero files:

- ws\_zipcode\_zipcode\_detailsservice.4gl 

  - It contains the Genero functions to connect to the server in SOAP over HTTP.
  - Take a look at the file if you are interested in Genero HTTP and XML low-level APIs.
- ws\_zipcode\_zipcode\_detailsservice.inc 

  - It contains the Genero data types and variables used for XML serialization.
  - Take a look at the file if you are interested in Genero XML to BDL variable mapping.

Option `-noFacets` is required for this demo because the I4GL CHAR data type will
be generated as string in Genero, which can lead to an XML serialization error if not present.
