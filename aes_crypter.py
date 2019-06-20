# coding: utf8
import six
import base64
import hashlib

from Crypto.Cipher import AES


def utf8(chars):
    """convert str to bytes"""
    return chars.encode("utf8") if isinstance(chars, six.text_type) else chars


def to_unicode(value):
    """Converts a string argument to a unicode string.

    If the argument is already a unicode string or None, it is returned
    unchanged.  Otherwise it must be a byte string and is decoded as utf8.
    """
    if isinstance(value, (six.text_type, type(None))):
        return value
    if not isinstance(value, bytes):
        raise TypeError("Expected bytes, unicode, or None; got %r" % type(value))
    return value.decode("utf-8")


class AesCrypter(object):
    def __init__(self, key):
        self.key = hashlib.sha256(utf8(key)).digest()
        self.iv = self.key[:16]

    def encrypt(self, data):
        data = self.pkcs7padding(utf8(data))
        cipher = AES.new(self.key, AES.MODE_CBC, self.iv)
        encrypted = cipher.encrypt(data)
        return to_unicode(base64.b64encode(encrypted))

    def decrypt(self, data):
        data = base64.b64decode(data)
        cipher = AES.new(self.key, AES.MODE_CBC, self.iv)
        decrypted = cipher.decrypt(data)
        decrypted = self.pkcs7unpadding(decrypted)
        return to_unicode(decrypted)

    def pkcs7padding(self, data):
        bs = AES.block_size
        padding = bs - len(data) % bs
        padding_text = chr(padding) * padding
        return utf8(data) + utf8(padding_text)

    def pkcs7unpadding(self, data):
        data = to_unicode(data)
        length = len(data)
        unpadding = ord(data[length - 1])
        return data[0 : length - unpadding]

