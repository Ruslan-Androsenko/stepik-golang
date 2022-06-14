package structures

import (
	"fmt"
	"strings"
)

// StepEight
/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции,
как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true),
если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры
с именем testStruct в функции main, в дальнейшем программа проверит результат.

Удачи!
*/
func StepEight() {
	testStruct(new(Hero))
}

func testStruct(hero *Hero) {
	var action int

	for {
		showMenu()
		fmt.Print("\nВыберите действие >_ ")
		fmt.Scan(&action)

		switch action {
		case 1:
			hero.ShowState()
		case 2:
			hero.RefillAmmo()
		case 3:
			hero.RefillPower()
		case 4:
			hero.ShootAtEnemies()
		case 5:
			hero.RideABike()
		case 6:
			hero.Activate()
		case 7:
			hero.Deactivate()

		case 0: // Выход из игры
			starsLine := strings.Repeat("*", 50)
			fmt.Printf("\n%s\n", starsLine)
			fmt.Print("\t\tGame Over! \n\t\tI will be back!")
			fmt.Printf("\n%s\n", starsLine)
			return
		}
	}
}

func showMenu() {
	fmt.Printf("\n%s\n", strings.Repeat("=", 50))
	fmt.Println("*** Меню игры ***")

	fmt.Println("1. Показать состояние персонажа.")
	fmt.Println("2. Пополнить боеприпасы.")
	fmt.Println("3. Восполнить энергию.")
	fmt.Println("4. Стрелять по врагам.")
	fmt.Println("5. Ехать на байке.")
	fmt.Println("\n6. Активировать персонажа.")
	fmt.Println("7. Деактивировать персонажа.")
	fmt.Println("0. Выход из игры.")
}
