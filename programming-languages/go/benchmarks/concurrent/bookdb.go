package concurrent

import "sync"

type Book struct {
	id     int
	author string
	name   string
}

type BookDB interface {
	GetByID(id int) (*Book, bool)
	Save(book *Book)
	Snapshot() map[int]*Book
}

// BookDBWithSyncMap uses sync.Map to handle concurrency.
type BookDBWithSyncMap struct {
	bookMap sync.Map
}

var _ BookDB = (*BookDBWithSyncMap)(nil)

func NewBookDBWithSyncMap() *BookDBWithSyncMap {
	return &BookDBWithSyncMap{
		bookMap: sync.Map{},
	}
}

func (db *BookDBWithSyncMap) GetByID(id int) (*Book, bool) {
	if res, ok := db.bookMap.Load(id); ok {
		return res.(*Book), true
	}
	return nil, false
}

func (db *BookDBWithSyncMap) Save(book *Book) {
	db.bookMap.Store(book.id, book)
}

func (db *BookDBWithSyncMap) Snapshot() map[int]*Book {
	snapshot := make(map[int]*Book)
	db.bookMap.Range(func(key, value interface{}) bool {
		snapshot[key.(int)] = value.(*Book)
		return true
	})
	return snapshot
}

// BookDBWithLock uses sync.RWMutex to handle concurrency.
type BookDBWithLock struct {
	mu sync.RWMutex

	bookMap map[int]*Book
}

var _ BookDB = (*BookDBWithLock)(nil)

func NewBookDBWithLock() *BookDBWithLock {
	return &BookDBWithLock{
		bookMap: make(map[int]*Book),
	}
}

func (db *BookDBWithLock) GetByID(id int) (*Book, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	book, ok := db.bookMap[id]
	return book, ok
}

func (db *BookDBWithLock) Save(book *Book) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.bookMap[book.id] = book
}

func (db *BookDBWithLock) Snapshot() map[int]*Book {
	db.mu.Lock()
	defer db.mu.Unlock()

	snapshot := make(map[int]*Book, len(db.bookMap))
	for k, v := range db.bookMap {
		snapshot[k] = v
	}
	return snapshot
}

// BookDBWithChannelSync uses channel internally to handle concurrency. But it also meets BookDB's interfaces
// which forces synchronization.
type BookDBWithChannelSync struct {
	bookMap map[int]*Book

	saveChan     chan *SaveBookRequest
	readChan     chan *ReadBookRequest
	snapshotChan chan *SnapshotRequest
	quit         chan bool
}

var _ BookDB = (*BookDBWithChannelSync)(nil)

type ReadBookRequest struct {
	id         int
	resultChan chan *Book
}

type SaveBookRequest struct {
	book       *Book
	resultChan chan bool
}

type SnapshotRequest struct {
	resultChan chan map[int]*Book
}

func NewBookDBWithChannelSync() *BookDBWithChannelSync {
	db := &BookDBWithChannelSync{
		bookMap:      make(map[int]*Book),
		saveChan:     make(chan *SaveBookRequest),
		readChan:     make(chan *ReadBookRequest),
		snapshotChan: make(chan *SnapshotRequest),
		quit:         make(chan bool),
	}
	go db.loop()
	return db
}

func (db *BookDBWithChannelSync) loop() {
	for {
		select {
		case <-db.quit:
			return
		case req := <-db.readChan:
			book, _ := db.bookMap[req.id]
			req.resultChan <- book
		case req := <-db.saveChan:
			db.bookMap[req.book.id] = req.book
			req.resultChan <- true
		case req := <-db.snapshotChan:
			snapshot := make(map[int]*Book, len(db.bookMap))
			for k, v := range db.bookMap {
				snapshot[k] = v
			}
			req.resultChan <- snapshot
		}
	}
}

func (db *BookDBWithChannelSync) Save(book *Book) {
	res := make(chan bool)
	db.saveChan <- &SaveBookRequest{
		book:       book,
		resultChan: res,
	}
	<-res
}

func (db *BookDBWithChannelSync) GetByID(id int) (*Book, bool) {
	res := make(chan *Book)
	db.readChan <- &ReadBookRequest{
		id:         id,
		resultChan: res,
	}
	book := <-res
	if book == nil {
		return nil, false
	}
	return book, true
}

func (db *BookDBWithChannelSync) Snapshot() map[int]*Book {
	res := make(chan map[int]*Book)
	db.snapshotChan <- &SnapshotRequest{
		resultChan: res,
	}
	return <-res
}

func (db *BookDBWithChannelSync) Stop() {
	db.quit <- true
}

// BookDBWithChannelAsync uses channel internally to handle concurrency. It's API does not return value
// directly to caller. Also, it's channel is buffered channel not the blocking channel.
type BookDBWithChannelAsync struct {
	bookMap map[int]*Book

	saveChan     chan *SaveBookCommand
	readChan     chan *ReadBookCommand
	snapshotChan chan *SnapshotCommand
	quit         chan bool
}

type SaveBookCommand struct {
	wg *sync.WaitGroup

	book *Book
}

func (c *SaveBookCommand) execute(db *BookDBWithChannelAsync) {
	db.bookMap[c.book.id] = c.book
	c.wg.Done()
}

type ReadBookCommand struct {
	wg *sync.WaitGroup

	id int
}

func (c *ReadBookCommand) execute(db *BookDBWithChannelAsync) {
	_ = db.bookMap[c.id]
	c.wg.Done()
}

type SnapshotCommand struct {
	wg *sync.WaitGroup
}

func (c *SnapshotCommand) execute(db *BookDBWithChannelAsync) {
	snapshot := make(map[int]*Book, len(db.bookMap))
	for k, v := range db.bookMap {
		snapshot[k] = v
	}
	c.wg.Done()
}

func NewBookDBWithChannelAsync(bufferSize int) *BookDBWithChannelAsync {
	db := &BookDBWithChannelAsync{
		bookMap:      make(map[int]*Book),
		saveChan:     make(chan *SaveBookCommand, bufferSize),
		readChan:     make(chan *ReadBookCommand, bufferSize),
		snapshotChan: make(chan *SnapshotCommand, bufferSize),
		quit:         make(chan bool),
	}
	go db.loop()
	return db
}

func (db *BookDBWithChannelAsync) loop() {
	for {
		select {
		case <-db.quit:
			return
		case cmd := <-db.saveChan:
			cmd.execute(db)
		case cmd := <-db.readChan:
			cmd.execute(db)
		case cmd := <-db.snapshotChan:
			cmd.execute(db)
		}
	}
}

func (db *BookDBWithChannelAsync) Save(book *Book, wg *sync.WaitGroup) {
	db.saveChan <- &SaveBookCommand{
		book: book,
		wg:   wg,
	}
}

func (db *BookDBWithChannelAsync) GetByID(id int, wg *sync.WaitGroup) {
	db.readChan <- &ReadBookCommand{
		id: id,
		wg: wg,
	}
}

func (db *BookDBWithChannelAsync) Snapshot(wg *sync.WaitGroup) {
	db.snapshotChan <- &SnapshotCommand{
		wg: wg,
	}
}

func (db *BookDBWithChannelAsync) Stop() {
	db.quit <- true
}
