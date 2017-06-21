package raindrops

/*import "fmt"

const testVersion = 3
/*

/*Convert converts an integer to "raindrop string"

    If the number has 3 as a factor, output 'Pling'.
	    If the number has 5 as a factor, output 'Plang'.
		    If the number has 7 as a factor, output 'Plong'.
			    If the number does not have 3, 5, or 7 as a factor,
					just pass the number's digits straight through.
*/
/*
func Convert(num int) string {
	dict := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	s := ""
	facs := []int{3, 5, 7}
	isDivisible := func(x, y int) bool {
		return x%y == 0
	}
	for _, factor := range facs {
		if isDivisible(num, factor) {
			s += dict[factor]
		}
	}
	if s == "" {
		return fmt.Sprintf("%d", num)
	}
	return s
*/
