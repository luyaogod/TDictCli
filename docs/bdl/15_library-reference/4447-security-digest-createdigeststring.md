---
title: "security.Digest.CreateDigestString"
source: "fgl-topics/c_gws_SecurityDigest_CreateDigestString.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Digest methods > security.Digest.CreateDigestString"
type: "concept"
---

# security.Digest.CreateDigestString

> Creates a SHA1 digest from the given string.

## Syntax

```
CreateDigestString(
   toDigest STRING,
   randomBase64 STRING )
  RETURNS STRING
```

1. toDigest defines the password to be digested.
2. randomBase64 defines a random string in Base64.

## Usage

Use this method to compute the SHA1 digest from the string in toDigest and an
optional random Base64 form string, and return it in a string encoded in Base64 form.

The random value must be a valid Base64 string. You typically generate this value with the [`security.RandomGenerator.CreateRandomString()`](4411-security-randomgenerator-createrandomstring.md "Creates a random base64 string.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
DEFINE password, digest STRING
...
LET digest =
    security.Digest.CreateDigestString(
       password,
       security.RandomGenerator.CreateRandomString(16) )
```
