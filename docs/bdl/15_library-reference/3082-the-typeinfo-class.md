---
title: "The TypeInfo class"
source: "fgl-topics/c_fgl_ClassTypeInfo.html"
breadcrumb: "Library reference > Built-in packages > The base package > The TypeInfo class"
type: "concept"
description: "The base.TypeInfo class creates a DOM node from a structured program variable. You can use this class to do program variables introspection, in order to get the data type names and values in a DOM ..."
---

# The TypeInfo class

The `base.TypeInfo` class creates a DOM node from a structured program
variable.

You can use this class to do program variables introspection, in order to get the data type names
and values in a DOM node, that can be traversed with the [XML
utility classes](3280-the-om-package.md "These topics cover the built-in classes of the om package"), or to be serialized in a file for export purpose.

The `base.TypeInfo` class does not have to be instantiated: you can directly use
the `create()` method.

A serialized DOM node created by `base.TypeInfo` class looks for example as
follows:

```
<?xml version="1.0"? encoding="ISO-8859-1">
<Record>
  <Field type="INTEGER" value="234" name="key"/>
  <Field type="CHAR(20)" value="Johnson" name="lastname"/>
  <Field type="DATE" value="12/24/1962" name="birthdate"/>
</Record>
```

Steps to use the class:

- Define a variable with the `om.DomNode` type.
- Create a DOM node with `base.TypoInfo.create( var )` and
  assign it to the DOM node variable.
- Use the created DOM node.

For example, to convert a list of database records to XML, fetch rows from a database table in a
structured array, specify the array as the input parameter for the
`base.TypeInfo.create()` method to create a new `base.DomNode` object,
and serialize the resulting DOM node to a file by using the `node.writeXml()` method.
You can then pass the resulting file to any application that is able to read XML for input.

> **Note:**
>
> Consider using the JSON interface to serialize and deserialize program variables.

## Child topics

- [base.TypeInfo methods](3083-base-typeinfo-methods.md)
