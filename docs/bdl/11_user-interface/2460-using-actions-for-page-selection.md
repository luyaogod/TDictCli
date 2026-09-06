---
title: "Using actions for page selection"
source: "fgl-topics/c_fgl_ui_folders_page_action.html"
breadcrumb: "User interface > User interface programming > Folders > Using actions for page selection"
type: "concept"
---

# Using actions for page selection

> An action event handler can be associated to a given folder page.

A folder `PAGE` can define an [`ACTION`](1758-action-attribute.md "The ACTION attribute defines the action associated with the form item.") attribute, to bind an `ON ACTION` action handler, and
detect that the folder page is selected.

This solution may be used in some rare cases, for example when no form item can get the focus
implicitly.

The `ACTION` attribute on folder pages was initially provided before multiple
dialogs were introduced, to control the folder pages with a set of singular dialogs. With this
solution, the code must exit the current singular dialog when the page action fires, and restart the
singular dialog controlling the target folder page. This practice is now strongly discouraged and
should be replaced by a `DIALOG` statement.

Avoid `NEXT FIELD` usage in the `ON ACTION` block of a folder page
action: Use the default focus management, based on `TABINDEX` attributes to control
which element must get the focus when a folder page is selected.
