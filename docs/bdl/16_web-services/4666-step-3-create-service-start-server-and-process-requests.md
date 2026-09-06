---
title: "Step 3: Create service, start server and process requests"
source: "fgl-topics/c_gws_server_tutorial_013.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 2: Writing a server using third-party WSDL (the fglwsdl tool) > Step 3: Create service, start server and process requests"
type: "concept"
---

# Step 3: Create service, start server and process requests

> Code to start the Genero Web Services (GWS) Server.

Create your own `MAIN` module that calls the function from the generated
.4gl file to create the service, then starts the Genero Web Services Server and
manages requests as in [Step 5: Start the GWS server and process requests](4660-step-5-start-the-gws-server-and-process-requests.md "Code to start the Genero Web Services (GWS) Server.") of
[Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.").

```
# example2main.4gl file         -- contains the MAIN program block

IMPORT com
IMPORT FGL example1Service      -- import the generated service file

MAIN
   DEFINE create_status INTEGER

   -- call the function generated in example1Service.4gl
   CALL example1Service.Createexample1Service()                                    
      RETURNING create_status
   IF create_status <> 0 THEN
      DISPLAY "error"
   ELSE
      # Start the server and manage requests
      CALL ManageService()
   END IF

END MAIN

FUNCTION ManageService()
   DEFINE ret INTEGER
   CALL com.WebServiceEngine.start()
   WHILE TRUE
   # continue as in Step 5 of Example 1 
   ...
END FUNCTION
```

## Backward compatibility for globals

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your GWS client app, you have a reference in the [`GLOBALS`](4601-step-2-import-the-stub-file-legacy.md) statement at the
top of the .4gl module that links the stub file.

```
# example2main.4gl file            -- contains the MAIN program block

IMPORT com

GLOBALS "example1Service.inc"      -- use the generated globals file

MAIN
   DEFINE create_status INTEGER

   CALL Createexample1Service()    -- call the function generated in example1Service.4gl 

# continue as in the example above
# ...
```
