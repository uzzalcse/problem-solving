package main

import (
    "testing"
)

func TestPropertyDivider(t *testing.T) {
    tests := []struct {
        name         string
        properties   map[string]int
        limit        int
        priorityList string
        want        map[string]int
    }{
        {
            name: "Three properties with limit 1",
            properties: map[string]int{
                "11": 30,
                "24": 35,
                "12": 35,
            },
            limit: 1,
            priorityList: "12-11-24",
            want: map[string]int{
                "11": 0,
                "24": 0,
                "12": 1,
            },
        },
        {
            name: "Three properties with limit 100",
            properties: map[string]int{
                "11": 30,
                "24": 35,
                "12": 35,
            },
            limit: 100,
            priorityList: "12-11-24",
            want: map[string]int{
                "11": 30,
                "24": 35,
                "12": 35,
            },
        },
        {
            name: "Three properties with limit 7",
            properties: map[string]int{
                "11": 30,
                "24": 35,
                "12": 35,
            },
            limit: 7,
            priorityList: "12-11-24",
            want: map[string]int{
                "11": 2,
                "24": 2,
                "12": 3,
            },
        },
        {
            name: "Equal percentages with priority",
            properties: map[string]int{
                "A": 50,
                "B": 50,
            },
            limit: 5,
            priorityList: "A-B",
            want: map[string]int{
                "A": 3,
                "B": 2,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := property_divider(tt.properties, tt.limit, tt.priorityList)
            
            // Check if total allocation equals limit
            total := 0
            for _, v := range got {
                total += v
            }
            if total != tt.limit {
                t.Errorf("Total allocation = %v, want %v", total, tt.limit)
            }

            // Check if results match expected values
            for k, v := range tt.want {
                if got[k] != v {
                    t.Errorf("property_divider() for property %v = %v, want %v", k, got[k], v)
                }
            }
        })
    }
}

func TestEdgeCases(t *testing.T) {
    t.Run("Zero limit", func(t *testing.T) {
        properties := map[string]int{
            "11": 30,
            "24": 35,
            "12": 35,
        }
        got := property_divider(properties, 0, "12-11-24")
        for k, v := range got {
            if v != 0 {
                t.Errorf("Expected 0 for property %v, got %v", k, v)
            }
        }
    })

    t.Run("Empty priority list", func(t *testing.T) {
        properties := map[string]int{
            "11": 50,
            "24": 50,
        }
        got := property_divider(properties, 2, "")
        total := 0
        for _, v := range got {
            total += v
        }
        if total != 2 {
            t.Errorf("Total allocation = %v, want 2", total)
        }
    })

    t.Run("Property not in priority list", func(t *testing.T) {
        properties := map[string]int{
            "11": 50,
            "99": 50,
        }
        got := property_divider(properties, 2, "11")
        total := 0
        for _, v := range got {
            total += v
        }
        if total != 2 {
            t.Errorf("Total allocation = %v, want 2", total)
        }
    })
}

func TestFractionHandling(t *testing.T) {
    tests := []struct {
        name         string
        properties   map[string]int
        limit        int
        priorityList string
    }{
        {
            name: "Small fractions",
            properties: map[string]int{
                "A": 33,
                "B": 33,
                "C": 34,
            },
            limit: 10,
            priorityList: "A-B-C",
        },
        {
            name: "Large fractions",
            properties: map[string]int{
                "A": 66,
                "B": 34,
            },
            limit: 3,
            priorityList: "A-B",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := property_divider(tt.properties, tt.limit, tt.priorityList)
            
            // Verify total allocation
            total := 0
            for _, v := range got {
                total += v
            }
            if total != tt.limit {
                t.Errorf("Total allocation = %v, want %v", total, tt.limit)
            }

            // Verify each property got some allocation
            for k := range tt.properties {
                if _, exists := got[k]; !exists {
                    t.Errorf("Property %v not found in result", k)
                }
            }
        })
    }
}