---
title: "Executing Java code with GMA"
source: "fgl-topics/c_fgl_JavaBridge_gma.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA"
type: "concept"
description: "On Android™ devices, Genero apps can use the Java interface. The GMA executes a program in a JVM process and therefore does not require more resources to execute Java code. JDK 17 is required to build ..."
---

# Executing Java code with GMA

On Android™ devices, Genero apps can use the Java
interface.

The GMA executes a program in a JVM process and therefore does not require more resources to
execute Java code.

JDK 17 is required to build Android apps. For the latest information regarding system requirements and
Java support, please refer to the Supported platforms and databases document, available
on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).

We distinguish the following use cases where the Java interface of Genero can be used in GMA:

- Use classes from the [standard Java or Android Java
  library](2694-standard-java-and-android-library-usage.md "You can use Java classes that are part of the standard Java library and Android Java library.").
- Implement and use [user-defined Java
  classes](2695-implement-java-user-extensions-in-gma.md "A GMA app can execute custom Java code."), requiring GMA packaging.
- Implement and execute a [user-defined Android
  activity](2696-implement-android-activities-in-gma.md "Android activities can be bundled with your GMA app and called from the Genero code."), requiring GMA packaging.

Java may also be used to extend the GMA front-end with user-defined front calls. For details, see
[Implement front call modules for GMA](2720-implement-front-call-modules-for-gma.md "Custom front call modules for the Android front-end are implemented by using the API for GMA front calls in Java.").

## Child topics

- [Standard Java and Android library usage](2694-standard-java-and-android-library-usage.md): You can use Java classes that are part of the standard Java library and Android Java library.
- [Implement Java user extensions in GMA](2695-implement-java-user-extensions-in-gma.md): A GMA app can execute custom Java code.
- [Implement Android activities in GMA](2696-implement-android-activities-in-gma.md): Android activities can be bundled with your GMA app and called from the Genero code.
- [Packaging custom Java extensions for GMA](2697-packaging-custom-java-extensions-for-gma.md): Custom Java extension must be integrated in the GMA to run on Android devices.
