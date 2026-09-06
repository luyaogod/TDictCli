---
title: "com.WebService.registerInputHttpVariable"
source: "fgl-topics/c_gws_ComWebService_registerInputHttpVariable.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.registerInputHttpVariable"
type: "concept"
---

# com.WebService.registerInputHttpVariable

> Registers the record variable for HTTP input.

## Syntax

```
registerInputHttpVariable(
   headers RECORD )
```

1. headers defines the HTTP
   input record variable with the following
   structure:

   ```
   RECORD
     verb  STRING,
     url   STRING,
     headers DYNAMIC ARRAY OF RECORD
       name  STRING,
       value  STRING
     END RECORD
   END RECORD
   ```

## Usage

The `registerInputHttpVariable()` method registers a program variable with a
specific structure, that will be filled with the HTTP request headers when a Web Operation
arrives.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.WebService.registerOutputHttpVariable](3764-com-webservice-registeroutputhttpvariable.md "Registers the record variable for HTTP output.")
