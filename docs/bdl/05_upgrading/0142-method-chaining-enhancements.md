---
title: "Method chaining enhancements"
source: "fgl-topics/c_fgl_Migrate_to_400_chained_calls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Method chaining enhancements"
type: "concept"
---

# Method chaining enhancements

> Values returned from functions can be used to chain with method calls.

Genero BDL version 3.10 introduced method chained calls like this:

```
MAIN
    DEFINE s STRING
    LET s = "ABC"
    DISPLAY s.subString(1,2).toLowerCase()
END MAIN
```

Starting with version 4.00, it is now possible to chain method calls, with the return value of a
function:

```
MAIN
    DISPLAY foo().subString(1,2).toLowerCase()
END MAIN

FUNCTION foo() RETURNS STRING
    RETURN "abc"
END FUNCTION
```

Example using a `DICTIONARY OF STRING`:

```
MAIN
    DISPLAY foo().getLength()
    DISPLAY foo()["abc"].toUpperCase()
END MAIN

FUNCTION foo() RETURNS DICTIONARY OF STRING
    DEFINE dic DICTIONARY OF STRING
    LET dic["abc"] = "xxxx"
    LET dic["def"] = "yyyyyy"
    RETURN dic
END FUNCTION
```

Using built-in functions like `upshift()`:

```
MAIN
    DISPLAY upshift("abc").getLength()
END MAIN
```

Using `base.Channel`:

```
MAIN
    DISPLAY my_open("file1").readLine()
    DISPLAY my_open("file2").readLine()
END MAIN

FUNCTION my_open(fn STRING) RETURNS base.Channel
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    TRY
        CALL ch.openFile(fn,"r")
    END TRY
    RETURN ch
END FUNCTION
```

## Related links

**Related concepts**  

[OOP support](../09_advanced-features/0942-oop-support.md "Describes Object Oriented Programming basics in the language.")
