---
title: "base.Channel.create"
source: "fgl-topics/c_fgl_ClassChannel_create.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.create"
type: "concept"
---

# base.Channel.create

> Create a new channel object.

## Syntax

```
base.Channel.create()
  RETURNS base.Channel
```

## Usage

Use the `base.Channel.create()` class method to create a channel object.

The new created object must be assigned to a program variable defined with the
`base.Channel` type.

## Example

```
DEFINE ch base.Channel 
LET ch = base.Channel.create()
```

For a complete example, see [Example 1: Using record-formatted data file](3009-example-1-using-record-formatted-data-file.md).
