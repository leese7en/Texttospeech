# Texttospeech

一个基于 GO 与 百度文字转语音接口，实现高速文章转语音的软件  
A text-to-speech software based on GO and Baidu text-to-speech interface

### fixed

1. 通过增加 sync.RWMutex 锁的方式，避免了再超高并发情况下 Map 被同时写入导致报错的情况
1. 爬虫失败时的错误处理

## 参数:

| 参数 | 默认值  | 类型   | 备注             |
|------|---------|--------|------------------|
| name | default | string | txt/输出音频名称 |
| s    | 6       | int    | 音色             |

* 如果不需要改变的话不需要添加环境变量保持默认值即可.

## 食用方法

wz___ 分别为已经编译好的各个平台可直接使用的二进制文件

``` bash
# 例如
./wzmac -name="default" -s=6
```

#### 像这样

![Alt text](https://github.com/Glovecc/Texttospeech/blob/master/user.jpg)

然后便会在当前目录下生成一个以 name参数.mp3 命名的媒体文件，里面的内容是按照 txt 内的顺序排列的。

## 注意事项

#### txt格式

如 default.txt ,通过 \n 进行分隔

#### 最后

程序本身只供技术交流，侵删。
