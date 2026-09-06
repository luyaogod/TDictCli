---
title: "File I/O statements and APIs"
source: "fgl-topics/c_fgl_MigI4GL_030.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > File I/O statements and APIs"
type: "concept"
---

# File I/O statements and APIs

> Both I4GL and Genero BDL support accessing files on operating systems. The base.Channel built-in class supported by Genero BDL can also open streams to subprocesses and sockets.

IBM®
Informix® 4GL version **7.50.xC4** introduced file
manipulation instructions to access files on the operating system running the application. These
instructions can be used to open, read, write, seek and close
files:

```
MAIN
   DEFINE fd1, fd2 INTEGER, v1,v2 VARCHAR(10)
   OPEN FILE fd1 FROM "/tmp/file1" OPTIONS (READ, FORMAT="CSV")
   OPEN FILE fd2 FROM "/tmp/file2" OPTIONS (WRITE, APPEND, CREATE, FORMAT="CSV")
   READ FROM fd1 INTO v1, v2
   SEEK ON fd2 TO 0 FROM LAST INTO v1
   WRITE TO fd2 USING v1, v2
   CLOSE FILE fd1
   CLOSE FILE fd2
END MAIN
```

Genero Business Development Language (BDL) implements file I/O support with the
*base.Channel* built-in class. This class implements file access, but it can
also open streams to subprocesses (i.e. pipes) and sockets.

## Related links

**Related concepts**  

[The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
