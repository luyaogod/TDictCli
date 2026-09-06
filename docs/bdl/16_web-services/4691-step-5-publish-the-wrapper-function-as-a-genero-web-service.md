---
title: "Step 5: Publish the wrapper function as a Genero web service"
source: "fgl-topics/c_gws_i4gl_migration_guide_007.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 5: Publish the wrapper function as a Genero web service"
type: "concept"
---

# Step 5: Publish the wrapper function as a Genero web service

> Based on the details in the configuration file (.4cf) file, create a function that registers your Web service with the Genero Web Service server.

Use the [COM APIs](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") to publish
the I4GL function as a Web service based on I4GL .4cf configuration file to get
a compatible Genero Web service.

To create a new BDL function in charge of the service publication, you will need the following
elements of the I4GL .4cf configuration file:

- The name of the service that is defined in the `SERVICENAME` entry
- The namespace of the service that is defined as http://www.ibm.com/
  concatenated to the `FUNCTION NAME`
- The name of the function to be published that is defined in the `FUNCTION NAME`
  entry

For example, the I4GL zipcode demo has one function published as a Doc/Literal
service.

```
FUNCTION create_zipcode_details_web_service()
  DEFINE serv com.WebService
  DEFINE op   com.WebOperation

  #
  # Create the web service based on the entries of the .4cf file
  #   SERVICENAME: The name of service is 'ws_zipcode'
  #   FUNCTION NAME: The namespace of the service is built from
  #                  the base url 'http://www.ibm.com/' concatenated to
  #                  the NAME of the I4GL function 'zipcode_details'
  #
  LET serv = com.WebService.CreateWebService("ws_zipcode",
              "http://www.ibm.com/zipcode_details")

  #
  # Create and publish the Doc/Literal web function based on
  #  step 2, step 3 and step 4
  #  and from the FUNCTION NAME defined in the .4cf file
  #
  LET op = com.WebOperation.CreateDOCStyle("zipcode_details_g",
            "zipcode_details",
            zipcode_details_in,
            zipcode_details_out)
  CALL serv.publishOperation(op,NULL)

  #
  # Register the service into the SOAP engine
  #
  CALL com.WebServiceEngine.RegisterService(serv)

END FUNCTION
```

I4GL supports only Doc/Literal services.

Genero Web Services can contain several BDL functions in the same service. In other words,
you can group several I4GL services in the same Genero service.
