---
title: "Understanding the preprocessor"
source: "fgl-topics/c_fgl_preprocessor_002.html"
breadcrumb: "Programming tools > Source preprocessor > Understanding the preprocessor"
type: "concept"
---

# Understanding the preprocessor

> This is an introduction to the code preprocessor.

The preprocessor is used to transform your sources before compilation.
It allows you to include other files and to define macros that will be
expanded when used in the source. It behaves similar to the C preprocessor,
with some differences.

> **Important:**
>
> It is recommended to avoid using the preprocessor
> if there is an alternative in the native language. For example, instead
> of defining program constants with an `&define` macro,
> use the [`CONSTANT`](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.")
> instruction. Other language features such as [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") increase code readability and modular
> programming, without the need of a preprocessor.

The preprocessor transforms files as follows:

- The source file is read and split into lines.
- Continued lines are merged into one long line if it is part of
  a preprocessor definition.
- Comments are not removed unless they appear in a macro definition.
- Each line is split into a list of lexical tokens.

The preprocessor options can be passed as [compilers
command options](2564-compilers-command-line-options.md "Preprocessor options can be used with fglcomp and fglform compilers.").

The preprocessor implements the following features:

1. [File inclusion](2565-file-inclusion.md "The &include directive instructs the preprocessor to include a file.")
2. [Conditional compilation](2566-conditional-compilation.md "Integrate code lines conditionally.")
3. Macro definition and expansion. There are different kind of macros:
   - [Simple macros](2567-simple-macro-definition.md "A simple macro is identified by its name and body.")
   - [Function macros](2568-function-macro-definition.md "Function macros are preprocessor macros which can take arguments.")
   - [Predefined macros](2569-predefined-macros.md "A set of predefined preprocessor macros are available.")
4. Macros can be defined with operators for:
   - [Stringification](2570-stringification-operator.md "Transforms a preprocessor macro element to a string.")
   - [Concatenation](2571-concatenation-operator.md "Concatenates two parameters of a preprocessor macro.")
5. You can [undefine macros](2572-undefining-a-macro.md "Undefines a preprocessor macro.").

If a preprocessing directive is invalid, the compilers will generate an .err
file with the preprocessing error included in the source file at the line position where the problem
exists. When using the `-M` option, preprocessor errors will be printed to stderr,
like regular compiler errors.
