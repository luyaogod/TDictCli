---
title: "os.Path.glob"
source: "fgl-topics/c_fgl_ext_os_path_glob.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.glob"
type: "concept"
---

# os.Path.glob

> Returns a list of files matching the specified pattern.

## Syntax

```
os.Path.glob( pattern STRING )
  RETURNS DYNAMIC ARRAY OF STRING
```

1. pattern is a string representing file names, using wildcard characters
   `*`, `?` and `[]`as the `MATCHES`
   operator.

## Usage

The glob() method produces a `DYNAMIC ARRAY OF STRING` containing a list of
pathnames of the file system that match the pattern provided as parameter.

If no file or directory path matches the pattern, the `glob()` method returns an
empty array.

The pattern can contain wildcard as follows:

- `?` : matches any single char at this position (except the path separator).
- `*` : matches any string (not including the path separator).
- `[ set-of-chars ]` : matches any character enclosed by the
  brackets. Can define a range of characters with the hyphen (
  `[a-z]` ). The matching is inverted when starting with a
  `'^'`: `[^0-9]` matches any non
  digit.

Globbing is applied on each of the components of a pathname separately. The path separator (Unix
'/', Windows '\' or '/') splits the pattern in subpatterns. A pattern in the form
`"pattern1/pattern2"` finds files matching `pattern2` in each
directory matching `pattern1`. Thus `glob()` limits the result to
directory-names if the pattern ends with a path separator.

| glob pattern | Description |
| --- | --- |
| `*.4gl` | Files in current dir ending with '.4gl' |
| `m*` | Files in current dir starting with the character 'm' |
| `[Mm]akefile` | The file Makefile or makefile in current dir |
| `[0-9]*` | Files in current dir starting with a digit |
| `subdir/*.4gl` | Files ending with '.4gl' in the directory 'subdir' |
| `subdir[12]/*.4gl` | Files ending with '.4gl' in subdir1 or subdir2 |
| `../*.4gl` | Files ending with '.4gl' in the parent directory |
| `/Users/Me/*.4gl` | Files ending with '.4gl' in the directory 'Users/Me' |
| `*/` | All subdirectories (at level 1) |

## Example

```
IMPORT os
MAIN
    DEFINE files DYNAMIC ARRAY OF STRING
    DEFINE x INTEGER
    LET files = os.Path.glob("[A-Z]*.4gl")
    FOR x=1 TO files.getLength()
         DISPLAY "File: ", files[x]
    END FOR
END MAIN
```
