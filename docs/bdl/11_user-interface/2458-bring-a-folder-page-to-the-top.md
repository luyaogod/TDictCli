---
title: "Bring a folder page to the top"
source: "fgl-topics/c_fgl_ui_folders_multiple.html"
breadcrumb: "User interface > User interface programming > Folders > Bring a folder page to the top"
type: "concept"
---

# Bring a folder page to the top

> How to make a folder page current?

To bring a folder page to the top by program:

- If the target folder page has at least one focusable field, use the `NEXT FIELD`
  instruction to give the focus to one of the active fields of the page. `NEXT FIELD
  first-col` can also be used for example to a `TABLE`
  controlled by a `DISPLAY ARRAY`.
- If the folder page has no focusable element, use the
  `ui.Form.ensureFieldVisible()` or the `ui.Form.ensureElementVisible()`
  method.

For more details, see [Giving the focus to a form element](2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.").
