---
title: "Step 3: Create a BDL RECORD for the output parameters"
source: "fgl-topics/c_gws_i4gl_migration_guide_005.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 3: Create a BDL RECORD for the output parameters"
type: "concept"
---

# Step 3: Create a BDL RECORD for the output parameters

> Define a BDL record for the output message of the Web function.

Add another modular BDL record where all members map to one of your I4GL Web service output
parameter, and keep the parameter order as defined in I4GL .4cf file.

You must then specify the Web service output message name via the Genero XML attribute called
`XMLName`, and assign it to the FUNCTION NAME as defined in the I4GL
.4cf file concatenated to response.

For example, in the I4GL zipcode demo there are two parameters: `city` and`state`. Genero Web Services supports complex data types as output parameters, so add
the following record at the beginning of the Genero
file:

```
DEFINE zipcode_details_out RECORD ATTRIBUTES(XMLName="zipcode_detailsresponse)
  city CHAR(100),
  state CHAR(100)
  END RECORD
```
