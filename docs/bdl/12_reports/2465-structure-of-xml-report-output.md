---
title: "Structure of XML report output"
source: "fgl-topics/c_fgl_reports_008.html"
breadcrumb: "Reports > XML output for reports > Structure of XML report output"
type: "concept"
description: "The generated XML output contains the structure of the formatted pages, with page header, page trailer, and group sections. Every PRINTX instruction will generate a <Print> node with a list of <Item> ..."
---

# Structure of XML report output

The generated XML output contains the structure of the formatted pages, with page header,
page trailer, and group sections.

Every [`PRINTX`](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.") instruction will
generate a `<Print>` node with a list of `<Item>` nodes
containing the data. The XML processor can use this structure to format and render the output as
needed.

The output of an XML report will have the following node structure:

```
<Report ...>
  <PageHeader pageNo="...">
     ...
  </PageHeader>
  <Group>
    <BeforeGroup>
      <Print name="...">
        <Item name="..." type="..." value="..." isoValue="..." />
        <Item name="..." type="..." value="..." isoValue="..." />
        ...
      </Print>
      ...
    </BeforeGroup>
    <OnEveryRow>
      <Print name="...">
        <Item name="..." type="..." value="..." isoValue="..." />
        <Item name="..." type="..." value="..." isoValue="..." />
        ...
      </Print>
      ...
    </OnEveryRow>
    ...
    <AfterGroup>
      <Print name="...">
        <Item name="..." type="..." value="..." isoValue="..." />
        <Item name="..." type="..." value="..." isoValue="..." />
        ...
      </Print>
    </AfterGroup>
    ...
  </Group>
  ...
  <OnLastRow ...>
     ...
  </OnLastRow>
  
  <PageTrailer ...>
     ...
  </PageTrailer>
  
</Report>
```

## Related links

**Related concepts**  

[START REPORT](2470-start-report.md "The START REPORT instruction initializes a report execution.")
