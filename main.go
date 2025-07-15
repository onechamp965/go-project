package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var money int = 10000 // 돈
	var choice string // 사용자 선택

	fmt.Println("카드 게임에 오신 것을 환영합니다.")

	for {
		fmt.Println("게임을 시작할까요? (y/n)") 
		fmt.Scanln(&choice) // 게임의 시작 여부 결정

		if choice == "n" { // 게임 종료 및 현재 보유금 출력
			fmt.Println("게임을 종료합니다. 현재 보유금액은", money, "원입니다.")
			break
		}

		if choice != "y" && choice != "n" { // 잘못된 입력 판단
			fmt.Println("잘못된 입력입니다. y 또는 n을 입력하세요.")
			continue
		}

		for {
			fmt.Println("베팅할 금액을 입력해주세요. 현재 보유금액은", money, "원입니다.") // 베팅 및 현재 보유금 출력

			var bet int // 베팅금액 선언
			fmt.Scanln(&bet) // 베팅금액 입력

			if bet <= 0 { // 베팅할 금액이 0보다 작은지 판단
				fmt.Println("0원 이하로는 베팅할 수 없습니다.")
				continue
			}
			if bet > money { // 베팅할 금액이 보유금보다 많은지 판단
				fmt.Println("보유금액보다 많은 금액은 베팅할 수 없습니다.")
				continue
			}

			myCard := rand.Intn(10) + 1 // 내 카드 선언
			opponentCard := rand.Intn(10) + 1 // 상대 카드 선언

			fmt.Println("당신의 카드:", myCard)
			fmt.Println("상대방의 카드:", opponentCard)

			if myCard > opponentCard {
				fmt.Println(">> 당신이 이겼습니다.") // 내 승리
				money += bet
			} else if myCard < opponentCard {
				fmt.Println(">> 상대방이 이겼습니다.") // 상대 승리
				money -= bet
			} else {
				fmt.Println(">> 비겼습니다.") // ㄲㅂ
			}

			fmt.Println("현재 보유금액:", money) // 현재 보유금 출력
			fmt.Println() // 줄바꿈

			if money <= 0 {
				fmt.Println("보유금액이 0원이 되어 게임을 더 이상 진행할 수 없습니다.") // 돈이 0일때 출력
				return
			}

			fmt.Println("한 판 더 하시겠습니까? (y/n)") // 응 씨발 안해
			fmt.Scanln(&choice)
			if choice == "n" {
				fmt.Println("게임을 종료합니다. 현재 보유금액은", money, "원입니다.") // fuck
				return
			}
		}
	}
}
