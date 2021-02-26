package main

import (
	"chinachess/battle"
	"chinachess/piece"
	"chinachess/web"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/manifoldco/promptui/screenbuf"
)

func main() {
	battleStart()
	// printCanMove()
}

func battleStart() {
	battle.ResetBoard()
	battle.InitBattle(true, promptLayout)
	battle.GBattle.Print()

	wb := web.New(battle.GBattle, true)
	wb.Listen(":8090")

	for {
		side := battle.GBattle.SideNow()
		var item []*piece.PiecePoint
		if side == piece.SideRed {
			item = battle.GBattle.Red.PieceNow
		} else {
			item = battle.GBattle.Black.PieceNow
		}

		prompt1 := promptui.Select{
			Label:        "choose piece",
			Items:        item,
			IsVimMode:    true,
			HideSelected: true,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "\U0001F336 {{ .Name | cyan }} ({{.Point.X}},{{.Point.Y}})[{{.}}]",
				Inactive: "  {{ .Name | cyan }} ({{.Point.X}},{{.Point.Y}})",
				Selected: "\U0001F336 {{ .Name | red | cyan }}",
				Details:  "",
			},
		}

		i1, _, err := prompt1.Run()
		if err != nil {
			fmt.Println("move failed:", err)
			break
		}

		pp := item[i1]
		canMoves := pp.CanMove(&battle.GBattle.Map)
		prompt2 := promptui.Select{
			Label: "move piece",
			Items: canMoves,
			// Templates: &promptui.SelectTemplates{
			// 	Label:    "{{ . }}?",
			// 	Active:   "\U0001F336 ({{.Point.X}},{{.Point.Y}})",
			// 	Inactive: "  ({{.Point.X}},{{.Point.Y}})",
			// 	Selected: "\U0001F336 ({{.Point.X}},{{.Point.Y}})",
			// 	// Help:    "帮助",
			// },
		}

		i2, _, err := prompt2.Run()
		if err != nil {
			fmt.Println("move failed:", err)
			break
		}
		battle.GBattle.Move(item[i1], canMoves[i2])
		battle.GBattle.Switch()
		if battle.GBattle.IsOver() {
			break
		}
		battle.GBattle.Print()
	}
}

func flush() {
	sb := screenbuf.New(os.Stdout)
	for {
		sb.Reset()
		sb.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
		sb.Flush()
		time.Sleep(1 * time.Second)
	}
}

func printCanMove() {
	battle.ResetBoard()
	battle.InitBattle(true, nil)

	for _, pn := range battle.GBattle.Red.PieceNow {
		canMoves := pn.CanMove(&battle.GBattle.Map)
		_ = canMoves
		fmt.Printf("原来= 【%v】canMoves = %v\n", *pn, canMoves)
	}
	for _, pn := range battle.GBattle.Black.PieceNow {
		canMoves := pn.CanMove(&battle.GBattle.Map)
		_ = canMoves
		fmt.Printf("原来= 【%v】canMoves = %v\n", *pn, canMoves)
	}
}
