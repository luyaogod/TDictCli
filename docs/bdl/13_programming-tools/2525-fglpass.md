---
title: "fglpass"
source: "fgl-topics/c_gws_tools_fglpass.html"
breadcrumb: "Programming tools > Command reference > fglpass"
type: "concept"
---

# fglpass

> The fglpass tool allows you to encrypt passwords.

## Syntax

```
fglpass [options]
```

1. options are described in fglpass options.

## Options

| Command | Description |
| --- | --- |
| `-V` | Displays version information of the tool. |
| `-Vssl` | Displays OpenSSL version. |
| `-h` | Displays options for the tool. |
| `-e` | Encrypt the password with a RSA key or certificate and encode it in BASE64 form. |
| `-d` | Decode the BASE64 form of the password and decrypt it with a RSA private key. |
| `-w cert` | Windows® certificate name to encrypt the password (Windows only) |
| `-c cert` | File of the PEM-encoded certificate to encrypt the password. |
| `-k key` | File of the PEM-encoded private key to encrypt or decrypt the password. |
| `-enc64 file` | File to be BASE64 encoded (result to stdout) |
| `-dec64 file` | BASE64 encoded file to be decoded (result to stdout) |
| `-agent:port files` | Start password agent on specified port to serve the list of private key files. |
| `-gid` | When executing fglpass in agent mode (with `-agent` option), allows authentication to be performed for all users belonging to the group of current users executing the command.This option requires the FGLPROFILE entry [`security.global.agent.gid=true`](../16_web-services/4916-fglprofile-entries-for-web-services.md) for fglrun. |

## Usage

The fglpass command line tool allows you to:

- Encrypt a password using a RSA key or X.509 certificate and encode it in BASE64 form.
- Run a password agent that returns (in a protected way) the passwords that grant access to the
  different private keys used in all your applications.
- Encode a file in BASE64 form and decode it back.

For security reasons, it is recommended to avoid storing clear passwords in a file, or leave
private keys unprotected without a password. The fglpass command can be used to
encrypt passwords.

If an error occurs, the execution status of the command will be non-zero. This allows for the
detection of errors in scripts and makefiles when the command is run from those contexts.

## Related links

**Related concepts**  

[Encryption, BASE64 and password agent with fglpass tool](../16_web-services/4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent.")
