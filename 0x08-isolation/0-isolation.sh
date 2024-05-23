###################################################
# terminal #1                                     #
###################################################

### open a terminal window
docker exec -it mysql8 mysql -uroot -ppassword123 fnbank

#inside a `mysql` session
select @@trasaction_isolation; #
select @@global.trasaction_isolation; #
set session transaction isolation level read uncommitted; # read uncommitted
begin;


###################################################
# terminal #2                                     #
###################################################

### open another terminal window
docker exec -it mysql8 mysql -uroot -ppassword123 fnbank

# inside a `mysql` session
select @@trasaction_isolation; #
select @@global.trasaction_isolation; #
set session transaction isolation level read uncommitted; # read uncommitted
begin;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. each has 100.0 KES. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # remove KES 10.00 from account 1
select * from accounts where id = 1; # account 1 has 90.00 KES

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 90.00 KES; dirty read

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
commit;

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
commit;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level read committed; #
select @@transaction_isolation; #read committed
begin;

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level read committed;
select @@transaction_isolation; #read committed
begin;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 90.00 KES the rest have 100.00 KES. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 90.00 KES; the rest have 100.00 KES. account_id 1, 2 and 3 respectively
select * from accounts where id = 1; # account 1 has 90.00 KES

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # remove KES 10.00 from account 1

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 90.00 KES, not 80.00; no dirty read
select * from accounts where balance >= 90; # account 1 has 90.00 KES; it is returned also

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
commit; # account 1 has 80.00 KES now


###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 80.00 KES now. recall that the previous query returned balance = 90.00 KES; non-repeatable read
select * from accounts where balance >= 90; # account 1 has 80.00 KES; it is not returned. this is a phantom read

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
commit;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level repeatable read; #
select @@transaction_isolation; #repeatable read
begin;

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level repeatable read;
select @@transaction_isolation; #repeatable read
begin;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 80.00 KES the rest have 100.00 KES. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 80.00 KES; the rest have 100.00 KES. account_id 1, 2 and 3 respectively
select * from accounts where id = 1; # account 1 has 80.00 KES
select * from accounts where balance >= 80; # account 1 has 80.00 KES; it is returned also

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # remove KES 10.00 from account 1
select * from accounts; #  account 1 has 70.00 KES; the rest have 100.00 KES
commit; # account 1 has 70.00 KES now

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 80.00 KES, not 70.00; no dirty read
select * from accounts where balance >= 80; # account 1 has 70.00 KES, however, it is returned because session 2 still thinks the balance is 80.00; this is repeatable read achieved, no phantom read
select * from accounts where balance >= 80; # account 1 is returned; repeatable read achieved, no phantom read
update accounts set balance = balance - 10 where id = 1; # remove KES 10.00 from account 1
select * from accounts; #  account 1 has 60.00 KES; the rest have 100.00 KES

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
commit; #account 1 has 60.00 KES now

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level serializable; #
select @@transaction_isolation; #serializable
begin;

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level serializable; #
select @@transaction_isolation; #serializable
begin;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 60.00 KES the rest have 100.00 KES. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 60.00 KES; the rest have 100.00 KES. account_id 1, 2 and 3 respectively
select * from accounts where id = 1; # account 1 has 60.00 KES

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # query is blocked
### LOCK WAIT TIMEOUT error
### restart the transactions
begin; #
select * from accounts where id = 1; # account 1 has 60.00 KES
update accounts set balance = balance - 10 where id = 1; # remove 10.00 KES from account 1

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # error 1213 (40001): deadlock

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
rollback; # roll back
begin; # begin
select * from accounts where id = 1; # account 1 has 60.00 KES

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 60.00 KES

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
update accounts set balance = balance - 10 where id = 1; # remove 10.00 KES from account 1

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
commit; # lock is released

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts where id = 1; # account 1 has 50.00 KES
commit; #account 1 has 50.00 KES now

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level serializable; #
select @@transaction_isolation; #serializable
begin;

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
set session transaction isolation level serializable; #
select @@transaction_isolation; #serializable
begin;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
select * from accounts; ## 3 accounts. account 1 has 50.00 KES the rest have 100.00 KES. account_id 1, 2 and 3 respectively
select sum(balance) from accounts; # 250 (50 + 100 + 100)
insert into accounts (owner, balance, currency) values ("four", 250, "KES"); #create a new account
select * from accounts; # 4 accounts last one has 250.00 KES

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select * from accounts; # query blocks until session 1 releases the lock (commits)

###################################################
# terminal #1                                     #
####################################################
#inside a `mysql` session
commit; # commit; query in session 2 unblocks. returns 4 rows as expected

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
select sum(balance) from accounts; # 500 (50 + 100 + 100 + 250)
insert into accounts (owner, balance, currency) values ("five", 500, "KES"); #create a new account
select * from accounts; # 5 accounts last one has 500.00 KES
commit; #commit successful; no duplicate records. serialisation anomaly eliminated

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
begin; #
select * from accounts; # 5 accounts; last one has 500.00 KES
select sum(balance) from accounts; # 1000 (50 + 100 + 100 + 250 + 500)

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
begin; #
select * from accounts; # 5 accounts; last one has 500.00 KES
select sum(balance) from accounts; # 1000 (50 + 100 + 100 + 250 + 500)

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
insert into accounts (owner, balance, currency) values ("six", 1000, "KES"); #query blocks until session 2 releases share lock

###################################################
# terminal #2                                     #
###################################################
#inside a `mysql` session
insert into accounts (owner, balance, currency) values ("six", 1000, "KES"); #error 1213 (40001): deadlock. interestingly, the deadlock releases the lock, therefore, the query on session 1 unblocks
rollback;

###################################################
# terminal #1                                     #
###################################################
#inside a `mysql` session
commit;
select * from accounts; # 6 accounts; last one has 1000.00 KES