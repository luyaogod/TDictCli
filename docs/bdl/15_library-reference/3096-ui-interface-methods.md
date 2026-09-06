---
title: "ui.Interface methods"
source: "fgl-topics/c_fgl_ClassInterface_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods"
type: "concept"
---

# ui.Interface methods

> Methods of the ui.Interface class

| Name | Description |
| --- | --- |
| ui.Interface.frontCall( moduleName STRING, functionName STRING, [ parameters ], [ results ] ) | `ui.Interface.frontCall` performs a function call to the current front-end. |
| ui.Interface.filenameToURI( path STRING ) RETURNS STRING | Converts a filename to a URI to be used as a web component image resource. |
| ui.Interface.getChildCount() RETURNS INTEGER | Get the number of children in a parent container. |
| ui.Interface.getChildInstances( name STRING ) RETURNS INTEGER | Get the number of child instances for a given program name. |
| ui.Interface.getContainer() RETURNS STRING | Get the parent container of the current program. |
| ui.Interface.getDocument() RETURNS om.DomDocument | Returns the DOM document of the abstract user interface tree. |
| ui.Interface.getFrontEndName() RETURNS STRING | Returns the type of the front-end currently in use. |
| ui.Interface.getFrontEndVersion() RETURNS STRING | Returns the version of the front-end currently in use. |
| ui.Interface.getName() RETURNS STRING | Returns the name of the program. |
| ui.Interface.getImage() RETURNS STRING | Returns the icon image of the program. |
| ui.Interface.getText() RETURNS STRING | Returns the title of the program. |
| ui.Interface.getType() RETURNS STRING | Returns the type of the program. |
| ui.Interface.getRootNode() RETURNS om.DomNode | Get the root DOM node of the abstract user interface. |
| ui.Interface.getUniversalClientName() RETURNS STRING | Returns the name of the front-end used for Universal Rendering. |
| ui.Interface.getUniversalClientVersion() RETURNS STRING | Returns the version of the front-end used for Universal Rendering. |
| ui.Interface.loadActionDefaults( path STRING ) | Load the default action defaults file. |
| ui.Interface.loadStartMenu( path STRING ) | Load the start menu file. |
| ui.Interface.loadToolBar( path STRING ) | Load a default/global toolbar file for all forms of the program. |
| ui.Interface.loadTopMenu( path STRING ) | Load a default/global topmenu file for all forms of the program. |
| ui.Interface.loadStyles( path STRING ) | Load the presentation styles file. |
| ui.Interface.refresh() | Synchronize the user interface with the front-end. |
| ui.Interface.setContainer( name STRING ) | Define the parent container for the current program. |
| ui.Interface.setImage( image STRING ) | Defines the icon image of the program. |
| ui.Interface.setName( name STRING ) | Define the name of the current program for the front-end. |
| ui.Interface.setText( title STRING ) | Defines the title for the program. |
| ui.Interface.setSize( h STRING, w STRING ) | Specify the initial size of the window container. |
| ui.Interface.setType( type STRING ) | Defines the type of the program for the front-end. |
