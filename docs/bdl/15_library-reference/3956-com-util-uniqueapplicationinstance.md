---
title: "com.Util.UniqueApplicationInstance"
source: "fgl-topics/c_gws_ComUtil_UniqueApplicationInstance.html"
breadcrumb: "Library reference > Extension packages > The com package > Helper classes > The Util class > Util methods > com.Util.UniqueApplicationInstance"
type: "concept"
---

# com.Util.UniqueApplicationInstance

> Checks that the calling application is the only one to run.

## Syntax

```
UniqueApplicationInstance(
   file STRING)
  RETURNS INTEGER
```

1. file specifies the string for the lock file.

## Usage

This method checks that the calling application is the only one to run, by trying to get an
exclusive lock on the given file.

If the lock could be set, the method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") (1); otherwise, it returns [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") (0) and updates `status` with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number.").
