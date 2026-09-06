---
title: "INSTANCEOF"
source: "fgl-topics/c_fgl_operators_INSTANCEOF.html"
breadcrumb: "Language basics > Operators > List of expression elements > Data type operators > INSTANCEOF"
type: "concept"
---

# INSTANCEOF

> The INSTANCEOF checks the class of an object.

## Syntax

```
variable INSTANCEOF class
```

1. variable is a variable referencing a Java object.
2. class is a Java class.

## Usage

The `INSTANCEOF` operator evaluates to
`TRUE` if the object reference is of the specified class.

## Example

```
IMPORT JAVA java.lang.Object 
IMPORT JAVA java.lang.StringBuffer 
IMPORT JAVA java.lang.Number 
MAIN
  DEFINE o java.lang.Object 
  DEFINE sb java.lang.StringBuffer 
  LET sb = StringBuffer.create()
  LET o = sb 
  DISPLAY sb INSTANCEOF java.lang.StringBuffer  -- shows 1
  DISPLAY o INSTANCEOF java.lang.StringBuffer   -- shows 1
  DISPLAY o INSTANCEOF java.lang.Number         -- shows 0
END MAIN
```

## Related links

**Related concepts**  

[The CAST operator](../14_extending-the-language/2689-the-cast-operator.md "The CAST operator")

[The Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
