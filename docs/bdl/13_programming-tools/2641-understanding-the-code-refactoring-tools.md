---
title: "Understanding the code refactoring tools"
source: "fgl-topics/c_fgl_refactor_basics.html"
breadcrumb: "Programming tools > Source refactoring > Understanding the code refactoring tools"
type: "concept"
---

# Understanding the code refactoring tools

> This is an introduction to the source code refactoring tools.

For backward compatibility, Genero BDL is a permissive language. For example, keywords and
symbols are not case-sensitive by default. However, if programmers don't follow a strict coding
convention, permissive languages typically make the source difficult to read and
maintain:

```
-- Bad programming, but it compiles...
MAIN
    DEFINE CustId INTEGER
    LET custId = 998
    select custname from Customer WHERE customer_num = custid
    LET CUSTID = 999
    SELECT custname FROM CUSTOMER WHERE customer_num = CUSTID
END MAIN
```

The Genero BDL compiler provides several source refactoring options, to increase code readability
and maintainability.

## Related links

**Related concepts**  

[Source code beautifier](2636-source-code-beautifier.md "Reformat the source code for better readability.")
