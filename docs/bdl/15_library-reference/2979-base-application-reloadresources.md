---
title: "base.Application.reloadResources"
source: "fgl-topics/c_fgl_ClassApplication_reloadResources.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.reloadResources"
type: "concept"
---

# base.Application.reloadResources

> Resets FGLRESOURCEPATH and reloads localized string resources.

## Syntax

```
base.Application.reloadResources(
    newResourcePath STRING)
```

1. newResourcePath is a list of directories to search for string resource files.

## Usage

The `reloadResources()` method overwrites the search path defined by the [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.") environment variable, to find
program resource files in directories that are only known at runtime.

> **Warning:**
>
> The `reloadResources()` method, to reset the FGLRESOURCEPATH
> environment variable in order to find program resource files in a different directory, must only be
> used at the beginning of the program execution.

This method is typically used to define a search path for localized string files when a program
starts, to let the end user change the current application language. It avoids starting a new
application (via `RUN`), after the end user selects a language in a parent
program.

The runtime behaves as if FGLRESOURCEPATH had been set to this value from the start.

Pay attention to the path separator, which is specific to the operating system. See
FGLRESOURCEPATH reference for more details.

The method does the following:

1. Resets the environment variable FGLRESOURCEPATH with the specified value.
2. Reloads already loaded string localization files (.42s)
3. Reloads the default action defaults file (default.4ad)

Notes:

- Reloading resources has no effect on displayed forms: only forms displayed after reloading
  resources will use the new reloaded strings.
- Presentation Styles (.4st) are not reloaded.

  The writing direction of a language/script is defined with the [`UserInterface.reverse`](../11_user-interface/1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.") presentation style
  attribute:

  ```
  <Style name="UserInterface">
    <StyleAttribute name="reverse" value="yes" />
  </Style>
  ```

  Since presentation styles are not reloaded, it is not possible to switch
  between scripts having different writing directions. To change the writing direction, the program
  must be restarted.
- Reloading resources has no effect on .42m modules that are already loaded
  (any %"string" will not be localized again). For this reason, the `reloadResources()`
  method must be called at a very early stage of the program.

  Note that the debugger (fglrun -d) loads all program modules immediately when
  starting. Therefore, reloading resources has no effect on localized strings in
  .42m modules when debugging.
- Reloading resources has no effect on files loaded by methods such as [`ui.Form.loadToolBar`](3153-ui-form-loadtoolbar.md "Load the form toolbar."), [`ui.From.loadTopMenu`](3154-ui-form-loadtopmenu.md "Load the form topmenu."), [`ui.form.loadActionDefaults`](3152-ui-form-loadactiondefaults.md "Load form action defaults."),
  [`ui.Interface.loadActionDefaults`](3112-ui-interface-loadactiondefaults.md "Load the default action defaults file."), [`ui.Interface.loadStartMenu`](3113-ui-interface-loadstartmenu.md "Load the start menu file."),
  [`ui.Interface.loadStyles`](3114-ui-interface-loadstyles.md "Load the presentation styles file."),
  [`ui.Interface.loadToolBar`](3115-ui-interface-loadtoolbar.md "Load a default/global toolbar file for all forms of the program."),
  [`ui.Interface.loadTopMenu`](3116-ui-interface-loadtopmenu.md "Load a default/global topmenu file for all forms of the program.").

## Related links

**Related concepts**  

[Loading localized strings at runtime](../09_advanced-features/0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.")
