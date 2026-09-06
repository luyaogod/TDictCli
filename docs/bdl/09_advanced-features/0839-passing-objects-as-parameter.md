---
title: "Passing objects as parameter"
source: "fgl-topics/c_fgl_runtime_stack_params_java_object.html"
breadcrumb: "Advanced features > Runtime stack > Passing objects as parameter"
type: "concept"
description: "Like other object oriented programming languages, objects of built-in classes or Java classes are passed by reference. It would not make much sense to pass an object by value, actually. The runtime ..."
---

# Passing objects as parameter

Like other object oriented programming languages, objects of built-in classes or Java classes are
passed by reference. It would not make much sense to pass an object by value, actually. The
runtime pushes the reference of the object on the stack (i.e. the object handler is passed
by value), and the reference is then popped to the receiving object variable in the
function. The function can then be used to manipulate the original object.

```
MAIN
  DEFINE ch base.Channel 
  LET ch = base.Channel.create()
  CALL open(ch,arg_val(1))
  CALL ch.close()
END MAIN

FUNCTION open(x base.Channel, fn STRING) RETURNS ()
  CALL x.openFile(fn,"r")
END FUNCTION
```

## Related links

**Related concepts**  

[OOP support](0942-oop-support.md "Describes Object Oriented Programming basics in the language.")

[The Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
