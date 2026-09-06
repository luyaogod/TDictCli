---
title: "security.PBKDF2 methods"
source: "fgl-topics/r_gws_SecurityPBKDF2_methods.html"
breadcrumb: "Library reference > Extension packages > The security package > The PBKDF2 class > PBKDF2 methods"
type: "reference"
---

# security.PBKDF2 methods

> Methods of the security.PBKDF2 class.

| Name | Description |
| --- | --- |
| security.PBKDF2.GenerateKey( password STRING, salt STRING, hash STRING, iter INTEGER, keySize INTEGER ) RETURNS STRING | Generates a password of a given size based on a human readable password using Password-Based Key Derivation Function 2 (PBKDF2) |
| security.PBKDF2.CheckKey( password STRING, salt STRING, hash STRING, iter INTEGER, hashedkey STRING ) RETURNS BOOLEAN | Validates a hashed key. |
