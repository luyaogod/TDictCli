---
title: "Import a Java class"
source: "fgl-topics/c_fgl_JavaBridge_010.html"
breadcrumb: "Extending the language > The Java interface > Getting started with the Java interface > Import a Java class"
type: "concept"
description: "In order to use a Java class in your program code, you must first import the class with the IMPORT JAVA instruction: IMPORT JAVA java.util.regex.Pattern This will import the specified Java class into ..."
---

# Import a Java class

In order to use a Java class in your program code, you must first import the class with the
[`IMPORT JAVA`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instruction:

```
IMPORT JAVA java.util.regex.Pattern
```

This will import the specified Java class into the current program module. Object references
can now be defined for this class.

A Java inner class can be referenced by importing the outer class:

```
IMPORT JAVA java.lang.Character
MAIN
    DISPLAY Character.UnicodeBlock.ARROWS
    DISPLAY Character.UnicodeScript.BRAILLE
END MAIN
```
