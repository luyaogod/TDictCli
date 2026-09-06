---
title: "ui.Form methods"
source: "fgl-topics/c_fgl_ClassForm_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods"
type: "concept"
---

# ui.Form methods

> Methods of the ui.Form class.

| Name | Description |
| --- | --- |
| ui.Form.setDefaultInitializer( initializer STRING ) | Define the default initializer for all forms. |
| ui.Form.setDefaultInitializerFunction( initializer FUNCTION( form ui.Form ) RETURNS () ) | Define the default initializer for all forms. |
| ui.Form.displayTo( value { base-type \| RECORD }, formFieldName STRING, screenLine INTEGER, attributes STRING ) | Displays values to form fields or screen arrays. |

| Name | Description |
| --- | --- |
| ensureElementVisible( name STRING ) | Ensure the visibility of a form element. |
| ensureFieldVisible( name STRING ) | Ensure visibility of a form field. |
| findNode( tagName STRING, name STRING ) RETURNS om.DomNode | Search for a child node in the form. |
| getNode() RETURNS om.DomNode | Get the DOM node of the form. |
| loadActionDefaults( path STRING ) | Load form action defaults. |
| loadToolBar( path STRING ) | Load the form toolbar. |
| loadTopMenu( path STRING ) | Load the form topmenu. |
| setElementComment( name STRING, comment STRING ) | Set the comment/hint of form elements. |
| setElementHidden( name STRING, hidden INTEGER ) | Show or hide form elements. |
| setElementImage( name STRING, image STRING ) | Change the image of form elements. |
| setElementStyle( name STRING, style STRING ) | Change the style of form elements. |
| setElementText( name STRING, text STRING ) | Change the text of form elements. |
| setFieldComment( name STRING, comment STRING ) | Set the comment/hint of a form field. |
| setFieldHidden( name STRING, hidden INTEGER ) | Show or hide a form field. |
| setFieldStyle( name STRING, style STRING ) | Change the style of a form field. |
