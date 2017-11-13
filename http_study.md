# http学习心得

### 1.GET请求

get请求有好几种方式，

1.直接使用net/http包内函数请求

```go
import(
	"net/http"
)
...
resp,err:=http.Get("http://www.baidu.com")
```

2.利用http.client结构体来请求

```go
import(
	"net/http"
)
...
clt:=http.Client{}
resp,err:=clt.Get("http://www.baidu.com")
```

3.最本质的请求方式

看源代码可以发现，以上两种方式都是用了一下这个最本质的请求方式，即使用http.NewRequest函数。

```go
req,err:=http.NewRequest("GET","http://baidu.com",nil)
//然后http.Client  结构体Do方法
//http.DefaultClient可以换成另外一个http.Client(自己定义)
resp,err:=http.DefaultClient.Do(req)
```

Go的get请求表面上有好几种请求方式，实则只有一种:

######  就是使用http.NewRquest函数获得request实体，然后利用http.Client结构体的Do方法，将request实体传入Do方法中。

### 2.POST请求

和get请求一样，post请求也有好多种方式，但本质还是使用了http.NewRquest函数和http.Client的Do方法。

1.使用net/http包带的post方法

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
body:=strings.NewReader(data,Encode())
resp,err:=http.Post("xxxxxx","application/x-www-form-urlencoded",body)
```

或者还可以

```go
import(
	"net/http"
  	"net/url"
)
...
var r http.Request
r.ParseForm()
r.Form.Add("xxx","xxx")
body:=strings.NewReader(r.Form.Encode())
http.Post("xxxx","application/x-www-form-urlencoded",body)
```

2.使用实例化的http.Client的post方法

###### 其实本质上直接使用包函数 和 实例化http.Client是一样的，包函数的底层也仅仅是实例化了一个DefaultClient,然后调用DefaultClient的方法。

下面给出使用实例化的http client的post方法：

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
body:=strings.NewReader(data.Encode())
clt:=http.Client{}
resp,err:=clt.Post("xxxxxx","application/x-www-form-urlencoded",body)
```

还有

```go
import(
	"net/http"
  	"net/url"
)
...
var r http.Request
r.ParseForm()
r.Form.Add("xxx","xxx")
body:=strings.NewReader(r.Form.Encode())
clt:=http.Client{}
clt.Post("xxxx","application/x-www-form-urlencoded",body)
```

简单的，仅限于表单

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
clt:=http.Client{}
clt.PostForm("xxxx",data)
```

3.使用net/http包的NewRequest函数

###### 其实不管是get方法也好，post方法也好，所有的get,post的http请求形式，最终都是会调用net/http包的NewRequest函数，多种多样的请求形式，也仅仅是封装的不同而已。

```go
import(
	"net/http"
  	"net/url"
)
...
data:=url.Values{"start":{"0"},"offset":{"xxxx"}}
body:= strings.NewReader(data.Encode())

req,err:=http.NewRequest("POST","xxxxx",body)
req.Header.Set("Content-Type","application/x-www-form-urlencoded")

clt:=http.Client{}
clt.Do(req)
```

###  3.添加request header

net/http包没有封装  直接使用请求带header的get或者post的方法，所以，要想请求中带header，只能使用NewRequest方法。

```go
import(
	"net/http"
)
...
req,err:=http.NewRequest("POST","xxxxx",body)
//此处还可以写req.Header.Set("User-Agent","myClient")
req.Header.Add("User-Agent","myClient")

clt:=http.Client{}
clt.Do(req)
```

##### 需要值得注意的一点是:再"添加"header操作的时候，req.Header.Add和req.Header.Set都可以,但是在修改操作的时候，只能使用req.Header.Set。这俩个方法是有区别的，Golang底层Header的实现是一个map[string][]string,req.Header.Set方法，如果原来Header中没有值，那么是没问题的，如果有值，会将原来的值替换掉。而req.Header.Add的话，是在原来值的基础上，再append一个值，比如，原来header的值是"s",我req.Header.Add的话，变成了[s a]。但是，获取header值的方法req.Header.Get却只取第一个，所以，如果原来有值，重新req.Header.Add一个新值的话，req.Header.Get得到的值不变。

### 4.打印response响应

```go
import (
	"net/http"
	"net/url"
	"io/ioutil"
)
...
content, err := ioutil.ReadAll(resp.Body)
respBody := string(content)
```

### 5.使用cookie

在Golang中使用http proxy，也必须构造自己的`http.client`，需要将`http.client`结构体的一个属性`Transport`自己实例化好。

###### 当使用环境变量$http_proxy或$HTTP_PROXY作为代理时(即全局代理)

```go
//从环境变量$http_proxy或$HTTP_PROXY中获取HTTP代理地址
func GetTransportFromEnvironment() (transport *http.Transport) {
	transport = &http.Transport{Proxy : http.ProxyFromEnvironment}
	return
}
```

###### 当使用自己搭建http代理时

参数`proxy_addr`即代理服务器IP端口号，例如：”[http://xxx.xxx.xxx.xxx:6000“，注意，必须加上"http](http://xxx.xxx.xxx.xxx:6000%E2%80%9C%EF%BC%8C%E6%B3%A8%E6%84%8F%EF%BC%8C%E5%BF%85%E9%A1%BB%E5%8A%A0%E4%B8%8A"http/)“

```go
func GetTransportFieldURL(proxy_addr *string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, error := url_i.Parse(*proxy_addr)
	if error != nil{
		fmt.Println(error.Error())
	}
	transport = &http.Transport{Proxy : http.ProxyURL(url_proxy)}
	return
}
```

