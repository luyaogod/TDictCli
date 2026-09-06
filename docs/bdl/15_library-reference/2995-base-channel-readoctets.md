---
title: "base.Channel.readOctets"
source: "fgl-topics/c_fgl_ClassChannel_readOctets.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.readOctets"
type: "concept"
---

# base.Channel.readOctets

> Read a given number of bytes and return as a character string.

## Syntax

```
readOctets(
   length INTEGER )
  RETURNS STRING
```

1. length is the number of bytes to read, not the number of characters.

## Usage

After opening the channel object, call the `readOctets()` method to read a given
number of bytes from the channel. The bytes are returned as a character string.

> **Important:**
>
> The bytes read with `readOctets()` must match the [application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."). In multibyte encoding (UTF-8), if a
> sequence of bytes does represent a valid character, the resulting string will get ? question mark
> for invalid characters and data will be lost.

A valid use case of the method is the HTTP protocol. Reading HTML content with
`readLine()` is not possible. The body consists of multiple lines, and the last line
might not be terminated by a line-terminator, and the stream gets no EOF:

```
HTTP/1.0 200 OK
Date:  Wed, 16 Apr 2014 18:50:51 GMT
Content-Type: text/html
Content-Length: 1354

<html>
<body>
<h1>My title</h1>
  :
</body>
</html>
```

The `readOctets()` function returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") if end of file is reached.

> **Note:**
>
> To distinguish empty lines from `NULL`, you must use the
> `STRING` data type. If you use a `CHAR` or `VARCHAR`,
> you will get `NULL` for empty lines. To properly detect end of file, use the [`isEof()`](2988-base-channel-iseof.md "Detect the end of a file.") method.

If `readOctets()` cannot read as many bytes as specified by the parameter, the
method returns `NULL`. Before reading the actual bytes with a
`readOctets()` call, find the number of data bytes to read from stream source, as
shown in the example below.

Error [-6346](4483-genero-bdl-errors.md) is thrown, if the channel fails to read data.

## Example

> **Note:**
>
> This example uses UTF-8 encoding.

The main program:

```
MAIN
    DEFINE len INTEGER
    DEFINE chunk STRING
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    CALL ch.openFile("file.txt","r")
    WHILE TRUE
        LET len = ch.readOctets(3) -- 3 digits for length
        IF ch.isEof() THEN EXIT WHILE END IF
        LET chunk = ch.readOctets(len)
        DISPLAY len USING "##&"," ", NVL(chunk,"(NULL)")
    END WHILE
    CALL ch.close()
END MAIN
```

The data file (contains 3 digits for byte length of following
string):

```
003abc006forêt000001x
```

Program output (note that forêt needs 3+2+1=6 bytes in
UTF-8):

```
$ fglcomp -M readOctets.4gl  && fglrun readOctets.42m
  3 abc
  6 forêt
  0 
  1 x
```

## Related links

**Related concepts**  

[Setup a TCP socket channel](3007-setup-a-tcp-socket-channel.md "Setup a TCP socket channel")
