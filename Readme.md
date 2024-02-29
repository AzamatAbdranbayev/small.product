
### Запуск через 

```
docker compose up -d
```


Здесь надо добавить токен, который был отправлен ранее в лс
https://github.com/AzamatAbdranbayev/small.product/blob/4ce56dffc029503a18fa6754217c3ee768cb2707/docker-compose.yml#L28

### [Swagger](http://localhost:8080/swagger/index.html#/)


### go test (покрыл частично и модульные и интег. к Postgres)

```go
?       github.com/AzamatAbdranbayev/small.product/internal/constant    [no test files]
?       github.com/AzamatAbdranbayev/small.product/internal/errapp      [no test files]
        github.com/AzamatAbdranbayev/small.product/internal/handler             coverage: 0.0% of statements
        github.com/AzamatAbdranbayev/small.product/docs         coverage: 0.0% of statements
        github.com/AzamatAbdranbayev/small.product/cmd          coverage: 0.0% of statements
        github.com/AzamatAbdranbayev/small.product/internal/app         coverage: 0.0% of statements
        github.com/AzamatAbdranbayev/small.product/config               coverage: 0.0% of statements
=== RUN   TestProduct_CheckId
=== RUN   TestProduct_CheckId/valid_uuid
=== RUN   TestProduct_CheckId/invalid_uuid
--- PASS: TestProduct_CheckId (0.00s)
    --- PASS: TestProduct_CheckId/valid_uuid (0.00s)
    --- PASS: TestProduct_CheckId/invalid_uuid (0.00s)
=== RUN   TestProduct_CheckMaxPrice
=== RUN   TestProduct_CheckMaxPrice/valid_price
=== RUN   TestProduct_CheckMaxPrice/invalid_price
--- PASS: TestProduct_CheckMaxPrice (0.00s)
    --- PASS: TestProduct_CheckMaxPrice/valid_price (0.00s)
    --- PASS: TestProduct_CheckMaxPrice/invalid_price (0.00s)
=== RUN   TestProduct_CheckValidName
=== RUN   TestProduct_CheckValidName/invalid_name
=== RUN   TestProduct_CheckValidName/valid_name
--- PASS: TestProduct_CheckValidName (0.00s)
    --- PASS: TestProduct_CheckValidName/invalid_name (0.00s)
    --- PASS: TestProduct_CheckValidName/valid_name (0.00s)
PASS
coverage: 57.1% of statements
ok      github.com/AzamatAbdranbayev/small.product/internal/models      0.003s  coverage: 57.1% of statements
        github.com/AzamatAbdranbayev/small.product/internal/service             coverage: 0.0% of statements
=== RUN   TestDbQueries
=== RUN   TestDbQueries/create_valid_product_name
=== RUN   TestDbQueries/create_invalid_product_name
=== RUN   TestDbQueries/create_user_1_
=== RUN   TestDbQueries/create_user_2
--- PASS: TestDbQueries (0.14s)
    --- PASS: TestDbQueries/create_valid_product_name (0.01s)
    --- PASS: TestDbQueries/create_invalid_product_name (0.05s)
    --- PASS: TestDbQueries/create_user_1_ (0.01s)
    --- PASS: TestDbQueries/create_user_2 (0.05s)
PASS
coverage: 9.4% of statements
ok      github.com/AzamatAbdranbayev/small.product/internal/repo        0.147s  coverage: 9.4% of statements
=== RUN   TestCheckValidUuid
=== RUN   TestCheckValidUuid/valid_uuid
=== RUN   TestCheckValidUuid/invalid_uuid
--- PASS: TestCheckValidUuid (0.00s)
    --- PASS: TestCheckValidUuid/valid_uuid (0.00s)
    --- PASS: TestCheckValidUuid/invalid_uuid (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/AzamatAbdranbayev/small.product/pkg/helpers  0.004s  coverage: 100.0% of statements

```