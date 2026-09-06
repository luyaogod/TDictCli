---
title: "reflect.Value methods"
source: "fgl-topics/c_fgl_ext_reflect_Value_methods.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods"
type: "concept"
---

# reflect.Value methods

> Methods for the reflect.Value class.

| Name | Description |
| --- | --- |
| reflect.Value.copyOf( val any-type ) RETURNS reflect.Value | Creates a new `reflect.Value` object from a copy of an expression. |
| reflect.Value.valueOf( val any-type ) RETURNS reflect.Value | Creates a new `reflect.Value` object as a reference to the original variable. |

| Name | Description |
| --- | --- |
| appendArrayElement() | Appends a new element to an array. |
| assignToVariable( var any-type ) | Assigns this `reflect.Value` to a variable. |
| canAssignToVariable( var any-type ) RETURNS BOOLEAN | Checks if this `reflect.Value` can be assigned to a variable. |
| clear( ) | Removes all elements of the `reflect.Value` referencing a collection. |
| deleteArrayElement( index INTEGER ) | Deletes an element of an array. |
| getArrayElement( index INTEGER ) RETURNS reflect.Value | Returns an element of an array. |
| getCurrentValue( ) RETURNS any-type | Returns the actual (typed) value of this value object. |
| getDictionaryElement( key STRING ) RETURNS reflect.Value | Returns an element of a dictionary. |
| getDictionaryKeys( ) RETURNS DYNAMIC ARRAY OF STRING | Returns the keys of a dictionary. |
| getField( index INTEGER ) RETURNS reflect.Value | Returns a field by index for a record. |
| getFieldByName( name STRING ) RETURNS reflect.Value | Returns a field of a record, from the member name. |
| getInterfaceValue( ) RETURNS reflect.Value | Returns the value record referenced by this interface variable. |
| getLength( ) RETURNS INTEGER | Returns the number of elements in an array. |
| getType() RETURNS reflect.Type | Returns the `reflect.Type` object of a `reflect.Value` object. |
| hasKey( key STRING ) RETURNS BOOLEAN | Checks if existence of an element in a dictionary. |
| initializeToNull( ) | Initializes this `reflect.Value` to `NULL`. |
| insertArrayElement( index INTEGER ) | Inserts a new element into an array. |
| isNull( ) RETURNS BOOLEAN | Checks if this `reflect.Value` is `NULL`. |
| removeDictionaryElement( key STRING ) | Deletes an element of a dictionary. |
| set( x reflect.Value ) | Assigns the specified value to this value object. |
| toString( ) RETURNS STRING | Converts this `reflect.Value` to a character string. |
