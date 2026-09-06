---
title: "com.WebServiceEngine.Flush"
source: "fgl-topics/c_gws_ComWebServiceEngine_Flush.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.Flush"
type: "concept"
---

# com.WebServiceEngine.Flush

> Forces the Web Service engine to immediately flush the response of the web service operation.

## Syntax

```
com.WebServiceEngine.Flush()
  RETURNS INTEGER
```

## Usage

The `com.WebServiceEngine.flush()` class method allows for the return of the
response inside a high-level web service operation, before the end of the web service function.

When this method is used, any other web operation output parameter changes are ignored.

The
status returned by the method provides information about the execution of the
last web operation. A return status of zero means OK. For a complete list of error codes, see [Error codes of com.WebServicesEngine](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods.")
> **Note:**
>
> The return status of the `com.WebServiceEngine.flush()` method has the same
> meaning as a status returned by [`com.WebServiceEngine.ProcessServices()`](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services."), with the additional status code
> [-32](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods."), meaning that the flush method was
> called outside of a web operation execution context.

> **Note:**
>
> `com.WebServiceEngine.ProcessServices()` and [`com.WebServiceEngine.HandleRequest()`](3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.") can return the status code of [-31](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods."), meaning that the flush function has been
> called in the last executed web operation.

## Example

In this code example, the `flush()` method is used to force the response of the
web service
operation.

```
DEFINE echoBoolean_in, echoBoolean_out RECORD
    a_boolean  BOOLEAN ATTRIBUTES(XMLName="Boolean")
  END RECORD

MAIN
  DEFINE ret INTEGER
  ...
  WHILE true
    LET ret = com.WebServiceEngine.ProcessServices(-1)
    CASE ret
      WHEN 0
        DISPLAY "Request automatically processed."
      WHEN -31
        DISPLAY "Operation has been flushed."
      ...
  END WHILE
  ...
END MAIN

FUNCTION echoBoolean()
  DEFINE ret INTEGER
  -- Assign output parameter with input parameter
  LET echoBoolean_out.a_boolean = echoBoolean_in.a_boolean
  -- Immediate flush of web operation
  LET ret = com.WebServiceEngine.flush()
  IF ret != 0 THEN
     DISPLAY "ERROR Code : ",ret
     EXIT PROGRAM (1)
  END IF
  -- Changing the output parameters after flush() would have no effect.
END FUNCTION
```

## Related links

**Related concepts**  

[com.WebServiceEngine.HandleRequest](3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")

[com.WebServiceEngine.ProcessServices](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.")
