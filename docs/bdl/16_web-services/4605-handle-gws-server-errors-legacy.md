---
title: "Handle GWS server errors (legacy)"
source: "fgl-topics/c_gws_wsError_record_legacy.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Handle GWS server errors > Handle GWS server errors (legacy)"
type: "concept"
---

# Handle GWS server errors (legacy)

> When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.

If you generated your client stub files using the --legacy option, the
generated globals file (.inc) contains the definition for the
`wsError` record.

To include the globals file in your client application, see [Step 2: Import the stub file (legacy)](4601-step-2-import-the-stub-file-legacy.md).

The `wsError` record is defined as
follows:

```
DEFINE wsError RECORD
  code STRING,         -- Short description of the error
  codeNS STRING,       -- The namespace of the error code
  description STRING,  -- Long description of the error
  action STRING        -- internal "SOAP action"
END RECORD
```
