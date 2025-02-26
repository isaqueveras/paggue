package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/isaqueveras/leaf"
)

func main() {
	// 1 - context
	// 2 - total number of workers
	// 3 - total items per page
	// 4 - interval at which the publish function will be executed
	queue := leaf.New(context.Background(), 100, 100, time.Second)
	defer queue.Wait()

	queue.Publish(publish)
	queue.Consume(consume)
}

func consume(ctx context.Context) error {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	value := leaf.GetData(ctx).(time.Time)
	log.Printf("[consumer] - Value: %v\n", value.String())
	return nil
}

func publish(ctx context.Context) (interface{}, error) {
	page := leaf.GetPage(ctx)

	log.Printf("[publisher] - Page: %d Offset: %d ItemsPerPage %d | Cursor: %d/%d",
		page.GetCurrentPage(), page.GetOffset(), page.GetItemsPerPage(), page.GetStartCursor(), page.GetEndCursor())

	if page.GetCurrentPage() == 1 {
		page.SetCursor(230)
	}

	if page.GetCurrentPage() == 3 {
		page.SetCursor(12)
	}

	if page.GetCurrentPage() == 5 {
		leaf.Stop(ctx)
	}

	return time.Now(), nil
}
