---
title: "cordova.listPlugins"
source: "fgl-topics/c_fgl_frontcall_cordova_listPlugins.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.listPlugins"
type: "concept"
---

# cordova.listPlugins

> Returns the list of available Cordova plugins.

## Syntax

```
ui.Interface.frontCall("cordova", "listPlugins",
  [ ], [plugins] )
```

1. plugins - Is a `DYNAMIC ARRAY OF STRING`, to hold the list of
   available plugins.

## Usage

The `listPlugins` front call fetches the names of available Cordova plugins
bundled with the app.

The result variable must be of type `DYNAMIC ARRAY OF
STRING`.

## Example

```
DEFINE plugins DYNAMIC ARRAY OF STRING
CALL ui.interface.frontcall( "cordova", "listPlugins", [], [plugins] )
```
