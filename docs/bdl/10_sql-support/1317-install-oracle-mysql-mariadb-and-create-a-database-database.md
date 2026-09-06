---
title: "Install Oracle MySQL/MariaDB and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagmys_003.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Installation (Runtime Configuration) > Install Oracle® MySQL/MariaDB and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: Install the Oracle® MySQL Server (or MariaDB) on your computer. Configure the server with the ..."
---

# Install Oracle MySQL/MariaDB and create a database - database configuration/design tasks

If you are tasked with installing and configuring the database, here is a list of steps to be
taken:

1. Install the Oracle® MySQL Server (or MariaDB) on your computer.
2. Configure the server with the appropriate storage engine.

   In order to have transaction support by default, you must use a storage engine that supports
   transactional tables, such as `INNODB`. In recent versions of MySQL, this is the
   default storage engine.
3. Consider setting the `sql_mode` configuration parameter to get the appropriate
   behavior of the MySQL server:
   1. When the `STRICT_TRANS_TABLES` option is set in the `sql_mode`
      parameter, numeric data truncation/overflow will produce SQL error -1264, and character strings too
      large for the target column will produce SQL error -1406.

      Without the `STRICT_TRANS_TABLES` option, MySQL will only produce the SQL
      warning -1265 (data truncated) when a character string is too large. However, numeric data
      truncation/overflow is also relaxed and produces warning -1264, when you might expect to get an SQL
      error and avoid invalid numeric values in your database. Therefore, it is recommended to use the
      `STRICT_TRANS_TABLES` option. See [CHAR and VARCHAR data types](1327-char-and-varchar-data-types.md) for more
      details.
   2. Blank padding of fetched `CHAR` data can be controlled with the
      `PAD_CHAR_TO_FULL_LENGTH` option of the `sql_mode` parameter.

      You can use this parameter to get `CHAR` values padded with blanks to their
      full length, but the result of the SQL `LENGTH()` function will be different since
      trailing blanks are significant for that function in MySQL. See [The LENGTH() function](1343-the-length-function.md) for more details.
   3. If your SQL statements use the `||` double-pipe concatenation operator, define
      the `PIPES_AS_CONCAT` option in the `sql_mode` parameter.

      See [String concatenation operator](1346-string-concatenation-operator.md) for more details.
4. The mysqld process must be started to listen to database client connections.
   See MySQL documentation for more details about starting the database server process.
5. Create a database user dedicated to your application, the application administrator.

   Connect as the MySQL root user, create the application administrator user and grant all
   privileges to this
   user:

   ```
   $ mysql -u root -p
   ...
   mysql> create user 'mysuser'@'localhost' identified by 'password';
   mysql> grant all privileges on *.* to 'mysuser'@'localhost';
   ```
6. Connect as the application administrator and create a MySQL database with the `CREATE
   DATABASE` statement, and specify the character set to be used for this database:

   ```
   $ mysql -u mysuser -p
   ...
   mysql> create database mydatabase
                 character set utf8mb4
                 collate utf8mb4_0900_as_cs;
   ```

   MySQL collation names include modifiers such as `_ai/_as` and
   `_ci/_cs`, to indicate if you want accent and character case sensitivity in your
   database. Consider to use the right modifiers, to have for example `WHERE 'é'='É'`
   evaluate to true or false. Use the `_bin` modifier (as in
   `utf8mb4_bin`), to distinguish any character, and when the sort order can be based on
   the binary value of characters. To check the client charset, database charset and
   collation:

   ```
   mysql> SELECT @@character_set_client, @@character_set_database, @@collation_database;
   ```
7. Create the application tables.

   Convert Informix® data types to MySQL data types.
   See [SQL types mapping: Oracle MySQL](1325-sql-types-mapping-oracle-mysql.md) for more details.
8. Configuring MySQL server and client for a secure connection through TLS/SSL
   1. MySQL server with OpenSSL support can automatically generate missing certificates and key files
      at startup. Check your data dir for "ca.pem",
      "server-cert.pem" and "server-key.pem" files.
   2. MySQL server parameter to enable encryption are basically:

      ```
      [mysqld]
      require_secure_transport=ON
      ssl_ca=ca.pem
      ssl_cert=server-cert.pem
      ssl_key=server-key.pem
      ```

      See MySQL documentation for more details about server side TLS/SSL
      configuration.
   3. Client side, locate the MySQL configuration file and add the following entries under the
      `[client]` section:

      ```
      [client]
      default-character-set=utf8
      ssl-mode=VERIFY_CA
      ssl-ca=/var/myapp/certificates/ca.pem
      ssl-cert=/var/myapp/certificates/client-cert.pem
      ssl-key=/var/myapp/certificates/client-key.pem
      ```
   4. Test with mysql command:

      ```
      mysql --host=myhost.mydomain.com \
         --user=user-name --password='pswd-or-token' dbname
      ```
   5. Test with a BDL program:

      Define the MySQL client configuration file in your FGLPROFILE
      file:

      ```
      dbi.database.mydb.mys.config = "/var/myapp/client/mysql_ssl.cnf"
      ```

      SQL Connection code:

      ```
      DEFINE dsrc, opts STRING
      LET dsrc = "test1@myhost.mydomain.com:3310"
      CONNECT TO SFMT("db1+driver='dbmmys',resource='mydb',source='%1'",dsrc)
            USER "mysuser" USING "fourjs"
      ```

      To check that the connection is encrypted:

      ```
      DEFINE name, value STRING
      PREPARE s1 FROM "SHOW SESSION STATUS LIKE 'Ssl_cipher'"
      EXECUTE s1 INTO name, value
      DISPLAY "Cipher: ", name, ": ", value
      ```

      Expected output:

      ```
      Cipher: Ssl_cipher: TLS_AES_256_GCM_SHA384
      ```
