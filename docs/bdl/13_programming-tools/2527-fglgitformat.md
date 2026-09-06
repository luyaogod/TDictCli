---
title: "fglgitformat"
source: "fgl-topics/c_fgl_tools_fglgitformat.html"
breadcrumb: "Programming tools > Command reference > fglgitformat"
type: "concept"
---

# fglgitformat

> The fglgitformat tool reformats the source lines changed in a GIT history.

## Syntax

```
fglgitformat [options] [commit] [ file [...] ]
```

1. commit is a GIT commit specification (like `HEAD~2`). Uses
   `HEAD` by default.
2. file is a filename to process. When no file is specified, reformat changes
   for all modified files.
3. options are described in Table 1.

## Options

| Option | Description |
| --- | --- |
| `-h` or `--help` | Displays options for the tool. |
| `-i` or `--inplace` | Reformats the source files instead of writing the diff with reformated code to stdout. |
| `-v` or `--verbose` | Output the new diff to stdout. Ineffective without `-i/--inplace`. |

## Usage

The fglgitformat command line tool reformats changes registered in a GIT
project history, by using [the formatter (`fglcomp
--format`)](2638-syntax-of-the-code-beautifier-tool.md "The source code beautifier option of fglcomp is --format.") to beautify the modified lines.

By default, the diff text is written to the (stdout) standard output stream.
Use the `--inplace` option to modify directly the source files referenced by the GIT
diff. Optionally, combine `--inplace` with `--verbose` to print the
changes to stdout:

```
fglgitformat --inplace --verbose
```

When specifying a commit such as `HEAD~2`, the tool reformats
all changes relative to the commit, plus the changes staged for the next
commit:

```
fglgitformat --inplace HEAD~2
```

When no commit is provided, fglgitformat only reformats the
changes staged for the next commit (relative to the `HEAD`).

By default, all modified .4gl source files are processed. To reformat the
changes in specific files, these can be provided as
arguments:

```
fglgitformat --inplace common/sql_utils.4gl common/gui_utils.4gl
```

> **Tip:**
>
> If the code modifications to be reformatted are part of the last commit, after
> reformatting with fglgitformat, consider using git commit
> --amend, to merge the reformatting with your last commit. Or, commit the reformatting as a
> separate commit, and reorganize the history with interactive rebase (git rebase
> -i).

## Related links

**Related concepts**  

[Source code beautifier](2636-source-code-beautifier.md "Reformat the source code for better readability.")

[fglformatdiff](2528-fglformatdiff.md "The fglformatdiff tool reformats source lines of a unified diff text.")
