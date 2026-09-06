---
title: "Folder style attributes"
source: "fgl-topics/r_fgl_presentation_styles_folder_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Folder style attributes"
type: "reference"
---

# Folder style attributes

> Folder presentation style attributes apply to FOLDER tab elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `collapserPosition`

Indicates the position of the collapser icon, when the `position` style attribute
is `"accordion"`.

Values can be `"left"` (default), `"right"`.

## `lateRendering`

Controls the rendering mode of folder tabs.

Values can be:

- `"no"` (default): When displaying a folder, all folder pages are created and
  measured to compute the folder size. This can impact performances with folders having many complex
  pages.
- `"yes"`: When displaying a folder, only the current folder page is created and
  measured to compute the folder size. The layout of other pages will be computed when displayed. This
  mode improves performances.

## `navigationArrows`

The `navigationArrows` style attribute controls left and right navigation arrows
for the view, to let the end user navigate between pages of the folder.

Navigation arrows are only displayed with the regular folder rendering (these do not appear with
accordion folder)

Values can be `"yes"`, `"no"` (default).

## `navigationDots`

The `navigationDots` style attribute controls the navigation dots created for the
view, to let the end user navigate between pages of the folder.

Navigation dots are only displayed with the regular folder rendering (these do not appear with
accordion folder)

Values can be `"yes"`, `"no"` (default).

## `position`

Defines the position and type of the folder tabs.

Values can be:

- `"top"` (default): The folder tabs are displayed on the top of folder pages.
- `"left"`: The folder tabs are displayed on the left of folder pages.
- `"right"`: The folder tabs are displayed on the right of folder pages.
- `"bottom"`: The folder tabs are displayed on the bottom of folder pages.
- `"accordion"`: The folder pages can be collapsed and expanded.

## Related links

**Related concepts**  

[FOLDER item type](1691-folder-item-type.md "Defines a layout area to hold folder pages.")
