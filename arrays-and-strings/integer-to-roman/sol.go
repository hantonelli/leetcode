package integer_to_roman

// X 10
// L 50
// C 100
// D 500
// M 1000

var (
	mapping = map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
		200: "CC",
		300: "CCC",
		400: "CD",
		500: "D",
		600: "DC",
		700: "DCC",
		800: "DCCC",
		900: "CM",
	}
)

func intToRoman(num int) string {
	res := ""
	for 0 < num-num%1000 {
		res += "M"
		num -= 1000
	}
	cent := (num - num%100) / 100
	if 0 < cent {
		res += mapping[cent*100]
		num -= cent * 100
	}
	des := (num - num%10) / 10
	if 0 < des {
		res += mapping[des*10]
		num -= des * 10
	}
	if num != 0 {
		res += mapping[num]
	}
	return res
}
