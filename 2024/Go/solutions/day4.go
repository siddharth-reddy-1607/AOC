package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findXMAS(grid [][]byte,r,c,rows,cols int) int{
    xmas := 0
    if c+3 < cols{
        if string(grid[r][c:c+4]) == "XMAS"{
            xmas += 1
        }
    }
    if c-3 >= 0{
        if string(grid[r][c-3:c+1]) == "SAMX"{
            xmas += 1
        }
    }
    if r+3 < rows{
        val := "X"
        for i := 1; i <= 3; i++{
            val += string(grid[r+i][c])
        }
        if val == "XMAS"{
            xmas += 1
        }
        val = "X"
        if c+3 < cols{
            val += string(grid[r+1][c+1]) + string(grid[r+2][c+2]) + string(grid[r+3][c+3])
            if val == "XMAS"{
                xmas += 1
            }
        }
        val = "X"
        if c-3 >= 0{
            val += string(grid[r+1][c-1]) + string(grid[r+2][c-2]) + string(grid[r+3][c-3])
            if val == "XMAS"{
                xmas += 1
            }
        }
    }
    if r-3 >= 0{
        val := "X"
        for i := 1; i <= 3; i++{
            val += string(grid[r-i][c])
        }
        if val == "XMAS"{
            xmas += 1
        }
        val = "X"
        if c+3 < cols{
            val += string(grid[r-1][c+1]) + string(grid[r-2][c+2]) + string(grid[r-3][c+3])
            if val == "XMAS"{
                xmas += 1
            }
        }
        val = "X"
        if c-3 >= 0{
            val += string(grid[r-1][c-1]) + string(grid[r-2][c-2]) + string(grid[r-3][c-3])
            if val == "XMAS"{
                xmas += 1
            }
        }
    }
    return xmas
}
func findX_MAS(grid [][]byte,r,c int,rows,cols int) int{
    if c-1 < 0 || c+1 == cols || r-1 < 0 || r+1 == rows{
        return 0
    }
    if grid[r-1][c-1] == 'M' && grid[r+1][c-1] == 'M' && grid[r-1][c+1] == 'S' && grid[r+1][c+1] == 'S'{
        return 1
    }
    if grid[r-1][c-1] == 'S' && grid[r+1][c-1] == 'S' && grid[r-1][c+1] == 'M' && grid[r+1][c+1] == 'M'{
        return 1
    }
    if grid[r-1][c-1] == 'S' && grid[r+1][c-1] == 'M' && grid[r-1][c+1] == 'S' && grid[r+1][c+1] == 'M'{
        return 1
    }
    if grid[r-1][c-1] == 'M' && grid[r+1][c-1] == 'S' && grid[r-1][c+1] == 'M' && grid[r+1][c+1] == 'S'{
        return 1
    }
    return 0
}
func Day4(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Panicf("Error openeng file: %v",err)
    }
    defer file.Close()
    grid := [][]byte{}
    r := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan(){
        line := scanner.Text()
        grid = append(grid, []byte{})
        for idx := range(len(line)){
            grid[r] = append(grid[r], line[idx])
        }
        r += 1
    }
    rows,cols := len(grid),len(grid[0])
    part1 := 0
    part2 := 0
    for r := range(rows){
        fmt.Println(string(grid[r]))
    }

    for r := range(rows){
        for c := range(cols){
            if grid[r][c] == 'X'{
                part1 += findXMAS(grid,r,c,rows,cols)
            }        
            if grid[r][c] == 'A'{
                part2 += findX_MAS(grid,r,c,rows,cols)
            }
        }
    }
    fmt.Printf("Part 1: %v,Part 2: %v",part1,part2)
}
