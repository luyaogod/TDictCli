---
title: "Example 1: Split a directory path"
source: "fgl-topics/c_fgl_ClassStringTokenizer_example_1.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > Examples > Example 1: Split a directory path"
type: "concept"
description: "Note: These examples are provided to illustrate the base.StringTokenizer class. For file system path management, consider using the os.Path extension class . Split a UNIX™ directory path using slash ..."
---

# Example 1: Split a directory path

> **Note:**
>
> These examples are provided to illustrate the `base.StringTokenizer` class. For
> file system path management, consider using the [`os.Path` extension class](3697-the-os-path-class.md "The os.Path class provides functions to manipulate files and directories on the machine where the program executes.").

## Split a UNIX™ directory path using slash

Program
code:

```
MAIN
  DEFINE tok base.StringTokenizer 
  LET tok = base.StringTokenizer.create("/home/tomy","/")
  WHILE tok.hasMoreTokens()
    DISPLAY tok.nextToken()
  END WHILE
END MAIN
```

Output:

```
home
tomy
```

## Split a Microsoft™ Windows® directory path using backslash

Note that you must escape the backslash:

```
MAIN
  DEFINE tok base.StringTokenizer 
  LET tok = base.StringTokenizer.create("C:\\My Documents\\My Pictures","\\")
  WHILE tok.hasMoreTokens()
    DISPLAY tok.nextToken()
  END WHILE
END MAIN
```

Output:

```
C:
My Documents
My Pictures
```
