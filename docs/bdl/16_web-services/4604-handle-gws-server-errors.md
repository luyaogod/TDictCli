---
title: "Handle GWS server errors"
source: "fgl-topics/c_gws_wsError_record.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Handle GWS server errors"
type: "concept"
---

# Handle GWS server errors

> When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.

The `wsError` record is defined in the [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module. To use it in your GWS client application, you must import it with an
`IMPORT FGL WSHelper` statement in your client [module](4599-step-1-import-the-com-library-of-the-gws-package.md).

```
DEFINE wsError RECORD
  code STRING,         -- Short description of the error
  codeNS STRING,       -- The namespace of the error code
  description STRING,  -- Long description of the error
  action STRING        -- internal "SOAP action"
END RECORD
```

## Related links

1. [Handle GWS server errors (legacy)](4605-handle-gws-server-errors-legacy.md)

   When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.
