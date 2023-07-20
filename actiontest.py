import requests
import chardet

url="http://baidu.com"
for i in range(10):
    resp=requests.get(url)
    content = resp.content
    encoding = chardet.detect(content)['encoding']
    print(encoding)
    resp.encoding=encoding
    print(resp.text)
    print(url)
    print()
print("测试结束")

