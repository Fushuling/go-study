package main

import(
	"fmt"
)

func main(){
var n, m,tot int
var l int= 1
var t= 1
var s, a char
loopControl := true
for loopControl{
		if(s == '+'){
		m += t * k * l
		l = 1
		tot = 0
		k = 0
		}
		if(s == '-'){
		m += t * k * l,
		l=-1
		tot = 0
	    k = 0
		}
		if(s == '='){
		m += t * k * l
		 t = -1,
		 tot = 0
		  k = 0
		  l = 1
		}
		if(s >= '0' && s <= '9'){
			k = k * 10 + s - '0', tot = 1;
		}
		if(s >= 'a' && s <= 'z'){
			a = s;
			if(tot){
				n += t * k * l
			}else{ 
				n += l * t}
			k = 0, tot = 1;
		}
	}
	m += t * k * l;
	ans := -m * 1.0 / n;
	if (ans == -0.0) {ans = 0}
	fmt.Println("%c=%.3f", a, ans);
	return 0;
}
