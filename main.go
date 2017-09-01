package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/genzi/ttt_brain"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("www/*.html"))
}

var tttGame *ttt_brain.TicTacToe

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/receive", receiveAjax)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tttGame = ttt_brain.New()
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	//	var player int = 1
	if r.Method == "POST" {

		ajaxPostData := r.FormValue("ajax_post_data")
		fmt.Println(ajaxPostData)

		var move int
		move, _ = strconv.Atoi(ajaxPostData)

		if tttGame.HumanMove(move) == true {
			tttGame.AiMove()
			w.Write([]byte(ConvertBoardToString(tttGame.GetBoard())))
		}

		if isBoardFull() == true || tttGame.Win() != 0 {
			w.Write([]byte(ConvertBoardToString(tttGame.GetBoard())))
			switch tttGame.Win() {
			case 0:
				w.Write([]byte("draw"))
			case 1:
				w.Write([]byte("lose"))
			case -1:
				w.Write([]byte("win"))
			}
		}
	}
}

func ConvertBoardToString(intArray []int) []byte {
	byteArray := make([]byte, len(intArray))

	for i := 0; i < len(byteArray); i++ {

		byteArray[i] = ttt_brain.BoardStateToChar(intArray[i])
	}
	return byteArray
}

func isBoardFull() bool {
	board := ConvertBoardToString(tttGame.GetBoard())
	for i := 0; i < 9; i++ {
		if board[i] == ' ' {
			return false
		}
	}
	return true
}
