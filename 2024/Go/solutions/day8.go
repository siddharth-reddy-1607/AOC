package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AntinodeProperty int

const (
    outOfBounds  AntinodeProperty =  iota
    seenAntinode
    validAntinode
)
func isValidAntinode(r,c int,seen [][]int,grid [][]byte) AntinodeProperty{
    rows,cols := len(grid),len(grid[0])
    if r < 0 || r >= rows || c < 0 || c >= cols{
        return outOfBounds
    }
    if seen[r][c] == -1{
        return seenAntinode
    }
    seen[r][c] = -1
    return validAntinode

}
func getCountOfValidAntinodes(a []int,b []int,seen [][]int,grid [][]byte) int{
    validAntinodes := 0
    rowDist := abs(a[0] - b[0])
    colDist := abs(a[1]-b[1])
    r1,r2,c1,c2 := -1,-1,-1,-1
    if a[0] < b[0]{
        if a[1] < b[1]{
            r1 = a[0] - rowDist
            c1 = a[1] - colDist
            r2 = b[0] + rowDist
            c2 = b[1] + colDist
        }else{
            r1 = a[0] - rowDist
            c1 = a[1] + colDist
            r2 = b[0] + rowDist
            c2 = b[1] - colDist
        }
    }else{
        if a[1] < b[1]{
            r1 = a[0] + rowDist
            c1 = a[1] - colDist
            r2 = b[0] - rowDist
            c2 = b[1] + colDist
        }else{
            r1 = a[0] + rowDist
            c1 = a[1] + colDist
            r2 = b[0] - rowDist
            c2 = b[1] - colDist
        }
    }
    res := isValidAntinode(r1,c1,seen,grid) 
    if res == validAntinode{
        // fmt.Printf("(%v,%v) is a valid antinode\n",r1,c1)
        validAntinodes += 1
    }
    res = isValidAntinode(r2,c2,seen,grid) 
    if res == validAntinode{
        validAntinodes += 1
    }
    return validAntinodes
}

func getCountOfValidAntinodesWithoutDistance(a, b []int, seen [][]int, grid [][]byte) int{
    validAntinodesWithoutDistance := 0
    rowDist := abs(a[0] - b[0])
    colDist := abs(a[1] - b[1])
    if a[0] < b[0]{
        if a[1] < b[1]{
            for i := 1;;i++{
                r := a[0] - (i*rowDist)
                c := a[1] - (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
            for i := 1;; i++{
                r := b[0] + (i*rowDist)
                c := b[1] + (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
        }else{
            for i := 1;;i++{
                r := a[0] - (i*rowDist)
                c := a[1] + (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
            for i := 1;; i++{
                r := b[0] + (i*rowDist)
                c := b[1] - (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
        }
    }else{
        if a[1] < b[1]{
            for i := 1;;i++{
                r := a[0] + (i*rowDist)
                c := a[1] - (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
            for i := 1;; i++{
                r := b[0] - (i*rowDist)
                c := b[1] + (i*colDist)
                res := isValidAntinode(r,c,seen,grid)
                if res == outOfBounds{
                    break
                }
                if res == seenAntinode{
                    continue
                }
                // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                validAntinodesWithoutDistance += 1
            }
        }else{
            if a[1] < b[1]{
                for i := 1;;i++{
                    r := a[0] + (i*rowDist)
                    c := a[1] + (i*colDist)
                    res := isValidAntinode(r,c,seen,grid)
                    if res == outOfBounds{
                        break
                    }
                    if res == seenAntinode{
                        continue
                    }
                    // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                    validAntinodesWithoutDistance += 1
                }
                for i := 1;; i++{
                    r := b[0] - (i*rowDist)
                    c := b[1] - (i*colDist)
                    res := isValidAntinode(r,c,seen,grid)
                    if res == outOfBounds{
                        break
                    }
                    if res == seenAntinode{
                        continue
                    }
                    // fmt.Printf("(%v,%v) is a valid antinode\n",r,c)
                    validAntinodesWithoutDistance += 1
                }
            }
        }
    }
    return validAntinodesWithoutDistance
}


func Day8(){
    part1,part2 := 0,0
    file,err := os.Open("../input.txt")
    if err != nil {
        log.Fatalf("Error while opening the file: %v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)


    hashmap := make(map[byte][][]int)
    grid := [][]byte{}
    for scanner.Scan(){
        grid = append(grid,[]byte(scanner.Text()))
    }
    rows,cols := len(grid),len(grid[0])
    for r := range(rows){
        for c := range(cols){
            chr := grid[r][c]
            if chr == '.'{
                continue
            }
            if _,found := hashmap[chr]; !found{
                hashmap[chr] = [][]int{}
            }
            hashmap[chr] = append(hashmap[chr], []int{r,c})
        }
    }
    seenAntinodes := make([][]int,rows,rows)
    for r := range(rows){
        seenAntinodes[r] = make([]int,cols,cols)
    }
    seenAntinodesWithoutDistance := make([][]int,rows,rows)
    for r := range(rows){
        seenAntinodesWithoutDistance[r] = make([]int,cols,cols)
    }
    fmt.Println(rows,cols)
    for _,positions := range(hashmap){
        for i := range(len(positions)){
            for j := i+1; j < len(positions); j++{
                // fmt.Printf("Getting antinodes between points %v,%v\n",positions[i],positions[j])
                part1 += getCountOfValidAntinodes(positions[i],positions[j],seenAntinodes,grid)
                part2 += getCountOfValidAntinodesWithoutDistance(positions[i],positions[j],seenAntinodesWithoutDistance,grid)
            }
        }
        if len(positions) > 1{
            for _,pos := range(positions){
                if seenAntinodesWithoutDistance[pos[0]][pos[1]] == -1{
                    continue
                }
                seenAntinodesWithoutDistance[pos[0]][pos[1]] = -1
                part2 += 1
            }
        }
    }
    fmt.Printf("Part 1 : %v, Part 2 : %v\n",part1,part2,)
}
