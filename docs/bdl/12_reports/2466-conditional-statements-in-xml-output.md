---
title: "Conditional statements in XML output"
source: "fgl-topics/c_fgl_reports_004.html"
breadcrumb: "Reports > XML output for reports > Conditional statements in XML output"
type: "concept"
description: "Principle If PRINTX commands are used inside program flow control instructions like IF , CASE , FOR , FOREACH and WHILE , the XML output will contain additional nodes to identify such conditional ..."
---

# Conditional statements in XML output

## Principle

If [`PRINTX`](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.") commands are used
inside program flow control instructions like `IF`, `CASE`,
`FOR`, `FOREACH` and `WHILE`, the XML output will
contain additional nodes to identify such conditional print instructions.

That information can be useful to process an XML report output.

## FOR … END FOR

```
<For>
  <ForItem>
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </ForItem>
  ...
</For>
```

## WHILE … END WHILE

```
<While>
  <WhileItem>
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </WhileItem>
  ...
</While>
```

## FOREACH … END FOREACH

```
<Foreach>
  <ForeachItem>
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </ForeachItem>
  ...
</Foreach>
```

## CASE … END CASE

```
<Case>
  <When id="position">
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </When>
  ...
</Case>
```

## IF … THEN … ELSE … END IF

```
<If>
  <IfThen>
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </IfThen>
  <IfElse>
    <Print name="...">
      <Item name="..." type="..." value="..." isoValue="..." />
    </Print>
    ...
  </IfElse>
</If>
```
