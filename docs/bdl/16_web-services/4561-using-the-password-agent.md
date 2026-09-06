---
title: "Using the password agent"
source: "fgl-topics/c_gws_ssl_fglpass_006.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Using the password agent"
type: "concept"
---

# Using the password agent

> Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.

The fglpass tool can be started as an agent with the
`-agent:tcp-port` option, to allow BDL applications requiring
passwords, to grant access to private keys without providing the passwords.

Passwords are provided once for each private key at the fglpass agent
startup.

The fglpass agent has the following fglrun authentication
methods:

- By default, the agent allows only fglrun to access the passwords, if the OS
  user executing fglpass and fglrun are the same.
- On UNIX® platforms, the
  fglpass agent can be started with the `-gid` option, to allow all
  users belonging to the OS group of users executing the fglpass
  program:

  ```
  fglpass -gid -agent:4242 myprivate1.pem myprivate2.pem ...
  ```

  For fglrun programs requiring the group-based agent authentication method, you
  need to set the following FGLPROFILE entry:

  ```
  security.global.agent.gid=true
  ```

Authentication and data encryption are performed between the BDL application and the agent to
guarantee passwords confidentiality, and the passwords are also stored in encrypted form in the
agent memory.

1. To start the password agent at port number `4242` and to serve the BDL
   applications with the passwords of the private key RSAKey1.pem and
   DSAKey2.der, specify the option `-agent`, followed by a colon,
   followed by the port number where it will be reachable, followed by the list of private keys the
   agent will handle for all BDL
   applications.

   ```
   fglpass -agent:4242 RSAKey1.pem DSAKey2.der
   ```
2. The agent prompts you to silently enter the password of the different keys (the passwords are not
   displayed to the console while being typed). In this example, you
   have:

   ```
   Enter pass phrase for RSAKey1.pem:
   ```

   Followed by:

   ```
   Enter pass phrase for DSAKey2.der:
   ```
3. Once all keys have been processed, the following message is displayed to notify that the agent is
   ready.

   ```
   Agent started
   ```
4. To enable one BDL application to use the password agent capability,
   set the entry called [security.global.agent](4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") in
   the FGLPROFILE file with the port number of the agent.

   In our
   example, with value 4242:

   ```
   security.global.agent = "4242"
   ```
