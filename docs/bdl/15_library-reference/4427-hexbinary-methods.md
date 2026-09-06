---
title: "security.HexBinary methods"
source: "fgl-topics/r_gws_SecurityHexBinary_methods.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods"
type: "reference"
---

# security.HexBinary methods

> Methods of the security.HexBinary class.

| Name | Description |
| --- | --- |
| security.HexBinary.FromBase64( val64 STRING ) RETURNS STRING | Converts a base64 string to the hexadecimal equivalent. |
| security.HexBinary.FromByte( data BYTE ) RETURNS STRING | Encodes BYTE data in hexadecimal. |
| security.HexBinary.FromString( clearVal STRING ) RETURNS STRING | Encodes a given string in hexadecimal. |
| security.HexBinary.FromStringWithCharset( clearVal STRING, charset STRING ) RETURNS STRING | Encodes a given string in hexadecimal, based on a given charset. |
| security.HexBinary.LoadBinary( path STRING ) RETURNS STRING | Reads binary data from a file and converts it to hexadecimal. |
| security.HexBinary.SaveBinary( path STRING, hexBinData STRING ) | Decodes an hexadecimal strings and writes the binary data to a file. |
| security.HexBinary.ToBase64( hexBinVal STRING ) RETURNS STRING | Converts an hexadecimal string to the base64 equivalent |
| security.HexBinary.ToByte( hex STRING, ret BYTE ) | Decodes an hexadecimal string into a BYTE variable. |
| security.HexBinary.ToString( hexVal STRING ) RETURNS STRING | Decodes an hexadecimal string to a clear, human-readable string. |
| security.HexBinary.ToStringWithCharset( hexVal STRING, charset STRING ) RETURNS STRING | Decodes an hexadecimal string to a clear, human-readable string, based on a given charset. |
| security.HexBinary.Xor( hexVal1 STRING, hexVal2 STRING ) RETURNS STRING | Computes the exclusive disjunction between two hexadecimal encoded strings. |
