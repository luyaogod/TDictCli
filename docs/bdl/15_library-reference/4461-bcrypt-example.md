---
title: "Example: Using security.BCrypt methods"
source: "fgl-topics/c_gws_SecurityBCrypt_example.html"
breadcrumb: "Library reference > Extension packages > The security package > The BCrypt class > BCrypt example"
type: "concept"
---

# Example: Using security.BCrypt methods

> This example creates (and checks) a hash password as BCrypt results.

```
IMPORT security

MAIN
   DEFINE salt STRING
   DEFINE hashed_pass STRING
   LET salt = security.BCrypt.GenerateSalt(12)
  
   CALL security.BCrypt.HashPassword(arg_val(1), salt) RETURNING hashed_pass
   DISPLAY "Hashed password is :",hashed_pass  
  
   IF security.BCrypt.CheckPassword(arg_val(1), hashed_pass) THEN
     DISPLAY "OK: password check done."
   ELSE
     DISPLAY "KO: password check failed."
   END IF
END MAIN
```
