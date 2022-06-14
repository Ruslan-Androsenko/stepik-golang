package structures

import "fmt"

const MAXIMUM_AMMO int = 50
const MAXIMUM_POWER int = 100

// RefillAmmo Пополнить боеприпасы
func (this *Hero) RefillAmmo() {
	this.beforeAction()

	if this.On {
		var ammo int

		for ammo <= 0 || ammo > MAXIMUM_AMMO {
			fmt.Print("\nHow much ammo do you want to add? >_ ")
			fmt.Scan(&ammo)

			if ammo == 0 {
				break
			}
		}

		if this.Ammo+ammo > MAXIMUM_AMMO {
			this.Ammo = MAXIMUM_AMMO
		} else {
			this.Ammo += ammo
		}
	}
}

// RefillPower Восполнить энергию
func (this *Hero) RefillPower() {
	this.beforeAction()

	if this.On {
		var power int

		for power <= 0 || power > MAXIMUM_POWER {
			fmt.Print("\nHow much power do you want to add? >_ ")
			fmt.Scan(&power)

			if power == 0 {
				break
			}
		}

		if this.Power+power > MAXIMUM_POWER {
			this.Power = MAXIMUM_POWER
		} else {
			this.Power += power
		}
	}
}

// ShootAtEnemies Стрелять по врагам
func (this *Hero) ShootAtEnemies() {
	this.beforeAction()

	if this.On {
		var counter int

		for counter <= 0 || counter > MAXIMUM_AMMO {
			fmt.Print("\nHow much shots do you want to take? >_ ")
			fmt.Scan(&counter)

			if counter == 0 {
				break
			}
		}

		for shoot := 1; shoot <= counter; shoot++ {
			if this.Shoot() {
				fmt.Print("Piu")
			} else {
				fmt.Println("\n\nClip in your weapon is empty.")
				break
			}
			if shoot < counter {
				fmt.Print("-")
			}
		}
	}
}

// RideABike Ехать на байке
func (this *Hero) RideABike() {
	this.beforeAction()

	if this.On {
		var counter int

		for counter <= 0 || counter > MAXIMUM_POWER {
			fmt.Print("\nHow far do you want to drive? >_ ")
			fmt.Scan(&counter)

			if counter == 0 {
				break
			}
		}

		for distance := 1; distance <= counter; distance++ {
			if this.Drive() {
				fmt.Print("Brr")
			} else {
				fmt.Println("\n\nYour energy is gone.")
				break
			}
			if distance < counter {
				fmt.Print("-")
			}
		}
	}
}

// ShowState Вывести состояние персонажа
func (this *Hero) ShowState() {
	this.beforeAction()

	if this.On {
		fmt.Printf("\nammo: %d, power: %d \n", this.Ammo, this.Power)
	}
}

func (this *Hero) beforeAction() {
	if !this.On {
		fmt.Println("\nHero is not active.")
		this.Activate()
	}
}

// Activate Активация персонажа
func (this *Hero) Activate() {
	if !this.On {
		var answer string
		fmt.Print("Are you want his to active? (y/n) >_ ")
		fmt.Scan(&answer)

		this.On = answer == "y" || answer == "yes" || answer == "Yes"
	} else {
		fmt.Println("\nThe hero is already activated.")
	}
}

// Deactivate Деактивация персонажа
func (this *Hero) Deactivate() {
	if this.On {
		var answer string
		fmt.Print("Are you want his to deactivate? (y/n) >_ ")
		fmt.Scan(&answer)

		this.On = answer == "n" || answer == "no" || answer == "No"
	} else {
		fmt.Println("\nThe hero is already deactivated.")
	}
}
