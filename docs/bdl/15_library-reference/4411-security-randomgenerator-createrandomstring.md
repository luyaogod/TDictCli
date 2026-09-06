---
title: "security.RandomGenerator.CreateRandomString"
source: "fgl-topics/c_gws_SecurityRandomGenerator_CreateRandomString.html"
breadcrumb: "Library reference > Extension packages > The security package > The RandomGenerator class > RandomGenerator methods > security.RandomGenerator.CreateRandomString"
type: "concept"
---

# security.RandomGenerator.CreateRandomString

> Creates a random base64 string.

## Syntax

```
security.RandomGenerator.CreateRandomString(
   size INTEGER )
  RETURNS STRING
```

1. size defines the size of the random string.

## Usage

This method generates a random string of binary data of size bytes long and
returns it in a `STRING` encoded in a Base64 form.

The size must be greater than 0.

Use this function when randomness is required, such as in
`xml.CryptoKey.deriveKey()` or [`security.Digest.CreateDigestString()`](4447-security-digest-createdigeststring.md "Creates a SHA1 digest from the given string.").

This method is based on OpenSSL, using `/dev/random` on UNIX® and `CryptGenRandom()` on Microsoft®
Windows®, which are based on the
cryptographically secure pseudo-random number generator (CSPRNG) specifications.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
