---
title: "Internationalization"
source: "fgl-topics/c_fgl_intro_BDL_007.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Internationalization"
type: "concept"
---

# Internationalization

> The language supports single-byte and multibyte internationalization.

The language supports single-byte such as ISO-8859-1, as well as multibyte character sets such
as BIG5 or UTF-8.

Length semantics to define variables and manipulate character string data can be based on byte
or character units.

Labels and messages can be separated from programs and forms, to customize your application for
specific subsets of for the user population, whether it is for a particular language or a
particular business segment.

The source files (4gl, per, 4ad,
and so on) can be written in a specific encoding. However, we recommend you to keep sources in
ASCII, and store locale-dependent strings in external strings files (`str`).

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
