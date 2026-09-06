---
title: "com.WebService.registerOutputHttpVariable"
source: "fgl-topics/c_gws_ComWebService_registerOutputHttpVariable.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.registerOutputHttpVariable"
type: "concept"
---

# com.WebService.registerOutputHttpVariable

> Registers the record variable for HTTP output.

## Syntax

```
registerOutputHttpVariable(
   headers RECORD )
```

1. headers defines the HTTP input record
   variable with the following
   structure:

   ```
   RECORD
     code INTEGER,
     desc STRING,
     headers DYNAMIC ARRAY OF RECORD
       name  STRING,
       value  STRING
     END RECORD
   END RECORD
   ```

## Usage

The `registerOutputHttpVariable()` method registers a program variable with a
specific structure, that will be used to fill the HTTP response headers when a Web Operation is
completed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.WebService.registerInputHttpVariable](3762-com-webservice-registerinputhttpvariable.md "Registers the record variable for HTTP input.")
