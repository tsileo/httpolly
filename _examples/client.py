import requests
from subprocess import Popen, PIPE

resp = requests.post('http://localhost:8015', json={'text': 'Salut tu vas bien ?', 'voice_id': 'Celine'})
resp.raise_for_status()
p = Popen(['madplay', '-'], stdin=PIPE, stderr=PIPE)
p.communicate(input=resp.content)
p.wait()
