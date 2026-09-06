---
title: "Order of precedence for database user specification"
source: "fgl-topics/c_fgl_Connections_021.html"
breadcrumb: "SQL support > Database connections > Database user authentication > Order of precedence for database user specification"
type: "concept"
description: "Database user login can be specified with different methods, as shown in this table. Precedence order if defined from top to bottom: Table 1. Database user login methods Connection Instruction ..."
---

# Order of precedence for database user specification

Database user login can be specified with different methods, as shown in this table. Precedence
order if defined from top to bottom:

| Connection Instruction | FGLPROFILE | Effect |
| --- | --- | --- |
| CONNECT TO "dbname" USER "user" USING "pswd"orDEFINE db STRING LET db = "dbname+username='username', password='pswd'" DATABASE db | N/A (ignored) | The user information in theUSER/USINGclause of the CONNECT TOinstruction or in the connection string of the DATABASEinstruction are used to identify the actual user.Connection string can also be used withCONNECT TO. |
| DATABASE dbnameorCONNECT TO "dbname" | No specificdbi.*entry | No user login and password is provided to the database server. Usually, the Operating System authentication takes place. |
| DATABASE dbnameorCONNECT TO "dbname" | dbi.default.userauth.callback = "fx" | Callback functionfxis called to get user name and password when connection instruction is executed. |
| DATABASE dbnameorCONNECT TO "dbname" | dbi.database.dbname.username = ... dbi.database.dbname.password = ... | The [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") default user name and password are used to connect to the database server.**Important:**NOT RECOMMENDED IN PRODUCTION! |
