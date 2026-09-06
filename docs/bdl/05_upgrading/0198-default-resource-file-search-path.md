---
title: "Default resource file search path"
source: "fgl-topics/c_fgl_Migrate_to_310_form_search.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Default resource file search path"
type: "concept"
---

# Default resource file search path

> Search rules for program resource files have been enhanced in 3.10.

## Default search in application directory

Starting with version 3.10, if resource files such as .42f form files are
not found in the current working directory, or in the directories specified in the FGLRESOURCEPATH
environment variable, or in $FGLDIR/lib, the runtime system also does a lookup
in the directory where the `MAIN`
.42m module or the .42r program file resides.

As a result, if all program files are in the directory where the main program module resides,
there is no need to set FGLRESOURCEPATH.

For more details about resource file directory search, see [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

## Form file path using absolute path

If the filename specified in `OPEN FORM` / `OPEN WINDOW WITH FORM`
is an absolute path, the runtime system must not search for the form in FGLRESOURCEPATH.

Before version 3.10, to load a form file specified with an absolute path in `OPEN
FORM` or `OPEN WINDOW WITH FORM`, the runtime system tried to find the file
by appending the absolute path to the directories defined by FGLRESOURCEPATH.

This could end up in loading unexpected resource files.

For example, with:

```
OPEN FORM f FROM "/dir/form"
```

The form "/dir/form" was searched in each element of FGLRESOURCEPATH.

Since version 3.10, FGLRESOURCEPATH is only used, if the filename is not an absolute
path.

For more details, see [WITH FORM clause](../11_user-interface/1574-with-form-clause.md "Creating a window object with a form.").

## Related links

**Related concepts**  

[Providing the image resource](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.")

[os.Path.pathType](../15_library-reference/3734-os-path-pathtype.md "Checks if a path is a relative path or an absolute path.")
