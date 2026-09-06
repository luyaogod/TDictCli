---
title: "Configure VIM for Genero BDL"
source: "fgl-topics/c_fgl_CodeEditing_005.html"
breadcrumb: "Programming tools > Source code edition > Configure VIM for Genero BDL"
type: "concept"
description: "The VIM editor VIM is a powerful source code editor well known by Unix programmers. VIM is also available on Microsoft Windows®. When using VIM, the following features are available with fglcomp and ..."
---

# Configure VIM for Genero BDL

## The VIM editor

VIM is a powerful source code editor well known by Unix programmers. VIM is also available on
Microsoft Windows®.

When using VIM, the following features are available with [fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") and [fglform](2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") compilers:

- Syntax highlighting
- Code completion (TAB while in insert mode)
- Code formatting at different levels (F2, F3,
  F4)
- Go to definition using gd command (gd, CTRL-O,
  CTRL-I)
- Genero BDL jump-to tags (CTRL-], CTRL-T)

To enable these features, you need to setup you VIM configuration file.

> **Important:**
>
> Genero code editing features require VIM version 7 or + (with the Omni Completion
> feature)

## Configuring VIM for Genero BDL

Perform the following steps to enable code-completion for Genero in VIM:

1. Locate the VIM resource file for your operating system user:
   - On UNIX® platforms, the VIM resource file
     is ~/.vimrc.
   - On Windows platforms, the VIM
     resource file is %USERPROFILE%\\_vimrc.
2. Add the following lines to make VIM find Genero VIM files in
   $FGLDIR/vimfiles:

   ```
   let generofiles=expand($FGLDIR . "/vimfiles")
   if isdirectory(generofiles)
      let &rtp=generofiles.','.&rtp
   endif
   ...
   ```
3. Add the next line to enable syntax highlighting in VIM:

   ```
   syntax on
   ```
4. Add the following lines to associate source file extensions to the corresponding syntax
   definition file:

   ```
   autocmd BufNewFile,BufRead *.per setlocal filetype=per
   ```

   In
   fact, the .4gl file type `"fgl"` is usually detected in the
   default configuration files of VIM (like /usr/share/vim/vim80/filetype.vim). If
   this file type is not detected by default by your VIM installation, add the next line in your
   .vimrc
   file:

   ```
   autocmd BufNewFile,BufRead *.4gl setlocal filetype=fgl
   ```
5. Disable case sensitivity for keywords by adding the following
   line:

   ```
   let fgl_ignore_case=1
   ```

   Genero BDL is not case sensitive and allows to
   use language keywords for identifiers (`DEFINE name STRING`). Common style guidelines
   and the default VIM syntax highlighting encourage the use of uppercase keywords. This avoids
   highlighting of keywords used as identifiers. Case sensitive highlighting can be disabled by setting
   `fgl_ignore_case` to 1 in the VIM resource file. When using this option, symbols like
   variable names matching language keywords (`define name string`) will be
   highlighted.
6. To enable lowercase keywords in code completion proposals, add the following
   line:

   ```
   let fgl_lowercase_keywords=1
   ```

   Setting
   `fgl_lowercase_keywords=1` implies implicitly
   `fgl_ignore_case=1`.
7. Add the following lines to define function keys and commands that reformat the code using
   fglcomp --format:
   - F2: formats the whole source file (notice the `%` sign before
     `FglFormat`):

     ```
     autocmd BufNewFile,BufRead *.4gl map <buffer> <F2> :%FglFormat<CR>
     ```
   - F3: formats lines changed since last git commit (requires a git
     repository):

     ```
     autocmd BufNewFile,BufRead *.4gl map <buffer> <F3> :FglGitFormat<CR>
     ```
   - F4: in command mode: formats a range, in edit mode: formats the current
     line:

     ```
     autocmd BufNewFile,BufRead *.4gl map <buffer> <F4> :FglFormat<CR>
     autocmd BufNewFile,BufRead *.4gl imap <buffer> <F4> <c-o> :FglFormat<CR>
     ```
8. To format the source code lines from a given mark to the current line, use this kind of VIM
   command (here using the mark "a"):

   ```
   :'a,.FglFormat
   ```
9. If you want to reformat the source code with [`fglcomp
   --format`](2636-source-code-beautifier.md "Reformat the source code for better readability.") each time you save the file, add the following
   lines:

   ```
   " Formats the file whenever Vim writes it
   autocmd BufWritePost *.4gl call FglWritePostHook()
   function! FglWritePostHook()
      silent! :!fglcomp --format --fo-inplace %
      if v:shell_error
         redraw!
      endif
      edit
   endfun
   ```

   The auto-formatting may not be done, if the source code contains errors. In
   this case the source file is left untouched, but it will be saved.
10. If you want an F6 shortcut to format the function you are
    in:

    ```
    autocmd BufNewFile,BufRead *.4gl map <buffer> <F6>
    \ mx:?^[ \t]*\<[Ff][uU][Nn][Cc][tT][iI][oO][nN]\>?,
    \/^[ \t]*\<[Ee][nN][dD]\>[ \t][ \t]*\<[Ff][Uu][Nn][Cc][Tt][Ii][Oo][Nn]\>
    \/FglFormat<CR>'x
    ```
11. The VIM `gf` command can open another file, from the current edited file, when
    the cursor is on a string that identifies the filename. This command does not recognize
    .4gl sources by default. To open another .4gl file with
    `gf`, when the cursor is on the module prefix of instructions such as `CALL
    module_name.function_name()`, add the following
    lines:

    ```
    set suffixesadd=.4gl

    function! LoadModule(fname)
        let x = stridx(a:fname,'.')
        if x > 1
            let fn = a:fname[0:x-1]
        else
            let fn = a:fname
        endif
        let fn = fn . ".4gl"
        if filereadable(fn)
            return fn
        else
            return a:fname
        endif
    endfunction

    set includeexpr=LoadModule(v:fname)
    ```
12. To make the jump-to-tag CTRL-] / CTRL-T commands
    Genero language specific, add the following
    line:

    ```
    autocmd FileType fgl map <buffer> <C-]> :FglTag<CR>
    ```

## Using VIM on Microsoft Windows platforms

On Windows platforms, you typically
install the GVim (Graphical VIM) software.

Some versions of VIM for Windows
may use different configuration files and locations (\_vimrc or
.vimrc?). Refer to the VIM documentation, to make sure that you use the proper
files.

When using the command line version VIM, you may want to
add:

```
color shine
```

## Using code completion with Genero and VIM

First make sure the Genero environment is set (FGLDIR, PATH).

Open a .4gl or .per file, start to edit the file with
VIM.

On Windows platforms, if you start
GVim from the icon, the Genero environment may not be set. As result
fglcomp/fglform cannot be called from VIM. You need to set the
Genero BDL environment before starting VIM.

When in insert mode, press CTRL-X + CTRL-O, to get
a list of language elements to complete the instruction syntax or expression.

For convenience, `TAB` can also be used to get the completion list as with the
CTRL-X + CTRL-O key combinations. However,
TAB will only do code completion, if the edit cursor is after some non-space
characters. If only spaces are present left of the cursor position, `TAB` adds
indentation characters.

Use the F2 function key to indent/beautify the current
.4gl source code, with the fglcomp --format beautifier
tool.

For more details about VIM, see <http://www.vim.org>.

## Related links

**Related concepts**  

[Source code completion](2543-source-code-completion.md "Source code completion")
