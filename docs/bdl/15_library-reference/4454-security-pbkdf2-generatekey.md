---
title: "security.PBKDF2.GenerateKey"
source: "fgl-topics/c_gws_SecurityPBKDF2_GenerateKey.html"
breadcrumb: "Library reference > Extension packages > The security package > The PBKDF2 class > PBKDF2 methods > security.PBKDF2.GenerateKey"
type: "concept"
---

# security.PBKDF2.GenerateKey

> Generates a password of a given size based on a human readable password using Password-Based Key Derivation Function 2 (PBKDF2)

## Syntax

```
security.PBKDF2.GenerateKey(
   password STRING,
   salt STRING,
   hash STRING,
   iter INTEGER,
   keySize INTEGER )
  RETURNS STRING
```

1. password defines the human readable password to derive using
   the PBKDF2 method.
2. salt defines the base64 random value created
   using `Security.RandomGenerator.CreateRandomString()`. Can be
   `NULL`.
3. hash defines the hash operation. By default, it is
   "SHA1". Valid values include:
   - SHA1
   - SHA224
   - SHA256
   - SHA384
   - SHA512
   - MD5
4. iter defines the number of iterations to compute the derived
   password. This value must be greater than or equal to zero (>=0).
5. keySize defines the size in bytes of the
   returned key. Must be greater than zero (>0)

## Usage

This method generates a password of a given size based on a human readable password using PBKDF2.

In practice, with the same salt value and the same human readable password, the same key can be
regenerated in another application and therefore used as symmetric key to decrypt data encrypted in
the initial application.

This method may raise exception [-15700](4483-genero-bdl-errors.md) (operation failed) or [-15701](4483-genero-bdl-errors.md) (invalid parameter).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

For an example using `PBKDF2` methods, see [Example: Using security.PBKDF2 methods](4455-pbkdf2-example.md "This example generates a key size of 128-bits based on a given password.").
