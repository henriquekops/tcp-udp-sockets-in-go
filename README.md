# Tcp + Udp sockets in Golang!
> Authors: Henrique Kops (@henriquekops) & Carlo Laitano (@carloj123)

This project is a simple representation of socket level communication using *.net*
golang's package for transporting files through TCP and UDP protocols.

### How to use

At poject's root directory setup the project:

```
# for linux
$ make setup-linux

# for windows
$ make setup-windows
```

Then, alter .env file:

```
$ nano .env

# alter the fields
# HOST=<HOST:string>
# PORT=<PORT:string>
```

Then run the project with:

```
$ ./main
```

When missing arguments, the below message should be printed:

```
$ Usage of ./main:
    -mode string
      	Start mode - {client, server} (default "server")
    -network string
      	Transport layer - {tcp, udp} (default "tcp")
```

Feel comfortable for re-building the project with the command:

```
$ go build main.go
```

![Be a ninja and use golang](https://blog.marcelocavalcante.net/img/go_mascot3.png)

Enjoy! ;)