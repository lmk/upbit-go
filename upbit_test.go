package upbit

import (
	"fmt"
	"log"
)

var client *Client

const (
	accessKey = ""
	secretKey = ""
)

func setUp() {
	client = NewClient(accessKey, secretKey)
}

func ExampleGetMarkets() {
	setUp()

	markets, err := client.Markets()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(markets[0].Market)

	// Output:
	// KRW-BTC
}

func ExampleGetMinuteCandles() {
	setUp()

	candles, err := client.MinuteCandles(1, "KRW-BTC", map[string]string{
		"count": "1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 1
	// KRW-BTC
}

func ExampleGetMinuteCandlesWithDate() {
	setUp()

	candles, err := client.MinuteCandles(1, "KRW-BTC", map[string]string{
		"count": "5",
		"to":    "2021-02-17 13:00:00",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	for _, candle := range candles {
		fmt.Println(candle)
	}
	// Output:
	// 5
	// &{KRW-BTC 2021-02-17T12:59:00 2021-02-17T21:59:00 5.6154e+07 5.6216e+07 5.6153e+07 5.6202e+07 1613566799895 5.5939488761622e+08 9.95936628 1}
	// &{KRW-BTC 2021-02-17T12:58:00 2021-02-17T21:58:00 5.6166e+07 5.617e+07 5.6153e+07 5.6153e+07 1613566740326 7.2022304480913e+08 12.82496912 1}
	// &{KRW-BTC 2021-02-17T12:57:00 2021-02-17T21:57:00 5.6169e+07 5.62e+07 5.6165e+07 5.6166e+07 1613566680374 5.7768898576905e+08 10.28399756 1}
	// &{KRW-BTC 2021-02-17T12:56:00 2021-02-17T21:56:00 5.615e+07 5.6294e+07 5.6091e+07 5.6186e+07 1613566620104 6.4990655501209e+08 11.56963224 1}
	// &{KRW-BTC 2021-02-17T12:55:00 2021-02-17T21:55:00 5.6025e+07 5.6171e+07 5.6e+07 5.615e+07 1613566560096 1.01642957584032e+09 18.13659205 1}
}

func ExampleWrongUnitGetMinuteCandles() {
	setUp()

	_, err := client.MinuteCandles(2, "KRW-BTC")

	fmt.Println(err)

	// Output:
	// Invalid unit
}

func ExampleDayCandles() {
	setUp()

	candles, err := client.DayCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleWeekCandles() {
	setUp()

	candles, err := client.WeekCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleMonthCandles() {
	setUp()

	candles, err := client.MonthCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleTradeTicks() {
	setUp()

	tradeTicks, err := client.TradeTicks("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(tradeTicks))
	fmt.Println(tradeTicks[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleGetTickers() {
	setUp()

	ticks, err := client.Ticker("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)

	// Output:
	// 2
	// KRW-BTC KRW-TRX
}

func ExampleOrderbooks() {
	setUp()

	ticks, err := client.Orderbooks("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)
	fmt.Println(len(ticks[0].OrderbookUnits))

	// Output:
	// 2
	// KRW-BTC KRW-TRX
	// 10
}

func ExampleAccounts() {
	setUp()

	accounts, err := client.Accounts()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(accounts[0].Currency)

	// Output:
	// KRW
}

func ExampleOrderChange() {
	setUp()

	orderChance, err := client.OrderChance("KRW-BTC")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(orderChance.Market.Id)

	// Output:
	// KRW-BTC
}

// func ExampleSell() {
// 	setUp()
//
// 	order, err := client.Order(
// 		strconv.Itoa(int(util.TimeStamp())),
// 		"bid",
// 		"BTC-TRX",
// 		"0.000003",
// 		"1000",
// 		"limit",
// 	)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
//
// 	fmt.Println(order)
//
// 	// Output:
// 	// BTC-TRX
// }
