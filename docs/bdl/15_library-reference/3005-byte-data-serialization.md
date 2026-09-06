---
title: "BYTE data serialization"
source: "fgl-topics/c_fgl_ClassChannel_byte_serialization.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > BYTE data serialization"
type: "concept"
description: "When using channels, binary data stored in BYTE variables is serialized in hexadecimal. For example, with the following program: MAIN DEFINE ch base.Channel DEFINE rec RECORD id INT, image BYTE END ..."
---

# BYTE data serialization

When using channels, binary data stored in [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") variables is serialized in hexadecimal.

For example, with the following program:

```
MAIN
    DEFINE ch base.Channel
    DEFINE rec RECORD
               id INT,
               image BYTE
           END RECORD
    LET ch = base.Channel.create()
    CALL ch.openFile("myfile.txt","w")
    LET rec.id = 99999
    LOCATE rec.image IN FILE "mydata.bin"
    CALL ch.write(rec)
    CALL ch.close()
END MAIN
```

On Unix, create the binary file like this:

```
$ echo -n -e '\x8f\x9f\xff' > mydata.bin
```

Then compile and run the
program:

```
$ fglcomp ch.4gl && fglrun ch.42m
```

The resulting text file will contain the following serialized
data:

```
99999|8f9fff|
```
