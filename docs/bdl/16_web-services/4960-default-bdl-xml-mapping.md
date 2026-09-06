---
title: "Default BDL/XML mapping"
source: "fgl-topics/c_gws_XML_attributes_006.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > Default BDL/XML mapping"
type: "concept"
---

# Default BDL/XML mapping

> By default, Genero Web Services converts BDL variables to standard XML Schema (XSD) data types for use in web service messages.

The XML data types conform to the standard XML Schema Definition (XSD):

In addition, the [Web Service Style](4673-choosing-a-web-services-style.md "Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.")
that you use determines what default [XMLName](5037-xmlname.md) attributes are
assigned to variables.

| Data type of BDL variable | Default XML data type |
| --- | --- |
| `BYTE` | xsd:base64binary |
| `CHAR` | xsd:string |
| `DATE` | xsd.date |
| `DATETIME YEAR TO FRACTION(1-5)` | xsd:dateTime |
| `DATETIME YEAR TO SECOND` | xsd:dateTime |
| `DATETIME YEAR TO HOUR` | xsd:dateTime |
| `DATETIME YEAR TO MINUTE` | xsd:dateTime |
| `DATETIME YEAR TO YEAR` | xsd:gYear |
| `DATETIME YEAR TO MONTH` | xsd:gYearMonth |
| `DATETIME YEAR TO DAY` | xsd:date |
| `DATETIME MONTH TO MONTH` | xsd:gMonth |
| `DATETIME MONTH TO DAY` | xsd:gMonthDay |
| `DATETIME DAY TO DAY` | xsd:gDay |
| `DATETIME HOUR TO HOUR` | xsd:time |
| `DATETIME HOUR TO MINUTE` | xsd:time |
| `DATETIME HOUR TO SECOND` | xsd:time |
| `DATETIME HOUR TO FRACTION(1-5)` | xsd:time |
| `DECIMAL` | xsd:decimal |
| `FLOAT` | xsd:double |
| `INTEGER` | xsd:int |
| `INTERVAL` | xsd:duration |
| `SMALLFLOAT` | xsd:float |
| `SMALLINT` | xsd:short |
| `STRING` | xsd:string |
| `TEXT` | xsd:string |
| `VARCHAR` | xsd:string |
| `TINYINT` | xsd:byte |
| `BIGINT` | xsd:long |
| `BOOLEAN` | xsd:boolean |
