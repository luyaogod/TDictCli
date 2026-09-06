---
title: "Step 4: Modify the I4GL web service function call"
source: "fgl-topics/c_gws_i4gl_migration_guide_016.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Step 4: Modify the I4GL web service function call"
type: "concept"
---

# Step 4: Modify the I4GL web service function call

> Rename I4GL function name names to Genero Web service function names.

The Genero Web service function name is defined in the generated .4gl file
and must be used instead of the I4GL function name.

For example, in the I4GL zipcode demo, the Web service function name is
`cons_ws_zipcode`, which must be renamed to `zipcode_details` as
follows:

```
FUNCTION func_cons_ws_zipcode()
   DEFINE state_rec RECORD
            pin CHAR(10),
            city CHAR(100),
            state CHAR(100)
         END RECORD;
   #         
   # Genero web service status returning
   # whether web function call was successful or not
   #
   DEFINE soapstatus INTEGER  

   #
   # I4GL web service function name is 'cons_ws_zipcode'
   # CALL cons_ws_zipcode("97006")
   #   RETURNING state_rec.city, state_rec.state
   # Genero web service function name is 'zipcode_details'
   CALL zipcode_details("97006")
     RETURNING soapstatus, state_rec.city, state_rec.state
   ...
END FUNCTION
```

In Genero Web Services there is an additional returned parameter,
`soapstatus`. If it contains 0 the operation was a success, otherwise an error
occurred.
