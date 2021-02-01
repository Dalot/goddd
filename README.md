===== Not suited for production =====

Building example project with DDD

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
```

You can also use `-cmd fresh`, it will wipe the database and seed it again.
``` 
$ go build main.go && ./main.exe -cmd fresh
```

And now to run the frontend app, 
```
$ npm install
$ npm run serve
```
