---
title: "security.RandomGenerator.CreateRandomNumber"
source: "fgl-topics/c_gws_SecurityRandomGenerator_CreateRandomNumber.html"
breadcrumb: "Library reference > Extension packages > The security package > The RandomGenerator class > RandomGenerator methods > security.RandomGenerator.CreateRandomNumber"
type: "concept"
---

# security.RandomGenerator.CreateRandomNumber

> Generates an 8-byte strong random number.

## Syntax

```
security.RandomGenerator.CreateRandomNumber()
  RETURNS BIGINT
```

## Usage

This method generates an 8-byte strong random number and returns it as a
`BIGINT`.

The generated number can then be used for advanced cryptographic features.

This method is based on OpenSSL, using `/dev/random` on UNIX® and `CryptGenRandom()` on Microsoft®
Windows®, which are based on the
cryptographically secure pseudo-random number generator (CSPRNG) specifications.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
