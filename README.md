# Alarm

A dead simple cmdline alarm with notification for Linux and Mac

## Dependencies

### Linux

`libnotify` and a notification server. If you are using a Desktop Environment,
you probably already have one. If you need a stand-alone notification server for
minimal system, I recommend [dunst](https://github.com/knopwob/dunst)

### Mac

Any version equal or higher than OS X 10.7

## Installing

If you have [go](http://golang.org) installed:

``` shell
go install github.com/luizbranco/alarm
```

Or downloading the [binaries](http://github.com/luizbranco/alarm/releases)

## Usage

``` shell
alarm title [description] [time] &

alarm Pizza 15m &
alarm "Pomodoro Timer" "Long break" 1h30m &
```

### Notify when a long process finishes running

Since `alarm` time defaults to 0 seconds, you can run after any command without
passing the time.

``` shell
slow process; alarm done
```
