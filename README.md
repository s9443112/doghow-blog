# doghow-blog
Practice golang 

# Install 
First to clone it down 
```
$ git clone https://github.com/s9443112/doghow-blog.git
```

If you have install govendor before and you can ...
```
$ govendor sync 
```

To start run before you have init config.yaml
```
$ touch config.yaml
```

And insert 
```js
appname: doghow_blog

db:
  host: "0.0.0.0:3306"
  user: "root"
  dbname: ""
  password: ""
```

And Go init database first 
```
$ go run index.go --dbinit 
```

Final We can start run server!
```
$ go run index.go --runserver
```
