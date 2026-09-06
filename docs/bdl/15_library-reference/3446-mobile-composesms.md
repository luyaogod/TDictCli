---
title: "mobile.composeSMS"
source: "fgl-topics/c_fgl_frontcall_mobile_composesms.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.composeSMS"
type: "concept"
---

# mobile.composeSMS

> Sends an SMS text to one or more phone numbers.

## Syntax

```
ui.Interface.frontCall("mobile", "composeSMS",
  [ recipients, content ],
  [ result ] )
```

1. recipients - A list of phone numbers, separated by commas. While the list
   uses commas to separate the phone numbers in the list, the list itself is enclosed in a single set
   of quotes.
2. content - The SMS message.
3. result - Holds a status message.

## Usage

The "`composeSMS`" front call opens the default messaging app with the information
provided as parameters, to let the end user send an SMS text to one or more phone numbers.

On a macOS desktop, the `composeSMS` front call will open the "Messages" app.

Consider using global phone numbers with a `+` plus sign, as described in
[[RFC3966]](http://www.ietf.org/rfc/rfc3966.txt).

The returned result string can take one of the following values:

- "`ok`": The SMS app could be started.
- "`failed`": The SMS could not be started.

Error [-6333](4483-genero-bdl-errors.md) is raised, if
there is no permission to compose an SMS on the mobile phone.

## Example

```
DEFINE result STRING
CALL ui.Interface.frontCall("mobile", "composeSMS",
     ["+332781211,+339956789", "This is the SMS text"],
     [result])
```
