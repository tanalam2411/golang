 
 ```bash
http_handlers$ go test -v
=== RUN   TestDoubleHandler
=== RUN   TestDoubleHandler/double_of_two
=== RUN   TestDoubleHandler/missing_value
=== RUN   TestDoubleHandler/not_a_number
--- PASS: TestDoubleHandler (0.00s)
    --- PASS: TestDoubleHandler/double_of_two (0.00s)
    --- PASS: TestDoubleHandler/missing_value (0.00s)
    --- PASS: TestDoubleHandler/not_a_number (0.00s)
=== RUN   TestRouting
--- PASS: TestRouting (0.10s)
PASS
ok  	golang/series_3/src/16_unit_testing_http_servers/http_handlers	0.108s
 
```


Test http Routing:
```bash
http_handlers$ go test -v -run Routing
=== RUN   TestRouting
--- PASS: TestRouting (0.00s)
PASS
ok  	golang/series_3/src/16_unit_testing_http_servers/http_handlers	0.004s
```

If we change url:
```bash
http_handlers$ go test -v -run Routing
=== RUN   TestRouting
--- FAIL: TestRouting (0.00s)
    main_test.go:91: expected status OK; got 404 Not Found
    main_test.go:101: expected an integer; got 404 page not found
FAIL
exit status 1
FAIL	golang/series_3/src/16_unit_testing_http_servers/http_handlers	0.004s
```