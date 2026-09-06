---
title: "security.Digest.DoBase64Digest"
source: "fgl-topics/c_gws_SecurityDigest_DoBase64Digest.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods > security.Digest.DoBase64Digest"
type: "concept"
---

# security.Digest.DoBase64Digest

> Creates a digest of the buffered data and returns the result in base64 format.

## Syntax

```
DoBase64Digest()
  RETURNS STRING
```

## Usage

This method processes the digest on all data previously added to the context and encodes it in
base64.

After that call, the internal buffer is cleaned and ready to be populated again with new data to
be digested.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
