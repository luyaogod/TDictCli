---
title: "WSHelper.FindQueryStringValue()"
source: "fgl-topics/c_gws_wshelper_findquerystringvalue.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library > WSHelper.FindQueryStringValue()"
type: "concept"
---

# WSHelper.FindQueryStringValue()

> Get a query string value by name.

## Syntax

```
FUNCTION FindQueryStringValue( 
   query STRING,
   name STRING )
  RETURNS STRING
```

1. query is the query string where to look for a value.
2. name is the name of the value wanted.

## Usage

Get a query string value by name.

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
