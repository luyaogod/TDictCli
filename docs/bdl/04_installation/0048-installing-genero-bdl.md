---
title: "Installing Genero BDL"
source: "fgl-topics/c_fgl_installation_010.html"
breadcrumb: "Installation > Installing Genero BDL"
type: "concept"
---

# Installing Genero BDL

> This section provides Genero BDL installation instructions.

Different forms of installation programs are provided, as individual packages are bundled with
other Genero components. Refer to the appropriate installation guide for a detailed description of
the installation procedure. Do not hesitate to contact your support center if you need help.

The installation and licensing of Genero products requires you to read and
accept the End User License Agreement, which can be found on the Four Js website at
<https://4js.com/end-user-license-agreements/>.

The Genero BDL produce is controlled by license number. You need to register a license before
using the product. For more information, refer to the Four Js License Manager User Guide.

After installing and licensing a Genero BDL package, it is recommended that you:

1. Set the [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.") environment variable to the BDL
   installation directory.
2. Add $FGLDIR/bin to the PATH environment variable, in order to run compilers
   and runtime system tools from the command line.
3. Set the [database client software environment](../10_sql-support/1062-database-client-environment.md "To connect to a database server, Genero BDL programs use vendor's database client software.") (for
   example, INFORMIXDIR, etc)
4. Set access path to database client software DLLs (PATH), or UNIX® shared libraries (LD\_LIBRARY\_PATH, SHLIB\_PATH, LIBPATH)
5. Depending on the database server you want to connect to, set up the correct [database driver](../10_sql-support/1071-connection-parameters.md "This section describes the different parameters which need to be specified in order to connect to a database.") in [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files"). The default database driver is
   selected during the installation procedure.
6. Depending on what rendering mode you want to use (text mode or graphical mode), you will have to
   set environment variables such as [FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program."),
   [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application."), TERM, INFORMIXTERM.
7. If your application uses [C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code."), a C
   compiler is required and you must recompile your C-Extensions as shared libraries.
8. If your application uses the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs."), setup
   the required [JDK environment](../14_extending-the-language/2658-java-software-requirements-for-fgl.md).

> **Tip:**
>
> The Genero BDL package installer creates an environment file named envcomp,
> in the installation directory. Run the envcomp script the set the minimum
> required environment for Genero BDL. On Microsoft™ Windows®, the environment file is named
> envcomp.bat.

## Related links

**Related concepts**  

[General BDL upgrade guide](../05_upgrading/0085-general-bdl-upgrade-guide.md "These topics describe general considerations when upgrading to a new version of Genero BDL.")

[Genero environment variables](../07_configuration/0506-genero-environment-variables.md "Genero environment variables")

[Operating system environment variables](../07_configuration/0496-operating-system-environment-variables.md "Describes some well-known system environment variables that are used by Genero software components.")

**Related tasks**  

[Install Genero Mobile for Android](0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.")

[Install Genero Mobile for iOS (single version)](0051-install-genero-mobile-for-ios-single-version.md "To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install a unique GMI version/package into FGLDIR.")

[Install Genero Mobile for iOS (multiple versions)](0052-install-genero-mobile-for-ios-multiple-versions.md "To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install multiple GMI versions/packages in parallel into different GMIDIR directories.")
