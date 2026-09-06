---
title: "security.BCrypt.GenerateSalt"
source: "fgl-topics/c_gws_SecurityBCrypt_GenerateSalt.html"
breadcrumb: "Library reference > Extension packages > The security package > The BCrypt class > BCrypt methods > security.BCrypt.GenerateSalt"
type: "concept"
---

# security.BCrypt.GenerateSalt

> Generates the encoded value needed as input to the HashPassword method.

## Syntax

```
security.BCrypt.GenerateSalt(
   cost INTEGER )
  RETURNS  STRING
```

1. cost defines the number of rounds of hashing to apply. The
   default value is 10. This value must be between 4 and 30. It represents 2^cost iteration. An
   iteration above 14 may take several minutes to compute.
   > **Warning:**
   >
   > Using a high cost
   > value for the salt is very CPU consuming, and can really slow down the application depending on
   > the system it is running. Be forewarned that this is expected!

## Usage

This method generates the encoded value needed as input to the `HashPassword`
method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

For an example using `BCrypt` methods, see [Example: Using security.BCrypt methods](4461-bcrypt-example.md "This example creates (and checks) a hash password as BCrypt results.").

## Related links

**Related concepts**  

[security.BCrypt.CheckPassword](4458-security-bcrypt-checkpassword.md "Checks the hash password.")

[security.BCrypt.HashPassword](4459-security-bcrypt-hashpassword.md "Creates a hash password.")
