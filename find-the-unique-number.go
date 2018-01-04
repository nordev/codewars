/*
There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
It’s guaranteed that array contains more than 3 numbers.

The tests contain some very huge arrays, so think about performance.
*/

package kata

func FindUniq(arr []float32) float32 {
  one_number := arr[0]
  one_number_cnt := 1
  var other_number float32
  other_number_cnt := 0
  for i := 1; i < len(arr); i++ {
    if (arr[i] == one_number){
      one_number_cnt  += 1
    } else{
      other_number = arr[i]
      other_number_cnt +=1
    }
    if (one_number_cnt > 1 && other_number_cnt == 1){
      return other_number
    }else if (other_number_cnt > 1 && one_number_cnt == 1){
      return one_number
    }
    
  } 
}
