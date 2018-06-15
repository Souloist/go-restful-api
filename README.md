# Introduction to creating APIs in Go

## Setting Up Database Schema and Connections

Before you run the server, assuming you are using Postgres, run the following:

```
psql <db_name>
\i migrations/tables.sql
```

The application also retrieves database credentials from environment variables so set the following:

```
$ export DB_HOST=localhost
$ export DB_PORT=5432
$ export DB_USERNAME=<username>
$ export DB_PASSWORD=<password>
$ export DB_NAME=<dbname>
```