# isolation
* the thrid property in ACID
    - Atomicity &rarr; all operations in a transaction either complete or the whole transaction fails
    - Concurrency &rarr; the state of the db must be valid after a transaction completes or fails; all constraints must be satisfied
    - Isolation &rarr; concurrent transactions must not affect each other
    - Durability &rarr; data written by a successful transaction must be stored in persistent storage
* idea is this: concurrent transactions must not affect each other
### how transactions may affect each other
#### read phenomena 
1. dirty read 💩
    * transaction reads data written by another concurrent, uncommmitted transaction
    * we do not know if the other transaction will, eventually, be committed or rolled back; we might end up using incorrect data in case of rollback
2. non-repeatable read 🤡
    * transaction reads the same row twice and sees different values because said row has been modified by another committed transaction
3. phantom read 👻
    * transaction re-executes a query to find rows that satisfy a condition and sees a diffrent set of rows because said rows have been modified by another committed transaction
    * non-repeatable read for multiple rows
4. serialisation anomaly ❗
    * the result of a group of concurrent, committed transactions is impossible to achieve because there is no way to run the queries of said transactions in any order without overlapping
### how to solve the problem
#### standard isolation levels
* ANSI has four standard isolation levels
    - level 1: read, uncommitted transactions
    - level 2: read, committed transactions
    - level 3: repeatable read
    - level 4: serialisable
#### read uncommitted isolation level
* a transaction can see data written by uncommitted transactions; this is how/why dirty reads occur
#### read committed isolation level
* a transaction can see data written by committed transactions
* dirty reads are not possible
#### repeatable reads isolation level
* a read query always returns the same result no matter how many times it is executed
* the condition should hold even after other concurrent transactions commit changes that satisfy the query
#### serialisable isolation level
* achieves the same result when transactions are executed serially, in the same order, instead of concurrently
* there exists at least one way to order the concurrent transactions w/o overlap
### isolation levels and read phenomena
#### mySQL
* start a mySQL docker container

    ```bash
        docker exec -it mysql8 mysql -uroot -ppassword123 fnbank
    ```

* in the mySQL environment, run the following

    ```mysql
        select @@trasaction_isolation;
        select @@global.trasaction_isolation;
        set session transaction isolation level read uncommitted;
    ```

* you can set the isolation level you want
* 

    ||read uncommitted|read committed|read repeatable|serialisable|
    |:---:|:---:|:---:|:---:|:---:|
    |dirty read|||||
    |non-repeatable|||||
    |phantom read|||||
    |serialisation anomaly|||||

#### postgres

    ||read uncommitted|read committed|read repeatable|serialisable|
    |:---:|:---:|:---:|:---:|:---:|
    |dirty read|||||
    |non-repeatable|||||
    |phantom read|||||
    |serialisation anomaly|||||