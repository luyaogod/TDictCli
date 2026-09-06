---
title: "Syntax of start menu files (.4sm)"
source: "fgl-topics/c_fgl_startmenus_003.html"
breadcrumb: "User interface > User interface programming > Start menus > Syntax of start menu files (.4sm)"
type: "concept"
---

# Syntax of start menu files (.4sm)

> A start menu file contains a tree of XML elements defining the application menu to start programs.

Start menus are defined in a .4sm file with the following XML syntax:

```
<StartMenu [ startmenu-attribute="value"[...] ] >
  group[...]
</StartMenu>
```

where group is:

```
<StartMenuGroup group-attribute="value"
    [...]>{ <StartMenuSeparator/>| <StartMenuCommand
     command-attribute="value"
    [...] />|
    group}
    [...]
</StartMenuGroup>
```

1. startmenu-attributedefines a property of the
   `StartMenu`.
2. command-attribute defines a property of a
   `StartMenuCommand`.
3. group-attribute defines a property of a
   `StartMenuGroup`.

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the StartMenu, can be omitted. |
| `text` | `STRING` | Defines the text to be displayed as title. |

| Attribute | Type | Description |
| --- | --- | --- |
| `disabled` | `INTEGER` | Indicates if the group must be disabled (grayed, cannot be selected). |
| `hidden` | `INTEGER` | Indicates if the group is hidden or visible. |
| `image` | `STRING` | Defines the icon to be used for this group. |
| `name` | `STRING` | Identifies the start menu group, can be omitted. |
| `text` | `STRING` | Defines the text to be displayed for this group. |

| Attribute | Type | Description |
| --- | --- | --- |
| `disabled` | `INTEGER` | Indicates if the item must be disabled (grayed, cannot be selected). |
| `comment` | `STRING` | Specifies the comment to be shown for this command. |
| `exec` | `STRING` | Defines the command to be executed when the user selects this command. |
| `hidden` | `INTEGER` | Indicates if the command is hidden or visible. |
| `image` | `STRING` | Defines the icon to be used for this command. |
| `name` | `STRING` | Identifies the StartMenu item, can be omitted. |
| `text` | `STRING` | Defines the text to be displayed for this command. |
| `waiting` | `INTEGER` | Defines if the command must be started without waiting (0, default) or waiting (1). |

| Attribute | Type | Description |
| --- | --- | --- |
| `name` | `STRING` | Identifies the StartMenu separator, can be omitted. |
