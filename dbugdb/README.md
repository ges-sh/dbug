# dbugdb

[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/ges-sh/dbug/dbugdb)

[dbugdb](https://godoc.org/github.com/ges-sh/dbug/dbugdb) is an sql.DB wrapper, which logs executed queries.

##### Usage example
```
  db, _ := sql.Open("postgres", databaseUrl)

  debugDB := dbugdb.New(db, logrus.StandardLogger())

  debugDB.QueryRowContext(context.Background(), "query", args...) 
  // prints query in logger debug level
```