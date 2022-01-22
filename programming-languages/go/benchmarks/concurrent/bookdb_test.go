package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

const (
	numReadInstances     = 2000
	numWriteInstances    = 1000
	numSnapshotInstances = 100
	numBookCounts        = 1000
)

func TestBasicFunctionalityForBookDBWithSyncMap(t *testing.T) {
	doTestFunctionalities(t, NewBookDBWithSyncMap())
}

func TestBasicFunctionalityForBookDBWithChannelSync(t *testing.T) {
	doTestFunctionalities(t, NewBookDBWithChannelSync())
}

func TestBasicFunctionalityForBookDBWithLock(t *testing.T) {
	doTestFunctionalities(t, NewBookDBWithLock())
}

func doTestFunctionalities(t *testing.T, db BookDB) {
	// test empty
	if _, ok := db.GetByID(1); ok {
		t.Fatal("Book should not be found")
	}

	// test save and get
	book := &Book{
		id:     1,
		author: "damian",
		name:   "study golang",
	}
	db.Save(book)
	ret, ok := db.GetByID(book.id)
	if !ok {
		t.Fatal("book should be found")
	}
	if book != ret {
		t.Fatal("error finding")
	}

	// test snapshot
	snapshot := db.Snapshot()
	if len(snapshot) != 1 {
		t.Fatal("snapshot count should be 1")
	}
}

func BenchmarkBookDBWithSyncMapAllOperations(b *testing.B) {
	doBenchmarkBookDBRunAllOperations(b, NewBookDBWithSyncMap())
}

func BenchmarkBookDBWithSyncMapOnlyReadAndWrite(b *testing.B) {
	doBenchmarkBookDBRunReadAndWrite(b, NewBookDBWithSyncMap())
}

func BenchmarkBookDBWithChannelSyncAllOperations(b *testing.B) {
	doBenchmarkBookDBRunAllOperations(b, NewBookDBWithChannelSync())
}

func BenchmarkBookDBWithChannelSyncOnlyReadAndWrite(b *testing.B) {
	doBenchmarkBookDBRunReadAndWrite(b, NewBookDBWithChannelSync())
}

func BenchmarkBookDBWithLockAllOperations(b *testing.B) {
	doBenchmarkBookDBRunAllOperations(b, NewBookDBWithLock())
}

func BenchmarkBookDBWithLockOnlyReadAndWrite(b *testing.B) {
	doBenchmarkBookDBRunReadAndWrite(b, NewBookDBWithLock())
}

func doBenchmarkBookDBRunAllOperations(b *testing.B, db BookDB) {
	// init
	var startWG sync.WaitGroup
	startWG.Add(1)

	var endWG sync.WaitGroup
	endWG.Add(numWriteInstances)
	endWG.Add(numReadInstances)
	endWG.Add(numSnapshotInstances)

	// launch write instances
	for i := 0; i < numWriteInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				book := &Book{
					id:     id,
					author: fmt.Sprintf("author-%d", id),
					name:   fmt.Sprintf("name-%d", id),
				}
				db.Save(book)
			}
			endWG.Done()
		}()
	}

	// launch read instances
	for i := 0; i < numReadInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				db.GetByID(id)
			}
			endWG.Done()
		}()
	}

	// launch snapshot instances
	for i := 0; i < numSnapshotInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				_ = db.Snapshot()
			}
			endWG.Done()
		}()
	}

	// start
	b.ResetTimer()
	startWG.Done()
	endWG.Wait()
}

func doBenchmarkBookDBRunReadAndWrite(b *testing.B, db BookDB) {
	// init
	var startWG sync.WaitGroup
	startWG.Add(1)

	var endWG sync.WaitGroup
	endWG.Add(numWriteInstances)
	endWG.Add(numReadInstances)

	// launch write instances
	for i := 0; i < numWriteInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				book := &Book{
					id:     id,
					author: fmt.Sprintf("author-%d", id),
					name:   fmt.Sprintf("name-%d", id),
				}
				db.Save(book)
			}
			endWG.Done()
		}()
	}

	// launch read instances
	for i := 0; i < numReadInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				db.GetByID(id)
			}
			endWG.Done()
		}()
	}

	// start
	b.ResetTimer()
	startWG.Done()
	endWG.Wait()
}

func BenchmarkBookDBWIthChannelAsyncAllOperations(b *testing.B) {
	// init
	db := NewBookDBWithChannelAsync(8192)
	var startWG sync.WaitGroup
	startWG.Add(1)

	var endWG sync.WaitGroup
	endWG.Add(b.N * numWriteInstances)
	endWG.Add(b.N * numReadInstances)
	endWG.Add(b.N * numSnapshotInstances)

	// launch write instances
	for i := 0; i < numWriteInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				book := &Book{
					id:     id,
					author: fmt.Sprintf("author-%d", id),
					name:   fmt.Sprintf("name-%d", id),
				}
				db.Save(book, &endWG)
			}
		}()
	}

	// launch read instances
	for i := 0; i < numReadInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				db.GetByID(id, &endWG)
			}
		}()
	}

	// launch snapshot instances
	for i := 0; i < numSnapshotInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				db.Snapshot(&endWG)
			}
		}()
	}

	// start
	b.ResetTimer()
	startWG.Done()
	endWG.Wait()
}

func BenchmarkBookDBWIthChannelAsyncOnlyReadAndWrite(b *testing.B) {
	// init
	db := NewBookDBWithChannelAsync(8192)
	var startWG sync.WaitGroup
	startWG.Add(1)

	var endWG sync.WaitGroup
	endWG.Add(b.N * numWriteInstances)
	endWG.Add(b.N * numReadInstances)

	// launch write instances
	for i := 0; i < numWriteInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				book := &Book{
					id:     id,
					author: fmt.Sprintf("author-%d", id),
					name:   fmt.Sprintf("name-%d", id),
				}
				db.Save(book, &endWG)
			}
		}()
	}

	// launch read instances
	for i := 0; i < numReadInstances; i++ {
		go func() {
			startWG.Wait()
			for i := 0; i < b.N; i++ {
				id := rand.Intn(numBookCounts)
				db.GetByID(id, &endWG)
			}
		}()
	}

	// start
	b.ResetTimer()
	startWG.Done()
	endWG.Wait()
}
