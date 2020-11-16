# onepass
onepass: password manager for human
> 自制密码管理器，用着放心


## Test
```
go test -run TestGenPassword
```


## test
```
go build && onepass -v -f -s 8 app
go build && onepass app5
```

## build

```
go build -ldflags "-w"
# for linux
env GOOS=linux GOARCH=amd64 go build -ldflags "-w"
```

## use 
```
add onepass.exe folder to envrionment $PATH and you can use it from anywhere 


onepass {appName}  # if found, copy into clipboard. else new and copy

parameters:
-h help 
-v always verbose password
-s set password length when create or renew
-f force refresh password
-l list all apps . this will ignore other parameters 
```

## use demo
```
D:\GoProject\onepass(main -> origin)$ onepass  google
->Generated password for: google        size: 10
->Already Copied. You can use it by Ctrl + V or Win + V

D:\GoProject\onepass(main -> origin)$ onepass  google
->Found in Cache: google
->Already Copied. You can use it by Ctrl + V or Win + V

D:\GoProject\onepass(main -> origin)$ onepass  -f -s 20 google
->Found in Cache: google
->Force renewed password for: google    size: 20
->Already Copied. You can use it by Ctrl + V or Win + V

D:\GoProject\onepass(main -> origin)$ onepass  -v google
->Found in Cache: google
->Already Copied. You can use it by Ctrl + V or Win + V
KY37R911W3aJ623M254K

D:\GoProject\onepass(main -> origin)$ go build && onepass app6
->Generated password for: app6  size: 10
->Already Copied. You can use it by Ctrl + V or Win + V

D:\GoProject\onepass(main -> origin)$ go build && onepass
1 google
2 app
3 app6

```


