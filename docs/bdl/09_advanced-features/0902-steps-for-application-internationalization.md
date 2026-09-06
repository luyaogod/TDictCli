---
title: "Steps for application internationalization"
source: "fgl-topics/c_fgl_localized_strings_006.html"
breadcrumb: "Advanced features > Localization > Localized strings > Steps for application internationalization"
type: "concept"
---

# Steps for application internationalization

> Follow these steps to internationalize your application.

1. Identify the current character set used in your sources and make sure the application locale
   (LANG/LC\_ALL) is set correctly.
2. In .4gl sources, [add a %
   prefix](0904-localized-strings-in-program-sources.md "How to specify a localized string in .4gl and .per sources?") to the strings that must be localized (translated). For [parameterized messages](0907-organizing-string-resources.md "Good practice in use of localized strings."),
   replace concatenated strings by a `SFMT()` usage with
   `%n` placeholders for variable message parts.
3. In .per sources `LAYOUT` section, replace hard-coded form
   elements like text labels with static `LABEL` form items and define
   the `TEXT` attributes with a % prefix in the `ATTRIBUTES` section.
4. In XML resources, [add `<LStr />`
   elements](0905-localized-strings-in-xml-resource-files.md "In XML resource files, localized string specification must follow the XML syntax and therefore must be defined as an XML node.") under the elements where text attributes must be localized.
5. [Extract the strings](0906-extracting-strings-from-sources.md "Localized strings can be easily extracted from .4gl and .per source files.") from the
   .4gl sources with fglcomp -m and use
   fglform -m for .per sources.
6. [Organize](0907-organizing-string-resources.md "Good practice in use of localized strings.") the generated
   .str source string files (identify duplicated strings and put them in a common
   file).
7. At this point, the string identifiers (on the left) are the same as the string texts (on the
   right). These string identifiers can be used as is, or can be changed to clear ASCII identifiers
   such as `"customer.list.title"`. Using simple identifiers allows you to distinguish
   strings depending on the context and use ASCII encoding for your sources. Keeping string identifiers
   with the original text requires no source changes (except adding the % prefix), but makes sources
   dependent to a locale: If you want to support multiple languages, you must use UTF-8 in sources and
   at runtime.
8. When using simple ASCII identifiers, replace original strings with the new string identifiers.
   Strings to be replaced can be located by their % prefix. You can, for example, use a script with an
   utility like the *sed* UNIX™ command to read the
   .str files and apply the changes automatically.
9. Recompile the .4gl and .per sources (when using simple
   ASCII strings identifiers, sources are expected to be full ASCII now).
10. [Compile](0908-compiling-string-resource-files-str.md "The .str source string files must be compiled to .42s binary files, in order to be loaded by the runtime system.") the .str files
    in the locale used by these files,
11. [Setup FGLPROFILE](0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.")
    `fglrun.localization.*` entries, to let fglrun
    find the string resource files.
12. Run your programs to check whether the application displays the text properly.
13. Copy the existing .str files, and translate the string text into another
    language (make sure the locale is correct).
14. Compile the new .str files, and copy the .42s files
    into another distribution directory, defined with the FGLRESOURCEPATH environment variable.
15. Run your programs again, to check that texts and labels of the other language are displayed.
16. Next changes to the .per and .4gl source files are
    done in the ASCII locale, and .str string files must be edited with their
    specific locale.

## Related links

**Related concepts**  

[Application locale](0864-application-locale.md "The application locale defines the language and codeset for your application.")
