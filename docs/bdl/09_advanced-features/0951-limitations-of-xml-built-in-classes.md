---
title: "Limitations of XML built-in classes"
source: "fgl-topics/c_fgl_xml_utils_005.html"
breadcrumb: "Advanced features > XML support > Limitations of XML built-in classes"
type: "concept"
---

# Limitations of XML built-in classes

> Built-in XML classes have some limitations you must be aware of.

The [`om.* built-in XML classes`](../15_library-reference/3280-the-om-package.md "These topics cover the built-in classes of the om package") are
provided for convenience and design for basic XML usage, to help you manipulate XML content easily
without loading a complete external XML library such as Java XML classes or a C-based XML
libraries.

Below is a list of known limitations of the `om.*` built-in classes:

1. There is no automatic charset conversion done when reading of writing XML documents:

   The encoding used by an XML file (`<?xml ... encoding='UTF-8' ?>`) must
   match the [current application locale](0864-application-locale.md "The application locale defines the language and codeset for your application.") defined for the
   runtime.
2. There is no DTD (Document Type Definition) nor is there XSD (XML Schema Definition) support for
   XML Schema validation.

   For example, you can create the same attribute twice or set an invalid attribute value. You must
   take care to follow the definition of the XML document when using these classes.
3. When reading XML documents, the `<!DOCTYPE>` part (the document type
   declaration) must not contain sub-elements such as `<!ENTITY>`.

   If `<!DOCTYPE>` is present and contains no sub-elements, it will just be
   ignored.

For a complete XML support, use the full-featured XML classes provided
in the [web services extension](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.").
