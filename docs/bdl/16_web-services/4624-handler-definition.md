---
title: "Handler definition"
source: "fgl-topics/c_gws_handlers_client_007.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > The generated callback handlers > Handler definition"
type: "concept"
---

# Handler definition

> Describes the various callback handlers and how they are defined.

There are three kind of callbacks you must implement for each service generated with the
`-domHandler` option.

- The request handler that allows the modification of the entire SOAP request before it is sent
  over the net.

  It must be named `ServiceName_HandleRequest`, where
  ServiceName is the name of the service following the different prefix options
  used during generation.

  It must return `TRUE` if you want the caller function to continue normally or
  `FALSE` to return from the caller function with a SOAP error you can define via the
  [wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.")
  record.

  ```
  FUNCTION ServiceName_HandleRequest(operation,doc,header,body)
    DEFINE operation STRING          -- Operation name of the 
                                     -- request to be modified.
    DEFINE doc       xml.DomDocument -- Entire XML document of the request
    DEFINE header    xml.DomNode     -- XML node of the SOAP header 
                                     -- of the request
    DEFINE body      xml.DomNode     -- XML node of the SOAP body of the 
                                     -- request
    CASE operation
      WHEN "Add"
        ... -- Use the DOM APIs to modify the request of the Add operation
      WHEN "Sub"
        ... -- Use the DOM APIs to modify the request of the Sub operation
      OTHERWISE
        DISPLAY "No modification for operation :",operation
    END CASE
    RETURN TRUE -- Continue normally in Add_g() or Sub_g()
  END FUNCTION
  ```
- The response handler that allows the validation of the entire SOAP response before it is
  deserialized into the corresponding record.

  It must be named `ServiceName_HandleResponse`, where
  ServiceName is the name of the service following the different prefix options
  used during generation.

  It must return `TRUE` if you want the caller function to continue normally or
  `FALSE` to return from the caller function with a SOAP error you can define via the
  [wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.")
  record.

  ```
  FUNCTION ServiceName_HandleResponse(operation,doc,header,body)
    DEFINE operation STRING         -- Operation name of the 
                                    -- response to be checked.
    DEFINE doc      xml.DomDocument -- Entire XML document of the response
    DEFINE header   xml.DomNode     -- XML node of the SOAP header of 
                                    -- the response
    DEFINE body     xml.DomNode     -- XML node of the SOAP body of the 
                                    -- response
    CASE operation
      WHEN "Add"
        ... -- Use the DOM APIs to check the response of the Add operation
      WHEN "Sub"
        ... -- Use the DOM APIs to check the response of the Sub operation
      OTHERWISE
        DISPLAY "No verification for operation :",operation
    END CASE
    RETURN TRUE -- Continue normally in Add_g() or Sub_g()
  END FUNCTION
  ```
- The fault response handler that allows the verification of the entire SOAP fault response before
  it is deserialized into the wsError record.

  It must be named `ServiceName_HandleResponseFault`, where
  ServiceName is the name of the service according to the different prefix options
  used during generation. It must return `TRUE` if you want the caller function to
  continue normally or `FALSE` to return from the caller function with a SOAP error you
  can define via the **[wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.")**
  record.

  ```
  FUNCTION ServiceName_HandleResponseFault(operation,doc,header,body)
    DEFINE operation STRING          -- Operation name of the fault 
                                     -- response to be checked.
    DEFINE doc       xml.DomDocument -- Entire XML document of the fault response
    DEFINE header    xml.DomNode     -- XML node of the SOAP header of the 
                                     -- fault response
    DEFINE body      xml.DomNode     -- XML node of the SOAP body of the fault 
                                     -- response
    CASE operation
      WHEN "Add"
        ... -- Use the DOM APIs to verify the SOAP fault response 
            -- of the Add operation
      WHEN "Sub"
        ... -- Use the DOM APIs to verify the SOAP fault response 
            -- of the Sub operation
      OTHERWISE
        DISPLAY "No verification for operation :",operation
    END CASE
    RETURN TRUE -- Continue normally in Add_g() or Sub_g()
  END FUNCTION
  ```
