---
title: "XMLElementNillable"
source: "fgl-topics/c_gws_XML_attribute_XMLElementNillable.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLElementNillable"
type: "concept"
description: "The XMLElementNillable attribute specifies that all members of a RECORD have to be serialized as xsi:nil=\"true\" when NULL . This has the same effect as defining each individual element with the ..."
---

# XMLElementNillable

The XMLElementNillable attribute specifies that all members of a `RECORD` have to
be serialized as `xsi:nil="true"` when `NULL`. This has the
same effect as defining each individual element with the [XMLNillable](5023-xmlnillable.md) attribute.

The XMLElementNillable attribute can only be set on records defined by
`TYPE` or `DEFINE`.

The XMLElementNillable attribute supports the record defined with `LIKE` clause to
use the table definition from a database
schema:

```
DEFINE rec RECORD ATTRIBUTE(XMLElementNillable) LIKE customer.*
```

If
columns in the database table are allowed to be null, you must add the XMLElementNillable attribute
to the record definition.

The XMLElementNillable attribute is not inherited by sub records.

## Basic example

```
DEFINE rec RECORD ATTRIBUTE(XMLElementNillable)
  val1 STRING,
  val2 STRING
END RECORD

LET rec.val1 = "Hello"
LET rec.val2 = NULL
```

In the resulting XML document, "val2" is serialized with the
`xsi:nil="true"`:

```
<rec>
  <val1>Hello</val1>
  <val2 xsi:nil="true"/>
</rec>
```

## Example using a sub record

Since XMLElementNillable attribute is not inherited by sub records, you must set it on each
sub-record if needed:

```
DEFINE rec RECORD ATTRIBUTE(XMLElementNillable)
  val1 STRING,
  val2 STRING,
  subrec RECORD
    val3 STRING
  END RECORD
END RECORD
```

In the resulting XML document, for instance, only "val1" and "val2" are serialized as
`xsi:nil="true"` when `NULL`, but "val3" is not.

To have "val3" serialized as `xsi:nil="true"` when `NULL`,
define:

```
DEFINE rec RECORD ATTRIBUTE(XMLElementNillable)
  val1 STRING,
  val2 STRING,
  subrec RECORD ATTRIBUTE(XMLElementNillable)
    val3 STRING
  END RECORD
END RECORD
```

## Examples with XMLNillable and XMLOptional

The XMLElementNillable attribute does not allow values, so you cannot write
`XMLElementNillable="preferred"` as you do for XMLNillable.

If [XMLOptional](5022-xmloptional.md) is set on a record member,
XMLElementNillable is not taken into account.

If you need XMLNillable to take precedence over XMLOptional, you must then specify
`XMLNillable="preferred"` on the record member.

In the following sample, if "val2" is `NULL`, it will serialize to the tag not
being present because XMLOptional hides XMLElementNillable. On the other hand, if "val3" is
`NULL`, it will be serialized as `xsi:nil="true"`, because the record
member is defined with XMLOptional and XMLNillable (preferred):

```
DEFINE rec RECORD ATTRIBUTE(XMLElementNillable)
  val1 STRING,
  val2 STRING ATTRIBUTE(XMLOptional),
  val3 STRING ATTRIBUTE(XMLOptional,XMLNillable="preferred")
END RECORD

LET rec.val1 = "Hello"
LET rec.val2 = NULL
LET rec.val3 = NULL
```

Will be serialized as:

```
<rec>
  <val1>Hello</val1>
  <val3 xsi:nil="true"/>
</rec>
```
