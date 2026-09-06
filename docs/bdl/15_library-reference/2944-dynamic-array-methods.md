---
title: "DYNAMIC ARRAY methods"
source: "fgl-topics/c_fgl_Arrays_ARRAY_methods.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods"
type: "concept"
description: "Note: While designed for dynamic arrays, these methods can also be used on static arrays. However, the semantics differ when applied to static arrays; these variations are outside the scope of this ..."
---

# DYNAMIC ARRAY methods

> **Note:**
>
> While designed for dynamic arrays, these methods can also be used on static arrays. However, the
> semantics differ when applied to static arrays; these variations are outside the scope of this
> documentation.

| Name | Description |
| --- | --- |
| appendElement( ) | Adds a new element to the end of the array. |
| clear( ) | Removes all elements of the array. |
| copyTo( dst dynamic-array-type ) | Copies a complete array to the destination array passed as parameter. |
| deleteElement( index INTEGER ) | Removes an element from the array. |
| getLength( ) RETURNS INTEGER | Returns the length of the array. |
| insertElement( index INTEGER ) | Inserts a new element at the given index. |
| search( key STRING, value value-type ) RETURNS INTEGER | Scans the array to find an element that matches the search parameter. |
| searchRange( key STRING, value value-type, from INTEGER, to INTEGER ) RETURNS INTEGER | Scans the array to find an element that matches the search parameter. |
| sort( key STRING, reverse BOOLEAN ) | Sorts the rows in the array. |
| sortByComparisonFunction( key STRING, reverse BOOLEAN, compar FUNCTION ( s1 STRING, s2 STRING ) RETURNS INTEGER ) | Sorts the rows in the array by using a comparison function. |
