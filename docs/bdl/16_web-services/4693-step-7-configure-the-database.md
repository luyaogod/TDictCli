---
title: "Step 7: Configure the database"
source: "fgl-topics/c_gws_i4gl_migration_guide_009.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 7: Configure the database"
type: "concept"
description: "Based on the DATABASE entry in the I4GL .4cf configuration file, use the Genero instruction to connect to the Informix database at server startup."
---

# Step 7: Configure the database

> Based on the DATABASE entry in the I4GL .4cf configuration file, use the Genero instruction to connect to the Informix® database at server startup.

For example, in the I4GL zipcode demo the service accesses the database called: " i4glsoa". So
add the following instruction at the beginning of the server file created in [step
6](4692-step-6-create-the-server.md "Provide a file with a BDL function that starts your Web service with the Genero Web Service server instead of Axis."):

```
DATABASE i4glsoa
MAIN
...
END MAIN
```
