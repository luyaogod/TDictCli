---
title: "Controlling semantics of AND / OR operators"
source: "fgl-topics/c_fgl_programs_OPTIONS_SHORT_CIRCUIT.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Compilation) > Controlling semantics of AND / OR operators"
type: "concept"
---

# Controlling semantics of AND / OR operators

> The OPTIONS SHORT CIRCUIT defines the semantics of AND/OR operators.

When using `OPTIONS SHORT CIRCUIT` at the beginning of a module, the runtime
system will optimize the evaluation of boolean expressions involving [`AND`](../08_language-basics/0620-and.md "The AND operator is the logical intersection operator.") and [`OR`](../08_language-basics/0621-or.md "The OR operator is the logical union operator.") operators, by using the *short-circuit evaluation* method (also
called *minimal evaluation* method). This behavior is enabled for the whole module.

By default, the behavior of `AND` and `OR` operators is to evaluate
all operands on the left and right side of the operator. In fact this is not required. If the left
operand of the `AND` evaluates to `FALSE`, there is no need to
evaluate the right operand, because the result of the `AND` operator will be false,
anyway. Similarly, when the left operand of an `OR` expression evaluates to
`TRUE`, there is no need to evaluate the right operand, since the result of the
boolean expression will be true, anyway.

This method can improve performances and simplify programming. However, existing code may rely on
the fact that all parts of a boolean expression are evaluated, especially when calling functions
that do some processing. By using the short-circuit evaluation method, it is unsure that the
function used in the right operand of `AND`/`OR` will be called,
because it depends on the result of the left operand.

By using short-circuit evaluation, it is possible to reference a dynamic array in the same
boolean expression, after checking that the index is in the current array element
range:

```
IF x<=arr.getLength() AND arr[x].order_date > TODAY THEN
     ...
END IF
```

With the default `AND` semantics, in this code, the right operand is also
evaluated. If the `x` index is greater than the array length, new array elements will
be automatically created in the expression on the right of the `AND` operator. To
avoid this situation, you are forced to write the following code, when `OPTIONS SHORT
CIRCUIT` is not used:

```
IF x<=arr.getLength() THEN
  IF arr[x].order_date > TODAY THEN
      ...
  END IF
END IF
```

## Related links

**Related concepts**  

[Boolean expressions](../08_language-basics/0594-boolean-expressions.md "This section covers boolean expression evaluation rules.")
