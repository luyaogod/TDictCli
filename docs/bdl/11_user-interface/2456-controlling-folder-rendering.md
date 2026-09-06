---
title: "Controlling folder rendering"
source: "fgl-topics/c_fgl_ui_folders_rendering.html"
breadcrumb: "User interface > User interface programming > Folders > Controlling folder rendering"
type: "concept"
---

# Controlling folder rendering

> Folder rendering can be controlled by the use of presentation styles and folder attributes.

## Default rendering

By default, a `FOLDER` container displays with folder tabs on top:

![Form with folder tabs screenshot](../_images/gbc_foldertabs_3.jpg)

*Form with default Folder rendering*

## Folder tab positions

The position of the folder tabs can be define with the [`position`](1638-folder-style-attributes.md) style attribute:

```
  <Style name="Folder">
     <StyleAttribute name="position" value="right" />
  </Style>
```

With the above style definition, folders will display as in the following screenshot:

![Form with folder tabs on the right screenshot](../_images/gbc_foldertabs_4.jpg)

*Form with Folder Tabs on the right*

## Collapsible folder tabs

To display folder tabs as collapsible elements, use the `"accordion"` value in the
[`position`](1638-folder-style-attributes.md) style attribute:

```
  <Style name="Folder">
     <StyleAttribute name="position" value="accordion" />
  </Style>
```

The folder will then render as shown in the next screenshot, with collapsible pages:

![Form with collapsible folder tabs screenshot](../_images/gbc_foldertabs_5.jpg)

*Form with collpasible Folder Tabs*
