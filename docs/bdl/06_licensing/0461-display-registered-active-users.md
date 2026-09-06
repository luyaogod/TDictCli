---
title: "Display registered active users information"
source: "genero-install-topics/t_license_controller_display_active_users.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Display registered active users"
type: "task"
---

# Display registered active users information

> Use this procedure to display the registered active user list of the current license.

To display the registered active user list, execute the command:

```
fglWrt -a info users
```

This example displays output when the connection is direct to the DVM via the
Genero Desktop Client
(GDC):

```
C:\Program Files\FourJs\Genero_Studio_3_00_22\fgl\bin>fglWrt.exe -a info users
License      : THM#XXXXXXXX
License key  : KKKKKKKKKKKK
Product      : Four Js Universal Compiler
Type         : Development version
Users        : 1/5
    FrontEndId2: {f3619311-3f9b-4a93-bba5-5d9b4ec6d84b}
        GUI Server 127.0.0.1:0 - Process Id 5144
            host-name: 4js1
            host-addr: ::ffff:127.0.0.1
            user-name: user1
            app-id: {f4afc501-f7c9-c9c6-e16f-2b881a8bf3f0}
            cx-mode: direct
```

This example displays output when the connection is through a web server to
the
DVM:

```
C:\Program Files\FourJs\Genero_Studio_3_00_22\fgl\bin>fglWrt.exe -a info users
License      : THM#XXXXXXXX
License key  : KKKKKKKKKKKK
Product      : Four Js Universal Compiler
Type         : Development version
Users        : 1/5
    GUI Server 127.0.0.1:47274
        Process Id 5596
```

where:

**License**

This is the value of the license number.

**License Key**

This is the value of the license key.

**Product**

The name of the product being licensed is displayed. In relation to Genero product, this can be,
for example:

- Four J's Universal Compiler (for Genero Enterprise and Genero Mobile products)
- Four Js Genero Report Writer (for Genero Report Engine)
- Genero Report Writer for Java and C#
  (for Genero Report Writer products)

**Type**

This is the type of license. This can be runtime version or development version.

**Users**

This is the number of active users over the maximum number of users allowed.

Additional information may be displayed about the client. This is known as the
`clientInfo`, which is information sent to the DVM as part of the protocol between
DVM and the front-end. For details, go to [Client information](0460-client-information.md "The client sends information when connecting to the DVM, which helps to identify the client for a session.").

## Related links

**Related reference**  

[The FGLDIR/lock directory](0457-the-fgldir-lock-directory.md "When using a local license, the license controller uses the FGLDIR/lock directory to store information (number of active users). A user running a BDL program must have access rights to this directory. Permissions can be restricted to specific users with FGLWRTUMASK.")
