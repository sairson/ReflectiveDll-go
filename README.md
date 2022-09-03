# ReflectiveDll-go

# build

```
go build -buildmode=c-archive main.go
```

```
 gcc -shared -pthread -o example.dll ReflectiveDll.c main.a -lwinmm -lntdll -lws2_32
```


# execute

```
rundll32.exe .\example.dll,MyFunction
```
```
regsvr32.exe /s /u example.dll
```
```
regsvr32.exe /s /n /i example.dll
```
```
regsvr32.exe /s example.dll
```
