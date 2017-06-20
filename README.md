# HTTPolly

HTTPolly is an basic HTTP proxy for calling [Amazon Polly](https://docs.aws.amazon.com/polly/latest/dg/what-is.html) via a super-simple JSON endpoint and returns the MP3 directly.

## Installation

```bash
$ go install github.com/tsileo/httpolly
$ export AWS_SECRET_ACCESS_KEY=XXX
$ export AWS_ACCESS_KEY_ID=XXX
$ export AWS_REGION=XXX
$ $GOPATH/bin/httpolly
```

## Client example

```python
import requests
from subprocess import Popen, PIPE

resp = requests.post('http://localhost:8015', json={'text': 'Salut tu vas bien ?', 'voice_id': 'Celine'})
resp.raise_for_status()
p = Popen(['madplay', '-'], stdin=PIPE, stderr=PIPE)
p.communicate(input=resp.content)
p.wait()
```
