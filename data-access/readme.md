# Install mysql database:

1. Create docker container

```
docker run --name data-access-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
```

2. Container shell access

```
docker exec -it data-access-mysql bash
```

3. Log into DBMS

```
mysql -u root -p
```

4. Run SQL script

```
create database recordings;

use recordings;

source /a-tour-of-go/data-access/create-tables.sql;
```