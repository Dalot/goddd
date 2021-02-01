===== Not suited for production =====

Please use the default ports since they are hardcoded, :8080 for the go backend app and :8081 for the frontend app. 

```
$ git clone https://github.com/Dalot/goddd
$ cd goddd
```

Running a dabatase with docker
```
docker run --name mysql -e MYSQL_DATABASE=code_challenge -e MYSQL_ROOT_PASSWORD=secret -p 3306:3306 -d mysql:8.0.22
```

Build the executable and run -cmd migrate to migrate the database and seed it.
```
$ go build main.go && ./main.exe -cmd migrate
...
[GIN-debug] Listening and serving HTTP on :8080
```

You can also use `-cmd fresh`, it will wipe the database and seed it again.
``` 
$ go build main.go && ./main.exe -cmd fresh
...
[GIN-debug] Listening and serving HTTP on :8080
```

And now to run the frontend app, 
```
$ npm install
$ npm run dev 
...
Your application is running here: http://localhost:8081
```
