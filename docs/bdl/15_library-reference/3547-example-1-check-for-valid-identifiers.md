---
title: "Example 1: Check for valid identifiers"
source: "fgl-topics/c_fgl_ext_util_Regexp_example_1.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > Examples > Example 1: Check for valid identifiers"
type: "concept"
description: "The following program compiles a typical regular expression that checks for valid identifiers, starting with a letter or underscore, and then containing more letters, underscore or digits: IMPORT util ..."
---

# Example 1: Check for valid identifiers

The following program compiles a typical regular expression that checks for valid identifiers,
starting with a letter or underscore, and then containing more letters, underscore or digits:

```
IMPORT util
  
MAIN
    CALL testASCII("1_ident")
    CALL testASCII("ident_1")
    CALL testASCII("A_é")
    CALL testAnyChar("A_é")
END MAIN

FUNCTION testASCII(s STRING)
    DISPLAY s, COLUMN 30, IIF(isIdentifierASCII(s),"TRUE","FALSE")
END FUNCTION

FUNCTION testAnyChar(s STRING)
    DISPLAY s, COLUMN 30, IIF(isIdentifierAnyChar(s),"TRUE","FALSE")
END FUNCTION

PRIVATE DEFINE re1 util.Regexp
PRIVATE DEFINE re2 util.Regexp

PUBLIC FUNCTION isIdentifierASCII(s STRING) RETURNS BOOLEAN
    IF re1 IS NULL THEN
        LET re1 = util.Regexp.compile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
    END IF
    RETURN re1.matches(s)
END FUNCTION

PUBLIC FUNCTION isIdentifierAnyChar(s STRING) RETURNS BOOLEAN
    IF re2 IS NULL THEN
        LET re2 = util.Regexp.compile(`^[\p{L}_][\p{L}_\d]*$`)
    END IF
    RETURN re2.matches(s)
END FUNCTION
```

Ouput:

```
1_ident                      FALSE
ident_1                      TRUE
A_é                          FALSE
A_é                          TRUE
```
