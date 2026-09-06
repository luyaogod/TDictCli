---
title: "ui.DragDrop methods"
source: "fgl-topics/c_fgl_ClassDragDrop_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The DragDrop class > ui.DragDrop methods"
type: "concept"
---

# ui.DragDrop methods

> Methods of the ui.DragDrop class.

| Name | Description |
| --- | --- |
| addPossibleOperation( operation STRING ) | Add a possible operation. |
| dropInternal() | Perform built-in row drop in trees. |
| getBuffer() RETURNS STRING | Get drag & drop data from the buffer. |
| getLocationParent() RETURNS INTEGER | Get the index of the parent node where the object was dropped. |
| getLocationRow() RETURNS INTEGER | Get the index of the target row where the object was dropped. |
| getOperation() RETURNS STRING | Identify the type of operation on drop. |
| getSelectedMimeType() RETURNS STRING | Get the previously selected MIME type. |
| selectMimeType( mimeType STRING ) | Select the MIME type before getting the data. |
| setBuffer( buffer STRING ) | Set the text data of the dragged object. |
| setFeedback( feedback STRING ) | Define the appearance of the target during Drag & Drop. |
| setMimeType( mimeType STRING ) | Define the MIME type of the dragged object. |
| setOperation( operation STRING ) | Define the type of Drag & Drop operation. |
