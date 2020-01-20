package concurrent_singleton

var addChan chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

type singleton struct {}

var instance singleton

func GetInstance() *singleton{
	return &instance
}

func init() {
	var count int
	go func(addCh <-chan bool, getCount <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count ++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addChan, getCountCh, quitCh)
}

func (s *singleton) AddOne(){
	addChan <- true
}

func (s *singleton) GetCount() int{
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop(){
	quitCh <- true
	close(addChan)
	close(getCountCh)
	close(quitCh)
}