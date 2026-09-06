---
title: "base.Application.isMobile"
source: "fgl-topics/c_fgl_ClassApplication_isMobile.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.isMobile"
type: "concept"
---

# base.Application.isMobile

> Indicates if the application runs on a mobile device.

## Syntax

```
base.Application.isMobile()
  RETURNS BOOLEAN
```

## Usage

This class method can be called to check if the program code is running on a smartphone or tablet
device. The method will return `TRUE` if the program executes in standalone mode
(meaning the runtime system is on a mobile device).

## Example

```
MAIN
  MENU "test"
     COMMAND "check"
       MESSAGE SFMT("isMobile = %1", base.Application.isMobile())
  END MENU
END MAIN
```
