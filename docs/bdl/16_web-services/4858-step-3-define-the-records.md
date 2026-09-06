---
title: "Step 3: Define the records"
source: "fgl-topics/c_gws_rest_client_records.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Step 3: Define the records"
type: "concept"
---

# Step 3: Define the records

> In this step you define the records you need for the HTTP Request and Response and the processing of the data.

Declare a record to use for HTTP Request/Response status. The calculator demo application creates
a user-defined `TYPE` (`TYP_status`) to reference when defining
variables.

```
TYPE TYP_status RECORD  
        code INTEGER,  # The HTTP status code
        desc STRING    # The HTTP status description
    END RECORD
```

Declare
a public record called `info` to hold the pieces of information that make up the HTTP
request you send out and that comes back in the response from the Web service. You will use this
when processing the data and outputting it for display.

```
 PUBLIC DEFINE info RECORD 
    url     STRING,       #  URI of resource on the server
    verb    STRING,       #  HTTP method (POST, PUT, GET, or DELETE ) 
            #  For HTTP Request   
    reqtype STRING,       #  Request type identifier (Content-Type)
    request STRING,       #  Request data format (application/json, or application/xml)
    status  STRING,         
            #  For HTTP Response 
    resptype STRING,      #  Response type identifier (Accept)    
    response STRING,      #  Response data format (application/json, or application/xml)
    result RECORD         #  Record for runtime execution errors
        code INT,
        desc STRING
    END RECORD
END RECORD
```

Declare
two records, one (`add_in`) to hold operands for the add calculation and another
(`add_out`) to hold the result of the calculation and response status. Recall that
the type `TYP_status` was defined at the top of the
module.

```
DEFINE 
    add_in RECORD          
      a INTEGER,
      b INTEGER
    END RECORD,

    add_out RECORD         
      status TYP_status,  
      r INTEGER           
    END RECORD
```

At this stage we are ready to code the HTTP Request, [Step 4: Build the HTTP request](4859-step-4-build-the-http-request.md "In this step you code the HTTP request and perform the request to the server to add two numbers.")
