---
title: "localStorage.keys"
source: "fgl-topics/c_fgl_frontcall_localstorage_keys.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls > localStorage.keys"
type: "concept"
---

# localStorage.keys

> Returns the list of defined local storage keys.

## Syntax

```
ui.Interface.frontCall("localStorage", "keys",
   [], [key-list] )
```

1. key-list list of key names as a JSON array of strings.

## Usage

The `keys` function returns the current local storage keys
defined on the front-end platform.

The list of keys is returned in a string variable, as a JSON array of strings.

Convert the JSON array to a BDL dynamic array with [util.JSON.parse](3582-util-json-parse.md "Parses a JSON string and fills program variables with the values."):

```
IMPORT util
...

DEFINE key_list STRING
DEFINE key_array DYNAMIC ARRAY OF STRING

CALL ui.Interface.frontCall("localStorage", "keys", [], [key_list] )
CALL util.JSON.parse( key_list, key_array )
```
