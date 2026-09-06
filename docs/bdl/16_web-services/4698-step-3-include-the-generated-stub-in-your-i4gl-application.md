---
title: "Step 3: Include the generated stub in your I4GL application"
source: "fgl-topics/c_gws_i4gl_migration_guide_015.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Step 3: Include the generated stub in your I4GL application"
type: "concept"
---

# Step 3: Include the generated stub in your I4GL application

> Use a GLOBALS statement to specify the Web service your application uses.

Add in all I4GL files calling a Web service in the generated .inc stub with
a `GLOBALS` instruction.

For example, in the I4GL zipcode demo, only the clsoademo.4gl file uses Web
services. So add the following line at the beginning of the
file:

```
GLOBALS "ws_zipcode_zipcode_detailsservice.inc"
MAIN
...
END MAIN
```

This allows access to the Genero global variables and data types used in the Web service call, so
the Genero global `wsError` record can retrieve error codes if any.
