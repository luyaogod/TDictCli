---
title: "base.Application.isGWA"
source: "fgl-topics/c_fgl_ClassApplication_isGWA.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.isGWA"
type: "concept"
---

# base.Application.isGWA

> Indicates if the application runs in a browser.

## Syntax

```
base.Application.isGWA()
  RETURNS BOOLEAN
```

## Usage

This class method can be called to check if the program code is running in a browser as a GWA
application.

## Example

```
MAIN
  MENU "test"
     COMMAND "check"
       MESSAGE SFMT("isGWA = %1", base.Application.isGWA())
  END MENU
END MAIN
```
