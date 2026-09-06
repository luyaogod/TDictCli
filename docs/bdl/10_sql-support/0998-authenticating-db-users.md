---
title: "Authenticating DB users"
source: "fgl-topics/c_fgl_sql_programming_047.html"
breadcrumb: "SQL support > SQL programming > SQL security > Authenticating DB users"
type: "concept"
---

# Authenticating DB users

> Understanding how users are authenticated to the database server.

When connecting to a database server, the user must be identified
by the server. Once connected, the current user is authenticated and
identified by the db server, and the database system can then apply
specific privileges, audit user activity, and so on.

Database user authentication is typically achieved by specifying a login and password in the
`CONNECT TO` instruction. However, most database servers support additional user
authentication methods, trusted connections, LDAP authentication, Single Sign-On authentication, and
even specific pluggable authentication methods.

Follow these simple security patterns to avoid basic user authentication
problems:

- Make sure that application files installed on your production server have the appropriate file
  system permissions set. For regular users it is recommended that they have read-only
  access to program and resource files. If any OS user can replace a program file with
  another program, it could harm your database or retrieve sensitive private
  data.
- Each physical end user must have a specific database account.
  If several end users connect as the same db application account, they
  cannot be distinguished in the security and auditing system.
- For normal application users, always use database accounts with the minimum database privileges
  required to achieve the daily work (GRANT/REVOKE). For example, regular users are
  not given permission to execute Data Definition Language statements (drop
  tables).
- Instead of asking a name and password in a login dialog when an application starts, some
  applications hard code the db user names and passwords in the program code, in scripts or configurations
  files such as FGLPROFILE. This is not a good practice and must be avoided. If a login dialog is
  not appropriate, you must set up another user authentication method supported by the database
  server, such as Single Sign-On.

## Related links

**Related concepts**  

[Database user authentication](1090-database-user-authentication.md "Different database user authentication methods exist.")

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
