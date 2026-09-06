---
title: "Ignorable return of Java methods"
source: "fgl-topics/c_fgl_JavaBridge_021.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Ignorable return of Java methods"
type: "concept"
description: "Java allows you to ignore the return value of a method (as in C/C++): StringBuffer sb = new StringBuffer; sb.append(\"abc\"); -- returns a new StringBuffer object but is ignored In programs, you can ..."
---

# Ignorable return of Java methods

Java allows you to ignore the return value of a method (as in
C/C++):

```
StringBuffer sb = new StringBuffer; 
sb.append("abc");  -- returns a new StringBuffer object but is ignored
```

In programs, you can call a Java method and ignore the return
value:

```
IMPORT JAVA java.lang.StringBuffer
MAIN
  DEFINE sb StringBuffer
  LET sb = StringBuffer.create()
  LET sb = sb.append("abc")
  CALL sb.append("def")  -- typical usage
END MAIN
```
