# File-extraction

File-extraction 是根据文件最后修改时间，将文件提取至指定文件夹的小工具。

## 编译

```shell
go build file_extraction.go
```

## 配置

默认配置文件存放在 config/config.yaml

```shell
dir: "/root/file-test" #需要检测的目标文件夹
save_dir: "/root/new-file-test" #保存提取文件的文件

exclude_list:    #排除文件或文件夹
  - "/root/file-test/c"
  - "/root/file-test/a/a1.txt"
```

## 使用

```shell
file_extraction -h
Usage of file_extraction:
  -conf_path string
        Configuration file path (default "config/config.yaml")
  -end_time string
        Last modified end time, the default is the current time 
  -start_time string
        Last modification time start time, default is one hour ago
```

使用实例:

```shell
file_extraction -conf_path = "/root/config.yaml" -start_time="2018-01-01 00:00:01" -end_time="2018-01-02 00:00:01"
file_extraction -start_time="2018-01-01 00:00:01" -end_time="2018-01-02 00:00:01"
file_extraction -start_time="2018-01-01 00:00:01" 
file_extraction -end_time="2018-01-02 00:00:01"
file_extraction
```

## License

The file-extraction is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).
