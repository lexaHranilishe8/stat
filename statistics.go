package statistics

// Avg возвращает среднее значение массива чисел.
func Avg(nums []float64) float64 {
    if len(nums) == 0 { // исправил ошибку с делением на 0 
        return 0
    }
    var sum float64
    for _, n := range nums {
        sum += n
    }
    return sum / float64(len(nums))
}
