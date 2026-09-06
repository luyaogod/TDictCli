---
title: "reflect.Type methods"
source: "fgl-topics/c_fgl_ext_reflect_Type_methods.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods"
type: "concept"
---

# reflect.Type methods

> Methods for the reflect.Type class.

| Name | Description |
| --- | --- |
| reflect.Type.typeOf( val any-type ) RETURNS reflect.Type | Creates a new `reflect.Type` object representing a type. |

| Name | Description |
| --- | --- |
| getAttribute( name STRING ) RETURNS STRING | Returns the value of a definition attribute. |
| getElementType() RETURNS reflect.Type | Returns the type of the elements in an array or in a dictionary. |
| getFieldCount() RETURNS INTEGER | Returns the number of fields of record type. |
| getFieldName( index INTEGER ) RETURNS STRING | Returns the name of a member of a record type. |
| getFieldType( index INTEGER ) RETURNS reflect.Type | Returns the type of a member of a record type. |
| getFieldTypeByName( name STRING ) RETURNS reflect.Type | Returns the type of a member of a record type, from the member name. |
| getKind() RETURNS STRING | Returns the kind of a type. |
| getMethod( index INTEGER ) RETURNS reflect.Method | Returns a `reflect.Method` object of record with methods or an interface. |
| getMethodCount() RETURNS INTEGER | Returns the number of methods in a record with methods or in an interface. |
| hasAttribute( name STRING ) RETURNS BOOLEAN | Indicates if a definition attribute exists in a type. |
| isAssignableFrom( x reflect.Type ) RETURNS BOOLEAN | Checks if this type can be assigned from another type. |
| toString() RETURNS STRING | Returns the name of a type. |
