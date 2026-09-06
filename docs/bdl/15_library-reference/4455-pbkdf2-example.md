---
title: "Example: Using security.PBKDF2 methods"
source: "fgl-topics/c_gws_SecurityPBKDF2_example.html"
breadcrumb: "Library reference > Extension packages > The security package > The PBKDF2 class > PBKDF2 example"
type: "concept"
---

# Example: Using security.PBKDF2 methods

> This example generates a key size of 128-bits based on a given password.

The 128-bits key can then be used in `xml.CryptoKey`, for instance.

```
IMPORT security

MAIN 
  DEFINE salt     STRING
  DEFINE result  STRING
  LET salt =   security.RandomGenerator.CreateRandomString(8)
  TRY
    CALL security.PBKDF2.GenerateKey(arg_val(1), salt, "sha1", 1000, 16) RETURNING result
    DISPLAY "Generate Key of 128bits value is :",result  # 128/8==16
  CATCH
       IF status == -15700 THEN
         DISPLAY "Generation failed :",sqlca.sqlerrm
       ELSE
         IF status == -15701 THEN
           DISPLAY "Invalid parameter :",sqlca.sqlerrm
         ELSE
           DISPLAY "Unkown error :",status
         END IF
       END IF
  END TRY
END MAIN
```
