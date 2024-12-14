package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getPositionAftertSeconds(pos []int,vel []int,maxX,maxY int,t int) []int{
    newX := pos[0] + (t * vel[0]) + (t * maxX)
    newY := pos[1] + (t * vel[1]) + (t * maxY)
    newX %= maxX
    newY %= maxY
    return []int{newX,newY}
}

func getSafetyFactor(positions,velocities [][]int,maxX,maxY int,t int) (int,[][]int){
    q1,q2,q3,q4 := 0,0,0,0
    //       |
    //  Q1   | Q2
    // ------------
    //  Q3   | Q4
    //       |
    newPositions := [][]int{}
    for i := range(len(positions)){
        newPos := getPositionAftertSeconds(positions[i],velocities[i],maxX,maxY,t)
        newPositions = append(newPositions,newPos)
        if newPos[0] < (maxX/2) && newPos[1] < (maxY/2){
            q1 += 1
        }
        if newPos[0] > (maxX/2) && newPos[1] < (maxY/2){
            q2 += 1
        }
        if newPos[0] < (maxX/2) && newPos[1] > (maxY/2){
            q3 += 1
        }
        if newPos[0] > (maxX/2) && newPos[1] > (maxY/2){
            q4 += 1
        }
    }
    return q1*q2*q3*q4,newPositions
}
func printGrid(positions [][]int,maxX,maxY int){
    grid := make([][]int,maxY+1,maxY+1)
    for y := range(maxY){
        grid[y] = make([]int,maxX+1,maxX+1)
    }
    for _,position := range(positions){
        grid[position[1]][position[0]] = 1
    }
    for y := range(maxY){
        for x := range(maxX){
            if grid[y][x] == 0{
                fmt.Print(" ")
            }else{
                fmt.Print("*")
            }
        }
        fmt.Println()
    }
}

func Day14(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening the file: %v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    positions := [][]int{}
    velocities := [][]int{}
    for (scanner.Scan()){
        splitBySpace := strings.Split(scanner.Text()," ")
        pos := strings.Split(splitBySpace[0],",")
        posX,_ := strconv.Atoi(strings.Trim(pos[0],"p="))
        posY,_ := strconv.Atoi(pos[1])
        positions = append(positions, []int{posX,posY})
        vel := strings.Split(splitBySpace[1],",")
        velX,_ := strconv.Atoi(strings.Trim(vel[0],"v="))
        velY,_ := strconv.Atoi(vel[1])
        velocities = append(velocities, []int{velX,velY})
    } 
    maxX := 101
    maxY := 103
    part1,_ := getSafetyFactor(positions,velocities,maxX,maxY,100)
    fmt.Printf("Part 1:%v\n",part1)
    for i := range(500){
        safetyFactor,newPositions := getSafetyFactor(positions,velocities,maxX,maxY,(7500 + i))
        // fmt.Println(safetyFactor)
        if safetyFactor <= 110 * 110 * 100 * 100{
            fmt.Printf("------------------------------------------------------- Time %v -----------------------------------------------------------------------------}\n",i)
            printGrid(newPositions,maxX,maxY)
        }
    }
}
