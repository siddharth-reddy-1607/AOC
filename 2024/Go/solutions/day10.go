package solutions

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func getTrailHeadRating(r,c int,grid[][]int) int{
    directions := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    if grid[r][c] == 9{
        return 1
    }
    ways := 0
    for _,direction := range(directions){
        newR := r + direction[0]
        newC := c + direction[1]
        if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]){
            continue
        }
        if grid[newR][newC] - grid[r][c] == 1{
            ways += getTrailHeadRating(newR,newC,grid)
        }
    }
    return ways
}
func getTrailHeadScore(r,c int,grid [][]int) int{
    score := 0
    rows,cols := len(grid),len(grid[0])
    directions := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    queue := [][]int{}
    queue = append(queue, []int{r,c})
    queueHead := 0
    seen := make([][]int,rows,rows)
    for r := range(rows){
        seen[r] = make([]int,cols,cols)
    }
    for queueHead < len(queue){
        pos := queue[queueHead]
        queueHead += 1
        if grid[pos[0]][pos[1]] == 9{
            score += 1
            continue
        }
        for _,direction := range(directions){
            newX := pos[0] + direction[0]
            newY := pos[1] + direction[1]
            if newX < 0 || newX >= rows || newY < 0 || newY >= cols{
                continue
            }
            if grid[newX][newY] - grid[pos[0]][pos[1]] != 1{
                continue
            }
            if seen[newX][newY] == 1{
                continue
            }
            seen[newX][newY] = 1
            queue = append(queue, []int{newX,newY})
        } 
    }
    // fmt.Printf("Score from trail head at %v,%v is %v\n",r,c,score)
    return score
}

func Day10(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening the file: %v",err)
    }
    part1,part2 := 0,0
    defer file.Close()
    scanner := bufio.NewScanner(file)
    grid := [][]int{}
    for scanner.Scan(){
        line := scanner.Text()
        row := []int{}
        for idx := range(len(line)){
            row = append(row, int(line[idx] - '0'))
        }
        grid = append(grid, row)
    }
    rows,cols := len(grid),len(grid[0])
    trailHeads := [][]int{}
    for r := range(rows){
        for c := range(cols){
            if grid[r][c] == 0{
                trailHeads = append(trailHeads, []int{r,c})
            }
        }
    }
    for _,trailHead := range(trailHeads){
        part1 += getTrailHeadScore(trailHead[0],trailHead[1],grid)
        part2 += getTrailHeadRating(trailHead[0],trailHead[1],grid)
    }
    fmt.Printf("Part 1: %v, Part 2 : %v\n",part1,part2)
}
