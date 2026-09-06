---
title: "XML declaration added automatically"
source: "fgl-topics/c_fgl_Migrate_to_210_xml_decl.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.10 upgrade guide > XML declaration added automatically"
type: "concept"
---

# XML declaration added automatically

> The XML declaration is added automatically when writing XML files.

An XML file must start with a "Prolog" or "XML Declaration" defining the XML version and
character set used by the file:

```
<?xml version='1.0' encoding='ISO-8859-1' ?>
<root ...>
...
</root>
```

Starting with Genero version 2.10.05, the XML declaration is now added automatically when writing
XML files.

Before 2.10.05, a workaround allowed your you to write this header as a processing
instruction, but this solution was subject to mistakes: the non-ASCII characters written to
the XML file must match the encoding specification in the XML Declaration.

To avoid invalid character set definitions, the Genero BDL built-in classes now add the XML
Declaration with the `encoding` attribute defined depending on the current
locale used by the runtime system. The value written in the `encoding`
attribute is defined by the [charmap.alias](../09_advanced-features/0889-using-the-charmap-alias-file.md "The charmap.alias file can be used to map a system specific locale to a standard IANA locale.")
file.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")

[The om package](../15_library-reference/3280-the-om-package.md "These topics cover the built-in classes of the om package")
