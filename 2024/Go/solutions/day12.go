package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct{
    r int
    c int
}

var directions [4][2]int = [4][2]int{
    {0,-1},
    {-1,0},
    {0,1},
    {1,0},
}

var diagDirections [4][2]int = [4][2]int{
    {-1,-1},
    {1,-1},
    {1,1},
    {-1,1},
}

func ceil(a,b int) int{
    if a%b == 0{
        return a/b
    }
    return (a/b) + 1
}

func getRegionPerimeterAndArea(r,c int,grid []string,seen [][]int) (int,int,int){
    rows,cols := len(grid),len(grid[0])
    area := 0
    perimeter := 0
    corners := 0
    cornerMap := make(map[point]bool)
    queue := [][]int{}
    queue = append(queue, []int{r,c})
    queueIdx := 0
    seenInRegion := make([][]int,rows,rows)
    for r := range(rows){
        seenInRegion[r] = make([]int,cols,cols)
    }
    seenInRegion[r][c] = 1
    for queueIdx < len(queue){
        pos := queue[queueIdx]
        posInDiagMap := []int{pos[0]*2 + 1,pos[1]*2 + 1}
        queueIdx += 1
        for _,diagDirection := range(diagDirections){
            cornerMap[point{posInDiagMap[0] + diagDirection[0],posInDiagMap[1] + diagDirection[1]}] = true
        }
        perimeterContrib := 4
        area += 1
        for _,direction := range(directions){
            newR := pos[0] + direction[0]
            newC := pos[1] + direction[1]
            if newR < 0 || newR >= rows || newC < 0 || newC >= cols{
                continue
            }
            if grid[newR][newC] != grid[r][c]{
                continue
            }
            perimeterContrib -= 1
            if seen[newR][newC] == 1{
                continue
            }
            seen[newR][newC] = 1
            seenInRegion[newR][newC] = 1
            queue = append(queue, []int{newR,newC})
        }
        perimeter += perimeterContrib
    }
    for corner,_ := range(cornerMap){
        numberOfPositionsCornerIsAPartOf := 0
        positionsCornerIsAPartOf := []bool{false,false,false,false}
        for idx,diagDirection := range(diagDirections){
            rowInDiagMap := corner.r + diagDirection[0]
            colInDiagMap := corner.c + diagDirection[1]
            if rowInDiagMap < 0 || rowInDiagMap > 2*rows || colInDiagMap < 0 || colInDiagMap > 2*cols{
                continue
            }
            row := ceil(rowInDiagMap,2) - 1
            col := ceil(colInDiagMap,2) - 1
            if seenInRegion[row][col] == 1{
                numberOfPositionsCornerIsAPartOf += 1
                positionsCornerIsAPartOf[idx] = true
            }
        }
        if numberOfPositionsCornerIsAPartOf == 1{
            corners += 1
        }
        if numberOfPositionsCornerIsAPartOf == 2 && ((positionsCornerIsAPartOf[0] && positionsCornerIsAPartOf[2]) || (positionsCornerIsAPartOf[1] && positionsCornerIsAPartOf[3])){
            corners += 2
        }
        if numberOfPositionsCornerIsAPartOf == 3{
            corners += 1
        }
    }
    return area,perimeter,corners
}

func Day12(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening file: %v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    grid := []string{}
    for scanner.Scan(){
        grid = append(grid, scanner.Text())
    }
    rows,cols := len(grid),len(grid[0])
    seen := make([][]int,rows,rows) 
    for r := range(rows){
        seen[r] = make([]int,cols,cols)
    }
    part1 := 0
    part2 := 0
    for r := range(rows){
        for c := range(cols){
            if seen[r][c] == 1{
                continue
            }
            seen[r][c] = 1
            area,perimeter,corners := getRegionPerimeterAndArea(r,c,grid,seen)
            part1 += (area*perimeter)
            part2 += (area*corners)
        }
    }
    fmt.Printf("Part 1: %v, Part 2: %v\n",part1,part2)

}
