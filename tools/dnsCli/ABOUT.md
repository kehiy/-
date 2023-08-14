# DNS cli
this is a simple command line tool for linux!
simply get dns from the args then set args as you dns in `/etc/reslov.conf`
it have some defaults

# why?
I made this to change my dns in linux fast and eazy every time.

# How to use?
to use this dns manager
just make a build:

```bash
go build -o dns
```
change mod(in bin folder):
```bash
chmod +x dns
```

copy it into bin:
```bash
sudo cp dns /bin/
```

now you can use it with `dns` command passing 1 and 2 for shecan and 403 and other IPs for custom.
