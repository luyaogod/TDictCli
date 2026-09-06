---
title: "Step 5: Handle Genero Web Services errors"
source: "fgl-topics/c_gws_i4gl_migration_guide_017.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Step 5: Handle Genero Web Services errors"
type: "concept"
---

# Step 5: Handle Genero Web Services errors

> Describes how to code to check the error status returned by the I4GL Web service, and get details of the error.

I4GL web service errors are returned on a non-conventional SOAP fault that cannot be handled in
Genero. However, the errors are handled through the additional returned parameter
`soapstatus` that must be checked after each Web service call. If its value is not
zero, an error has occurred and can be retrieved via the global Genero `wsError`
record defined in the generated .inc file.

## Example

In the Genero Web Service you must check the soap status after each Web service call:

```
FUNCTION func_cons_ws_zipcode()
   DEFINE state_rec RECORD
             pin CHAR(10),
             city CHAR(100),
             state CHAR(100)
         END RECORD;
   #         
   # Genero web service status returning
   #  whether web function call was successful or not
   #
   DEFINE soapstatus INTEGER

   # Genero web service function call
   CALL zipcode_details("97006")
     RETURNING soapstatus, state_rec.city, state_rec.state
   # Check soap status for errors after zipcode_details call
   IF soapstatus<>0 THEN
      # Display error information from the server
      DISPLAY "Error:"
      DISPLAY "  code :",wsError.code
      DISPLAY "  ns   :",wsError.codeNS
      DISPLAY "  desc :",wsError.description
      DISPLAY "  actor:",wsError.action
   ELSE 
      # Display results
      DISPLAY "\n ------------------------- \n"
      DISPLAY "SUPPLIED ZIP CODE: 97006 \n"
      DISPLAY " ------------------------- \n"
      DISPLAY "RESPONSE FROM WEB SERVICE \n"
      DISPLAY " ------------------------- \n"
      DISPLAY " CITY:",state_rec.city
      DISPLAY "\n STATE:",state_rec.state
      DISPLAY "\n ======================== \n"
   END IF
   ...
END FUNCTION
```
