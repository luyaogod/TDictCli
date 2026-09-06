---
title: "xml.KeyStore.AddKey"
source: "fgl-topics/c_gws_XmlKeyStore_AddKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The KeyStore class > KeyStore methods > xml.KeyStore.AddKey"
type: "concept"
---

# xml.KeyStore.AddKey

> Registers in the keystore the given key by name for the application.

## Syntax

```
xml.KeyStore.AddKey( 
   global xml.CryptoX509 )
```

1. global defines the key object `xml.CryptoX509` to add to the
   keystore.

## Usage

This method adds a key to the keystore to be used for XML signature verification or XML
decryption when a key name has been specified in the XML [KeyInfo](http://www.w3.org/TR/xmldsig-core/#sec-KeyInfo) node and no
other key has been set in the `Signature` or `Encryption` object.

The CryptoKey must have the [KeyName](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
feature set, and the name must be unique in the application.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
