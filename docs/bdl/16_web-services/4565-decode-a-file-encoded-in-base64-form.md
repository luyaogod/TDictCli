---
title: "Decode a file encoded in BASE64 form"
source: "fgl-topics/c_gws_ssl_fglpass_010.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Decode a file encoded in BASE64 form"
type: "concept"
---

# Decode a file encoded in BASE64 form

> The fglpass tool can decode a BASE64 encoded file.

1. To decode a file encoded in BASE64 form,
   enter:

   ```
   fglpass -dec64 Base64filename
   ```

   The fglpass tool
   outputs the file in clear form on the console.

   ```
   security.global.agent      = "4242"
   crypto.id1.key             = "RSA1024Key.pem"
   crypto.id2.key             = "RSA2048Key.pem"
   crypto.id3.key             = "DSA1024Key.pem"
   crypto.id4.key             = "RSA512Protected.pem"
   crypto.id5.key             = "DSA512ReallyProtected.pem"
   ```

   - You don't have to remove the `BASE64 BEGIN` and `BASE64 END` tags,
     if they are present in the file, because the fglpass tool detects and removes
     them automatically.
   - You can redirect the output of the fglpass tool to a file. For
     example:

     ```
     fglpass -dec64 Base64MyFile > MyFile2
     ```

## Related links

**Related concepts**  

[Encode a file in BASE64 form](4564-encode-a-file-in-base64-form.md "The fglpass tool can encode a file in BASE64 form.")
