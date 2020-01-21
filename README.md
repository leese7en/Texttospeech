# Texttospeech

一个基于 GO 与 百度文字转语音接口，实现文章转语音的软件  
A text-to-speech software based on GO and Baidu text-to-speech interface

## 参数:

| 参数 | 默认值  | 类型   | 备注             |
|------|---------|--------|------------------|
| p    | false   | bool   | 是否使用并发转换 |
| name | default | string | txt/输出音频名称 |
| s    | 6       | int    | 音色             |

* 如果不需要改变的话不需要添加环境变量保持默认值即可.

## 使用方法

wz___ 分别为已经编译好的各个平台可直接使用的二进制文件

``` bash
# 例如
./wzmac -name="default" -p -s=6
```

#### 像这样

![Alt text](https://github.com/Glovecc/Texttospeech/blob/master/user.jpg)

## 注意事项

#### txt格式

如 default.txt ,通过 \n 进行分隔

#### 并发与默认的区别

在默认模式下，会生成一个以 name参数.mp3 命名的媒体文件，里面的内容是按照 txt 内的顺序排列的。
在并发模式下，由于协程先后顺序不可控，故采用生成一个以 name 参数 命名的文件夹，文件夹内是每段话的音频文件，以先后顺序的索引作为命名。
