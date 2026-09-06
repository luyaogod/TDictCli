---
title: "File inclusion"
source: "fgl-topics/c_fgl_preprocessor_004.html"
breadcrumb: "Programming tools > Source preprocessor > File inclusion"
type: "concept"
---

# File inclusion

> The &include directive instructs the preprocessor to include a file.

## Syntax

```
&include "filename"
```

1. filename is the file to be included during preprocessing.
   filename can be a relative or an absolute path. To specify a path, the slash
   (`/`) directory separator can be used for UNIX™
   and Windows® platforms.

## Usage

The file to be included is searched first in the directory containing the current file, then in
the directorie(s) provided with [`-I`
option](2564-compilers-command-line-options.md "Preprocessor options can be used with fglcomp and fglform compilers.").

The filename argument can be followed by spaces and comments.

The included file will be scanned and processed before continuing with the rest of
the current file.

Source: File A

```
First line 
&include "B"
Third line
```

Source: File B

```
Second line
```

Result (fglcomp -E
output):

```
& 1 "A"
First line 
& 1 "B"
Second line 
& 3 "A"
Third line
```

These
preprocessor directives inform the compiler of its current location
with special preprocessor comments, so the compiler can provide the
right error message when a syntax error occurs.

The preprocessor-generated comments use the following
format:

```
& number "filename"
```

where:

- number is the current line in the preprocessed file
- filename is the current file name

## Recursive inclusions

Recursive inclusions
are not allowed. Doing so will fail and output an error message.

The
following example is incorrect:

Source: File A

```
&include "B"
```

Source: File B

```
HELLO
&include "A"
```

`fglcomp -M A.4gl
output`

```
B.4gl:2:1:2:1:error:(-8029) Multiple inclusion of the source file 'A'.
```

Including
the same file several times is allowed:

Source: File A

```
&include "B"
&include "B"  -- correct
```

Source: File B

```
HELLO
```

Result (fglcomp -E
output):

```
& 1 "A"
& 1 "B"
HELLO
& 2 "A"
& 1 "B"
HELLO
```

## File path

In the `&include "filename"` macro,
filename can be a relative or an absolute path.

To specify a path, the slash (`/`) directory separator can be used for UNIX and Windows
platforms:

```
&include "../include/common.inc"
```
