---
title: "Controlling the UI with XML classes"
source: "fgl-topics/c_fgl_xml_utils_007.html"
breadcrumb: "Advanced features > XML support > Controlling the UI with XML classes"
type: "concept"
---

# Controlling the UI with XML classes

> The User Interface of a Genero application can be manipulated with the build-in XML API.

The runtime system represents the user interface of a program with a DOM tree.

User interface elements can be manipulated with the DOM and SAX built-in classes.

However, you must pay attention when modifying the AUI tree directly through the use of these
classes. Invalid node or attribute creation can lead to unpredictable results.

## Related links

**Related concepts**  

[The abstract user interface tree](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")

[DOM and SAX built-in classes](0950-dom-and-sax-built-in-classes.md "The DOM and SAX APIs both contain a set of built-in classes.")
