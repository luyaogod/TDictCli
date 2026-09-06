---
title: "FGLRESOURCEPATH"
source: "fgl-topics/c_fgl_EnvVariables_FGLRESOURCEPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLRESOURCEPATH"
type: "concept"
---

# FGLRESOURCEPATH

> Defines a list of paths for program resource files.

## FGLRESOURCEPATH definition

The FGLRESOURCEPATH environment variable is used to define the search paths for program resource
files:

1. [Form definition files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
   (.42f),
2. [Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") (.iem),
3. [Action defaults files](../11_user-interface/1600-action-defaults-files.md "Action defaults files allow to centralize action configuration parameters such as text, icon, accelerators and behavior options in XML format.")
   (.4ad),
4. [Presentation styles files](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")
   (.4st),
5. [Start menu files](../11_user-interface/2446-start-menus.md "Start menus define a tree of application programs that can be started.") (.4sm),
6. [Toolbar files](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.") (.4tb),
7. [Topmenu files](../11_user-interface/1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.") (.4tm),
8. [Localized strings files](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
   (.42s).

When the specified resource file is not an absolute path, the runtime system searches in
directories in the following order:

1. The current working directory.
2. A path defined in the FGLRESOURCEPATH (or DBPATH) environment variable.
3. The [$FGLDIR/lib](0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")
   directory.
4. The directory where the program file resides (the .42m module containing
   `MAIN` or the .42r program file).

The path separator is platform specific ( **":"** on UNIX™ platforms and **";"** on Windows®
platforms).

On mobile platforms, localized string files are found by default in the language sub directories
of the app directory. For more details, see [Loading localized strings at runtime](../09_advanced-features/0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.").

## FGLRESOURCEPATH versus DBPATH

For compatibility with Informix® 4GL,
DBPATH is used by default to search for resource files such as form files and XML files used by the
program.

However, DBPATH is also used by the Informix database
software to locate databases: Informix Dynamic Server
uses DBPATH to let you specify fallback servers if INFORMIXSERVER is not available, and former Informix Standard Engine needs DBPATH to find
.dbs database files.

This can be a problem when connecting from a machine where path format is not the same as on the
remote database server: It is not possible to mix UNIX and DOS
path formats in DBPATH.

To work around this Informix limitation,
FGLRESOURCEPATH can be used instead of DBPATH to specify the directories of program resource files.
You are then free to define DBPATH as Informix
requires.

## Related links

**Related concepts**  

[DBPATH](0513-dbpath.md "Defines a list of paths for Genero program resource files.")

[IBM Informix Dynamic Server](../10_sql-support/1063-ibm-informix-dynamic-server.md "IBM Informix Dynamic Server")
