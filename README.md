# onepass
onepass: password manager for human

## Test
```
go test -run TestGenPassword
```

## run 
```
go build && onepass -v -f -s 8 app
go build && onepass app5
```

## use 
```
onepass {appName}  # if found, copy into clipboard. else new and copy
-h help 
-v always verbose password
-s set password length when create or renew
-f force refresh password
```

