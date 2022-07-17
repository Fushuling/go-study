package main

import (
	"bufio"
	"fmt"
	"os"
)

var sz = [3][4]int{}

func move(x int, y int, dir int) {
	if dir == 1 { //向上
		if sz[x-1][y] != '*' { //上面一格如果不是障碍物
			sz[x-1][y] == 'C'
			sz[x][y] = '.'
			x = x - 1
		} else { //上一格是障碍物
			dir = 2
		}
	} else if dir == 2 { //向右
		if sz[x][y+1] != '*' { //右面一格如果不是障碍物
			sz[x][y+1] = 'C'
			sz[x][y] = '.'
			y = y + 1
		} else { //右一格是障碍物
			dir = 3
		}
	} else if dir == 3 { //向下
		if sz[x+1][y] != '*' { //下面一格如果不是障碍物
			sz[x+1][y] = 'C'
			sz[x][y] = '.'
			x = x + 1
		} else { //下一格是障碍物
			dir = 4
		}
	} else if dir == 4 { //向左
		if sz[x][y-1] != '*' { //左面一格如果不是障碍物
			sz[x][y-1] = 'C'
			sz[x][y] = '.'
			y = y - 1
		} else { //左一格是障碍物
			dir = 1
		}
	}
}

func main() {
	var x_c, y_c, x_f, y_f int //牛的横纵坐标，人的横纵坐标
	dir_c := 1                 //牛的方向，1向上，2向右，3向下，4向左
	dir_f := 1                 //人的方向
	var sum int                //花费时间
	//加上虚拟边框
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			sz[i][j] = '*'
		}
	}
	//初始化
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			reader := bufio.NewReader(os.Stdin)
			sz[i][j], _ = reader.ReadString('\n')
			if sz[i][j] == 'C' {
				x_c = i
				y_c = j
			}
			if sz[i][j] == 'F' {
				x_f = i
				y_f = j
			}
		}
	}
	//开始行走
	loopControl := true
	for loopControl {
		move(x_c, y_c, dir_c) //牛动
		move(x_f, y_f, dir_f) //人动
		sum++
		if x_c == x_f && y_c == y_f {
			fmt.Println(sum)
			break
		}
		if sum > 160000 { //要是能找到，总时间不可能超过160000
			break
		}
	}

}
