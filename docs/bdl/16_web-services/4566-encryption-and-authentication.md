---
title: "Encryption and authentication"
source: "fgl-topics/c_gws_ssl_001.html"
breadcrumb: "Web services > Security > Encryption and authentication"
type: "concept"
---

# Encryption and authentication

> A scenario involving a person (Georges) and his bank guides you through the concepts of secured communication, certificates, and certificate authorities.

## Secured communications

Secured communications are
important. If an application wants to send or receive messages from a financial, business, or
personnel application on the web, it must be able to authenticate the origin of the message, ensure
that no malicious application has altered the original message, and ensure that no third party
application can intercept the message.

Suppose that a person named Georges wants to send a
message to his bank to transfer some money on the Internet. In this scenario, he faces the following
concerns:

1. The **privacy** of the message, since it includes his account number and the transfer
   amount.
2. The **integrity** of the message, since someone might try to modify the original message or
   substitute a different message in order to transfer the money to another account.
3. The **authentication** of the message, since the bank must ensure that the message was sent
   from the right person.

Message privacy

To keep a message private, use a cryptographic algorithm - a
technique that transforms a message into an encrypted form unreadable except by those it is intended
for. Once it is in this form, the message may only be interpreted through the use of a secret key.
There are two kinds of cryptography algorithms: symmetric and
asymmetric.

Symmetric means the sender and the receiver of a message have to
share the same key used to encrypt a clear message into an encrypted form, and then to decrypt it
back into the original message. If that key is kept secret, nobody other than the sender and the
receiver can read the message. However, the task of choosing a private key before communicating can
be problematic.

Asymmetric means that there are two different keys working as a
key-pair. One key is used to encrypt a message, and the second one is used to decrypt the encrypted
message back into its original form. This solves the problem of key sharing in the symmetric
cryptography algorithm, and makes it possible to receive secure messages, simply by publishing the
key used to encrypt messages (the public key), and keeping secret the key used to
decrypt messages (the private key). Anyone can encrypt a message using the public key,
but only the owner of the private key can read it.

> **Important:**
>
> The use of an
> asymmetric key-pair (public and private key), allows Georges to send private messages to his bank,
> simply by using the bank's public key to encrypt a message. Only the owner of the corresponding
> private key (the bank in this scenario) is able to read it.

Message integrity

To
guarantee the integrity of a message, send a concise summary of the original message. The receiver
of the message can create its own summary and compare it to the sender's summary. If they are
similar, the message is considered intact, meaning that no third party has modified the original
message.

Such a summary is called a message digest and is based on hash
algorithms that produce a fixed-length representation of variable-length messages. Message digests
are designed to make it very difficult (if not impossible) to determine the original message from a
summary.

The message digest must be sent to the receiver in a secure way to assure the message
integrity. This is achieved with a digital signature authenticating the sender and containing the
sender's message digest.

> **Important:**
>
> The use of message digests allows Georges' bank to verify that no one
> has modified the original message he sent.

Message authentication

To authenticate a message, add a digital signature to
that message.

A digital signature is another message, created by encrypting the
message digest, along with some other information, with the sender's private key. Anyone with the
corresponding public key can decrypt the digital signature. If an application is able to decrypt it,
it means the owner of the private key was able to encrypt it, proving that the message comes from
this sender and not from someone else.

Once the sender has been authenticated, the receiver
can compare the message digest integrated into the digital signature to the one it created from the
message it receives, in order to check the message integrity.

> **Important:**
>
> The use
> of digital signatures allows Georges' bank to verify that the message really comes from him.

## Certificates

An SSL/TLS certificate is a kind of digital identity card that associates the public key with a
unique digital thumbprint identifying an individual, a server, or any other entity.

Now that Georges is able to send a secured message to his bank, there is still a problem. How can
Georges be sure that the server he is connected to is really the bank's server and not a malicious
server?

Georges must be sure that the public key he is using to encrypt his message corresponds to the
bank's private key. Similarly, the bank needs to verify that the message signature it receives
corresponds to Georges' signature.

To identify a remote peer, use a certificate - a kind of digital identity card that
associates the public key with a unique digital thumbprint identifying an individual, a server, or
any other entity (known as the subject). It also includes the identification and
signature of the Certificate Authority that
issued the certificate, and the period of time during which the certificate is valid. It may have
additional information (or extensions) as well as administrative information for the Certificate
Authority's use, such as a serial number.

A standard X.509 certificate contains the following standard fields:

