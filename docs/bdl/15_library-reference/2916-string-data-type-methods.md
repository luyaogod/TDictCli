---
title: "STRING data type methods"
source: "fgl-topics/c_fgl_datatypes_STRING_methods.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods"
type: "concept"
description: "Table 1. Object methods Name Description append ( str STRING ) RETURNS STRING Concatenates a string. equals ( str STRING ) RETURNS BOOLEAN Compares a string to the content of a string variable. ..."
---

# STRING data type methods

| Name | Description |
| --- | --- |
| append( str STRING ) RETURNS STRING | Concatenates a string. |
| equals( str STRING ) RETURNS BOOLEAN | Compares a string to the content of a string variable. |
| equalsIgnoreCase( str STRING ) RETURNS BOOLEAN | Makes a case-insensitive string comparison. |
| expandTabs( ) RETURNS STRING | Converts TAB characters to a 8 space characters. |
| getCharAt( index INTEGER ) RETURNS STRING | Returns the character at the specified position. |
| getIndexOf( str STRING, startIndex INTEGER ) RETURNS INTEGER | Returns the position of a substring. |
| getLength( ) RETURNS INTEGER | Returns the length of the current string. |
| getMultibyteLength( ) RETURNS INTEGER | Counts the number of bytes in the string. |
| matches( regex STRING ) RETURNS BOOLEAN | Tests if the string matches a regular expression. |
| replaceAll( regex STRING, replacement STRING ) RETURNS STRING | Replace all substrings matching a regular expression. |
| replaceFirst( regex STRING, replacement STRING ) RETURNS STRING | Replace a substring matching a regular expression. |
| split( regex STRING ) RETURNS DYNAMIC ARRAY OF STRING | Splits the current string around matches of the given regular expression. |
| subString( startIndex INTEGER, endIndex INTEGER ) RETURNS STRING | Returns a substring from start and end positions in a given string. |
| toLowerCase( ) RETURNS STRING | Returns the string converted to lower case. |
| toUpperCase( ) RETURNS STRING | Returns the string converted to upper case. |
| trim( ) RETURNS STRING | Removes leading and trailing blank space (ASCII 32) characters. |
| trimWhiteSpace( ) RETURNS STRING | Removes leading and trailing whitespace characters. |
| trimLeft( ) RETURNS STRING | Removes leading blank space (ASCII 32) characters. |
| trimLeftWhiteSpace( ) RETURNS STRING | Removes leading whitespace characters. |
| trimRight( ) RETURNS STRING | Removes trailing blank space (ASCII 32) characters. |
| trimRightWhiteSpace( ) RETURNS STRING | Removes trailing whitespace characters. |
