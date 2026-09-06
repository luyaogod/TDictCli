---
title: "Step 6: Compile and run the Genero client"
source: "fgl-topics/c_gws_i4gl_migration_guide_018.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Step 6: Compile and run the Genero client"
type: "concept"
---

# Step 6: Compile and run the Genero client

> Describes how to compile and run the client to test it.

Then simply compile your modified I4GL application for Genero and execute it. Your application
will then connect to the Web service; passing and returning the parameters as if they were only
simple BDL function calls.

For example, to compile your I4GL web service application for Genero, you must execute the
following commands:

```
$ fglcomp -M ws_zipcode_zipcode_detailsservice.4gl
$ fglcomp -M clsoademo.4gl
$ fgllink -o clsoademo.42r clsoademo.42m ws_zipcode_zipcode_detailsservice.42m
$ fglrun clsoademo.42r
```
