---
title: "Step 2: Create a BDL RECORD for the input parameters"
source: "fgl-topics/c_gws_i4gl_migration_guide_004.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 2: Create a BDL RECORD for the input parameters"
type: "concept"
---

# Step 2: Create a BDL RECORD for the input parameters

> Define a BDL record for the input message of the Web function.

Add a new modular BDL record where all members map to one of your I4GL Web service input
parameter, and keep the parameter order as defined in I4GL .4cf file.

You must then specify the Web service input message name via the Genero XML attribute called
`XMLName`, and assign it to the FUNCTION NAME as defined in the I4GL
.4cf file.

For example, in the I4GL zipcode demo there is only one parameter:  `pin`.

Genero Web Services supports complex data types as input parameters, so add the following
record at the beginning of the Genero
file:

```
DEFINE zipcode_details_in RECORD ATTRIBUTES(XMLName="zipcode_details")
  pin CHAR(10)
  END RECORD
```
