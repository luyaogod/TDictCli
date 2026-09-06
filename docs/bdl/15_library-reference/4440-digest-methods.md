---
title: "security.Digest methods"
source: "fgl-topics/r_gws_SecurityDigest_methods.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods"
type: "reference"
---

# security.Digest methods

> Methods of the security.Digest class.

| Name | Description |
| --- | --- |
| security.Digest.CreateDigest( algo STRING ) RETURNS security.Digest | Defines a new digest context by specifying the algorithm to be used. |

| Name | Description |
| --- | --- |
| AddData( toDigest BYTE ) | Adds data from a `BYTE` variable to the digest buffer. |
| AddBase64Data ( toDigest STRING ) | Adds data in base64 format to the digest buffer. |
| AddHexBinaryData( toDigest STRING ) | Adds data in hexadecimal format to the digest buffer. |
| AddStringData( toDigest STRING ) | Adds a data string to the digest buffer. |
| AddStringDataWithCharset( toDigest STRING, charset STRING ) | Adds a data string to the digest buffer, after converting to the specified character set. |
| CreateDigestString( toDigest STRING, randomBase64 STRING ) RETURNS STRING | Creates a SHA1 digest from the given string. |
| DoBase64Digest() RETURNS STRING | Creates a digest of the buffered data and returns the result in base64 format. |
| DoHexBinaryDigest() RETURNS STRING | Creates a digest of the buffered data and returns the result in hexadecimal format. |
