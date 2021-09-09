package cmd

import "github.com/fatih/color"

const (
	OK int = iota
	ERROR
	WARNING
	SUCCESS
	FATAL
	INFO
)

/*
0 => OK
1 => ERROR
2 => WARNING
3 => SUCCESS
4 => FATAL
5 => INFO
*/
func NewFlag(flag int) string {
	switch flag {
	case OK:
		return color.GreenString("[ OK ] ")
	case ERROR:
		return color.RedString("[ ERROR ] ")
	case WARNING:
		return color.YellowString("[ WARN ] ")
	case SUCCESS:
		return color.GreenString("[ SUCCESS ] ")
	case FATAL:
		return color.RedString("[ FATAL ] ")
	case INFO:
		return color.BlueString("[ INFO ] ")
	default:
		return color.HiCyanString("[ UNKNOW ] ")
	}
}
