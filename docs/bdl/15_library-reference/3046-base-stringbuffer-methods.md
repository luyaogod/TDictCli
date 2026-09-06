---
title: "base.StringBuffer methods"
source: "fgl-topics/c_fgl_ClassStringBuffer_methods.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods"
type: "concept"
description: "Table 1. Class methods Name Description base.StringBuffer.create () RETURNS base.StringBuffer Create a string buffer object. Table 2. Object methods Name Description append ( str STRING ) Append a ..."
---

# base.StringBuffer methods

| Name | Description |
| --- | --- |
| base.StringBuffer.create() RETURNS base.StringBuffer | Create a string buffer object. |

| Name | Description |
| --- | --- |
| append( str STRING ) | Append a string at the end of the current string. |
| clear() | Clear the string buffer. |
| equals( str STRING ) RETURNS BOOLEAN | Compare strings (case sensitive). |
| equalsIgnoreCase( str STRING ) RETURNS BOOLEAN | Compare strings (case insensitive) |
| getCharAt( index INTEGER ) RETURNS STRING | Return the character at a specified position. |
| getIndexOf( str STRING, startIndex INTEGER ) RETURNS INTEGER | Return the position of a substring. |
| getLength() RETURNS INTEGER | Return the length of a string. |
| insertAt( index INTEGER, str STRING ) | Insert a string at a given position. |
| replace( oldStr STRING, newStr STRING, occurrences INTEGER ) | Replace one string with another. |
| replaceAt( index INTEGER, length INTEGER, str STRING ) | Replace part of a string with another string. |
| subString( startIndex INTEGER, endIndex INTEGER ) RETURNS STRING | Return the substring at the specified position. |
| toLowerCase() | Converts the string in the buffer to lower case. |
| toUpperCase() | Converts the string in the buffer to upper case. |
| toString() RETURNS STRING | Create a `STRING` from the string buffer. |
| trim() | Remove leading and trailing blank space (ASCII 32) characters. |
| trimWhiteSpace() | Remove leading and trailing whitespace characters. |
| trimLeft() | Removes leading blank space (ASCII 32) characters. |
| trimLeftWhiteSpace() | Removes leading whitespace characters. |
| trimRight() | Removes trailing blank space (ASCII 32) characters. |
| trimRightWhiteSpace() | Removes trailing whitespace characters. |
