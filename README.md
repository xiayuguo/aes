# aes
AES/CBC/PKCS7Padding 加/解密的实现(兼容`python2`和`python3`)


## Third party dependence
```shell
pip install pycrypto
```

## Quick Start
- Python2
```python
>>> from aes_crypter import AesCrypter
>>> a = AesCrypter("hugo")
>>> a.encrypt("ABC")
'QYqJdNQIZ5j5q4iqIhuAsg=='
>>> a.decrypt("QYqJdNQIZ5j5q4iqIhuAsg==")
u'ABC'
>>>
```
- Python3
```python
>>> from aes_crypter import AesCrypter
>>> a = AesCrypter("hugo")
>>> a.encrypt("ABC")
b'QYqJdNQIZ5j5q4iqIhuAsg=='
>>> a.decrypt("QYqJdNQIZ5j5q4iqIhuAsg==")
'ABC'
>>> 
```

