---
title: "security.BCrypt.CheckPassword"
source: "fgl-topics/c_gws_SecurityBCrypt_CheckPassword.html"
breadcrumb: "Library reference > Extension packages > The security package > The BCrypt class > BCrypt methods > security.BCrypt.CheckPassword"
type: "concept"
---

# security.BCrypt.CheckPassword

> Checks the hash password.

## Syntax

```
security.BCrypt.CheckPassword(
   password STRING,
   hashedPass STRING )
  RETURNS  INTEGER
```

1. password defines the password to hash. The password is limited
   to 72 bytes.
2. hashedPass defines the hash password created by the
   `HassPassword` method.

## Usage

The method returns `TRUE`, the hashed password is valid.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

For an example using `BCrypt` methods, see [Example: Using security.BCrypt methods](4461-bcrypt-example.md "This example creates (and checks) a hash password as BCrypt results.").

## Related links

**Related concepts**  

[security.BCrypt.HashPassword](4459-security-bcrypt-hashpassword.md "Creates a hash password.")

[security.BCrypt.GenerateSalt](4460-security-bcrypt-generatesalt.md "Generates the encoded value needed as input to the HashPassword method.")
