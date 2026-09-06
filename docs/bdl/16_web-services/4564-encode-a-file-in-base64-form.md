---
title: "Encode a file in BASE64 form"
source: "fgl-topics/c_gws_ssl_fglpass_009.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Encode a file in BASE64 form"
type: "concept"
---

# Encode a file in BASE64 form

> The fglpass tool can encode a file in BASE64 form.

1. To encode the file MyFile in BASE64,
   enter:

   ```
   fglpass -enc64 MyFile
   ```

   The fglpass tool outputs the BASE64 form of the file to the
   console.

   ```
   BASE64 BEGIN
   c2VjdXJpdHkuZ2xvYmFsLmFnZW50ICAgICAgPSAiNDI0MiINCmNyeXB0by
   5pZDEua2V5ICAgICAgICAgICAgID0gIlJTQTEwMjRLZXkucGVtIg0KY3J5
   cHRvLmlkMi5rZXkgICAgICAgICAgICAgPSAiUlNBMjA0OEtleS5wZW0iDQ
   pjcnlwdG8uaWQzLmtleSAgICAgICAgICAgICA9ICJEU0ExMDI0S2V5LnBl
   bSINCmNyeXB0by5pZDQua2V5ICAgICAgICAgICAgID0gIlJTQTUxMlByb3
   RlY3RlZC5wZW0iDQpjcnlwdG8uaWQ1LmtleSAgICAgICAgICAgICA9ICJE
   U0E1MTJSZWFsbHlQcm90ZWN0ZWQucGVtIg0K
   BASE64 END
   ```

   - The BASE64 encoded file is the string between `BASE64 BEGIN` and `BASE64
     END`.
   - You can redirect the output of fglpass tool to a file. For
     example:

     ```
     fglpass -enc64 MyFile > Base64filename
     ```

## Related links

**Related concepts**  

[Decode a file encoded in BASE64 form](4565-decode-a-file-encoded-in-base64-form.md "The fglpass tool can decode a BASE64 encoded file.")
