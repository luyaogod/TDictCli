---
title: "security.Digest.CreateDigest"
source: "fgl-topics/c_gws_SecurityDigest_CreateDigest.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods > security.Digest.CreateDigest"
type: "concept"
---

# security.Digest.CreateDigest

> Defines a new digest context by specifying the algorithm to be used.

## Syntax

```
security.Digest.CreateDigest(
   algo STRING )
  RETURNS security.Digest
```

1. algo defines the digest algorithm to be used.

## Usage

Use this method to create and initialize a digest context to compute data digest based on the
specified algorithm.

Available digest algorithms are:

- `"SHA1"`
- `"SHA224"`
- `"SHA256"`
- `"SHA384"`
- `"SHA512"`
- `"MD5"`

> **Note:**
>
> MD5 digest is not recommended and may not be available on the system running Genero due to
> certain security configurations or the version of OpenSSL used.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
