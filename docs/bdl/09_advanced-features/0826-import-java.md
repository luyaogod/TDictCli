---
title: "IMPORT JAVA"
source: "fgl-topics/c_fgl_programs_IMPORT_JAVA.html"
breadcrumb: "Advanced features > Importing modules > IMPORT JAVA"
type: "concept"
---

# IMPORT JAVA

> The IMPORT JAVA instruction imports Java module elements.

## Syntax

```
IMPORT JAVA [ packagename . [...] ] classname
```

1. packagename and classname define the Java class to be
   imported.

## Usage

The `IMPORT JAVA` instruction can be used to import a Java class.

At runtime, the imported Java classes are only loaded on demand, when the program flow reaches an
instruction that uses the class. For example, when reaching the declaration of a variable defined to
reference an object of a Java class.

The name specified after the `IMPORT JAVA` instruction is
case-sensitive.

The CLASSPATH environment variable defines the directories for
Java packages.
See the Java documentation for more
details.

To mimic the Java import rules, the fglcomp compiler allows subsequent
`IMPORT JAVA` instructions with the same class name. However, it is recommended that
you avoid duplicating the same `IMPORT JAVA`
instructions:

```
IMPORT JAVA java.util.regex.Matcher
...
IMPORT JAVA java.util.regex.Matcher
```

## Related links

**Related concepts**  

[The Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
