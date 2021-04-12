# display-metrics-go #

## links ##

- [How To Change The Command Line Font Size](https://www.raspberrypi-spy.co.uk/2014/04/how-to-change-the-command-line-font-size/)

## snippets ##

```shell
if [ "/dev/tty1" = "$( tty )" ]; then
    echo "Execute display"
else
    echo "Ignore displat"
fi
```