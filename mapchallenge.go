package main

import "fmt"

func main() {
    // Create a map called 'studentGrades' which maps student names (string) to their grades (int).
    // Add at least three students with their respective grades.
    // Print the name and grade of each student in the map.
    
    // Add your solution here!
    
    for student, grade := range studentGrades {
        fmt.Printf("Student: %s, Grade: %d\n", student, grade)
    }
}
