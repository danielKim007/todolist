Postgresql Docker

1. 컨테이너 생성

    docker run --name tododb -p 5432:5432  -e POSTGRES_PASSWORD=1234Qwer -d postgres

2. postgresql 접속 확인
    ```
    KDS-2:~ kimdaesung$ docker exec -it tododb /bin/bash
    root@1c305d318414:/# psql -U postgres
    psql (13.1 (Debian 13.1-1.pgdg100+1))
    Type "help" for help.
   
    postgres=# 
    postgres=#
    ```
3. tododb(database) 생성 후 접속
    ```
    postgres=# Create database tododb
    postgres=# \c tododb
    You are now connected to database "tododb" as user "postgres".
    ```
4. table 생성
    ```
    CREATE TABLE todo_list
    (
        id SERIAL PRIMARY KEY,
        name VARCHAR NOT NULL
    );

    CREATE TABLE todo_item
    (
        id SERIAL PRIMARY KEY,
        todo_list_id INT NOT NULL REFERENCES todo_list ON DELETE CASCADE,
        text VARCHAR NOT NULL,
        done BOOLEAN NOT NULL DEFAULT FALSE
    );

    ```
5. DatabaseURL

    ```
    docker run --name tododb -p 5432:5432  -e POSTGRES_PASSWORD=1234Qwer -d postgres
    DATABASE_URL=postgres://my_app:my_password@localhost/my_database
    ```






