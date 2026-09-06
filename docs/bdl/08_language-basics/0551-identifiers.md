---
title: "Identifiers"
source: "fgl-topics/c_fgl_language_features_identifiers.html"
breadcrumb: "Language basics > Syntax features > Identifiers"
type: "concept"
---

# Identifiers

> A Genero BDL identifier is a sequence of characters used to identify a program entity.

An identifier must conform to the following rules:

- It must include at least one character, without any limitation in size.
- The characters must be alphanumeric, including the underscore (`_`) symbol.
- The initial character must be a letter or an underscore.
- Blanks, hyphens, and other non-alphanumeric characters are not allowed.
- Letters must be classified as alphabetic characters (`[:alpha:]` in the regex
  specification). For example, the Euros sign € is not allowed.
- Common identifiers are not case sensitive, so `my_Var` and `MY_vaR`
  both denote the same identifier. However, in some cases, identifiers are case sensitive (like action
  names in the AUI tree). It is recommended to always write identifiers in lower case to avoid
  mistakes.

Within non-English locales, BDL identifiers can include non-ASCII characters in identifiers, if
those characters are defined in the code set of the current locale. In multibyte East Asian locales
that support languages whose written form is not alphabet-based (such as Chinese, Japanese, or
Korean), an identifier does not need to begin with a letter. It is however recommended to program in
ASCII.

## Related links

**Related concepts**  

[User interface basics](../11_user-interface/1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface.")

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
