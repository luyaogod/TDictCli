---
title: "xml.Signature.appendReferenceTransformation"
source: "fgl-topics/c_gws_XmlSignature_appendReferenceTransformation.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.appendReferenceTransformation"
type: "concept"
---

# xml.Signature.appendReferenceTransformation

> Appends transformations related to the specified reference index.

## Syntax

```
appendReferenceTransformation(
   referenceIndex INTEGER,
   method STRING
   [, args ] )
```

1. referenceIndex defines the index in this Signature object.
2. method represents a URL as [transformation identifier](4296-transformation-identifier.md) of the
   transformation algorithm.
3. args is a variable list of parameters, to define transformation options,
   depending on the method.

## Usage

This method appends a reference transformation that is executed before any computation.

A transformation modifies the reference URI before signing or validating it. Several
transformations are executed one after another, and only once the last transformation has
been applied, is the reference really signed or verified.

Depending on the [transformation
identifier](4296-transformation-identifier.md), additional parameters are necessary. These can be passed as a regular,
comma-separated list of method parameters.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
