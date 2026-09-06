---
title: "Example 1: Start menu in XML format"
source: "fgl-topics/c_fgl_startmenus_010.html"
breadcrumb: "User interface > User interface programming > Start menus > Examples > Example 1: Start menu in XML format"
type: "concept"
description: "<StartMenu> <StartMenuGroup text=\"Ordering\" > <StartMenuCommand text=\"Orders\" exec=\"fglrun orders\" disabled=\"1\" /> <StartMenuCommand text=\"Customers\" exec=\"fglrun custs\" image=\"smiley\" /> ..."
---

# Example 1: Start menu in XML format

```
<StartMenu>
   <StartMenuGroup text="Ordering" >
      <StartMenuCommand text="Orders" exec="fglrun orders"
                    disabled="1" />
      <StartMenuCommand text="Customers" exec="fglrun custs"
                    image="smiley" />
      <StartMenuCommand text="Items" exec="fglrun items"
                    waiting="1" />
   <StartMenuCommand text="Reports" exec="fglrun reports"
                comment="Run reports" />
   </StartMenuGroup>
   <StartMenuGroup text="Configuration" >
      <StartMenuCommand text="Database" exec="fglrun dbseconf" />
      <StartMenuCommand text="Users" exec="fglrun userconf" />
      <StartMenuCommand text="Printers" exec="fglrun prntconf" />
   </StartMenuGroup>
</StartMenu>
```
