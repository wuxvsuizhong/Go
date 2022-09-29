package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func work(t *task, fop *fileop) {
	defer t.wg.Done()

	for {
		select {
		case <-t.ctx.Done():
			fop.filelock.Lock()
			_, err := fop.writer.WriteString(fmt.Sprintf("I'm %s routine, i received stop singnal...\n", t.ctx.Value("key")))
			if err != nil {
				fop.filelock.Unlock()
				return
			}
			// fmt.Printf("I'm %s routine, i received stop singnal...\n", t.ctx.Value("key"))
			fop.writer.Flush()
			fop.filelock.Unlock()
			return
		default:
			fop.filelock.Lock()
			_, err := fop.writer.WriteString(fmt.Sprintf("I'm %s routine, i'm working...\n", t.ctx.Value("key")))
			if err != nil {
				fop.filelock.Unlock()
				return
			}
			fop.writer.Flush()
			fop.filelock.Unlock()
			time.Sleep(time.Second * 3)
		}
	}
}

func getctx(ctxs *[]*con) func() context.Context {
	single_item := new(con)
	single_item.ctx, single_item.calcelfunc = context.WithCancel(context.Background())
	double_item := new(con)
	double_item.ctx, double_item.calcelfunc = context.WithCancel(context.Background())
	single_item.ctx = context.WithValue(single_item.ctx, "key", "single")
	double_item.ctx = context.WithValue(double_item.ctx, "key", "double")

	*ctxs = append(*ctxs, single_item)
	*ctxs = append(*ctxs, double_item)

	var flag = 0
	return func() context.Context {
		if flag == 0 {
			flag = 1
			return (*ctxs)[0].ctx
		}
		flag = 0
		return (*ctxs)[1].ctx
	}
}

type workfunc func(*task, *fileop)

type fileop struct {
	writer   *bufio.Writer
	filelock *sync.Mutex
}

func routineHandle(ctxs *[]*con, ch chan workfunc, wg *sync.WaitGroup, mctx context.Context) {
	defer wg.Done()

	fileObj, _ := os.OpenFile("./outlog.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)

	var flock sync.Mutex
	fop := fileop{
		writer:   wr,
		filelock: &flock,
	}

	ctxhandle := getctx(ctxs)
	for {
		select {
		case w := <-ch:
			t := task{wg, ctxhandle()}
			wg.Add(1)
			go w(&t, &fop)
		case <-mctx.Done():
			return
		default:
			time.Sleep(time.Second * 5)
			ch <- work

		}
	}
}

type task struct {
	wg  *sync.WaitGroup
	ctx context.Context
}

type con struct {
	ctx        context.Context
	calcelfunc context.CancelFunc
}

func main() {
	var wg sync.WaitGroup
	var ctxs []*con
	ch := make(chan workfunc, 4)
	wg.Add(1)
	maincontext, cancel := context.WithCancel(context.Background())
	go routineHandle(&ctxs, ch, &wg, maincontext)
	var control int
	for {
		fmt.Print("请输入数字(偶数停止value为偶数的routine，奇数同理，小于0退出所有routine):")
		fmt.Scanln(&control)
		if control < 0 {
			for _, c := range ctxs {
				c.calcelfunc()
			}
			cancel()

			break
		} else if control%2 == 1 {
			for _, c := range ctxs {
				if c.ctx.Value("key") == "single" {
					c.calcelfunc()
				}
			}
		} else {
			for _, c := range ctxs {
				if c.ctx.Value("key") == "double" {
					c.calcelfunc()
				}
			}
		}
	}

	wg.Wait()
}
