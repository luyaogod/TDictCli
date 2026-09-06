---
title: "Using the BYTE type"
source: "fgl-topics/c_fgl_JavaBridge_028.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Mapping native and Java data types > Using the BYTE type"
type: "concept"
description: "When calling a Java method with an expression evaluating to a BYTE , the runtime system converts the BYTE handle to an instance of the com.fourjs.fgl.lang.FglByteBlob class implemented in ..."
---

# Using the BYTE type

When calling a Java method with an expression evaluating to a [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds."), the runtime
system converts the BYTE handle to an instance of the
`com.fourjs.fgl.lang.FglByteBlob` class implemented
in $FGLDIR/lib/fgl.jar. You can then manipulate
the LOB from within the Java code.

You must add $FGLDIR/lib/fgl.jar to the class
path in order to compile Java code with `com.fourjs.fgl.lang.FglByteBlob` class.

The `com.fourjs.fgl.lang.FglByteBlob` class implements the following:

| Method | Description |
| --- | --- |
| void dispose() | Dereferences the underlying `BYTE` variable. |
| java.lang.String toString() | Returns the HEX string representing the binary data. |
| static FglByteBlob valueOf( java.lang.String val) | Creates a new `FglByteBlob` object from a String object representing the binary data in HEX format. |

In the Java code, you can
pass a `com.fourjs.fgl.lang.FglByteBlob` object
as in this example:

```
public static void useByte(FglByteBlob b) throws ParseException {
    String s = b.toString();
    ...
}
```

If you need to create a `com.fourjs.fgl.lang.FglByteBlob` object in your program,
you can use the `valueOf()` class method as in this
example:

```
IMPORT JAVA com.fourjs.fgl.lang.FglByteBlob
MAIN
  DEFINE jbyte com.fourjs.fgl.lang.FglByteBlob
  LET jbyte = FglByteBlob.valueOf("0FA5617BDE")
  DISPLAY jbyte.toString()
END MAIN
```
