---
title: "BDL 2.51 new features"
source: "fgl-topics/fgl_whatsnew_251.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.51 new features"
type: "topic"
---

# BDL 2.51 new features

> Features added in 2.51 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.51 upgrade guide](0226-bdl-2-51-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.51.").

> **Important:**
>
> Most of the new features of BDL 2.51 have been added for Genero Mobile. The
> features designed for Genero Mobile may not be supported by desktop and web-browser front-ends in
> the coming releases.

Prior new features guide: [BDL 2.50 new features](0066-bdl-2-50-new-features.md "Features added in 2.50 releases of the Genero Business Development Language.").

## Genero Mobile V 1.0 (FGL 2.51.06)

| Overview | Reference |
| --- | --- |
| Remote debugging through network TCP socket | See [Debugging on a mobile device](../13_programming-tools/2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device."). |
| The channel methods `openServerSocket()` and `readOctets()` | See [base.Channel.openServerSocket](../15_library-reference/2992-base-channel-openserversocket.md "Open a TCP server socket channel."), [base.Channel.readOctets](../15_library-reference/2995-base-channel-readoctets.md "Read a given number of bytes and return as a character string."). |
| The `sort()` method of ARRAY variables. | See [DYNAMIC ARRAY.sort](../15_library-reference/2953-dynamic-array-sort.md "Sorts the rows in the array."). |
| Datetime-related utility methods. | See [util.Datetime methods](../15_library-reference/3487-util-datetime-methods.md "Methods for the util.Datetime class."). |
| String-related utility methods. | See [util.Strings methods](../15_library-reference/3550-util-strings-methods.md "Methods for the util.Strings class."). |
| Write to stdout with `om.XmlWriter.createFileWriter(NULL)`. | See [om.XmlWriter.createFileWriter](../15_library-reference/3380-om-xmlwriter-createfilewriter.md "Creates an om.SaxDocumentHandler object writing to a file."). |

| Overview | Reference |
| --- | --- |
| FGLPROFILE settings to define environment variables | See [Setting environment variables in FGLPROFILE (mobile)](../07_configuration/0495-setting-environment-variables-in-fglprofile-mobile.md). |
| The method `base.Application.isMobile()` | See [base.Application.isMobile](../15_library-reference/2978-base-application-ismobile.md "Indicates if the application runs on a mobile device."). |
| FGL Java class to access Android™ JVM context | See [Standard Java and Android library usage](../14_extending-the-language/2694-standard-java-and-android-library-usage.md "You can use Java classes that are part of the standard Java library and Android Java library."). |
| VCard utility functions. | See [VCard: VCF file format module](../15_library-reference/2815-vcard-vcf-file-format-module.md). |

| Overview | Reference |
| --- | --- |
| URL-based Web Components | See [Using a URL-based web component](../11_user-interface/2385-using-a-url-based-web-component.md "This section describes how to add a URL-based web component to your application."). |
| The `DATETIMEEDIT` form item type | See [DATETIMEEDIT item type](../11_user-interface/1689-datetimeedit-item-type.md "Defines a line-edit with a calendar widget to pick a datetime."). |
| Dialog-level action attribute definitions with `ON ACTION name ATTRIBUTES()`. | See [Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes."). |
| New `ON SELECTION CHANGE` control block. | See [Multiple row selection](../11_user-interface/2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records."). |

| Overview | Reference |
| --- | --- |
| `START DIALOG` / `TERMINATE DIALOG` / `fgl_eventLoop()` | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| Window `TYPE` attribute in `OPEN WINDOW` instruction. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| `DISPLAY ARRAY` attributes for list views handling: `ACCESSORYTYPE`, `DETAILACTION`, `DOUBLECLICK`. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| The `DISCLOSUREINDICATOR` action attribute. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |
| The `ROWBOUND` action attribute. | See [ROWBOUND action attribute](../11_user-interface/2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list."). |
| The `KEYBOARDHINT` form field attribute. | See [KEYBOARDHINT attribute](../11_user-interface/1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly."). |
| List filter with `DISPLAY ARRAY` dialog. | See [List reduce filter](../11_user-interface/2327-list-reduce-filter.md "The reduce filter allows a user to limit the row set in the list by using a filter."). |
| Method `ui.Interface.getFrontEndName()` can now return `GMI` or `GMA` | See [ui.Interface.getFrontEndName](../15_library-reference/3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use."). |
| Front-end functions for Genero Mobile (GMA / GMI) | See [Genero Mobile common front calls](../15_library-reference/3441-genero-mobile-common-front-calls.md "This section describes common front calls provided by all mobile front-ends."), [Genero Mobile Android front calls](../15_library-reference/3463-genero-mobile-android-front-calls.md "This section describes front calls specific to the Android platform."), [Genero Mobile iOS front calls](../15_library-reference/3469-genero-mobile-ios-front-calls.md "This section describes front calls specific to the iOS platform."). |
| Navigation bar button colors and background colors for iOS device (`iosTintColor` , `iosNavigationBarTintColor`, `iosToolBarTintColor`, `iosTabBarTintColor`) - provided as Window class style attributes. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |

| Overview | Reference |
| --- | --- |
| Simplified database driver specification | See [New database driver name specification](0228-new-database-driver-name-specification.md "Allows database driver specification without target database version information."). |
| Support for SQL Server 2014 | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| Support for Oracle Database 12c | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| Support for PostgreSQL 9.3 | See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| Better support for DATETIME types with SQLite | See [DATETIME types with SQLite](0231-datetime-types-with-sqlite.md "Better support for Informix DATETIME types emulation within SQLite."). |
| `STRING` typed variables can be used in SQL statements. | See [STRING](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation."). |
| Upgrade notes for database drivers | See [Database drivers changes](0227-database-drivers-changes.md "Desupported database drivers."). |

## Genero Mobile V 1.1 (FGL 2.51.07)

| Overview | Reference |
| --- | --- |
| Implementing C-Extensions on iOS / GMI. | See [Implementing C-Extensions for GMI](../14_extending-the-language/2717-implementing-c-extensions-for-gmi.md "This section describes how to program C-Extensions for the GMI VM."). |
| Using Java interface for Android / GMA. | See [Executing Java code with GMA](../14_extending-the-language/2693-executing-java-code-with-gma.md). |
| Implementing customer front calls for GMA. | See [Implement front call modules for GMA](../14_extending-the-language/2720-implement-front-call-modules-for-gma.md "Custom front call modules for the Android front-end are implemented by using the API for GMA front calls in Java."). |
| Presentation styles are now supported by mobile front-ends. | See [Style attributes reference](../11_user-interface/1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element."). |
| GMA bundles zxing for Android. | See [mobile.scanBarCode](../15_library-reference/3459-mobile-scanbarcode.md "Allow the user to scan a barcode with a mobile device"). |

| Overview | Reference |
| --- | --- |
| Complete support of Web Services on Android mobile devices.(Web Services are partly supported on iOS mobile devices) | See [Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL."). |

| Overview | Reference |
| --- | --- |
| Presentation styles are now supported by mobile front-ends. | See [Style attributes reference](../11_user-interface/1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element."). |
| GMA bundles zxing for Android. | See [mobile.scanBarCode](../15_library-reference/3459-mobile-scanbarcode.md "Allow the user to scan a barcode with a mobile device"). |
