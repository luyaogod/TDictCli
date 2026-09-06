---
title: "What is JSON?"
source: "fgl-topics/c_fgl_json_utils_basics.html"
breadcrumb: "Advanced features > JSON support > What is JSON?"
type: "concept"
---

# What is JSON?

> JSON (JavaScript Object Notation) is a well known lightweight data-interchange format for JavaScript.

A JSON string (or object) is a comma-separated list of name/value pairs, with a
`:` colon separating the key and the value. The list of name/value pairs is enclosed
in `{}` curly brackets.

The names are delimited by double-quotes.

The value can be a single numeric value, a double-quotes string, an array, or a sub-element.

Arrays are defined by a comma-separated list of values enclosed in `[]` square
brackets.

Sub-elements are defined inside {} curly brackets and define name/value pairs.

For
example:

```
{
   "cust_num":865234,
   "cust_name":"McCarlson",
   "order_ids":[234,3456,24656,34561],
   "address": {
          "street":"34, Sunset Bld",
          "city":"Los Angeles",
          "state":"CA"
        }
}
```

For more details, see <http://www.json.org>.
