#Using latest version of Mariadb as BASE
FROM mariadb:latest

#COPYING SQL SCRIPT FOR TABLE CREATION AND DUMMY RECORDS
COPY setup.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]