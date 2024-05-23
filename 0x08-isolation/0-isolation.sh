docker exec -it mysql8 mysql -uroot -ppassword123 fnbank

################################################################
# inside a `msql` environment
################################################################

select @@trasaction_isolation;
select @@global.trasaction_isolation;
set session transaction isolation level read uncommitted; #or whatever isolation level you want