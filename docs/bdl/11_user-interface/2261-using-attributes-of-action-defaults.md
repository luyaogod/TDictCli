---
title: "Using attributes of action defaults"
source: "fgl-topics/c_fgl_prog_dialogs_action_defaults.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Using attributes of action defaults"
type: "concept"
description: "Purpose of action defaults Action defaults allow you to define default attributes for common action. These defaults can be overwritten with form item attributes, or with dialog action handler ..."
---

# Using attributes of action defaults

## Purpose of action defaults

Action defaults allow you to define default attributes for common action. These
defaults can be overwritten with form item attributes, or with dialog action handler attributes
(only for default action views).

Centralize action attributes with action defaults, to avoid specifying them in all the source
files that define the same action view and action handler. For example, you can specify the default
text, image and keyboard accelerator for elements like push `BUTTON`, `TOOLBAR` items, `TOPMENU` options.

Common action defaults are typically defined in a global action defaults
(.4ad) file, while form specific actions are configured with form action
defaults in the `ACTION
DEFAULTS` section of the .per form specification file.

## Global action defaults file

Global action defaults are defined in an XML file with the
.4ad extension. By default, the runtime system searches for a file named
default.4ad in several directories as described in [the FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files."). If no
file was found, standard action default settings are loaded from the
$FGLDIR/lib/default.4ad file.
> **Important:**
>
> Global action defaults must be defined in a unique file; you cannot combine
> several `4ad` files.

It is possible to use localized strings in action default attributes such as
`TEXT` and `COMMENT`, by using `LStr` XML
elements:

```
<ActionDefaultList>
  <ActionDefault name="yes" text="Yes">
     <LStr text="common.yes"/>
  </ActionDefault>
  ...
```

A global action defaults file can be loaded dynamically by program with the [`ui.Interface.loadActionDefaults()`](../15_library-reference/3112-ui-interface-loadactiondefaults.md "Load the default action defaults file.") method. Use this solution if you need
different global action defaults depending on the program. However, it is recommended that you
consider using a single global default actions file, to get the same decoration and keyboard
accelerators in all your programs.

## Form specific action defaults

Action defaults can be defined at the form level in the `ACTION DEFAULTS` section.
When action defaults are defined in the form file, action views get the attributes defined locally
for this
form:

```
ACTION DEFAULTS
  ACTION print (TEXT="Print",
                IMAGE="printer",
                COMMENT="Print the current record",
                ACCELERATOR=CONTROL-P)
END
```

It is possible to use localized strings in action default attributes such as
`TEXT` and
`COMMENT`:

```
ACTION print (TEXT=%"common.print")
```

Form specific action defaults files can be loaded dynamically by program with the [`ui.Form.loadActionDefaults()`](../15_library-reference/3152-ui-form-loadactiondefaults.md "Load form action defaults.")
method. Use this solution if you cannot change your .per form definition files
to define `ACTION DEFAULTS` section. The `loadActionDefaults()` method
of a form object is typically used in a [generic form initializer function](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.").

## Action defaults are applied only once

Decoration attributes (like `TEXT`, `IMAGE`) of an action view
will automatically be set with the value defined in the action defaults to all new action views of a
newly-created form, if there is no explicit value specified in the element definition for that
attribute. Decoration action default attributes are applied only once, to newly-created form
elements. Dynamic changes are not reapplied to action views. For example, if you first load a
toolbar, then you load a global action defaults file, the attributes of the toolbar items will not
be updated with the last loaded action defaults. If you dynamically create action views (like
`TopMenu` or `ToolBar`), the action defaults are not applied, so you
must set all decoration attributes by hand.

## Action defaults and sub-dialog actions

The action default attributes to be applied are selected following the name of the action. In some
situations, the action view can be bound to an action handler by specifying a sub-dialog and/or
field name prefix. For those views, the action defaults defined with the corresponding action
name will be used to set the attributes with the default values. In other words, the prefix
will be ignored. For example, if an action view is defined with the name
`custlist.append`, it will get the action defaults defined for the
`append` action.

## Functional attributes

Functional attributes (like `VALIDATE`,
`ACCELERATOR`) can only
be defined in action defaults, or in `ON ACTION` dialog action handlers with the
`ATTRIBUTES` clause.
Functional attributes take effect for a given action when the dialog becomes active.
