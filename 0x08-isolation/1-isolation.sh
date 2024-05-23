###################################################
# terminal #1                                     #
###################################################

### open a terminal window
docker exec -it postgres16 psql -U root -d fnbank

#inside a `postgresql` session
show trasaction isolation level; # default is read repeatable
begin; #
set trasaction isolation level read uncommitted; #
show transaction isolation level; # read uncommitted



###################################################
# terminal #2                                     #
###################################################

### open another terminal window
docker exec -it postgres16 psql -U root -d fnbank

# inside a `postgresql` session
show trasaction isolation level; # default is read repeatable
begin; #
set trasaction isolation level read uncommitted; #
show transaction isolation level; # read uncommitted

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 3 accounts. each has 100.0 KES. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
update accounts set balance = balance - 10 where id = 1 returning *; # remove KES 10.00 from account 1
select * from accounts where id = 1; # account 1 has 90.00 KES

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 100.00 KES; no dirty read because postgresql works this way

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
commit; #account 1 has 90.00 KES now

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 90.00 KES; no dirty read because postgresql works this way
commit;

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level read committed; #
show transaction isolation level; # read committed

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level read committed; #
show transaction isolation level; # read committed

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 3 accounts. account 1 has 90.00 KES. the rest have 100.0 KES each. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 90.00 KES
select * from accounts where balance >= 90; # account 1 has 90.00 KES; it is returned also

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
update accounts set balance = balance - 10 where id = 1 returning *; # remove KES 10.00 from account 1

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 90.00 KES because the transaction in session/terminal 1 is not committed

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
commit; #account 1 has 80.00 KES now

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 80.00 KES because the transaction in session/terminal 1 is committed; non-repeatable read
select * from accounts where balance >= 90; # account 1 has 80.00 KES; it is not returned; phantom read
commit;

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level repeatable read; #
show transaction isolation level; # repeatable read

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level repeatable read; #
show transaction isolation level; # repeatable read

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 3 accounts. account 1 has 80.00 KES. the rest have 100.0 KES each. account_id 1, 2 and 3 respectively

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 80.00 KES
select * from accounts where balance >= 80; # account 1 has 80.00 KES; it is returned also

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
update accounts set balance = balance - 10 where id = 1 returning *; # remove KES 10.00 from account 1
commit; # account 1 has 70.00 KES now

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts where id = 1; # account 1 has 80.00 KES even though the transaction in session/terminal 1 is committed; non-repeatable read eliminated
select * from accounts where balance >= 80; # account 1, according to session 2, has 80.00 KES; it is returned also; phantom read eliminated
update accounts set balance = balance - 10 where id = 1 returning *; # ERROR: could not serialize access due to concurrent update
rollback; # account 1 has 70.00 KES

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level repeatable read; #
show transaction isolation level; # repeatable read

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level repeatable read; #
show transaction isolation level; # repeatable read

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 3 accounts. account 1 has 70.00 KES. the rest have 100.0 KES each. account_id 1, 2 and 3 respectively
select sum(balance) from accounts; ## 3 accounts. account; # one row with the value 270 (70 + 100 + 100)
insert into accounts (owner, balance, currency) values ("four", 270, "KES") returning *; ## create a new account
select * from accounts; ## 3 accounts. account 1 has 70.00 KES; ## 4 accounts. account 1 has 70.00 KES, accounts 2 and 3 have 100.00 KES each and account 4 has 270.00 KES. account_id 1, 2, 3 and 4 respectively

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
select * from accounts; ##  ## 3 accounts. account 1 has 70.00 KES. the rest have 100.0 KES each. account_id 1, 2 and 3 respectively. we just created another account; where is it?
select sum(balance) from accounts; ## 3 accounts. account; # one row with the value 270 (70 + 100 + 100)
insert into accounts (owner, balance, currency) values ("four", 270, "KES") returning *; ## create a new account
select * from accounts; ## 4 accounts, not 5; wtf is going on? serialisation anomaly, perhaps?

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
commit; #successful commit

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
commit; #successful commit
select * from accounts; ## 5 accounts, not 4; serialisation anomaly confirmed

###################################################
# terminal #1                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level serializable; #
show transaction isolation level; # serializable

###################################################
# terminal #2                                     #
###################################################
# inside a `postgresql` session
begin; #
set trasaction isolation level serializable; #
show transaction isolation level; # serializable

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 5 accounts: account 1 has 70.00, 2&3 have 100.00 each and 4&5 have 270.00 each. account_id 1, 2, 3, 4, 5 respectively
select sum(balance) from accounts; # one row with the value 810 (70 + 100 + 100 + 270 + 270)
insert into accounts (owner, balance, currency) values ("five", 810, "KES") returning *; # create a new account (notice account_id = 6)

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
select * from accounts; ## 5 accounts, not 6: account 1 has 70.00, 2&3 have 100.00 each and 4&5 have 270.00 each. account_id 1, 2, 3, 4, 5 respectively
select sum(balance) from accounts; # one row with the value 810, not 1620 (70 + 100 + 100 + 270 + 270)
insert into accounts (owner, balance, currency) values ("five", 810, "KES") returning *; # create a new account (notice account_id = 7)

###################################################
# terminal #1                                     #
###################################################
#inside a `postgresql` session
commit; #successful commit

###################################################
# terminal #2                                     #
###################################################
#inside a `postgresql` session
commit; #ERROR: could not serialize access due to read/write dependencies among transactions
select * from accounts; ## 6 accounts, not 7; account_id 7 was not written to DB. serialisation anomaly eliminated
