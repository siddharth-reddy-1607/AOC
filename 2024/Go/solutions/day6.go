package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Returns cellsVisisted while playing the game. Return true if placing a obstacle causes a cycle
func playGame(grid [][]byte,placeNewObstacle bool,r,c,guardR,guardC int) (int,bool){
    if placeNewObstacle{
        grid[r][c] = '#'
        defer func(){
            grid[r][c] = '.'
        }()
    }
    rows := len(grid)
    cols := len(grid[0])
    directions := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    curDirIdx := 0
    seenWithDir := make([][][4]int,rows,rows)
    for row := range(rows){
        seenWithDir[row] = make([][4]int,cols,cols)
    }
    seen := make([][]int,rows,rows)
    for row := range(rows){
        seen[row] = make([]int,cols,cols)
    }
    cellsVisisted := 0
    x := guardR 
    y := guardC
    newX := x
    newY := y
    for (true){
        for (grid[newX][newY] != '#'){
            x = newX 
            y = newY
            seenWithDir[newX][newY][curDirIdx] = -1
            if seen[newX][newY] == 0{
                cellsVisisted += 1
            }
            seen[newX][newY] = -1
            newX = x + directions[curDirIdx][0]
            newY = y + directions[curDirIdx][1]
            if newX < 0 || newX == rows || newY < 0 || newY == cols{
                return cellsVisisted,false
            }
        }
        curDirIdx = (curDirIdx + 1)%4
        newX = x + directions[curDirIdx][0]
        newY = y + directions[curDirIdx][1]
        if newX < 0 || newX == rows || newY < 0 || newY == cols{
            return cellsVisisted,false
        }
        if seenWithDir[newX][newY][curDirIdx] == -1{
            return 0,true
        }
    }
    return -1,false
}
func Day6(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening file: %v",err)
    }
    grid := [][]byte{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        grid = append(grid,[]byte(scanner.Text()))
    }
    rows,cols := len(grid),len(grid[0])
    guard := []int{-1,-1}
    for r := range(rows){
        for c := range(cols){
            if grid[r][c] == '^'{
                guard[0] = r
                guard[1] = c
            }
        }
    }
    part1,_ := playGame(grid,false,0,0,guard[0],guard[1])
    part2 := 0
    for r := range(rows){
        for c := range(cols){
            if grid[r][c] == '#' || grid[r][c] == '^'{
                continue
            }
            _,val := playGame(grid,true,r,c,guard[0],guard[1])
            if val{
                part2 += 1
            }
        }
    }
    fmt.Printf("Part 1 : %v,Part 2: %v\n",part1,part2)
}
