package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var money int = 10000
	var choice string

	fmt.Println("카드 게임에 오신 것을 환영합니다.")

	for {
		fmt.Println("게임을 시작할까요? (y/n)")
		fmt.Scanln(&choice)

		if choice == "n" {
			fmt.Println("게임을 종료합니다. 현재 보유금액은", money, "원입니다.")
			break
		}

		if choice != "y" && choice != "n" {
			fmt.Println("잘못된 입력입니다. y 또는 n을 입력하세요.")
			continue
		}

		for {
			fmt.Println("베팅할 금액을 입력해주세요. 현재 보유금액은", money, "원입니다.")

			var bet int
			fmt.Scanln(&bet)

			if bet <= 0 {
				fmt.Println("0원 이하로는 베팅할 수 없습니다.")
				continue
			}
			if bet > money {
				fmt.Println("보유금액보다 많은 금액은 베팅할 수 없습니다.")
				continue
			}

			myCard := rand.Intn(10) + 1
			opponentCard := rand.Intn(10) + 1

			fmt.Println("당신의 카드:", myCard)
			fmt.Println("상대방의 카드:", opponentCard)

			if myCard > opponentCard {
				fmt.Println(">> 당신이 이겼습니다.")
				money += bet
			} else if myCard < opponentCard {
				fmt.Println(">> 상대방이 이겼습니다.")
				money -= bet
			} else {
				fmt.Println(">> 비겼습니다.")
			}

			fmt.Println("현재 보유금액:", money)
			fmt.Println()

			if money <= 0 {
				fmt.Println("보유금액이 0원이 되어 게임을 더 이상 진행할 수 없습니다.")
				return
			}

			fmt.Println("한 판 더 하시겠습니까? (y/n)")
			fmt.Scanln(&choice)
			if choice == "n" {
				fmt.Println("게임을 종료합니다. 현재 보유금액은", money, "원입니다.")
				return
			}
		}
	}
}
