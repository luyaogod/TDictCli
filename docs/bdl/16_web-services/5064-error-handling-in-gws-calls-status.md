---
title: "Error handling in GWS calls (status)"
source: "fgl-topics/c_gws_error_handling.html"
breadcrumb: "Web services > Reference > Error handling in GWS calls (status)"
type: "concept"
---

# Error handling in GWS calls (status)

> When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.

By default, the program will stop if an exception is thrown. You can trap the GWS errors with a
`WHENEVER ERROR` handler or with a `TRY`/`CATCH` block.
In the example, the `readTextRequest()` API is surrounded by a
`TRY`/`CATCH`
block:

```
DEFINE req com.HttpServiceRequest,
       data STRING
...
LET req = com.WebServiceEngine.getHTTPServiceRequest(5)
...
TRY
   ...
   CALL req.readTextRequest() RETURNING data
   ...
CATCH
   CALL show_err(SFMT("Unexpected HTTP request read exception: %1", status))
END TRY
```

For some errors, a human-readable description of the error code is available in the
`sqlca.sqlerrm` register.

## Related links

**Related concepts**  

[Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.")

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
