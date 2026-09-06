---
title: "Implement Java user extensions in GMA"
source: "fgl-topics/c_fgl_JavaBridge_gma_user_ext.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA > Implement Java user extensions in GMA"
type: "concept"
---

# Implement Java user extensions in GMA

> A GMA app can execute custom Java code.

JDK 17 is required to build Android™ apps. For the latest information regarding system requirements and
Java support, please refer to the Supported platforms and databases document, available
on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).

In order to execute Java user code on the mobile device, the compiled Java classes
need to be available to the Genero runtime system. They can then be imported with the `IMPORT
JAVA classname` instruction.

When executing the Genero program on a computer in development mode, define the CLASSPATH to your
.jar files. This allows the JVM loaded by the Genero runtime system find the
appropriate Java classes.

When executing the Genero program on a mobile device, the compiled user Java classes must be
included in the mobile app Android package
(.apk), which is created in the Genero Studio deployment
procedure.

## Related links

**Related concepts**  

[Packaging custom Java extensions for GMA](2697-packaging-custom-java-extensions-for-gma.md "Custom Java extension must be integrated in the GMA to run on Android devices.")
