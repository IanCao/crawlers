FROM mysql:5.7.26

ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

COPY setup.sh /mysql/setup.sh
COPY init.sql /mysql/init.sql
COPY privileges.sql /mysql/privileges.sql

CMD ["sh", "/mysql/setup.sh"]
