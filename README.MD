# Kong and Go Services

## Docs
Register Go Service to Kong Api Gateway https://konghq.com/ using Resty Http Client and using Echo Http Framework 
https://echo.labstack.com/ and Gin https://gin-gonic.com/

## Install
#### A. Requirement
1. Kong Installed or on Cloud Server (install on local https://konghq.com/install/)
2. Go Version => 1.13
3. Dep Dependency Manager Tool Golang https://golang.github.io/dep/

#### B. Installation
##### 1. clone this repo 
```$xslt
    git clone https://github.com/huzairuje/kong_with_service
```

##### 2. cd to the directory
```$xslt
    cd ~/GO/SRC/DIRECTORY/kong_with_service
```

##### 3. cd to user service
```$xslt
    cd ~/GO/SRC/DIRECTORY/kong_with_service/user_service
```
##### 4. install the dependency
```$xslt
    dep ensure
```
##### 5. cd to garage service
```$xslt
    cd ~/GO/SRC/DIRECTORY/kong_with_service/garage_service
```
##### 6. install the dependency
```$xslt
    dep ensure
```

#### C. Run The Server
##### 1. make sure the kong api gateway is running
or you can test it with access the admin API with
```$xslt
    curl -i -X POST \
    --url http://localhost:8001/
```

or you can start the kong api gateway with
 ```$xslt
     sudo kong start
 ```
##### 2. run the user service
cd to user service
```$xslt
    cd ~/GO/SRC/DIRECTORY/kong_with_service/user_service
```
and run the service with
```$xslt
    go run main.go
```

##### 3. run the garage service
cd to user service
```$xslt
    cd ~/GO/SRC/DIRECTORY/kong_with_service/garage_service
```
and run the service with
```$xslt
    go run main.go
```

##### 4. Access The Service
a. access the user service
```$xslt
    curl -i -X GET \
    --url http://localhost:8000/api/users
```

b. access the garage service
```$xslt
    curl -i -X GET \
    --url http://localhost:8000/api/garages
```

you will get the response from kong api gateway.

