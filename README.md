Postgresql Docker

1. 컨테이너 생성

    docker run --name tododb -p 5432:5432  -e POSTGRES_PASSWORD=1234Qwer -d postgres

2. postgresql 접속

    KDS-2:~ kimdaesung$ docker exec -it tododb /bin/bash
    root@1c305d318414:/# psql -U postgres
    psql (13.1 (Debian 13.1-1.pgdg100+1))
    Type "help" for help.
    
    postgres=# 
    postgres=# 
3. DatabaseURL

```
docker run --name tododb -p 5432:5432  -e POSTGRES_PASSWORD=1234Qwer -d postgres
DATABASE_URL=postgres://my_app:my_password@localhost/my_database
```
    





