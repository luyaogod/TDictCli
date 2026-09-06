---
title: "json.JSONReader.setInputCharset"
source: "fgl-topics/c_gws_jsonJSONReader_setInputCharset.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.setInputCharset"
type: "concept"
---

# json.JSONReader.setInputCharset

> Defines the charset used on the input stream.

## Syntax

```
setInputCharset(
   charset STRING )
```

1. charset defines the character set to use.

## Usage

Defines the character set used when decoding on the input stream.

The character set in use may be defined by the [LANG
or LC\_ALL](../07_configuration/0497-lc-all-or-lang.md "Defines the current application locale on UNIX platforms.") variables in both Linux® and
Windows®. In Linux, you can check the system locale with the
localctl command, especially on Debian-based systems.

In Windows, when the LANG
environment variable is undefined, the language and character set defaults to the system locale
which is defined by the system's regional settings. When the LANG environment variable is defined,
the Genero compilers and runtime system will use this variable.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[Example: Reading a JSON file](3661-example-reading-a-json-file.md "This example shows two ways to read a JSON file.")
