---
title: "xml.CryptoKey.deriveKey"
source: "fgl-topics/c_gws_XmlCryptoKey_deriveKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.deriveKey"
type: "concept"
---

# xml.CryptoKey.deriveKey

> Derives the symmetric or HMAC CryptoKey object using the given method identifier and concatenating the optional label, the mandatory seed value and the optional created date as initial random value.

## Syntax

```
deriveKey(
   url STRING,
   label STRING,
   seed STRING,
   created STRING,
   offset INTEGER,
   bytes INTEGER )
```

1. url defines the [identifier](4223-derived-keys.md) of the algorithm to
   apply to the password and its inputs.
2. label defines the optional label
   input.
3. seed defines the mandatory seed input as a valid Base64 string representing
   random binary data obtained with the [security.RandomGenerator.CreateRandomNumber](4410-security-randomgenerator-createrandomnumber.md "Generates an 8-byte strong random number.") helper method.
4. created defines the optional created
   date input.
5. offset defines the number of bytes
   the resulting octet stream must be shifted to obtain the derived key.
6. bytes defines the number of bytes of
   the resulting derived key.

## Usage

If it is a symmetric key, the size can be 0, or must match the original key depending on the
[identifier](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.") of the key type.

See [Derived keys](4223-derived-keys.md) for more
details.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
