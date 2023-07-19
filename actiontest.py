import requests


url="http://baidu.com"
for i in range(10):
    resp=requests.get(url)
    print(resp.text)
    print(url)
    print()
print("测试结束")

