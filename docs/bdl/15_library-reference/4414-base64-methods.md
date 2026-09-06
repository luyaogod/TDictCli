---
title: "security.Base64 methods"
source: "fgl-topics/r_gws_SecurityBase64_methods.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods"
type: "reference"
---

# security.Base64 methods

> Methods of the security.Base64 class.

| Name | Description |
| --- | --- |
| security.Base64.FromByte( data BYTE ) RETURNS STRING | Encodes the given BYTE data in base64. |
| security.Base64.FromHexBinary( hexBinVal STRING ) RETURNS STRING | Decodes the given hexadecimal string to base64. |
| security.Base64.FromString( clearVal STRING ) RETURNS STRING | Encodes the given string in base64. |
| security.Base64.FromStringWithCharset( clearVal STRING, charset STRING ) RETURNS STRING | Encodes the given string in base64, based on a given charset. |
| security.Base64.LoadBinary( path STRING ) RETURNS STRING | Reads data from a file and encodes to base64. |
| security.Base64.SaveBinary( path STRING, base64Data STRING ) | Decodes the given base64 string and writes the data to a file. |
| security.Base64.ToByte( val64 STRING, ret BYTE ) | Decodes the given base64 string into a BYTE. |
| security.Base64.ToHexBinary( val64 STRING ) RETURNS STRING | Decodes the given base64 string to hexadecimal. |
| security.Base64.ToString( val64 STRING ) RETURNS STRING | Decodes the given base64 string. |
| security.Base64.ToStringWithCharset( val64 STRING, charset STRING ) RETURNS STRING | Decodes the given base64 string, based on a given charset. |
| security.Base64.Xor( clearVal1 STRING, clearVal2 STRING ) RETURNS STRING | Computes the exclusive disjunction between two base64 encoded strings. |
