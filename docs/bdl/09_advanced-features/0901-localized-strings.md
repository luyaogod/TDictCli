---
title: "Localized strings"
source: "fgl-topics/c_fgl_localized_strings_001.html"
breadcrumb: "Advanced features > Localization > Localized strings"
type: "concept"
---

# Localized strings

> Localized strings provide a means of writing applications in which the text of strings can be customized on site.

This string localization feature is a simple way to define external resource files which the
runtime system can search, in order to assign text to elements displayed by programs. It
can be used to implement internationalization in your application, or to use
site-specific text, for example, when business terms are specific to the territory where
the application is used.

The localized string resource files (.42s) are loaded at runtime and
shared by all fglrun processes. Localized strings are used to replace the
original strings found in the p-code modules ( .42m ), in the compiled form (
.42f ), and in any XML resource files loaded in the abstract user interface
tree ( .4ad , .4st , .4tb , etc).

Runtime system messages displayed during dialog execution can be localized with a predefined
string key format. See [Defining dialog messages with localized strings](0888-runtime-system-messages.md).

## Related links

**Related concepts**  

[Form specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[The abstract user interface tree](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")

## Child topics

- [Steps for application internationalization](0902-steps-for-application-internationalization.md): Follow these steps to internationalize your application.
- [Creating source string files](0903-creating-source-string-files.md): A source string file contains localized string definitions for a given language (or localization context).
- [Localized strings in program sources](0904-localized-strings-in-program-sources.md): How to specify a localized string in .4gl and .per sources?
- [Localized strings in XML resource files](0905-localized-strings-in-xml-resource-files.md): In XML resource files, localized string specification must follow the XML syntax and therefore must be defined as an XML node.
- [Extracting strings from sources](0906-extracting-strings-from-sources.md): Localized strings can be easily extracted from .4gl and .per source files.
- [Organizing string resources](0907-organizing-string-resources.md): Good practice in use of localized strings.
- [Compiling string resource files (.str)](0908-compiling-string-resource-files-str.md): The .str source string files must be compiled to .42s binary files, in order to be loaded by the runtime system.
- [Loading localized strings at runtime](0909-loading-localized-strings-at-runtime.md): Understand the rules for using localized strings at runtime.
- [Predefined application strings](0910-predefined-application-strings.md): The runtime system may need to display text to the user.
- [Best practices for localized strings](0911-best-practices-for-localized-strings.md): This section describes good practices to localize your application messages.
- [Example](0912-example.md): Here is an example using localized strings.
