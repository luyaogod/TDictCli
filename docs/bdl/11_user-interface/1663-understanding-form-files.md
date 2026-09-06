---
title: "Understanding form files"
source: "fgl-topics/c_fgl_FormSpecFiles_intro.html"
breadcrumb: "User interface > Form definitions > Form specification files > Understanding form files"
type: "concept"
---

# Understanding form files

> A form specification file is a source file that defines an application form providing for end user interaction with a program.

The form file defines the disposition, presentation (in other words the decoration), and
behavior of screen elements called form items.

The source file must have the .per file extension:
myform.per. Programs load the .42f compiled
version of the form files, and use interactive instructions (dialogs) to control the
form.

To compile a .per source file to a .42f format, use
the [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") form compiler. When a
`SCHEMA` is specified in the form file, fglform requires that the
[database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") already exist. Compiled form
files depend on both the source files and the database schema files.

Compiled forms will be loaded by the programs with the `OPEN FORM` or the
`OPEN WINDOW WITH FORM` instructions. The .42f form file is
searched for in several directories, as described in the [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.") reference topic.

Once a form is loaded, the program can manipulate forms to display or let the user edit data,
with interactive instructions such as `INPUT` or `DISPLAY
ARRAY`. Program variables are used as display and/or input buffers.

The content of a .per form file must follow a specific syntax as
described in [Form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

## Related links

**Related concepts**  

[What are dialog controllers?](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.")

[Form rendering](1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends.")

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
