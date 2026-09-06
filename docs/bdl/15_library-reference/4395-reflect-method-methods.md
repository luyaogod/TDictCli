---
title: "reflect.Method methods"
source: "fgl-topics/c_fgl_ext_reflect_Method_methods.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class > reflect.Method methods"
type: "concept"
---

# reflect.Method methods

> Methods for the reflect.Method class.

| Name | Description |
| --- | --- |
| getName() RETURNS STRING | Returns the name of the method. |
| getParameterCount() RETURNS INTEGER | Returns the number of parameters of a method. |
| getParameterType( index INTEGER ) RETURNS reflect.Type | Returns the type of a parameter of a method represented by this `reflect.Method` object. |
| getReturnCount() RETURNS INTEGER | Returns the number of values returned by a method. |
| getReturnType( index INTEGER ) RETURNS reflect.Type | Returns the type of a return value of a method. |
| getSignature() RETURNS STRING | Returns the signature of a method, as a string. |
