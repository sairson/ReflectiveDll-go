# ReflectiveDll-go

# build

```
go build -buildmode=c-archive main.go
```

```
 gcc -shared -pthread -o example.dll ReflectiveDll.c main.a -lwinmm -lntdll -lws2_32
```
