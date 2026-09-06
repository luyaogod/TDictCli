---
title: "Step 4: Create a BDL wrapper function"
source: "fgl-topics/c_gws_i4gl_migration_guide_006.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 4: Create a BDL wrapper function"
type: "concept"
---

# Step 4: Create a BDL wrapper function

> Create the wrapper function that uses BDL records to call the I4GL function.

Create a Genero BDL wrapper function without any parameters that will then use the input and
output record created at [Step 2](4688-step-2-create-a-bdl-record-for-the-input-parameters.md "Define a BDL record for the input message of the Web function.") and [Step 3](4689-step-3-create-a-bdl-record-for-the-output-parameters.md "Define a BDL record for the output message of the Web function.") to call the I4GL function passing it the parameters
retrieved from the records.

For example, in the I4GL zipcode demo there are 1 input and 2 output parameters. So the BDL
wrapper function must use these records to call the I4GL function as
follows:

```
FUNCTION zipcode_details_g()
	CALL zipcode_details(zipcode_details_in.pin) 
    RETURNING zipcode_details_out.city,zipcode_details_out.state
END FUNCTION
```
