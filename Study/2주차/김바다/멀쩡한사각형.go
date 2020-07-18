/* 실수 계산에서 오차때문에 값이 정확하게 나오지 않는 경우가 있음. 개선 필요*/
func solution(w int, h int) int64 {
    var answer int64 
    var sum int =0
    var dif int
    var i,before,after int
    if w==h {
        return int64(w)*int64(w-1)
    }
 
    if w<h {
        for i=1;i<=w;i++{
            after = int(float64(h)/float64(w)*float64(i))
            dif = after-before
            sum+=dif+1
            before=after
        }
    } else {
        for i=1;i<=h;i++{
            after = int(float64(w)/float64(h)*float64(i))
            dif = after-before
            sum+=dif+1
            before=after
        }
    }
    
    answer =int64(w)*int64(h) - int64(sum)
    return answer
}