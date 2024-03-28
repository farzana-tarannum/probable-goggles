package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
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
  getOrders("orders.json", &orders)
  discounts := []Discount{}
  getDiscounts("discounts.json", &discounts)
  mapDiscounts := make(map[string]float64)

  for _, discount := range discounts {
    mapDiscounts[discount.Discount] = discount.Amount
  }

  mapProducts := make(map[int]Product)
  products := []Product{}
  getProducts("products.json", &products)
  
  for _, product := range products {
    fmt.Printf("SKU: %d, Price: %f\n", product.Sku, product.Price)
    mapProducts[product.Sku] = product
  }


  
  for _, order := range orders { 
    // discount is not empty
    discount := 1.0
    if order.Discount != "" {
      fmt.Printf("Order ID: %d, Discount: %s\n", order.OrderId, order.Discount)
      discount = mapDiscounts[order.Discount]
      fmt.Printf("Discount: %f\n", discount)
      
    } else {
      fmt.Printf("Order ID: %d\n", order.OrderId)
    }

    for _, item := range order.Items {  
      fmt.Printf("SKU: %d, Quantity: %d\n", item.Sku, item.Quantity)
      product := mapProducts[item.Sku]
      fmt.Printf("Price: %f\n", product.Price)
      total := product.Price * float64(item.Quantity) 
      discountedTotal := total - (total * discount)

      fmt.Printf("Total: %f Discounted Total: %f\n", total, discountedTotal)

    }
  }
}

















