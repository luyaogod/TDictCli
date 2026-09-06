---
title: "Step 4: Build the HTTP request"
source: "fgl-topics/c_gws_rest_client_request.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Step 4: Build the HTTP request"
type: "concept"
---

# Step 4: Build the HTTP request

> In this step you code the HTTP request and perform the request to the server to add two numbers.

Include code inside a `TRY/CATCH` block to trap runtime exceptions. Assign two
values to fields in the `add_in` record as operands for the calculation. For details
about the `add_in` record we use in this topic, see [Step 3: Define the records](4858-step-3-define-the-records.md "In this step you define the records you need for the HTTP Request and Response and the processing of the data.").

```
TRY                    
    LET add_in.a = 1     
    LET add_in.b = 2      
CATCH
    ...
  
    ...

 END TRY
```

Recall we defined the variable `req` as an object of the
`com.HttpRequest` class earlier, see [Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md "The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages."). We instantiate this now by invoking the
`com.HttpRequest.Create` class method to create a URI for our HTTP request; adding
to the URI string:

```
http[s]://host:port/resource-name[?query-params]
```

1. host:port is the
   base URL from the server.
2. resource-name, the resource name or name of the server
   function, "add" in our case.
3. A question mark ("?"), to indicate a query string.
4. query-params is a query string that contains two
   key/value pairs in the format "field=value" separated by an ampersand ("&"), for example
   ("`a=1&b=2`"). They make up the query parameters or arguments required by the
   server's "add" function which will be extracted for processing by the Web service. We assign the
   values for the calculation from the "`add_in`" record.

```
LET req = com.HttpRequest.Create("http://localhost:8090/add?a=" || add_in.a || "&b=" || add_in.b)
```

In
our example, we anticipate the calculator Web server is on the localhost. For a
production environment you need to take care of URI parsing: making sure that the URI
matches what your service expects.

Set the HTTP method with the "GET" verb to request the service from the calculator Web resource.
We do this by calling the `com.HttpRequest` object's setMethod
function referenced by the `req` variable.

```
CALL req.setMethod("GET")
```

Next we add the `Accept` and `Content-Type` headers to the
`com.HttpRequest` object referenced by the `req` variable by calling
the setHeader function. These headers specify to the Web service how we intend to
receive and deliver content in the message body's media type, JSON or XML. Our preference is for
JSON.

```
CALL req.setHeader("Content-Type", "application/json")
CALL req.setHeader("Accept", "application/json")
```

Finally, we are ready to perform the HTTP request and call the `com.HttpRequest`
object's `doRequest` function.

```
CALL req.doRequest()
```

In the next step we handle the response from the Web service, [Step 5: Process the HTTP response](4860-step-5-process-the-http-response.md "In the final step you process the HTTP response; check for errors, and process the data.").
