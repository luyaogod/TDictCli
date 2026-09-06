---
title: "Example"
source: "fgl-topics/c_gws_security_library_example.html"
breadcrumb: "Library reference > Extension packages > The security package > The Digest class > Example"
type: "concept"
---

# Example

> Computing a hash value of a string.

Program example ComputeHash.4gl :

```
IMPORT SECURITY

MAIN

  DEFINE result STRING

  IF num_args() != 2 THEN
    DISPLAY "Usage: ComputeHash <string> <hashcode>"
    DISPLAY "  string: the string to digest"
    DISPLAY "  hashcode: SHA1, SHA512, SHA384, SHA256, SHA224, MD5"
  ELSE
    LET result = ComputeHash(arg_val(1), arg_val(2))
    IF result IS NOT NULL THEN
      DISPLAY "Hash value is :",result
    ELSE
      DISPLAY "Error"
    END IF
  END IF

END MAIN

FUNCTION ComputeHash(toDigest, algo)

  DEFINE toDigest, algo, result STRING
  DEFINE dgst security.Digest

  TRY
    LET dgst = security.Digest.CreateDigest(algo)
    CALL dgst.AddStringData(toDigest)
    LET result = dgst.DoBase64Digest()
  CATCH
    DISPLAY "ERROR : ", status, " - ", sqlca.sqlerrm
    EXIT PROGRAM(-1)
  END TRY

  RETURN result
END FUNCTION
```

Example execution:

$ fglrun ComputeHash "Hello World" SHA1

Hash value is: Ck1VqNd45QIvq3AZd8XYQLvEhtA=
