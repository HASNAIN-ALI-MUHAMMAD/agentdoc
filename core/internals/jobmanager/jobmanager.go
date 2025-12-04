package jobmanager

import (
	"context"
	"fmt"
	"sync"
)


var TaskManager = &BackgrondManager{
	funcCancels: make(map[string]context.CancelFunc),
}

type BackgrondManager struct{
	wg sync.WaitGroup
	mu sync.Mutex

	funcCancels map[string]context.CancelFunc
}

func (bm *BackgrondManager) RunTask(taskName string,task func()){
	bm.wg.Add(1)
	go func (t string,taskFunc func()){
		defer bm.wg.Done()
		fmt.Printf("Background task initializing: %v\n",taskName);
		taskFunc()
		fmt.Printf("Background task done: %v\n",taskName);	
	}(taskName,task)
}

func (bm *BackgrondManager) RunCancellableTask(taskId string,task func(ctx context.Context)){
	bm.wg.Add(1)
	ctx,cancel := context.WithCancel(context.Background())
	bm.mu.Lock()
	bm.funcCancels[taskId] = cancel
	bm.mu.Unlock()

	go func(t string,taskFunc func(ctx context.Context)){
		defer bm.wg.Done()
		defer bm.CleanUp(taskId)
		fmt.Printf("Cancellable background task initializing: %v\n",taskId);
		taskFunc(ctx)
		fmt.Printf("Cancellable background task done: %v\n",taskId);
	}(taskId,task)
}

func (bm *BackgrondManager) StopTask(taskId string){
	bm.mu.Lock()
	defer bm.mu.Unlock()
	if cancel,exists := bm.funcCancels[taskId]; exists{
		cancel()
		fmt.Printf("Background task cancelled: %v\n",taskId);
	}
}

func (bm *BackgrondManager) ShutDown(){
	bm.mu.Lock()
	defer bm.mu.Unlock()
	for _,cancel := range bm.funcCancels{
		cancel()
	}
	bm.wg.Wait()
}

func (bm *BackgrondManager) CleanUp(taskId string){
	bm.mu.Lock()
	defer bm.mu.Unlock()
	delete(bm.funcCancels,taskId)
}



func (bm *BackgrondManager) WaitAll(){
	bm.wg.Wait()
}