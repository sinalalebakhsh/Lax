package main

import (
	"fmt"
	"time"
	)

func main() {
	n := 4
	stock := 2
	buyers := 0

	for i := 1; i <= n; i++ {
		fmt.Println("وارد فروشگاه م یشود" , i ,"مشتری") 
		fmt.Println("دارید؟ PS5 م یپرسد آیا" , i ," مشتری") 
		fmt.Println(i ,"فروشنده در حال چک کردن موجودی برای مشتری") 
		time.Sleep(200 * time.Millisecond)
		if stock <= 0 {
			fmt.Println("م یگوید نه متأسفانه موجودی به پایان رسیده است" , i ," فروشنده به مشتری ") 
			continue
			}

		fmt.Println("عدد موجود است" , stock ," م یگوید بله" , i ," فروشنده به مشتری") 
		fmt.Println("صادر م یکند" , i ,"فروشنده فاکتور فروش را برای مشتری ") 
		time.Sleep(150 * time.Millisecond)
		fmt.Println("خریداری م یکند PS5 یک عدد" , i ,"مشتری") 
		
		fmt.Println()

		stock--
		buyers++
		}

		fmt.Println()
		fmt.Println(buyers ,"تعداد افرادی که خریداری کرده اند") 
		fmt.Println(stock ,":تعداد باقی مانده") 
}

