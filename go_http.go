package main

// 필요한 패키지 임포트
import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func defaultHnadler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Get 파라미터 및 정보 출력
	fmt.Println("default : ", r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("param : ", r.Form["test_param"])
	// Parameter 전체 출력
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 기본 출력
	// fmt.Fprintf(w, "Golang WebServer Working!")
	fmt.Fprintf(w, "안녕하세요sf")
}

func main() {
	// 기본 Url 핸들러 매소드 지정
	http.HandleFunc("/", defaultHnadler)
	// 서버 시작
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(9090)")
	}
}
