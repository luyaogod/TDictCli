---
title: "Step 8: Compile and run the Genero service"
source: "fgl-topics/c_gws_i4gl_migration_guide_010.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 8: Compile and run the Genero service"
type: "concept"
---

# Step 8: Compile and run the Genero service

> Describes how to compile and run the service to test it.

Compile and link the 2 Genero files created and run your Genero service. It will be directly
available for any client, and will provide the WSDL when requested via a HTTP GET with WSDL as query
string.

## Example

The Genero web service is accessible on URL: http://hostname:9876/ws\_zipcode
and can return the WSDL on URL: http://hostname:9876/ws\_zipcode?WSDL.

```
$ fglcomp -M genero_service.4gl
$ fglcomp -M genero_server.4gl
$ fgllink -o genero_zipcode genero_service.42m genero_server.42m
$ export FGLAPPSEVER=9876
$ fglrun genero_zipcode.42r
```

The host name depends on the machine your Genero application is started.

For deploying the service on production sites you will need the Genero application server (GAS)
to load-balance the service. See the Genero Application Server User Guide about Web
services when deployment is required.
