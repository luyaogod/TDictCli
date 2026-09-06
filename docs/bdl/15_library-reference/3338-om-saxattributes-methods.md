---
title: "om.SaxAttributes methods"
source: "fgl-topics/c_fgl_ClassSaxAttributes_methods.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods"
type: "concept"
---

# om.SaxAttributes methods

> Methods of the om.SaxAttributes class.

| Name | Description |
| --- | --- |
| copy( attr om.SaxAttributes ) RETURNS om.SaxAttributes | Clones an existing SAX attributes object. |
| create() RETURNS om.SaxAttributes | Create a new SAX attributes object. |

| Name | Description |
| --- | --- |
| addAttribute( name STRING, value STRING ) | Appends a new attribute to the end of the list. |
| clear() | Clears the SAX attribute list. |
| getLength() RETURNS INTEGER | Returns the number of attributes in the list. |
| getName( index INTEGER ) RETURNS STRING | Returns the name of an attribute by position. |
| getValue( name STRING ) RETURNS STRING | Returns the value of an attribute by name. |
| getValueByIndex( index INTEGER ) RETURNS STRING | Returns an attribute value by position. |
| removeAttribute( index INTEGER ) | Delete an attribute by position. |
| setAttributes( attr om.SaxAttributes ) | Clears the list and copies the attributes passed. |
