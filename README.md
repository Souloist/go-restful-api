# Example API in Go (Address Book)

## Setting Up Database Schema and Connections

Before you run the server, assuming you are using Postgres, run the following:

```
psql <db_name>
\i migrations/tables.sql
```

The application also retrieves database credentials from environment variables so set the following:

```
$ export DB_USERNAME=<username>
$ export DB_PASSWORD=<password>
$ export DB_URL=<dburl>
```
