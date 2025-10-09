package main

import (
	"errors"
	"fmt"
)

type Order struct {
	Customer
	items  map[int]*OrderItem
	Status OrderStatus
}

type OrderItem struct {
	ID       int
	Name     string
	Count    uint
	Quantity int
	Price    float64
}

type Customer struct {
	Name string
}

// вспомогательное, в голову другого не пришло
type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (o *Order) AddItem(NewItem *OrderItem) {
	if o.items == nil {
		o.items = make(map[int]*OrderItem)
	}
	fmt.Printf("%s добавлен в заказ", NewItem.Name)
	o.items[NewItem.ID] = NewItem
}

func (o *Order) RemoveItem(itemID int) error {
	value, ok := o.items[itemID]
	if ok {
		fmt.Println(value.Name + " был убран из заказа =>")
		delete(o.items, itemID)
		return nil
	} else {
		return errors.New("Такого нет")
	}
}

func (o *Order) UpdateStatus(newStatus OrderStatus) {
	o.Status = newStatus
}

func (o *Order) GetStatus() string {
	statusMap := map[OrderStatus]string{
		OrderStatusCreated:   "Создан",
		OrderStatusPaid:      "Оплачен",
		OrderStatusDelivered: "Доставлен",
		OrderStatusCancelled: "Отменен",
	}
	return statusMap[o.Status]
}

func (o *Order) TotalCost() float64 {
	var sum float64
	for product := range o.items {
		sum += float64(o.items[product].Quantity) * o.items[product].Price
	}
	return sum
}

func main() {
	order := &Order{
		Customer: Customer{Name: "Иван Иванов"},
		Status:   OrderStatusCreated,
	}

	item1 := &OrderItem{ID: 1, Name: "Ноутбук", Quantity: 1, Price: 50000}
	item2 := &OrderItem{ID: 2, Name: "Мышь", Quantity: 2, Price: 1500}

	order.AddItem(item1)
	order.AddItem(item2)

	fmt.Printf("\nтекущий статус %s\n", order.GetStatus())
	fmt.Printf("общая стоимость: %.2f руб.\n", order.TotalCost())

	order.UpdateStatus(OrderStatusPaid)
	fmt.Printf("новый статус %s\n", order.GetStatus())

	err := order.RemoveItem(1)
	if err != nil {
		fmt.Println("Ошибка", err)
	}

	fmt.Printf("общая стоимость (после удаления) %.2f руб.\n", order.TotalCost())

	err = order.RemoveItem(999)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
}
