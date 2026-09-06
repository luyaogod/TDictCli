---
title: "Character string encoding"
source: "fgl-topics/c_fgl_ClassChannel_char_encoding.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Character string encoding"
type: "concept"
description: "The runtime system stores and handles character strings in a given application locale , defined by environment settings. When reading or writing character strings from/to a channel, no character set ..."
---

# Character string encoding

The runtime system stores and handles character strings in a given [application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system."), defined by environment settings.

When reading or writing character strings from/to a channel, no character set conversion is done.
The other end of the channel must use the same encoding as the runtime system.

```
-- This source is using UTF-8 encoding
MAIN
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    CALL ch.openFile("myfile.txt","w")
    CALL ch.writeLine("abcéúíôû")
    CALL ch.close()
END MAIN
```

Compiling the source file and checking the .42m file encoding:

```
$ fglcomp ch.4gl && fglrun ch.42m
$ file myfile.txt
myfile.txt: UTF-8 Unicode text
```
