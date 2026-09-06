---
title: "Working with objects"
source: "fgl-topics/c_fgl_oop_006.html"
breadcrumb: "Advanced features > OOP support > Working with objects"
type: "concept"
---

# Working with objects

> This topic describes basic object usage in Genero BDL.

## Instantiating objects

In order to instantiate an object in your program:

1. Define an object variable using the class identifier.
2. Instantiate the object; this is usually done by invoking a class method.

An object variable only contains a reference to the object. For example, when passed to a
function, only the reference to the object is copied onto the stack.

In the following code example, the object referenced by the variable `n` is
instantiated using the `create()` class method of the DomDocument class. The object
referenced by the variable `b` is instantiated using the
`getDocumentElement()` object method of the DomDocument class. This method returns
the DomNode object that is the root node of the DomDocument object referenced by
`n`:

```
DEFINE n om.DomDocument, b DomNode
LET n = om.DomDocument.create("Stock")
LET b = n.getDocumentElement()
```

## Destroying objects

Objects created during program execution do not need to be explicitly destroyed. This is done
automatically by the runtime system, based on a reference
counter.

```
MAIN
  DEFINE d om.DomDocument
  LET d = om.DomDocument.create("Stock")  -- Reference counter = 1
END MAIN  -- d is removed, reference counter = 0 => object is destroyed.
```

When an object is referenced by several variables, an internal counter is incremented and
decremented:

```
MAIN
  DEFINE d1, d2 om.DomDocument
  LET d1 = om.DomDocument.create("Stock")  -- Reference counter = 1
  LET d2 = d1                              -- Reference counter = 2
  LET d1 = NULL                            -- Reference counter = 1
  LET d2 = NULL                            -- Reference counter = 0, object is destroyed
END MAIN
```

## Using object references within functions

Object references can be passed to and returned from functions.

In this example, the function creates the object and returns its reference on the
stack:

```
MAIN
  DEFINE x om.DomDocument
  LET x = createStockDomDocument()
END MAIN

FUNCTION createStockDomDocument()
  DEFINE d om.DomDocument
  LET d = om.DomDocument.create("Stock")  -- Reference counter = 1
  RETURN d
END FUNCTION  -- Reference counter is still 1 because d is on the stack
```

Another part of the program can get the result of that function and pass it as a parameter to
another function:

```
MAIN
  DEFINE x om.DomDocument
  LET x = createStockDomDocument()
  CALL writeStockDomDocument( x )
END MAIN

FUNCTION createStockDomDocument()
  DEFINE d  om.DomDocument
  LET d = om.DomDocument.create("Stock")
  RETURN d
END FUNCTION

FUNCTION writeStockDomDocument( d )
  DEFINE d  om.DomDocument
  DEFINE r om.DomNode
  LET r = d.getDocumentElement()
  CALL r.writeXml("Stock.xml")
END FUNCTION
```

## Invoking class and object methods

Class methods must be invoked with the package and class
name:

```
DEFINE d  om.DomDocument
LET d = om.DomDocument.create("Stock")
```

Object methods are invoked with the variable referencing the
object:

```
DEFINE ch base.Channel
LET ch = base.Channel.create()
CALL ch.openFile("myfile.txt","r")
```

If a method returns an object reference, it can be directly used to invoke another method of the
returned object:

```
DEFINE s STRING
LET s = "ABC"
DISPLAY s.subString(1,2).toLowerCase()
```
