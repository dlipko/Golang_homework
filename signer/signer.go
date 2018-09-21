package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	// группа для ожидания завершения всех задач
	wJobs := &sync.WaitGroup{}

	in := make(chan interface{}, MaxInputDataLen)
	out := make(chan interface{}, MaxInputDataLen)
	for _, j := range jobs {
		wJobs.Add(1)
		// запуск обертки над задачей
		go jobCover(j, in, out, wJobs)
		in = out
		out = make(chan interface{}, MaxInputDataLen)
	}
	wJobs.Wait()
}

// запускает задачу, ожидает её выполнения, закрывает канал
func jobCover(j job, in chan interface{}, out chan interface{}, wJob *sync.WaitGroup) {
	defer wJob.Done()
	j(in, out)
	// для завершения цикла for по этому каналу в следующей задаче
	close(out)
}

// запускает горутины для рассчета singleHash
func SingleHash(in, out chan interface{}) {
	// группа ожидания завершения рассчета singleHash
	// для всех входных данных
	wSingleHash := &sync.WaitGroup{}
	// мьютекс для запрета одновременного запуска DataSignerMd5
	mdMutex := &sync.Mutex{}

	for data := range in {
		// попытка проверки входных данных
		if number, ok := data.(int); ok {
			data = strconv.Itoa(number)
		}
		if data, ok := data.(string); ok {
			wSingleHash.Add(1)
			go countSingleHash(data, out, wSingleHash, mdMutex)
		}
	}

	wSingleHash.Wait()
}

// расчитывает crc32(data) + crc32(md5(data))
func countSingleHash(data string, out chan interface{}, 
						wSingleHash *sync.WaitGroup, mdMutex *sync.Mutex) {
	defer wSingleHash.Done()

	crcCh := make(chan string, 1)
	// расчет crc(data)  в отдельной горутине
	go coverCrc(crcCh, data)

	mdMutex.Lock()
	mdData := DataSignerMd5(data)
	mdMutex.Unlock()
	crcMdData := DataSignerCrc32(mdData)
	crcData := <-crcCh
	out <- (crcData + "~" + crcMdData)
}

// обертка над DataSignerCrc32 для запуска в отдельной горутине
func coverCrc(crcCh chan<- string, data string) {
	crcCh <- DataSignerCrc32(data)
}

// расчитывает crc32(th + data)
func MultiHash(in, out chan interface{}) {
	wMultiHach := &sync.WaitGroup{}
	for data := range in {
		wMultiHach.Add(1)
		go countMultiHash(data.(string), out, wMultiHach)
	}
	wMultiHach.Wait()
}

func countMultiHash(data string, out chan interface{}, wMultiHach *sync.WaitGroup) {
	defer wMultiHach.Done()

	th := 6;
	// группа для ожидания завершения расчета crc(th + data), th = 0..5
	wTh := &sync.WaitGroup{}
	dataCh := make(chan string, 6)
	numberCh := make(chan int, 6)
	mutex := &sync.Mutex{}


	for i := 0; i < th; i++ {
		wTh.Add(1)
		go coverCrcWaitGroup(i, data, dataCh, numberCh, mutex, wTh)
	}
	wTh.Wait()

	slice := make([]string, th, th)
	for i := 0; i < th; i++ {
		data := <-dataCh
		number := <-numberCh
		slice[number] = data
	}

	out <- strings.Join(slice, "")
}

// обертка над DataSignerCrc32 для запуска в отдельной горутине с waitGroup
func coverCrcWaitGroup(number int, data string, dataCh chan<- string,
					 numberCh chan<- int, mutex *sync.Mutex, wTh *sync.WaitGroup) {
	defer wTh.Done()

	outData := DataSignerCrc32(strconv.Itoa(number)+data)
	mutex.Lock()
	dataCh <- outData
	numberCh <- number
	mutex.Unlock()
}

func CombineResults(in, out chan interface{}) {
	var result []string
	for data := range in {
		result = append(result, data.(string))
	}

	sort.Strings(result)
	str := strings.Join(result, "_")
	out <- str
}
