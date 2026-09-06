---
title: "fglformatdiff"
source: "fgl-topics/c_fgl_tools_fglformatdiff.html"
breadcrumb: "Programming tools > Command reference > fglformatdiff"
type: "concept"
---

# fglformatdiff

> The fglformatdiff tool reformats source lines of a unified diff text.

## Syntax

```
[ command-producing-diff-output | ] fglformatdiff [options]
```

1. The fglformatdiff command gets input from stdin
   stream.
2. options are described in Table 1.

## Options

| Option | Description |
| --- | --- |
| `-h` or `--help` | Displays options for the tool. |
| `-i` or `--inplace` | Reformats the source files referenced in the diff text instead of writing the diff with reformated code to stdout. |
| `-p=num` or `--strip=num` | Strip the smallest prefix containing num leading slashes from the filename found in the diff data. |
| `-v` or `--verbose` | Output the new diff to stdout. Ineffective without `-i/--inplace`. |
| `-e` or `--assume-name=name` | editor support: assume this name when reading from a temporary file. |
| `-c` or `--cursor=pos` | editor support: define the the cursor position |

## Usage

The fglformatdiff command line tool reads a unified diff text from the
(stdin) standard input stream, calculates what lines have changed and calls
[the formatter (`fglcomp --format`)](2638-syntax-of-the-code-beautifier-tool.md "The source code beautifier option of fglcomp is --format.")
to reformat the modified lines.

> **Tip:**
>
> If GIT is your version control system, use [fglgitformat](2527-fglgitformat.md "The fglgitformat tool reformats the source lines changed in a GIT history.") instead of
> fglformatdiff: fglgitformat is designed for GIT.
> fglformatdiff is provided for other version control systems such as CVS and
> SVN.

The fglformatdiff tool is to be used after modifying sources, and before
committing the changes into the repository of the version control system.

> **Important:**
>
> The diff data must be in unified format and contain only modified lines
> (with no context lines). To get no context lines, use the diff `-U`
> / `--unified` option with the parameter value 0 (zero).

By default, fglformdiff writes the new diff result to the
(stdout) standard output stream, after reformatting the lines provided in the
input diff text. Use the `--inplace` option to modify directly the source files
referenced by the diff text. Additionally, you can use the `--verbose` option to
print the changes to
stdout:

```
svn diff --diff-cmd=diff -x-U0 | fglformatdiff --inplace --verbose
```

Use the `--strip=num` option to remove leading path elements of
the filenames of the diff text. For example, when in the root directory or a GIT project, a
git diff will produce filenames with a leading `a/` and
`b/`
prefixes:

```
$ git diff -U0
diff --git a/database/tools/ext.4gl b/database/tools/ext.4gl
...
--- a/database/tools/ext.4gl
+++ b/database/tools/ext.4gl
...
```

To process such diff text with the fglformatdiff, you must use the
`--strip=1`
option:

```
git diff -U0 | fglformatdiff --inplace --strip=1
```

## Related links

**Related concepts**  

[Source code beautifier](2636-source-code-beautifier.md "Reformat the source code for better readability.")

[fglgitformat](2527-fglgitformat.md "The fglgitformat tool reformats the source lines changed in a GIT history.")
