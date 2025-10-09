package main

import (
    "fmt"
    "sync"
    "time"
)

type Store struct {
    items map[string]int
    mu    sync.RWMutex
}

func (s *Store) Add(item string, quantity int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    s.items[item] += quantity
    fmt.Printf("поступление: %s +%d (остаток: %d)\n", item, quantity, s.items[item])
}

func (s *Store) Sell(item string, quantity int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if s.items[item] < quantity {
        fmt.Printf("продажа %s: недостаточно товара (%d < %d)\n", item, s.items[item], quantity)
        return false
    }
    
    s.items[item] -= quantity
    fmt.Printf("продажа: %s -%d (остаток: %d)\n", item, quantity, s.items[item])
    return true
}

func (s *Store) Status() {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    fmt.Println("\n склад ")
    for item, count := range s.items {
        fmt.Printf("%s: %d\n", item, count)
    }
}

func supplier(store *Store, item string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 1; i <= 3; i++ {
        store.Add(item, 10)
        time.Sleep(200 * time.Millisecond)
    }
}

func customer(store *Store, wg *sync.WaitGroup) {
    defer wg.Done()
    
    items := []string{"яблоки", "бананы", "апельсины"}
    for i := 1; i <= 4; i++ {
        for _, item := range items {
            store.Sell(item, 5)
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func main() {
    store := &Store{items: make(map[string]int)}
    var wg sync.WaitGroup
    

    store.Add("яблоки", 20)
    store.Add("бананы", 20)
    store.Add("апельсины", 20)
    

    wg.Add(3)
    go supplier(store, "яблоки", &wg)
    go supplier(store, "бананы", &wg)
    go supplier(store, "апельсины", &wg)
    
    
    wg.Add(2)
    go customer(store, &wg)
    go customer(store, &wg)
    
    wg.Wait()
    store.Status()
}
