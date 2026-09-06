---
title: "Read and write record data"
source: "fgl-topics/c_fgl_ClassChannel_readwrite_formatted_data.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Read and write record data"
type: "concept"
description: "When the channel is open, use the read() / write() methods to read and write data records where field values are separated by a delimiter defined by setDelimiter() . The LOAD / UNLOAD SQL instructions ..."
---

# Read and write record data

When the channel is open, use the [`read()`](2993-base-channel-read.md "Reads a list of data delimited by a separator from the channel.")/[`write()`](2997-base-channel-write.md "Writes a list of data delimited by a separator to the channel.") methods
to read and write data records where field values are separated by a delimiter
defined by [`setDelimiter()`](2996-base-channel-setdelimiter.md "Define the value delimiter for a channel.").

The [`LOAD`](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table.")/ [`UNLOAD`](../10_sql-support/1177-unload.md "Copies data from the database tables into a file.") SQL instructions follow the same
formatting rules as the `read()`/`write()` channel methods.

The input or output stream is text data where each line contains the string
representation of a record. Field values are separated by the delimiter character
defined.

For example, a formatted text file looks like this, when using a default pipe
(`|`) delimiter:

```
8712|David|Cosneski|24-12-1978|
3422|Frank|Zapinetti|13-04-1968|
323|Mark|Kelson|03-10-1988|
```

In the serialized data, empty fields (`||`) have a length of zero
and are considered as [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.").

The code in the example reads the above field-formatted data:

```
MAIN
  DEFINE ch base.Channel 
  DEFINE custinfo RECORD
            cust_num INTEGER,
            cust_fname VARCHAR(40),
            cust_lname VARCHAR(40),
            cust_bdate DATE
         END RECORD
  LET ch = base.Channel.create()
  CALL ch.setDelimiter("|")
  CALL ch.openFile("custinfo.txt","r")
  WHILE ch.read(custinfo)
    DISPLAY custinfo.*
  END WHILE
  CALL ch.close()
END MAIN
```

The backslash `\` is the escape character: When writing data with
`write()`, special characters like the backslash, line-feed or the delimiter
character are escaped. When reading data with `read()`, any escaped
`\char` character is converted to `char`.

The following code example writes a single field value where the character string contains a
backslash, the pipe delimiter and a line-feed character. The backslash is also the escape character
for string literals, therefore we need to double the backslash to get a backslash in the string,
while the line-feed character (`<lf>`) is represented by backslash-n
(`\n`) in string literals:

```
CALL ch.setDelimiter("|")
CALL ch.write("aaa\\bbb|ccc\nddd")   -- [aaa<bs>bbb|ccc<lf>ddd]
```

This code will produce the following text file:

```
aaa\\bbb\|ccc\
ddd|
```

When reading such a line back into memory with the `read()` method, all
escaped characters are converted back to the single character. In this example,
`\\` becomes `\`, `\|` becomes
`|` and `\<lf>` becomes `<lf>`.

When using the `read()`/`write()` methods, the escaped
line-feed (LF, \n) characters are written as BS + LF to the output, and when reading with
`read()`, BS + LF are detected and interpreted, to be restored as if the
value was assigned by a `LET` instruction, with the same string used in
the `write()` function.

If you want to write a LF as part of a value, the string must contain the backslash and
line-feed as two independent characters. You need to escape the backslash when you write
the string constant in the .4gl source file.

```
CALL ch.setDelimiter("|")
CALL ch.write("aaa\\\nbbb")  -- [aaa<bs><lf>bbb]
CALL ch.write("ccc\nddd")    -- [aaa<lf>bbb]
```

The code above would generate the following output:

```
aaa\
bbb|
ccc| 
ddd|
```

where the first two lines contain data for the same line, in the meaning of a Channel
record.

When you read these lines back with a `read()` call, you get the following
strings in memory:

```
Read 1: aaa<bs><lf>bbb 
Read 2: ccc 
Read 3: ddd
```

These reads would correspond to the following assignments when
using string constants:

```
LET s = "aaa\\\nbbb"
LET s = "ccc"
LET s = "ddd"
```

Data of a record can also be serialized as CSV (Comma Separated Values), when defining
`"CSV"` as delimiter value:

```
CALL ch.setDelimiter("CSV")
```

This CSV format is similar to the standard channel format, with the following differences:

- Values in the file might be surrounded with double quotes (").
- If a value contains a comma or a NEWLINE, it is not escaped; the value must be quoted
  in the file.
- Double-quote characters in values are doubled in the output file and the output value
  must be quoted.
- Backslash characters are not escaped and are read as is; the value must be quoted.
- Leading and trailing blanks are kept (no truncation).
- No ending delimiter is expected at the end of the record line.
