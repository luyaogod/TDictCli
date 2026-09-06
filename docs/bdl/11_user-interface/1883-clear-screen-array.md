---
title: "CLEAR SCREEN ARRAY"
source: "fgl-topics/c_fgl_record_display_CLEAR_SCREEN_ARRAY.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > CLEAR SCREEN ARRAY"
type: "concept"
---

# CLEAR SCREEN ARRAY

> The CLEAR SCREEN ARRAY instruction clears the values of all rows of the form list identified by the specified screen array.

## Syntax

```
CLEAR SCREEN ARRAY screen-array.*
```

1. screen-array is a screen array specified in the form.

## Usage

After executing a `DISPLAY ARRAY` or `INPUT ARRAY` instruction,
values remain in the form list identified by the [screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.").

> **Important:**
>
> The `CLEAR SCREEN ARRAY` instruction sets the [field modification flags](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.").

The `CLEAR SCREEN ARRAY` instruction automatically clears all rows of the list,
regardless of the view: a `TABLE`, `TREE`, `SCROLLGRID`,
or in a matrix of fields (an old-style/text-mode static screen array).

The `CLEAR SCREEN ARRAY` instruction replaces code which clears each individual
row through the use of a loop:

```
-- Clearing each row individually
FOR i=1 TO <screen-array-length>
  CLEAR screen-array[i].*
END FOR
-- Unique instruction to clear a list
CLEAR SCREEN ARRAY screen-array.*
```

Using the `CLEAR SCREEN ARRAY` instruction eliminates the need for calculating
the screen array length, a value which can change when using a `TABLE` container,
that can be resized.

The `CLEAR SCREEN ARRAY` instruction clears the field values and resets the TTY
attributes to `NORMAL`.

The `CLEAR SCREEN ARRAY` instruction is not needed if the program is always in the
context of a [dialog controlling the form
fields](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.").

## Example

```
...
   DISPLAY ARRAY cust_arr TO sa.*
   ...
   CLEAR SCREEN ARRAY sa.*
   ...
```

## Related links

**Related concepts**  

[CLEAR FORM](1882-clear-form.md "The CLEAR FORM instruction clears all fields in the current form.")

[CLEAR field-list](1884-clear-field-list.md "The CLEAR field-list instruction clears specific fields in the current form.")
