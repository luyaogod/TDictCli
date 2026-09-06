---
title: "WSHelper.SplitUrl()"
source: "fgl-topics/c_gws_wshelper_spliturl.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library > WSHelper.SplitUrl()"
type: "concept"
---

# WSHelper.SplitUrl()

> Splits a complete URL string into pieces.

## Syntax

```
FUNCTION SplitUrl( 
   url STRING )
  RETURNS ( scheme STRING,
            host STRING,
            port STRING,
            path STRING,
            query STRING )
```

1. url is a `STRING` with the URL to be split in pieces.
2. scheme is the URL scheme (http, https, file, and so on). It includes all
   parts before `://`
3. host is the hostname.
4. port : the port number, or `NULL` if there is none.

   If
   `NULL` is returned, there is always a default port depending on
   the scheme: 443 for HTTPS, and 80 for HTTP.
5. path is the path of the URL.
6. query is the query string of the URL, or `NULL` if there is
   none.

## Usage

Splits a complete URL string into pieces.

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
