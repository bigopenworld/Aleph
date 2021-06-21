package cmd

import "github.com/fatih/color"


const OK = 0
const ERROR = 1 
const WARNING = 2
const SUCCESS = 3
const FATAL = 4
const INFO = 5

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
	case OK : 
		return color.GreenString("[ OK ] ")
	case ERROR : 
		return color.RedString("[ ERROR ] ")
	case WARNING : 
		return color.YellowString("[ WARN ] ")
	case SUCCESS : 
		return color.GreenString("[ SUCCESS ] ")
	case FATAL : 
		return color.RedString("[ FATAL ] ")
	case INFO :
		return color.BlueString("[ INFO ] ")
	default : 
		return color.HiCyanString("[ UNKNOW ] ")
	}
}