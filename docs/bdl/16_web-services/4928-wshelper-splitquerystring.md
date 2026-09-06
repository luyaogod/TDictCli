---
title: "WSHelper.SplitQueryString()"
source: "fgl-topics/c_gws_wshelper_splitquerystring.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library > WSHelper.SplitQueryString()"
type: "concept"
---

# WSHelper.SplitQueryString()

> Splits the query string of a URL into an array or key-value pairs.

## Syntax

```
FUNCTION SplitQueryString( 
   query STRING )
  RETURNS WSHelper.WSQueryType
```

1. query is the query string to be split into a dynamic array.

## Usage

Split a given query string into a dynamic array of key-value pairs, defined as
`WSHelper.WSQueryType`.

The pieces are returned in dynamic array of key-value pairs defined as a
`WSHelper.WSQueryType`. See [WSHelper.WSQueryType](4925-wshelper-wsquerytype.md "The WSQueryType defines a dynamic array of key-value pairs that stores the query string of a URL.") for more
information regarding `WSQueryType`.

`NULL` may be returned if a value is not found.

In case of error, a `NULL` value will be returned.

## WSHelper functions example

```
IMPORT FGL WSHelper

MAIN

DEFINE val, scheme, host, port, path, query STRING
DEFINE ind INTEGER
DEFINE ret WSHelper.WSQueryType

CALL WSHelper.SplitQueryString("name1=val1&name2=val2&name3=val3")
  RETURNING ret

DISPLAY "Query:"
FOR ind = 1 TO ret.getLength()
  DISPLAY " param : ", ret[ind].name, " = ", ret[ind].value
END FOR

LET val = WSHelper.FindQueryStringValue("name1=val1&name2=val2","name1")
DISPLAY "QueryStringValue :", val

CALL WSHelper.SplitURL("https://cube.strasbourg.4js.com:3128/GWS-492/TestSplitURL/test1")
  RETURNING scheme, host, port, path, query

  DISPLAY "URL parts:"
  DISPLAY " scheme: ",scheme
  DISPLAY " host  : ",host
  DISPLAY " port  : ",port
  DISPLAY " path  : ",path
  DISPLAY " query : ",query

END MAIN
```

Output:

```
Query:
 param : name1 = val1
 param : name2 = val2
 param : name3 = val3
QueryStringValue :val1
URL parts:
 scheme: https
 host  : cube.strasbourg.4js.com
 port  : 3128
 path  : /GWS-492/TestSplitURL/test1
 query :
```
