#Using 10.6 version of Mariadb as BASE
FROM mariadb:10.6

#COPYING SQL SCRIPT FOR TABLE CREATION AND DUMMY RECORDS
COPY setup.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]