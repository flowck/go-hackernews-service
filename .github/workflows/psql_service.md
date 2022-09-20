Snippet to run Postgres on Github Action
 
```
    services:
      postgres:
        image: postgres:14
        env:
          POSTGRES_DB: go_hackernews_service_test
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: root
        ports:
          - 5432:5432
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
```