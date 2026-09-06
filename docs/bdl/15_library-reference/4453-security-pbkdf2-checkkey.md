---
title: "security.PBKDF2.CheckKey"
source: "fgl-topics/c_gws_SecurityPBKDF2_CheckKey.html"
breadcrumb: "Library reference > Extension packages > The security package > The PBKDF2 class > PBKDF2 methods > security.PBKDF2.CheckKey"
type: "concept"
---

# security.PBKDF2.CheckKey

> Validates a hashed key.

## Syntax

```
security.PBKDF2.CheckKey(
   password STRING,
   salt STRING,
   hash STRING,
   iter INTEGER,
   hashedkey STRING )
  RETURNS BOOLEAN
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
5. hashedkey defines the key created by the [`Security.PBKDF2.GenerateKey()`](4454-security-pbkdf2-generatekey.md "Generates a password of a given size based on a human readable password using Password-Based Key Derivation Function 2 (PBKDF2)")
   class method.

## Usage

This method validates a hashed key produced by the `security.PBKDF2.GenerateKey`
method.

The method returns `TRUE`, if the hashed key is valid, and returns
`FALSE`, if the hashed key is not valid.

This method may raise exception [-15700](4483-genero-bdl-errors.md) (operation failed) or [-15701](4483-genero-bdl-errors.md) (invalid parameter).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
