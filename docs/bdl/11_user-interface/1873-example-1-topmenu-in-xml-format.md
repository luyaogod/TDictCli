---
title: "Example 1: Topmenu in XML format"
source: "fgl-topics/c_fgl_topmenus_012.html"
breadcrumb: "User interface > Form definitions > Topmenus > Examples > Example 1: Topmenu in XML format"
type: "concept"
description: "<TopMenu> <TopMenuGroup text=\"Form\" style=\"mystyle\"> <TopMenuCommand name=\"help\" text=\"Help\" image=\"quest\" /> <TopMenuCommand name=\"quit\" text=\"Quit\" acceleratorName=\"Alt-F4\"/> </TopMenuGroup> ..."
---

# Example 1: Topmenu in XML format

```
<TopMenu>
   <TopMenuGroup text="Form" style="mystyle">
      <TopMenuCommand name="help" text="Help" image="quest" />
      <TopMenuCommand name="quit" text="Quit" acceleratorName="Alt-F4"/>
   </TopMenuGroup>
   <TopMenuGroup text="Edit">
      <TopMenuCommand name="accept" text="Validate" image="ok" />
      <TopMenuCommand name="cancel" text="Cancel" image="cancel" />
   </TopMenuGroup>
   <TopMenuGroup text="Records">
      <TopMenuCommand name="append" text="Add" image="add" />
      <TopMenuCommand name="delete" text="Remove" image="delete" />
      <TopMenuCommand name="update" text="Modify" image="change" />
      <TopMenuSeparator/>
      <TopMenuCommand name="search" text="Query" image="find" />
   </TopMenuGroup>
</TopMenu>
```
