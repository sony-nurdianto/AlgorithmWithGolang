package better_queue

import (
	"fmt"
	"sync"
	"time"
)

// Queue adalah struktur antrian yang aman untuk konkurensi
type Queue struct {
	items chan int      // Channel untuk menyimpan item antrian
	stop  chan struct{} // Channel untuk menghentikan antrian
	wg    sync.WaitGroup
}

// NewQueue membuat antrian baru dengan kapasitas tertentu
func NewQueue(capacity int) *Queue {
	return &Queue{
		items: make(chan int, capacity),
		stop:  make(chan struct{}),
	}
}

// Enqueue menambahkan item ke antrian
func (q *Queue) Enqueue(item int) {
	select {
	case q.items <- item: // Coba kirim item ke channel
		fmt.Printf("Enqueued: %d\n", item)
	case <-q.stop: // Jika antrian dihentikan, keluar
		return
	}
}

// Dequeue mengambil item dari antrian
func (q *Queue) Dequeue() int {
	select {
	case item := <-q.items: // Coba terima item dari channel
		fmt.Printf("Dequeued: %d\n", item)
		return item
	case <-q.stop: // Jika antrian dihentikan, keluar
		return -1
	}
}

// Stop menghentikan antrian
func (q *Queue) Stop() {
	close(q.stop) // Tutup channel stop untuk memberi sinyal berhenti
	q.wg.Wait()   // Tunggu semua goroutine selesai
}

func (q *Queue) StartWorker() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			select {
			case item := <-q.items:
				fmt.Printf("Processing: %d\n", item)
				time.Sleep(500 * time.Millisecond)
			case <-q.stop:
				fmt.Println("Worker stopped")
				return
			}
		}
	}()
}

func main() {
	// Buat antrian dengan kapasitas 5
	queue := NewQueue(5)

	// Mulai worker untuk memproses item
	queue.StartWorker()

	// Tambahkan item ke antrian
	for i := 1; i <= 10; i++ {
		queue.Enqueue(i)
		time.Sleep(200 * time.Millisecond)
	}

	// Beri waktu untuk memproses item
	time.Sleep(2 * time.Second)

	// Hentikan antrian
	queue.Stop()
	fmt.Println("Queue stopped")
}
