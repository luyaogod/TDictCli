---
title: "WSHelper.WSQueryType"
source: "fgl-topics/c_gws_wshelper_WSQueryType.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library > WSHelper.WSQueryType"
type: "concept"
---

# WSHelper.WSQueryType

> The WSQueryType defines a dynamic array of key-value pairs that stores the query string of a URL.

## Syntax

```
TYPE WSQueryType DYNAMIC ARRAY OF RECORD
    name  STRING,
    value STRING
  END RECORD
```

1. `name` is the name of a query argument.
2. `value` is the value of a query argument.

## Usage

The `WSHelper.WSQueryType` type defines a dynamic array of key-value pairs to
return the query string of a given URL.

For examples using this record, see [WSHelper.SplitQueryString()](4928-wshelper-splitquerystring.md "Splits the query string of a URL into an array or key-value pairs.") and [Examples using com.HttpServiceRequest methods](../15_library-reference/3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.").
