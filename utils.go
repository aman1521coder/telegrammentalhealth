package main
import(
	"strings"
)
var crisesworss=[]string {"sucide","kill myself","die","end my life"}

func IsDanger(text  string )bool{
lowertext:=strings.ToLower(text)
for _,word:=range crisesworss{
	if strings.Contains(lowertext,word){
		return true
}}
	return false
}
