---
title: "The Digest class"
source: "fgl-topics/c_gws_SecurityDigest.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class"
type: "concept"
---

# The Digest class

> The security.Digest class implements digest algorithms to process data.

This class is provided in the `security` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`security` package with:

```
IMPORT security
```

The class implements several methods that allow you to add data piece by piece and process
these data with a specified digest algorithm.

Steps to process data with a digest algorithm:

1. Define the digest algorithm with the [`security.Digest.CreateDigest`](4446-security-digest-createdigest.md "Defines a new digest context by specifying the algorithm to be used.")
   method.
2. Add data to the digest buffer with methods such as [`security.Digest.AddData`](4442-security-digest-adddata.md "Adds data from a BYTE variable to the digest buffer."), [`security.Digest.AddBase64Data`](4441-security-digest-addbase64data.md "Adds data in base64 format to the digest buffer."),
   [`security.Digest.AddHexBinaryData`](4443-security-digest-addhexbinarydata.md "Adds data in hexadecimal format to the digest buffer."),
   [`security.Digest.AddStringData`](4444-security-digest-addstringdata.md "Adds a data string to the digest buffer.").
3. When all data pieces are added, the buffer can be processed by calling
   methods like [`security.Digest.DoBase64Digest`](4448-security-digest-dobase64digest.md "Creates a digest of the buffered data and returns the result in base64 format.") or
   [`security.Digest.DoHexBinaryDigest`](4449-security-digest-dohexbinarydigest.md "Creates a digest of the buffered data and returns the result in hexadecimal format.").

Alternatively, a simple data string can be processed with the [`security.Digest.CreateDigestString`](4447-security-digest-createdigeststring.md "Creates a SHA1 digest from the given string.")
method.

> **Note:**
>
> MD5 digest is not recommended and may not be available on the system running Genero due to
> certain security configurations or the version of OpenSSL used.

## Child topics

- [security.Digest methods](4440-digest-methods.md): Methods of the security.Digest class.
- [Example](4450-example.md): Computing a hash value of a string.
