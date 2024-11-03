package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}	*/

package main

import "fmt"

func main() { 
	i := 0
	for i < 5  {
		fmt.Println(i)
		i++
	}
} */


package main

import "fmt"


func main() {
	var grade int
	fmt.Println("Enter your grade")
	fmt.Scan(&grade)
	//grade := 70
	if grade > 65 {
		fmt.Println("Student is passed")
	}
}*/


package main

import "fmt"


func main() {
	balance := -5
	if balance < 0 {
		fmt.Println("Please add amount else pelanty will be imposed")
	}
}*/

/*package main

import "fmt"

func main() {
	grade := 60
	if grade > 65 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}*/

/*package main

import "fmt"

func main() {
	balance := 5000
	if balance < 0 {
		fmt.Println("Penalty")
	} else if balance == 0 {
		fmt.Println("Balance is equal to 0, add funds soon.")
	} else   {
		fmt.Println("Balance is maintained")
	}
}*/


























package main

import "fmt"

func main() { 
	grade := 60
	if grade > 90 {
		fmt.Println("A grade")
	} else if grade > 80 {
		fmt.Println("B grade")
	} else if grade > 70 {
		fmt.Println("C grade")
	} else if grade > 60 {
		fmt.Println("D grade")
	} else {
		fmt.Println("Failed")
	}
}*/

package main

import "fmt"

func main() { 
	grade := 92
	if grade >= 65 {
		fmt.Println("Passed")
		if grade > 90 {
			fmt.Println("A")
 
		} else if grade > 80 {
			fmt.Println("B")
		} else if grade > 70 {
			fmt.Println("C")
		} else if grade >= 65 {
			fmt.Println("D")
	    }
	}  else {
		fmt.Println("Failing grade")
	}
} */

package main

import "fmt"

func main() {
    flavours := []string{"orange", "vanilla", "chocolate", "butterscotch"}

    for _, flav := range flavours {
        if flav == "orange" {
            fmt.Println(flav, "is my favorite")
            continue
        }

        if flav == "vanilla" {
            fmt.Println(flav, "is great")
            continue
        }

        if flav == "chocolate" {
            fmt.Println(flav, "is great")
            continue
        }

        if flav == "butterscotch" {
            fmt.Println(flav, "never tried")
        }
    }
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)


	for {
		var guess int
		fmt.Print("Enter a guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid guess: err:", err)
			continue
		}

		if guess > target {
			fmt.Println("Too high!")
			continue
		}

		if guess < target { 
			fmt.Println("Too low")
			continue
		}

		fmt.Println("You win!")
		break

}
} 


package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}


package main

import "fmt"

func main() {
	integers := make([]int, 10)
	fmt.Println(integers)
	
	for i:= range integers {
		integers[i] = i
	}

	fmt.Println(integers)
}*/








package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}



package main

import "fmt"

func main() {
	x := 10
	var p *int
	p = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Pointer of p:", p)
	fmt.Println("Value at p", *p)

	*p = 20
    fmt.Println("New value of x:", x)
} 
	*/


/*package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 5, 10
	fmt.Println("Before swap: x = ", x, " y = ", y)

	swap(&x, &y)
	fmt.Println("After swap: x = ", x, " y = ", y)

} 





package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
} 


