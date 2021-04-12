# display-metrics-go #

## links ##

- [How To Change The Command Line Font Size](https://www.raspberrypi-spy.co.uk/2014/04/how-to-change-the-command-line-font-size/)
- [How to setup Bluetooth on a Raspberry Pi 3](https://www.cnet.com/how-to/how-to-setup-bluetooth-on-a-raspberry-pi-3/)

## snippets ##

```shell
if [ "/dev/tty1" = "$( tty )" ]; then
    echo "Execute display"
else
    echo "Ignore displat"
fi
```
