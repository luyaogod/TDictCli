---
title: "util.Strings methods"
source: "fgl-topics/c_fgl_ext_util_Strings_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods"
type: "concept"
---

# util.Strings methods

> Methods for the util.Strings class.

| Name | Description |
| --- | --- |
| util.Strings.base64Decode( base64 STRING, filename STRING ) | Decodes a Base64 encoded string and writes the bytes to a file. |
| util.Strings.base64Encode( filename STRING ) RETURNS STRING | Converts the content of a file to a Base64 encoded string. |
| util.Strings.base64DecodeToHexString( base64 STRING ) RETURNS STRING | Decodes a base64 encoded string and returns the corresponding hexadecimal string. |
| util.Strings.base64DecodeToString( base64 STRING ) RETURNS STRING | Decodes a base64 encoded string and returns the corresponding string. |
| util.Strings.base64EncodeFromHexString( s STRING ) RETURNS STRING | Converts the hexadecimal string passed as parameter to a Base64 encoded string. |
| util.Strings.base64EncodeFromString( s STRING ) RETURNS STRING | Converts the string passed as parameter to a Base64 encoded string. |
| util.Strings.collate( s1 STRING, s2 STRING ) RETURNS INTEGER | Compares two strings using locale collation rules. |
| util.Strings.collateNumeric( s1 STRING, s2 STRING ) RETURNS INTEGER | Compares two strings using locale collation rules and sequences of numerical digits. |
| util.Strings.urlDecode( s STRING ) RETURNS STRING | Converts the URL-encoded string to a string in the current application locale. |
| util.Strings.urlEncode( source STRING ) RETURNS STRING | Converts a string from the current codeset to a URL-encoded string. |
