package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

type Order struct {
  OrderId  int      `json:"orderId"`
  Discount string   `json:"discount"`
  Items    []Item   `json:"items"`
}

type Item struct {	
  Sku      int      `json:"sku"`
  Quantity int      `json:"quantity"`
}

type Product struct {
  Sku      int      `json:"sku"`
  Price    float64  `json:"price"`
}

type Discount struct {
  Discount string   `json:"key"`
  Amount   float64  `json:"value"`
}

func getOrders(filename string, orders *[]Order) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
  defer file.Close()
  data, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  err = json.Unmarshal(data, orders)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

func getDiscounts(filename string, discounts *[]Discount) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  defer file.Close()

  data, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }

  err = json.Unmarshal(data, discounts)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

func getProducts(filename string, products *[]Product) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
		os.Exit(-1)
	}
   
  defer file.Close()

  data, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
		os.Exit(-1)
	}

  err = json.Unmarshal(data, products)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
   

func main() {

  orders := []Order{}
  getOrders("part/orders.json", &orders)
  discounts := []Discount{}
  getDiscounts("part/discounts.json", &discounts)
  mapDiscounts := make(map[string]float64)

  for _, discount := range discounts {
    mapDiscounts[discount.Discount] = discount.Amount
  }

  mapProducts := make(map[int]Product)
  products := []Product{}
  getProducts("part/products.json", &products)
 
 
  for _, product := range products {
//    fmt.Printf("SKU: %d, Price: %f\n", product.Sku, product.Price)
    mapProducts[product.Sku] = product
  }
  


  averageDiscount := 0.0 
  for _, order := range orders { 
    // discount is not empty
    discount := 0.0
    orderedPrice := 0.0
    orderedDiscountedPrice := 0.0
    if order.Discount != "" {

      splitDiscount := []string{}
      // split discount string by comma
      splitDiscount = strings.Split(order.Discount, ",")

      for _, coupon := range splitDiscount { 
	fmt.Printf("Order ID: %d, Discount: %s, Coupon: %s\n", order.OrderId, order.Discount, coupon)
	discount += mapDiscounts[coupon]
	fmt.Printf("Discount: %f\n", discount)
      }
      
    } else {
      discount = 1.0
      fmt.Printf("Order ID: %d\n", order.OrderId)
    }

    for _, item := range order.Items {  
      fmt.Printf("SKU: %d, Quantity: %d\n", item.Sku, item.Quantity)
      product := mapProducts[item.Sku]
      fmt.Printf("Price: %f\n", product.Price)
      total := product.Price * float64(item.Quantity) 
      orderedPrice += total
      discountedTotal := total - (total * discount)
      orderedDiscountedPrice += discountedTotal

      fmt.Printf("Total: %f Discounted Total: %f\n", total, discountedTotal)

    }
    fmt.Printf("Ordered Price: %f, Ordered Discounted Price: %f\n", orderedPrice, orderedDiscountedPrice)
    averageDiscount += (orderedPrice - orderedDiscountedPrice) / orderedPrice
  }
  averageDiscount = averageDiscount / float64(len(orders))
  fmt.Printf("Average Discount: %f\n", averageDiscount * 100.0)
  
}

















