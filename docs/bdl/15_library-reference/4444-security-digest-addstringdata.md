---
title: "security.Digest.AddStringData"
source: "fgl-topics/c_gws_SecurityDigest_AddStringData.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods > security.Digest.AddStringData"
type: "concept"
---

# security.Digest.AddStringData

> Adds a data string to the digest buffer.

## Syntax

```
AddStringData(
   toDigest STRING )
```

1. toDigest defines the human-readable character string to be
   added to the digest buffer.

## Usage

This method adds the specified string data to the digest buffer.

After adding all
the data pieces, the buffer can be processed by calling [`security.Digest.DoBase64Digest`](4448-security-digest-dobase64digest.md "Creates a digest of the buffered data and returns the result in base64 format.") or [`security.Digest.DoHexBinaryDigest`](4449-security-digest-dohexbinarydigest.md "Creates a digest of the buffered data and returns the result in hexadecimal format.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[security.Digest.AddStringDataWithCharset](4445-security-digest-addstringdatawithcharset.md "Adds a data string to the digest buffer, after converting to the specified character set.")
