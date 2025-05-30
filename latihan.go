package main

import "fmt"

func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // Inisialisasi baris pertama dan kolom pertama = 1
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }

    // Hitung jumlah jalur untuk setiap sel
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}

func main() {
    fmt.Println("Unique Paths (3x7):", uniquePaths(3, 7)) // Output: 28
    fmt.Println("Unique Paths (3x2):", uniquePaths(3, 2)) // Output: 3
}
