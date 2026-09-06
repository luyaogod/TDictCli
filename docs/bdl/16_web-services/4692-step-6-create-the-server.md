---
title: "Step 6: Create the server"
source: "fgl-topics/c_gws_i4gl_migration_guide_008.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service provider to Genero > Step 6: Create the server"
type: "concept"
---

# Step 6: Create the server

> Provide a file with a BDL function that starts your Web service with the Genero Web Service server instead of Axis.

I4GL uses Axis as server for its services, but Genero has its own server programmable via the
[COM library](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services."). Create a new file
and add the `IMPORT com` instruction at the beginning of the server file, then simply
create the main loop in BDL that will process any incoming HTTP request.

The port of the service defined in the I4GL .4cf configuration file (via the
`PORTNO` entry) can be reused by setting the FGLAPPSERVER environment variable to the
same value before running the server. However, only on development or for tests, on production
Genero Web Services requires an application server called GAS in charge of load balancing. See the
Genero Application Server User Guide for more
details about port configuration for deployment purpose.

For example, to migrate the I4GL zipcode demo, the service must be created in the server before
running the main loop as
follows:

```
MAIN
  DEFINE ret INTEGER
  DEFER INTERRUPT

  # Create zipcode_details service
  CALL create_zipcode_details_web_service()

  # Start the server on port set in FGLAPPSERVER
  #    (to be set to same value as PORTNO defined in the .4cf file)
  CALL com.WebServiceEngine.Start()

  # Handle any incoming request in a WHILE loop...
  WHILE TRUE
    LET ret = com.WebServiceEngine.ProcessServices(-1)
    CASE ret
      WHEN 0
        DISPLAY "Request processed."
      WHEN -1
        DISPLAY "Timeout reached."
      WHEN -2
        DISPLAY "Disconnected from application server."
        EXIT PROGRAM   # The Application server has closed the connection
      WHEN -3
        DISPLAY "Client Connection lost."
      WHEN -4
        DISPLAY "Server interrupted with Ctrl-C."
      WHEN -10
        DISPLAY "Internal server error."
      END CASE
      IF int_flag<>0 THEN
        LET int_flag=0
        EXIT WHILE
      END IF   
   END WHILE
END MAIN
```

With Genero Web Services, one server can contain several services. In other words, you can
put all your I4GL services into one server.

## Related links

**Related information**  

[Process requests](4660-step-5-start-the-gws-server-and-process-requests.md "Process requests")
