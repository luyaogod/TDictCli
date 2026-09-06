---
title: "WSHelper: library"
source: "fgl-topics/r_gws_wshelper_apis.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library"
type: "reference"
---

# WSHelper: library

> The WSHelper library provides functions and types for working with URLs.

| Types | Description |
| --- | --- |
| TYPE WSQueryType DYNAMIC ARRAY OF RECORD name STRING, value STRING END RECORD | The `WSQueryType` defines a dynamic array of key-value pairs that stores the query string of a URL. |
| TYPE WSServerCookiesType DYNAMIC ARRAY OF RECORD name STRING, # Cookie name value STRING, # Cookie value path STRING, # Cookie path (or null) domain STRING, # Cookie domain (or null) expires DATETIME YEAR TO SECOND, # Cookie expiration date (or null) httpOnly BOOLEAN, secure BOOLEAN sameSite STRING # Lax (default), Strict, or None (requires secure) END RECORD | The `WSHelper.WSServerCookiesType` defines a dynamic array for storing cookies. |

| Function Name | Description |
| --- | --- |
| FUNCTION SplitUrl( url STRING ) RETURNS ( scheme STRING, host STRING, port STRING, path STRING, query STRING ) | Splits a complete URL string into pieces. |
| FUNCTION FindQueryStringValue( query STRING, name STRING ) RETURNS STRING | Get a query string value by name. |
| FUNCTION SplitQueryString( query STRING ) RETURNS WSHelper.WSQueryType | Splits the query string of a URL into an array or key-value pairs. |
