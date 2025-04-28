# go-config

读取YAML配置，并转换为map

用法

yml

```yml
cli:
  name: tngDaemon
  node: transCodeNode
  server:
    port: 9090
    bind: 127.0.0.1
```

读取YML

```go
	pwd, _ := os.Getwd()
	cfg.Load(utils.CombinPath(pwd, "/config.yml"))
```

获取yml里面的配置

```go
node, _ := cfg.GetConfig("cli.node")//获取到的类型为字符串
server, _ := cfg.GetConfig("cli.server")//获取到的类型为map[string]interface{}
```

