package main

import (
    "fmt"
    "sort"
    "strings"
)

func property_divider(properties map[string]int, limit int, priorityList string) map[string]int {
    // Parse priority list and create priority map
    priority := strings.Split(priorityList, "-")
    priorityMap := make(map[string]int)
    for i, value := range priority {
        priorityMap[value] = i
    }

    // Calculate base allocations and fractions
    result := make(map[string]int)
    var alloted int

    type PropertyInfo struct {
        id        string
        fraction  float64
        priority  int
    }
    
    var propertyList []PropertyInfo
    
    // Calculate initial allocations and collect non-zero fractions
    for key, percentage := range properties {
        exactValue := (float64(percentage) * float64(limit)) / 100.0
        baseValue := int(exactValue)
        result[key] = baseValue
        alloted += baseValue
        
        fraction := exactValue - float64(baseValue)
        // Only collect properties with non-zero fractions
        if fraction > 0 {
            priority := len(priority) // Default to lowest priority if not in list
            if val, exists := priorityMap[key]; exists {
                priority = val
            }
            propertyList = append(propertyList, PropertyInfo{key, fraction, priority})
        }
    }

    // Calculate remaining units to distribute
    remaining := limit - alloted

    // If there are remaining units and properties with non-zero fractions
    if remaining > 0 && len(propertyList) > 0 {
        // Sort by fraction (descending) and then by priority
        sort.SliceStable(propertyList, func(i, j int) bool {
            if propertyList[i].fraction != propertyList[j].fraction {
                return propertyList[i].fraction > propertyList[j].fraction
            }
            return propertyList[i].priority < propertyList[j].priority
        })

        // Distribute remaining units only to properties with non-zero fractions
        for i := 0; i < remaining && i < len(propertyList); i++ {
            result[propertyList[i].id]++
        }
    }

    return result
}

func main() {
    properties := map[string]int{
        "11": 30,
        "24": 35,
        "12": 35,
    }

    // Test cases with different scenarios
    testCases := []struct {
        limit       int
        priorityList string
    }{
        {1, "12-11-24"}, 
        {100, "12-11-24"},
        {7, "12-11-24"},
        {20, "12-11-24"},
    }

    for _, tc := range testCases {
        fmt.Printf("\nLimit: %d, Priority: %s\n", tc.limit, tc.priorityList)
        result := property_divider(properties, tc.limit, tc.priorityList)
        
        // Print detailed distribution information
        fmt.Printf("Result: %v\n", result)
        
    }

	properties2 := map[string]int{
        "11": 55,
        "24": 45,
	}

	result2 := property_divider(properties2, 3, "24-11")

	fmt.Println(result2)

}