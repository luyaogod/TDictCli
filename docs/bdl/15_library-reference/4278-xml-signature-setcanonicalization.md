---
title: "xml.Signature.setCanonicalization"
source: "fgl-topics/c_gws_XmlSignature_setCanonicalization.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setCanonicalization"
type: "concept"
---

# xml.Signature.setCanonicalization

> Sets the canonicalization method to use for the signature.

## Syntax

```
setCanonicalization(
   url STRING )
```

1. url defines one of the four canonicalization [identifier](4296-transformation-identifier.md).

## Usage

This method sets the canonicalization method to use for the signature. The default value is the
[c14n](4296-transformation-identifier.md)method.

> **Note:**
>
> Windows® .NET default [c14n](4296-transformation-identifier.md) canonicalization method is not compatible with the W3C standard, therefore it is
> recommended to use the [exc-c14n](4296-transformation-identifier.md) method when interoperating with a Windows
> system.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
