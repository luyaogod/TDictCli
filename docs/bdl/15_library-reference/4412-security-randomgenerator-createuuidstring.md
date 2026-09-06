---
title: "security.RandomGenerator.CreateUUIDString"
source: "fgl-topics/c_gws_SecurityRandomGenerator_CreateUUIDString.html"
breadcrumb: "Library reference > Extension packages > The security package > The RandomGenerator class > RandomGenerator methods > security.RandomGenerator.CreateUUIDString"
type: "concept"
---

# security.RandomGenerator.CreateUUIDString

> Creates a new universal unique identifier (UUID).

## Syntax

```
security.RandomGenerator.CreateUUIDString()
  RETURNS STRING
```

## Usage

This method generates an universal unique identifier and returns the value as
`STRING`.

The generated string follows the UUID version 4 specification. Version 4 UUIDs represent a
128-bit long value in hexadecimal form:

```
xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
```

Where `x` is any hexadecimal digit and `y` is one of 8, 9, A, or B.
There are 4 hyphen ("-") symbols which make its length equal to 36 characters.

UUIDs are widely used for a variety of purposes:

- to generate unique random id.
- in cryptography and hashing applications.
- in generating random documents, addresses, etc.

> **Note:**
>
> UUIDs are unique across both space and time. In terms of space, they are based on a unique
> value in respect of all UUIDs, combined with the time generated. No centralized authority is
> required to administer them. While the probability that a UUID will be duplicated is not zero, the
> likelihood of it happening is negligible. For more information on the standard specification, see
> [rfc4122](http://tools.ietf.org/html/rfc4122).

> **Note:**
>
> This method replaces `com.Util.CreateUUIDString()`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
