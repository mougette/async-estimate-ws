# async-estimate-ws


## sqlite3 notes:

```brew install sqlite3```

### Init the db:
    ```$ sqlite3 mytestdb.db < init.sql```

### DB Changes:

#### Make changes through CLI:
    ```
    // Start sqlite3 CLI
    $ sqlite3 mytestdb.db
    // Through the cli you can do query or insert
    $ select * from users;
    // quit with
    $ .quit
    ```

#### Make changes through a file

Add your updates to a .sql file then run something like:

```$ sqlite3 mytestdb.db < newChanges.sql```