- Certificate version
- Serial number of the certificate
- The distinguished name of the certificate issuer
- The distinguished name of the certificate owner
- The validity period of the certificate
- The public key
- The digital signature of the issuer
- Signature algorithm used
- Zero or more certificate extensions

1. An example of a distinguished name is: **CN=Georges,E=georges@mycompany.com,OU=Sales,O=My
   Company Name,C=FR,S=France**
2. The CN (Common Name) of the distinguished name of the certificate owner corresponds to the
   certificate subject, and identifies the owner of that certificate.

## Certificate authorities

When a certificate authority signs a certificate, it is validating that the certificate is
valid.

Each time Georges sends a message to his bank, he will present his own certificate to the bank,
and will get the bank's certificate back. But as every one can create a certificate in the name of
Georges, a higher authority that confirms the validity of a certificate is necessary. The bank must
be sure it is Georges' certificate, and that no one else has taken his identity. Similarly, Georges
needs an authority that confirms that the certificate coming from the server is really the bank's
certificate.

The solution to validating a certificate is to sign it with a trusted certificate called
**certificate authority**. This is a certificate in which an application has total confidence
concerning the validity of the certificates it has signed. Before signing a certificate, a
certificate authority must proceed with a strict identification of the owner of that
certificate.

> **Important:**
>
> The private key associated to a Certificate Authority must be managed with care, as it is the
> entity in charge of the validity of all other certificates it has signed.

There are several companies (such as *VeriSign*, *GlobalSign* or *RSA Security*)
that have established themselves as certificate authorities and provide the following services over
the Internet:

- Verifying certificate requests
- Processing certificate requests
- Issuing and managing certificates

It is also possible to create your own Certificate Authority, but it is up to you to manage it
securely.

**Root Certificate Authority**

A Certificate Authority signed by itself is called a Root Certificate Authority, meaning that the
certificate issuer is the same as the certificate subject. Most of the time, such a certificate
belongs to a company established as a Certificate Authority, and is used to sign certificate
requests coming from different companies that want their own Certificate Authority. If a client
certificate is signed by a Certificate Authority previously signed by a Root Certificate Authority,
the client certificate can be validated by the Root Certificate Authority even if the Certificate
Authority is not present.

For example, if a company wants to buy a Certificate Authority from VeriSign, VeriSign signs that
Certificate Authority with its own Root Certificate Authority. The company can then create
certificates with the Certificate Authority provided by VeriSign and connect to secure servers
without providing them with their own Certificate Authority. The secure server, of course, has to
know the VeriSign Root Certificate Authority.

**Certificate Chains**

A certificate authority may issue a certificate for another certificate authority. This means
that when an application wants to examine the certificate of the issuer, it must check all parent
certificates of that issuer until it reaches one in which it has confidence.

The certificate chain corresponds to the number of parent certificate authorities allowed to
validate a certificate.

**Certificate Authority List**

A Certificate Authority List is a list of all certificate authorities considered as trusted by
one application, classified by order of importance. Each of these certificates allows the
authentication of a certificate presented to that application from a remote peer.

With most applications, the Certificate Authority List is a concatenated file of all certificate
authorities.

## Certificates and private keys storage

The entire concept of security is based on the publication of the public key, and the privacy of
the associated private key. For maximum security, it is critical to restrict the access of the
private key to the owner of the certificate and associated private key.

Some companies provide systems to manage certificates and private keys in complete security.

UNIX™ systems

As the UNIX system is already able to restrict the access
of a file to only one person, simply restrict access to the private key to the owner of that key to
achieve a good level of security. This provides enough security to allow a Genero Web Services
client to perform secured communications in the name of the certificate and private key owner,
because access to the private key file is granted only if the correct user has logged in.

Windows® systems

Windows has an integrated **key store** system to
manage certificates and private keys. It allows the registration and the storage of X.509
certificate authorities, as well as personal X.509 certificates and their associated private keys
accessible only if the correct user has logged in. It is recommended that you store the certificate
and associated private key in the Windows key store
instead of in files on the disk.

## Related links

**Related concepts**  

[Missing certificates](4577-missing-certificates.md "Identifying missing certificates.")

[Error: Peer certificate is issued by a company not in our CA list](4576-error-peer-certificate-is-issued-by-a-company-not-in-our-ca.md "When a client connects to a server using HTTPS, the client needs to trust the server it is in communication with. So the client needs to add the server's CAs (certificate authorities lists) to its trusted CAs.")
