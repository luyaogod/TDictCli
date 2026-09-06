---
title: "security.Digest.AddData"
source: "fgl-topics/c_gws_SecurityDigest_AddData.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods > security.Digest.AddData"
type: "concept"
---

# security.Digest.AddData

> Adds data from a BYTE variable to the digest buffer.

## Syntax

```
AddData(
   toDigest BYTE )
```

1. toDigest defines the binary data to be added to the digest buffer.

## Usage

Adds the binary data contained in the given `BYTE` to the digest context.

After adding all
the data pieces, the buffer can be processed by calling [`security.Digest.DoBase64Digest`](4448-security-digest-dobase64digest.md "Creates a digest of the buffered data and returns the result in base64 format.") or [`security.Digest.DoHexBinaryDigest`](4449-security-digest-dohexbinarydigest.md "Creates a digest of the buffered data and returns the result in hexadecimal format.").

> **Important:**
>
> The `BYTE` must be located in
> memory.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
