---
title: "Case sensitivity with Java"
source: "fgl-topics/c_fgl_JavaBridge_017.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Case sensitivity with Java"
type: "concept"
description: "The Java language is case-sensitive . Therefore, when you write the name of a Java package, class or method in a .4gl source, it must match the exact name as if you were writing a Java program. The ..."
---

# Case sensitivity with Java

The Java language is case-sensitive. Therefore, when you write the name of a Java package,
class or method in a .4gl source, it must match the exact name as if you were
writing a Java program. The [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler takes care
of this, and writes case-sensitive class and method names in the .42m p-code
modules.

```
IMPORT JAVA java.util.regex.Pattern
MAIN
  DEFINE p java.util.regex.PATTERN   -- Note the case error
END MAIN
| symbol 'PATTERN' not found in package 'regex'.
| See error number -8447.
```

With this code example, fglcomp will raise error [-8447](../15_library-reference/4483-genero-bdl-errors.md) because
"java/util/PATTERN" name cannot be found.
