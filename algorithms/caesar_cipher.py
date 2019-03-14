import string

class CaesarCipher:
    def __init__(self, shift):
        self.shift = shift

    def _transform(self, message, shift):
        transformed = list(message)
        for index, ch in enumerate(transformed):
            if ch.islower():
                transformed[index] = chr((ord(ch) - ord('a') + shift)%26 + ord('a'))
        return ''.join(transformed)

    def encrypt(self, message):
        return self._transform(message, self.shift)

    def decrypt(self, message):
        return self._transform(message, -self.shift)



cipher = CaesarCipher(3)
#message = "the eagle is in play; meet at joe's"
message = string.ascii_letters
coded = cipher.encrypt(message)
print('secret: ', coded)
answer = cipher.decrypt(coded)
print('Message: ', answer)
